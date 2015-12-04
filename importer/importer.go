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

var patientXPath = xpath.Compile("/cda:ClinicalDocument/cda:recordTarget/cda:patientRole/cda:patient")
var encounterXPath = xpath.Compile("//cda:encounter[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.23']")
var diagnosisXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.11']")
var timeLowXPath = xpath.Compile("cda:effectiveTime/cda:low/@value")
var timeHighXPath = xpath.Compile("cda:effectiveTime/cda:high/@value")
var lastNameXPath = xpath.Compile("cda:name/cda:family")
var firstNameXPath = xpath.Compile("cda:name/cda:given")
var birthTimeXPath = xpath.Compile("cda:birthTime/@value")
var genderXPath = xpath.Compile("cda:administrativeGenderCode/@code")
var raceXPath = xpath.Compile("cda:raceCode/@code")
var raceCodeSetXPath = xpath.Compile("cda:raceCode/@codeSystemName")
var ethnicityXPath = xpath.Compile("cda:ethnicGroupCode/@code")
var ethnicityCodeSetXPath = xpath.Compile("cda:ethnicGroupCode/@codeSystemName")
var codeXPath = xpath.Compile("cda:code/@code")
var codeCodeSetXPath = xpath.Compile("cda:code/@codeSystem")
var valueCodeXPath = xpath.Compile("cda:value/@code")
var valueCodeSetXPath = xpath.Compile("cda:value/@codeSystem")
var textXPath = xpath.Compile("cda:text")
var idRootXPath = xpath.Compile("cda:id/@root")
var idExtXPath = xpath.Compile("cda:id/@extension")

func main() {}

func Read_patient(path string) string {

	data, err := ioutil.ReadFile(path)
	util.CheckErr(err)

	doc, err := xml.Parse(data, nil, nil, 0, xml.DefaultEncodingBytes)
	util.CheckErr(err)
	defer doc.Free()

	xp := doc.DocXPathCtx()
	xp.RegisterNamespace("cda", "urn:hl7-org:v3")

	// fmt.Println("\nPatient Name:\n")

	patientElements, err := doc.Root().Search(patientXPath)
	util.CheckErr(err)
	patientElement := patientElements[0]
	patient := &models.Record{}

	ExtractDemographics(patient, patientElement)
	ExtractEncounters(patient, doc.Root())
	ExtractDiagnoses(patient, doc.Root())

	patientJSON, err := json.Marshal(patient)
	if err != nil {
		fmt.Println(err)
	}

	return string(patientJSON)

}

func ExtractDemographics(patient *models.Record, patientElement xml.Node) {
	patient.First = FirstElementContent(firstNameXPath, patientElement)
	patient.Last = FirstElementContent(lastNameXPath, patientElement)
	patient.Gender = FirstElementContent(genderXPath, patientElement)
	patient.Birthdate = GetTimestamp(birthTimeXPath, patientElement)
	patient.Race.Code = FirstElementContent(raceXPath, patientElement)
	patient.Race.CodeSet = FirstElementContent(raceCodeSetXPath, patientElement)
	patient.Ethnicity.Code = FirstElementContent(ethnicityXPath, patientElement)
	patient.Ethnicity.CodeSet = FirstElementContent(ethnicityCodeSetXPath, patientElement)
}

func ExtractBasicEntry(entry *models.Entry, entryElement xml.Node, codePath *xpath.Expression, codeSetPath *xpath.Expression, oid string) {
	//extract cda identifier
	entry.ID = models.CDAIdentifier{Root: FirstElementContent(idRootXPath, entryElement), Extension: FirstElementContent(idExtXPath, entryElement)}

	//extract codes
	code := FirstElementContent(codePath, entryElement)
	codeSystem := models.CodeSystemFor(FirstElementContent(codeSetPath, entryElement))
	entry.Codes = map[string][]string{
		codeSystem: []string{code},
	}

	//extract dates
	entry.StartTime = GetTimestamp(timeLowXPath, entryElement)
	entry.EndTime = GetTimestamp(timeHighXPath, entryElement)

	//extract description
	entry.Description = FirstElementContent(textXPath, entryElement)

}

func ExtractEncounters(record *models.Record, xmlNode xml.Node) {
	encounterElements, err := xmlNode.Search(encounterXPath)
	util.CheckErr(err)
	encounters := make([]models.Encounter, len(encounterElements))
	for i, encounterElement := range encounterElements {
		encounter := models.Encounter{}
		oid := "2.16.840.1.113883.3.560.1.79"
		ExtractBasicEntry(&encounter.Entry, encounterElement, codeXPath, codeCodeSetXPath, oid)
		encounters[i] = encounter
	}
	record.Encounters = encounters
}

func ExtractDiagnoses(record *models.Record, xmlNode xml.Node) {
	diagnosisElements, err := xmlNode.Search(diagnosisXPath)
	util.CheckErr(err)
	diagnoses := make([]models.Diagnosis, len(diagnosisElements))
	for i, diagnosisElement := range diagnosisElements {
		diagnosis := models.Diagnosis{}
		oid := "2.16.840.1.113883.3.560.1.2"
		ExtractBasicEntry(&diagnosis.Entry, diagnosisElement, valueCodeXPath, valueCodeSetXPath, oid)
		diagnoses[i] = diagnosis
	}
	record.Diagnoses = diagnoses
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
