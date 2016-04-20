package importer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
)

func main() {}

func Read_patient(path string) string {

	data, err := ioutil.ReadFile(path)
	util.CheckErr(err)

	doc, err := xml.Parse(data, nil, nil, 0, xml.DefaultEncodingBytes)
	util.CheckErr(err)
	defer doc.Free()

	xp := doc.DocXPathCtx()
	xp.RegisterNamespace("cda", "urn:hl7-org:v3")
	xp.RegisterNamespace("stdc", "urn:hl7-org:sdtc")

	var patientXPath = xpath.Compile("/cda:ClinicalDocument/cda:recordTarget/cda:patientRole/cda:patient")
	patientElements, err := doc.Root().Search(patientXPath)
	util.CheckErr(err)
	patientElement := patientElements[0]
	patient := &models.Record{}
	ExtractDemographics(patient, patientElement)

	//encounter performed
	var encounterPerformedXPath = xpath.Compile("//cda:encounter[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.23']")
	rawEncountersPerformed := ExtractSection(patientElement, encounterPerformedXPath, EncounterPerformedExtractor, "2.16.840.1.113883.3.560.1.79")
	patient.Encounters = make([]models.Encounter, len(rawEncountersPerformed))
	for i := range rawEncountersPerformed {
		patient.Encounters[i] = rawEncountersPerformed[i].(models.Encounter)
	}

	//encounter ordered
	var encounterOrderXPath = xpath.Compile("//cda:encounter[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.22']")
	rawEncounterOrders := ExtractSection(patientElement, encounterOrderXPath, EncounterOrderExtractor, "")
	for i := range rawEncounterOrders {
		patient.Encounters = append(patient.Encounters, rawEncounterOrders[i].(models.Encounter))
	}

	//diagnosis active
	var diagnosisActiveXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.11']")
	rawDiagnosesActive := ExtractSection(patientElement, diagnosisActiveXPath, DiagnosisActiveExtractor, "2.16.840.1.113883.3.560.1.2")
	patient.Diagnoses = make([]models.Diagnosis, len(rawDiagnosesActive))
	for i := range rawDiagnosesActive {
		patient.Diagnoses[i] = rawDiagnosesActive[i].(models.Diagnosis)
	}

	//diagnosis inactive
	var diagnosisInactiveXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.13']")
	rawDiagnosesInactive := ExtractSection(patientElement, diagnosisInactiveXPath, DiagnosisInactiveExtractor, "2.16.840.1.113883.3.560.1.2")
	for i := range rawDiagnosesInactive {
		patient.Diagnoses = append(patient.Diagnoses, rawDiagnosesInactive[i].(models.Diagnosis))
	}

	//lab results
	var labResultXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.40']")
	rawLabResults := ExtractSection(patientElement, labResultXPath, LabResultExtractor, "")
	patient.LabResults = make([]models.LabResult, len(rawLabResults))
	for i := range rawLabResults {
		patient.LabResults[i] = rawLabResults[i].(models.LabResult)
	}

	//lab orders
	var labOrderXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.37']")
	rawLabOrders := ExtractSection(patientElement, labOrderXPath, LabOrderExtractor, "2.16.840.1.113883.3.560.1.50")
	for i := range rawLabOrders {
		patient.LabResults = append(patient.LabResults, rawLabOrders[i].(models.LabResult))
	}

	//insurance provider
	var insuranceProviderXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.55']")
	rawInsuranceProviders := ExtractSection(patientElement, insuranceProviderXPath, InsuranceProviderExtractor, "2.16.840.1.113883.3.560.1.405")
	patient.InsuranceProviders = make([]models.InsuranceProvider, len(rawInsuranceProviders))
	for i := range rawInsuranceProviders {
		patient.InsuranceProviders[i] = rawInsuranceProviders[i].(models.InsuranceProvider)
	}

	// diagnostic study order
	var diagnosticStudyOrderXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.17']")
	rawDiagnosticStudyOrders := ExtractSection(patientElement, diagnosticStudyOrderXPath, DiagnosticStudyOrderExtractor, "2.16.840.1.113883.3.560.1.40")
	patient.Procedures = make([]models.Procedure, len(rawDiagnosticStudyOrders))
	for i := range rawDiagnosticStudyOrders {
		patient.Procedures[i] = rawDiagnosticStudyOrders[i].(models.Procedure)
	}

	// transfer from
	var transferFromXPath = xpath.Compile("//cda:encounter[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.81']")
	rawTransferFroms := ExtractSection(patientElement, transferFromXPath, TransferFromExtractor, "2.16.840.1.113883.3.560.1.71")
	for i := range rawTransferFroms {
		patient.Encounters = append(patient.Encounters, rawTransferFroms[i].(models.Encounter))
	}

	patientJSON, err := json.Marshal(patient)
	if err != nil {
		fmt.Println(err)
	}

	return string(patientJSON)

}

func ExtractSection(xmlNode xml.Node, sectionXpath *xpath.Expression, extractor EntryExtractor, oid string) []interface{} {
	sectionElements, err := xmlNode.Search(sectionXpath)
	util.CheckErr(err)

	entries := make([]interface{}, len(sectionElements))
	for i, entryElement := range sectionElements {
		entries[i] = ExtractEntry(entryElement, oid, extractor)
	}
	return entries
}

type EntryExtractor func(*models.Entry, xml.Node) interface{}

func ExtractEntry(entryElement xml.Node, oid string, extractor EntryExtractor) interface{} {
	var entry models.Entry

	//extract cda identifier
	var idRootXPath = xpath.Compile("cda:id/@root")
	var idExtXPath = xpath.Compile("cda:id/@extension")
	entry.ID = models.CDAIdentifier{Root: FirstElementContent(idRootXPath, entryElement), Extension: FirstElementContent(idExtXPath, entryElement)}

	//extract dates
	ExtractDates(&entry, entryElement)

	//extract description
	var textXPath = xpath.Compile("cda:text")
	entry.Description = FirstElementContent(textXPath, entryElement)

	//set oid
	entry.Oid = oid

	//create code map
	entry.Codes = map[string][]string{}

	fullEntry := extractor(&entry, entryElement)
	return fullEntry
}

func ExtractCodes(coded *models.Coded, entryElement xml.Node, codePath *xpath.Expression) {
	codeElements, err := entryElement.Search(codePath)
	util.CheckErr(err)
	for _, codeElement := range codeElements {
		coded.AddCodeIfPresent(codeElement)
		translationElements, err := codeElement.Search("cda:translation")
		util.CheckErr(err)
		for _, translationElement := range translationElements {
			coded.AddCodeIfPresent(translationElement)
		}
	}
}

func ExtractDates(entry *models.Entry, entryElement xml.Node) {
	var timeLowXPath = xpath.Compile("cda:effectiveTime/cda:low/@value")
	var timeHighXPath = xpath.Compile("cda:effectiveTime/cda:high/@value")
	entry.StartTime = GetTimestamp(timeLowXPath, entryElement)
	entry.EndTime = GetTimestamp(timeHighXPath, entryElement)
}

func ExtractReason(encounter *models.Encounter, entryElement xml.Node) {
	var reasonXPath = xpath.Compile("cda:entryRelationship[@typeCode='RSON']/cda:observation")
	reasonElements, err := entryElement.Search(reasonXPath)
	util.CheckErr(err)
	if len(reasonElements) > 0 {
		reasonElement := reasonElements[0]

		//extract reason value code
		var valueCodeXPath = xpath.Compile("cda:value/@code")
		var valueCodeSystemXPath = xpath.Compile("cda:value/@codeSystem")
		valueCode := FirstElementContent(valueCodeXPath, reasonElement)
		valueCodeSystem := models.CodeSystemFor(FirstElementContent(valueCodeSystemXPath, reasonElement))
		encounter.Reason.Code = valueCode
		encounter.Reason.CodeSystem = valueCodeSystem
		encounter.Reason.CodeSystemName = valueCodeSystem
	}
}

func ExtractValues(entry *models.Entry, entryElement xml.Node, valuePath *xpath.Expression) {
	valueElements, err := entryElement.Search(valuePath)
	util.CheckErr(err)
	if len(valueElements) > 0 {
		for _, valueElement := range valueElements {
			value := valueElement.Attribute("value")
			code := valueElement.Attribute("code")
			if value != nil {
				unit := valueElement.Attribute("unit").String()
				scalar, err := strconv.ParseInt(value.String(), 10, 64)
				util.CheckErr(err)
				entry.AddScalarValue(scalar, unit)
			} else if code != nil {
				val := models.ResultValue{}
				val.AddCodeIfPresent(valueElement)
				var timeLowXPath = xpath.Compile("cda:effectiveTime/cda:low/@value")
				var timeHighXPath = xpath.Compile("cda:effectiveTime/cda:high/@value")
				val.StartTime = GetTimestamp(timeLowXPath, entryElement)
				val.EndTime = GetTimestamp(timeHighXPath, entryElement)
				entry.Values = append(entry.Values, val)
			} else {
				unit := valueElement.Attribute("unit")
				valString := valueElement.Content()
				if unit == nil {
					entry.AddStringValue(valString, "")
				} else {
					entry.AddStringValue(valString, unit.String())
				}
			}
		}
	}
}

func FirstElementContent(xpath *xpath.Expression, xmlNode xml.Node) string {
	resultNodes, err := xmlNode.Search(xpath)
	util.CheckErr(err)
	if len(resultNodes) > 0 {
		firstNode := resultNodes[0]
		return firstNode.Content()
	}
	return ""
}

func GetTimestamp(xpath *xpath.Expression, xmlNode xml.Node) int64 {
	attrValue := FirstElementContent(xpath, xmlNode)
	if attrValue != "" {
		return TimestampToSeconds(attrValue)
	}
	return 0
}

func TimestampToSeconds(timestamp string) int64 {
	year, _ := strconv.ParseInt(timestamp[0:4], 10, 32)
	month, _ := strconv.ParseInt(timestamp[4:6], 10, 32)
	day, _ := strconv.ParseInt(timestamp[6:8], 10, 32)
	hour, _ := strconv.ParseInt(timestamp[8:10], 10, 32)
	minute, _ := strconv.ParseInt(timestamp[10:12], 10, 32)
	second, _ := strconv.ParseInt(timestamp[12:14], 10, 32)
	desiredDate := time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, time.UTC)
	return desiredDate.Unix()
}
