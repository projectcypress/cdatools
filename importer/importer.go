package importer

import (
	"encoding/json"
	"fmt"
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
	"io/ioutil"
	"strconv"
	"time"
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

	var patientXPath = xpath.Compile("/cda:ClinicalDocument/cda:recordTarget/cda:patientRole/cda:patient")
	patientElements, err := doc.Root().Search(patientXPath)
	util.CheckErr(err)
	patientElement := patientElements[0]
	patient := &models.Record{}

	ExtractDemographics(patient, patientElement)

	var encounterPerformedXPath = xpath.Compile("//cda:encounter[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.23']")
	rawEncountersPerformed := ExtractSection(patientElement, encounterPerformedXPath, EncounterPerformedExtractor, "2.16.840.1.113883.3.560.1.79")
	patient.Encounters = make([]models.Encounter, len(rawEncountersPerformed))
	for i := range rawEncountersPerformed {
		patient.Encounters[i] = rawEncountersPerformed[i].(models.Encounter)
	}

	var encounterOrderXpath = xpath.Compile("//cda:encounter[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.22']")
	rawEncounterOrders := ExtractSection(patientElement, encounterOrderXPath, EncounterOrderExtractor, "")
	for i := range rawEncounterOrders {
		patient.Encounters = append(patient.Encounters, rawEncounterOrders[i].(models.Encounter))
	}

	var diagnosisActiveXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.11']")
	rawDiagnosesActive := ExtractSection(patientElement, diagnosisActiveXPath, DiagnosisActiveExtractor, "2.16.840.1.113883.3.560.1.2")
	patient.Diagnoses = make([]models.Diagnosis, len(rawDiagnosesActive))
	for i := range rawDiagnosesActive {
		patient.Diagnoses[i] = rawDiagnosesActive[i].(models.Diagnosis)
	}

	var diagnosisInactiveXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.13']")
	rawDiagnosesInactive := ExtractSection(patientElement, diagnosisInactiveXPath, DiagnosisInactiveExtractor, "2.16.840.1.113883.3.560.1.2")
	patient.Diagnoses = make([]models.Diagnosis, len(rawDiagnosesInactive))
	for i := range rawDiagnosesInactive {
		patient.Diagnoses[i] = rawDiagnosesInactive[i].(models.Diagnosis)
	}

	patientJSON, err := json.Marshal(patient)
	if err != nil {
		fmt.Println(err)
	}

	return string(patientJSON)

}

func ExtractDemographics(patient *models.Record, patientElement xml.Node) {
	var firstNameXPath = xpath.Compile("cda:name/cda:given")
	patient.First = FirstElementContent(firstNameXPath, patientElement)
	var lastNameXPath = xpath.Compile("cda:name/cda:family")
	patient.Last = FirstElementContent(lastNameXPath, patientElement)
	var genderXPath = xpath.Compile("cda:administrativeGenderCode/@code")
	patient.Gender = FirstElementContent(genderXPath, patientElement)
	var birthTimeXPath = xpath.Compile("cda:birthTime/@value")
	patient.Birthdate = GetTimestamp(birthTimeXPath, patientElement)
	var raceXPath = xpath.Compile("cda:raceCode/@code")
	patient.Race.Code = FirstElementContent(raceXPath, patientElement)
	var raceCodeSetXPath = xpath.Compile("cda:raceCode/@codeSystemName")
	patient.Race.CodeSet = FirstElementContent(raceCodeSetXPath, patientElement)
	var ethnicityXPath = xpath.Compile("cda:ethnicGroupCode/@code")
	patient.Ethnicity.Code = FirstElementContent(ethnicityXPath, patientElement)
	var ethnicityCodeSetXPath = xpath.Compile("cda:ethnicGroupCode/@codeSystemName")
	patient.Ethnicity.CodeSet = FirstElementContent(ethnicityCodeSetXPath, patientElement)
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

	fullEntry := extractor(&entry, entryElement)
	return fullEntry
}

func ExtractCodes(entry *models.Entry, entryElement xml.Node, codePath *xpath.Expression, codeSetPath *xpath.Expression) {
	code := FirstElementContent(codePath, entryElement)
	codeSystem := models.CodeSystemFor(FirstElementContent(codeSetPath, entryElement))
	entry.Codes = map[string][]string{
		codeSystem: []string{code},
	}
}

func ExtractDates(entry *models.Entry, entryElement xml.Node) {
	var timeLowXPath = xpath.Compile("cda:effectiveTime/cda:low/@value")
	var timeHighXPath = xpath.Compile("cda:effectiveTime/cda:high/@value")
	entry.StartTime = GetTimestamp(timeLowXPath, entryElement)
	entry.EndTime = GetTimestamp(timeHighXPath, entryElement)
}

func ExtractSeverity(diagnosis *models.Diagnosis, entryElement xml.Node, severityCodeXPath *xpath.Expression, severityCodeSetXPath *xpath.Expression) {
	severityCode := FirstElementContent(severityCodeXPath, entryElement)
	severityCodeSystem := models.CodeSystemFor(FirstElementContent(severityCodeSetXPath, entryElement))
	diagnosis.Severity = map[string][]string{
		severityCodeSystem: []string{severityCode},
	}
}

func ExtractReason(encounter *models.Encounter, entryElement xml.Node) {
	var reasonXPath = xpath.Compile("cda:entryRelationship[@typeCode='RSON']/cda:observation")
	reasonElements, err := xmlNode.Search(xpath)
	util.CheckErr(err)
	if len(reasonElements) > 0 {
		reasonElement := resultNodes[0]
		encounter.Reason = *models.Reason

		//extract reason code
		var reasonCodePath = xpath.Compile("cda:code/@code")
		var reasonCodeSetPath = xpath.Compile("cda:code/@codeSystem")
		ExtractCodes(encounter.Reason, reasonElement, reasonCodePath, reasonCodeSetPath)

		//extract dates
		ExtractDates(encounter.Reason, reasonElement)

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

func EncounterPerformedExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	encounter := models.Encounter{}
	encounter.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:code/@code")
	var codeSetPath = xpath.Compile("cda:code/@codeSystem")
	ExtractCodes(&encounter.Entry, entryElement, codePath, codeSetPath)

	//set discharge time
	encounter.DischargeTime = encounter.Entry.EndTime

	//extract discharge disposition
	var dischargeDispositionCodeXPath = xpath.Compile("stdc:dischargeDispositionCode/@code")
	var dischargeDispositionCodeSystemXPath = xpath.Compile("stdc:dischargeDispositionCode/@codeSystem")
	dischargeDispositionCode := FirstElementContent(dischargeDispositionCodeXPath, entryElement)
	dischargeDispositionCodeSystemOid := FirstElementContent(dischargeDispositionCodeSystemXPath, entryElement)
	dischargeDispositionCodeSystem := models.CodeSystemFor(dischargeDispositionCodeSystemOid)
	encounter.DischargeDisposition = map[string][]string{
		"code":          dischargeDispositionCode,
		"codeSystem":    dischargeDispositionCodeSystem,
		"codeSystemOid": dischargeDispositionCodeSystemOid,
	}

	return encounter
}

func EncounterOrderExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	encounterOrder := models.Encounter{}
	encounterOrder.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:code/@code")
	var codeSetPath = xpath.Compile("cda:code/@codeSystem")
	ExtractCodes(&encounterOrder.Entry, entryElement, codePath, codeSetPath)

	//extract order specific dates
	var orderTimeXPath = xpath.Compile("cda:author/cda:time/@value")
	encounterOrder.StartTime = GetTimestamp(orderTimeXPath, entryElement)
	encounterOrder.EndTime = GetTimestamp(orderTimeXPath, entryElement)

	return encounterOrder
}

func DiagnosisActiveExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	diagnosisActive := models.Diagnosis{}
	diagnosisActive.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:value/@code")
	var codeSetPath = xpath.Compile("cda:value/@codeSystem")
	ExtractCodes(&diagnosisActive.Entry, entryElement, codePath, codeSetPath)

	//extract severity
	var severityCodeXPath = xpath.Compile("cda:entryRelationship/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.22.4.8']/cda:value/@code")
	var severityCodeSetXPath = xpath.Compile("cda:entryRelationship/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.22.4.8']/cda:value/@codeSystem")
	ExtractSeverity(&diagnosisActive, entryElement, severityXPath)

	return diagnosisActive
}

func DiagnosisInactiveExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	diagnosisInactive := models.Diagnosis{}
	diagnosisInactive.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:value/@code")
	var codeSetPath = xpath.Compile("cda:value/@codeSystem")
	ExtractCodes(&diagnosisActive.Entry, entryElement, codePath, codeSetPath)

	return diagnosisInactive
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
	desiredDate := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)
	return desiredDate.Unix()
}
