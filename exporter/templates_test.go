package exporter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
	"text/template"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
	"github.com/projectcypress/cdatools/fixtures"
	"github.com/projectcypress/cdatools/models"
	"github.com/stretchr/testify/assert"
)

func TestCodeTemplate(t *testing.T) {
	// tag name is "code", has preferred code, attribute is not "codes"
	preferredCode := models.Concept{Code: "my_code", CodeSystem: "SNOMED-CT"}
	codeDisplay := models.CodeDisplay{CodeType: "my_code_type", TagName: "code", Attribute: "my_attr", PreferredCode: preferredCode, ExtraContent: "my_extra_content=\"extra_content_value\""}
	rootNode := xmlCodeRootNode(codeDisplay)
	assertXPath(t, rootNode, "//code", map[string]string{"code": "my_code", "codeSystem": "2.16.840.1.113883.6.96", "my_extra_content": "extra_content_value"}, nil)
	assertNoXPath(t, rootNode, "//code/originalText")

	// tag name is not "code"
	codeDisplay = models.CodeDisplay{CodeType: "my_code_type", TagName: "other_tag_name"}
	assertXPath(t, xmlCodeRootNode(codeDisplay), "//other_tag_name", nil, nil)

	// tag name is "code", no preferred code, exclude null flavor true
	codeDisplay = models.CodeDisplay{CodeType: "my_code_type", TagName: "code", ExcludeNullFlavor: true}
	assertXPath(t, xmlCodeRootNode(codeDisplay), "//code", nil, []string{"excludeNullFlavor"})

	// tag name is "code", no preferred code, exclude null flavor false
	codeDisplay = models.CodeDisplay{CodeType: "my_code_type", TagName: "code", Attribute: "my_attr", ExcludeNullFlavor: false, ExtraContent: "extra_stuff"}
	assertXPath(t, xmlCodeRootNode(codeDisplay), "//code", map[string]string{"nillFlavor": "UNK"}, nil)

	// attribute is "codes"
	codeDisplay = models.CodeDisplay{CodeType: "my_code_type", Attribute: "codes"}
	rootNode = xmlCodeRootNode(codeDisplay)
	assertXPath(t, rootNode, "//code/originalText", nil, nil)
	assertContent(t, rootNode, "//code/originalText", "my lil description")

	// laterality is included
	ei := models.EntryInfo{}
	setMapDataCriteria(&ei)
	codeDisplay = models.CodeDisplay{CodeType: "my_code_type", TagName: "code",
		Laterality: models.Laterality{CodedConcept: models.CodedConcept{Code: "LATERALITY_CODE_1", CodeSystem: "2.16.840.1.113883.6.1"}, Title: ""}, MapDataCriteria: ei.MapDataCriteria}
	assertXPath(t, xmlCodeRootNode(codeDisplay), "//qualifier/value", map[string]string{"code": "LATERALITY_CODE_1"}, nil)
}

func xmlCodeRootNode(codeDisplay models.CodeDisplay) *xml.ElementNode {
	codeDisplay.Description = "my lil description"
	xmlString := generateXML("_code.xml", codeDisplay)
	// printXmlString(xmlString)
	return xmlRootNode(xmlString)
}

func TestReasonTemplate(t *testing.T) {
	// do not negate reason
	reason := models.CodedConcept{Code: "REASON_CODE_1", CodeSystem: "2.16.840.1.113883.6.1"} // specified in cms9_26.json
	rootNode := xmlReasonRootNode(reason, false)
	assertXPath(t, rootNode, "//entryRelationship", map[string]string{"typeCode": "RSON"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation", map[string]string{"classCode": "OBS", "moodCode": "EVN"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation/templateId", map[string]string{"root": "2.16.840.1.113883.10.20.24.3.88", "extension": "2014-12-01"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation/id", map[string]string{"root": "1.3.6.1.4.1.115"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation/statusCode", map[string]string{"code": "completed"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation/effectiveTime/low", map[string]string{"value": "197001010000+0000"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation/value", map[string]string{"xsi:type": "CD", "code": "REASON_CODE_1", "codeSystem": "2.16.840.1.113883.6.1", "sdtc:valueSet": "1.2.3.4.5.6.7.8.9.11"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation/code", map[string]string{"code": "77301-0", "codeSystem": "2.16.840.1.113883.6.1", "displayName": "reason", "codeSystemName": "LOINC"}, nil)

	// negate reason
	rootNode = xmlReasonRootNode(reason, true)
	assertXPath(t, rootNode, "//entryRelationship", nil, nil)

	// reason that is not specifed by a measure
	reason = models.CodedConcept{Code: "not_a_specified_code", CodeSystem: "¯\\_(ツ)_/¯"}
	xmlString := generateXML("_reason.xml", getReasonData(reason, false))
	assert.Equal(t, "", strings.TrimSpace(xmlString))
}

func xmlReasonRootNode(reason models.CodedConcept, negateReason bool) *xml.ElementNode {
	data := getReasonData(reason, negateReason)
	setMapDataCriteria(data)
	xmlString := generateXML("_reason.xml", *data)
	// printXmlString(xmlString)
	return xmlRootNode(xmlString)
}

func getReasonData(reason models.CodedConcept, negateReason bool) *models.EntryInfo {
	encounter := models.Encounter{}
	var stime = new(int64)
	*stime = 0
	if negateReason {
		encounter.Entry = models.Entry{NegationReason: reason}
	} else {
		encounter.Entry = models.Entry{Reason: reason}
	}
	encounter.StartTime = stime
	return &models.EntryInfo{EntrySection: &encounter}
}

func setMapDataCriteria(ei *models.EntryInfo) {
	var fieldOids = map[string][]string{"REASON": []string{"1.2.3.4.5.6.7.8.9.11"},
		"ORDINAL":  []string{"1.2.3.4.5.6.7.8.9.10"},
		"SEVERITY": []string{"1.2.3.4.5.6.7.8.9.13"},
		"ROUTE":    []string{"1.2.3.4.5.6.7.8.9.12"}}
	var resultOids = []string{"1.2.3.4.5.6.7.8.9.14"}
	// var vsOid = "1.2.3.4.5.6.7.8.9"
	ei.MapDataCriteria = models.Mdc{FieldOids: fieldOids, ResultOids: resultOids}
}

func TestResultValueTemplate(t *testing.T) {
	entryInfos := getResultValueData()

	// Codes included
	rootNode := xmlResultValueRootNode(entryInfos[0])
	assertXPath(t, rootNode, "//code", map[string]string{"code": "first", "codeSystem": "2.16.840.1.113883.6.96"}, nil)

	// Value is a scalar
	rootNode = xmlResultValueRootNode(entryInfos[1])
	assertXPath(t, rootNode, "//value", map[string]string{"xsi:type": "PQ", "value": "5.2", "unit": "Inches"}, nil)

	// Value is a scalar with no units
	rootNode = xmlResultValueRootNode(entryInfos[2])
	assertXPath(t, rootNode, "//value", map[string]string{"xsi:type": "PQ", "value": "5.3", "unit": "1"}, nil)

	// Value is a boolean
	rootNode = xmlResultValueRootNode(entryInfos[3])
	assertXPath(t, rootNode, "//value", map[string]string{"xsi:type": "BL", "value": "true"}, nil)

	// No values
	rootNode = xmlResultValueRootNode(entryInfos[4])
	assertXPath(t, rootNode, "//value", map[string]string{"xsi:type": "CD", "nullFlavor": "UNK"}, nil)
}

func xmlResultValueRootNode(eInfo models.EntryInfo) *xml.ElementNode {
	xmlString := generateXML("_result_value.xml", eInfo.EntrySection.GetEntry().WrapResultValues(eInfo.EntrySection.GetEntry().Values))
	return xmlRootNode(xmlString)
}

func getResultValueData() []models.EntryInfo {
	// Sample ResultValue objects to be embedded in the entries.
	expectedCodeDisplay := models.CodeDisplay{CodeType: "resultValue", PreferredCodeSets: []string{"SNOMED-CT"}}
	coded := models.Coded{Codes: map[string][]string{"codeSetA": []string{"third", "fourth"}, "SNOMED-CT": []string{"first"}}}

	// Several entries created to test different paths in the template
	var entries []models.Entry
	entries = append(entries, models.Entry{Values: [](models.ResultValue){models.ResultValue{Scalar: "2", Units: "", Coded: coded}}, CodeDisplays: [](models.CodeDisplay){expectedCodeDisplay}})
	entries = append(entries, models.Entry{Values: [](models.ResultValue){models.ResultValue{Scalar: "5.2", Units: "Inches"}}})
	entries = append(entries, models.Entry{Values: [](models.ResultValue){models.ResultValue{Scalar: "5.3", Units: ""}}})
	entries = append(entries, models.Entry{Values: [](models.ResultValue){models.ResultValue{Scalar: "true", Units: ""}}})
	entries = append(entries, models.Entry{})
	var entrySections []models.HasEntry
	for _, entry := range entries {
		entrySections = append(entrySections, &models.Encounter{Entry: entry})
	}
	entrySections = append(entrySections, nil)
	entryInfos := models.AppendEntryInfos([]models.EntryInfo{}, entrySections, models.Mdc{})

	return entryInfos
}

func TestOrdinalityTemplate(t *testing.T) {
	ordinality := models.Ordinality{
		CodedConcept: models.CodedConcept{Code: "ORDINAL_CODE_1", CodeSystem: "2.16.840.1.113883.6.1"}, Title: "Principal"}
	rootNode := xmlOrdinalityRootNode(ordinality)

	assertXPath(t, rootNode, "//priorityCode", map[string]string{"code": "ORDINAL_CODE_1", "codeSystem": "2.16.840.1.113883.6.1", "sdtc:valueSet": "1.2.3.4.5.6.7.8.9.10"}, nil)

}

func xmlOrdinalityRootNode(ordinality models.Ordinality) *xml.ElementNode {
	data := getOrdinalityData(ordinality)
	setMapDataCriteria(data)
	xmlString := generateXML("_ordinality.xml", *data)
	return xmlRootNode(xmlString)
}

func getOrdinalityData(ordinality models.Ordinality) *models.EntryInfo {
	procedure := models.Procedure{Ordinality: ordinality}
	return &models.EntryInfo{EntrySection: &procedure}
}

func TestMedicationDetailsTemplate(t *testing.T) {
	route := models.CodedConcept{Code: "ROUTE_CODE_1", CodeSystem: "2.16.840.1.113883.6.1"}
	neg := true
	dose := models.Scalar{Value: "1", Unit: "d"}

	rootNode := xmlMedicationDetailsRootNode(&route, &neg, dose)

	assertXPath(t, rootNode, "//routeCode", map[string]string{"code": "ROUTE_CODE_1", "codeSystem": "2.16.840.1.113883.6.1", "sdtc:valueSet": "1.2.3.4.5.6.7.8.9.12"}, nil)

	rootNode = xmlMedicationDetailsRootNode(nil, &neg, dose)
	assertXPath(t, rootNode, "//doseQuantity", map[string]string{"nullFlavor": "NA"}, nil)

	rootNode = xmlMedicationDetailsRootNode(nil, nil, dose)
	assertXPath(t, rootNode, "//doseQuantity", map[string]string{"value": "1", "unit": "d"}, nil)
}

func xmlMedicationDetailsRootNode(route *models.CodedConcept, neg *bool, dose models.Scalar) *xml.ElementNode {
	med := models.Medication{
		Route: route, Dose: dose, Entry: models.Entry{NegationInd: neg}}
	data := &models.EntryInfo{EntrySection: &med}
	setMapDataCriteria(data)
	xmlString := generateXML("_medication_details.xml", *data)
	return xmlRootNode(xmlString)
}

func TestMedicationDispenseTemplate(t *testing.T) {
	dispenseDate := int64(1420581600)
	quantityDispensed := models.Scalar{Value: "10", Unit: "c"}
	fulfillmentHistory := [](models.FulfillmentHistory){models.FulfillmentHistory{DispenseDate: &dispenseDate, QuantityDispensed: quantityDispensed}}

	rootNode := xmlMedicationDispenseRootNode(fulfillmentHistory)

	assertXPath(t, rootNode, "//entryRelationship/supply/effectiveTime", map[string]string{"value": "201501062200+0000"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/supply/quantity", map[string]string{"value": "10", "unit": "c"}, nil)
}

func xmlMedicationDispenseRootNode(fulfillmentHistory [](models.FulfillmentHistory)) *xml.ElementNode {
	expectedCodeDisplay := models.CodeDisplay{CodeType: "medicationDispense", PreferredCodeSets: []string{"SNOMED-CT"}}
	med := models.Medication{
		FulfillmentHistory: fulfillmentHistory, Entry: models.Entry{CodeDisplays: [](models.CodeDisplay){expectedCodeDisplay}}}
	data := &models.EntryInfo{EntrySection: &med}
	setMapDataCriteria(data)
	xmlString := generateXML("_medication_dispense.xml", *data)
	return xmlRootNode(xmlString)
}

// test _2.16.840.1.113883.10.20.24.3.23.xml
func TestEncounterPerformedTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.23"

	// only tests r2 compatable
	rootNode := xmlRootNodeForQrdaOid(qrdaOid)
	assertXPath(t, rootNode, "//entry/encounter", map[string]string{"classCode": "ENC", "moodCode": "ENV"}, nil)
	assertXPath(t, rootNode, "//entry/encounter/templateId[@root='2.16.840.1.113883.10.20.22.4.49']", nil, nil)
	assertXPath(t, rootNode, "//entry/encounter/templateId[@root='2.16.840.1.113883.10.20.24.3.23']", nil, nil)
	assertXPath(t, rootNode, "//entry/encounter/id", map[string]string{"root": "1.3.6.1.4.1.115"}, nil)
	assertContent(t, rootNode, "//entry/encounter/text", "Encounter, Performed: Encounter Inpatient")
	assertXPath(t, rootNode, "//entry/encounter/statusCode", map[string]string{"code": "completed"}, nil)
	assertXPath(t, rootNode, "//entry/encounter/effectiveTime", nil, nil)
	// assertXPath(t, rootNode, "//entry/encounter/effectiveTime/low", map[string]string{"value": "201408110415"}, nil)

	// test admit time vs start time for <low> tag. test discharge time vs end time for <high> tag
	ei := getDataForQrdaOid(qrdaOid)
	entrySection := ei.EntrySection.(*models.Encounter)
	entrySection.AdmitTime = new(int64)
	*entrySection.AdmitTime = 1262462640 // is time 2010 01 02 1504 in EST
	entrySection.StartTime = nil
	entrySection.DischargeTime = new(int64)
	*entrySection.DischargeTime = 1293998640 // is time 2011 01 02 1504 in EST
	entrySection.EndTime = nil
	ei.EntrySection = entrySection
	rootNode = xmlRootNodeForQrdaOidWithData(qrdaOid, ei)
	assertXPath(t, rootNode, "//entry/encounter/effectiveTime/low", map[string]string{"value": "201001022004+0000"}, nil)
	assertXPath(t, rootNode, "//entry/encounter/effectiveTime/high", map[string]string{"value": "201101022004+0000"}, nil)
	entrySection.AdmitTime = nil
	entrySection.StartTime = new(int64)
	*entrySection.StartTime = 1293998640 // is time 2011 01 02 1504 in EST
	entrySection.DischargeTime = nil
	entrySection.EndTime = new(int64)
	*entrySection.EndTime = 1262462640 // is time 2010 01 02 1504 in EST
	ei.EntrySection = entrySection
	rootNode = xmlRootNodeForQrdaOidWithData(qrdaOid, ei)
	assertXPath(t, rootNode, "//entry/encounter/effectiveTime/low", map[string]string{"value": "201101022004+0000"}, nil)
	assertXPath(t, rootNode, "//entry/encounter/effectiveTime/high", map[string]string{"value": "201001022004+0000"}, nil)

	// continue testing here
}

func TestCommunicationFromPatientToProviderTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.2"
	dataCriteriaName := "communication_patient_to_provider"
	entryName := "communication_patient_to_provider"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Communication{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/act/templateId", map[string]string{"root": "2.16.840.1.113883.10.20.24.3.2"}, nil)

	assertXPath(t, xrn, "//entry/act/effectiveTime/low", map[string]string{"value": "201311030815+0000"}, nil)
	assertXPath(t, xrn, "//entry/act/effectiveTime/high", map[string]string{"value": "201311030815+0000"}, nil)

	assertXPath(t, xrn, "//entry/act/code", map[string]string{"code": "315640000", "codeSystem": "2.16.840.1.113883.6.96"}, nil)

	assertNoXPath(t, xrn, "//entry/act/entryRelationship/observation/templateId[@root='2.16.840.1.113883.10.20.24.3.88']")
}

func TestCommunicationFromProviderToProviderTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.4"
	dataCriteriaName := "communication_provider_to_provider"
	entryName := "communication_provider_to_provider"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Communication{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/act/templateId", map[string]string{"root": "2.16.840.1.113883.10.20.24.3.4"}, nil)

	assertXPath(t, xrn, "//entry/act/effectiveTime/low", map[string]string{"value": "201405020815+0000"}, nil)
	assertXPath(t, xrn, "//entry/act/effectiveTime/high", map[string]string{"value": "201405020823+0000"}, nil)

	assertXPath(t, xrn, "//entry/act/code", map[string]string{"code": "312904009", "codeSystem": "2.16.840.1.113883.6.96"}, nil)

	assertXPath(t, xrn, "//entry/act/entryRelationship/observation/templateId", map[string]string{"root": "2.16.840.1.113883.10.20.24.3.88"}, nil)
}

func TestCommunicationFromProviderToPatientTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.3"
	dataCriteriaName := "communication_provider_to_patient"
	entryName := "communication_provider_to_patient"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Communication{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/act/templateId", map[string]string{"root": qrdaOid}, nil)
	assertXPath(t, xrn, "//entry/act/effectiveTime/low", map[string]string{"value": "201404251800+0000"}, nil)
	assertXPath(t, xrn, "//entry/act/effectiveTime/high", map[string]string{"value": "201404251800+0000"}, nil)
	assertXPath(t, xrn, "//entry/act/code", map[string]string{"code": "410264007", "codeSystem": "2.16.840.1.113883.6.96"}, nil)
}

func TestDeviceAppliedTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.7"
	dataCriteriaName := "device_applied"
	entryName := "device_applied"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Procedure{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/procedure", map[string]string{"classCode": "PROC", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/procedure/templateId[@root='2.16.840.1.113883.10.20.24.3.7']", nil, nil)
	assertContent(t, xrn, "//entry/procedure/text", "Device, Applied: Graduated compression stockings (GCS)")
	assertXPath(t, xrn, "//entry/procedure/effectiveTime/low", map[string]string{"value": "201504070801+0000"}, nil)
	assertXPath(t, xrn, "//entry/procedure/effectiveTime/high", map[string]string{"value": "201504070801+0000"}, nil)

	assertNoXPath(t, xrn, "//entry/act")
}

func TestDeviceOrderTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.9"
	dataCriteriaName := "device_order"
	entryName := "device_order"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.MedicalEquipment{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/supply", map[string]string{"classCode": "SPLY", "moodCode": "RQO"}, nil)
	assertXPath(t, xrn, "//entry/supply/templateId[@root='2.16.840.1.113883.10.20.24.3.9']", nil, nil)
	assertContent(t, xrn, "//entry/supply/text", "Device, Order: Intermittent pneumatic compression devices (IPC)")
	assertXPath(t, xrn, "//entry/supply/effectiveTime/low", map[string]string{"value": "201504060830+0000"}, nil)
	assertXPath(t, xrn, "//entry/supply/effectiveTime/high", map[string]string{"value": "201504060830+0000"}, nil)
}

func TestLaboratoryTestOrderTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.37"
	dataCriteriaName := "laboratory_test_order"
	entryName := "laboratory_test_order"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.LabResult{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/observation", map[string]string{"classCode": "OBS", "moodCode": "RQO"}, nil)
	assertXPath(t, xrn, "//entry/observation/templateId[@root='2.16.840.1.113883.10.20.22.4.44']", nil, nil)
	assertXPath(t, xrn, "//entry/observation/templateId[@root='2.16.840.1.113883.10.20.24.3.37']", nil, nil)
	assertContent(t, xrn, "//entry/observation/text", "Laboratory Test, Order: Pregnancy Test")
	assertXPath(t, xrn, "//entry/observation/author/templateId[@root='2.16.840.1.113883.10.20.22.4.119']", nil, nil)
}

func TestLaboratoryTestPerformedTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.38"
	dataCriteriaName := "laboratory_test_performed"
	entryName := "laboratory_test_performed"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.LabResult{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/observation", map[string]string{"classCode": "OBS", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/observation/templateId[@root='2.16.840.1.113883.10.20.24.3.38']", nil, nil)
	assertContent(t, xrn, "//entry/observation/text", "Laboratory Test, Performed: LDL-c")
	assertXPath(t, xrn, "//entry/observation/effectiveTime/low", map[string]string{"value": "201504060700+0000"}, nil)
	assertXPath(t, xrn, "//entry/observation/effectiveTime/high", map[string]string{"value": "201504060700+0000"}, nil)
}

func TestDiagnosticStudyOrderTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.17"
	dataCriteriaName := "diagnostic_study_order"
	entryName := "diagnostic_study_order"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Procedure{})
	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/observation", map[string]string{"classCode": "OBS", "moodCode": "RQO"}, nil)
	assertXPath(t, xrn, "//entry/observation/templateId[@root='2.16.840.1.113883.10.20.24.3.17']", map[string]string{"extension": "2014-12-01"}, nil)
	assertContent(t, xrn, "//entry/observation/text", "Diagnostic Study, Order: VTE Diagnostic Test")
	assertXPath(t, xrn, "//entry/observation/effectiveTime/low", map[string]string{"value": "201505050700+0000"}, nil)
	assertXPath(t, xrn, "//entry/observation/effectiveTime/high", map[string]string{"value": "201505050700+0000"}, nil)
	assertXPath(t, xrn, "//entry/observation/author/time/low", map[string]string{"value": "201505050700+0000"}, nil)
	assertXPath(t, xrn, "//entry/observation/author/time/high", map[string]string{"value": "201505050700+0000"}, nil)
}

func TestDiagnosticStudyPerformedTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.18"
	dataCriteriaName := "diagnostic_study_performed"
	entryName := "diagnostic_study_performed"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Procedure{})
	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/observation", map[string]string{"classCode": "OBS", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/observation/templateId[@root='2.16.840.1.113883.10.20.24.3.18']", map[string]string{"extension": "2014-12-01"}, nil)
	assertContent(t, xrn, "//entry/observation/text", "Diagnostic Study, Performed: Ct Scan Including Chest Diagnostic Test")
	assertXPath(t, xrn, "//entry/observation/effectiveTime/low", map[string]string{"value": "201505200800+0000"}, nil)
	assertXPath(t, xrn, "//entry/observation/effectiveTime/high", map[string]string{"value": "201505200810+0000"}, nil)
}

func TestPatientCharacteristicClinicalTrialParticipantTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.51"
	dataCriteriaName := "patient_characteristic_clinical_trial_participant"
	entryName := "patient_characteristic_clinical_trial_participant"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Encounter{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/observation", map[string]string{"classCode": "OBS", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/observation/templateId[@root='2.16.840.1.113883.10.20.24.3.51']", nil, nil)
	assertXPath(t, xrn, "//entry/observation/code", map[string]string{"code": "ASSERTION", "codeSystem": "2.16.840.1.113883.5.4"}, nil)
	assertXPath(t, xrn, "//entry/observation/effectiveTime/low", map[string]string{"value": "201507130830+0000"}, nil)
}

func TestPatientCharacteristicExpiredTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.54"
	dataCriteriaName := "patient_characteristic_expired"
	entryName := "patient_characteristic_expired"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Encounter{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/observation", map[string]string{"classCode": "OBS", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/observation/templateId[@root='2.16.840.1.113883.10.20.22.4.79']", nil, nil)
	assertXPath(t, xrn, "//entry/observation/templateId[@root='2.16.840.1.113883.10.20.24.3.54']", nil, nil)
	assertXPath(t, xrn, "//entry/observation/effectiveTime/low", map[string]string{"value": "201505010800+0000"}, nil)
}

func TestPatientCharacteristicPayerTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.55"
	dataCriteriaName := "patient_characteristic_payer"
	entryName := "patient_characteristic_payer"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.InsuranceProvider{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/observation", map[string]string{"classCode": "OBS", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/observation/templateId[@root='2.16.840.1.113883.10.20.24.3.55']", nil, nil)
	assertXPath(t, xrn, "//entry/observation/id", map[string]string{"root": "1.3.6.1.4.1.115"}, nil)
	assertXPath(t, xrn, "//entry/observation/code", map[string]string{"code": "48768-6", "codeSystemName": "LOINC", "codeSystem": "2.16.840.1.113883.6.1"}, nil)
}

func TestPatientCharacteristicObservationAssertionTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.103"
	dataCriteriaName := "patient_characteristic_observation_assertion"
	entryName := "patient_characteristic_observation_assertion"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Encounter{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/observation", map[string]string{"classCode": "OBS", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/observation/templateId[@root='2.16.840.1.113883.10.20.24.3.103']", nil, nil)
	assertXPath(t, xrn, "//entry/observation/id", map[string]string{"root": "1.3.6.1.4.1.115"}, nil)
	assertXPath(t, xrn, "//entry/observation/code", map[string]string{"code": "ASSERTION", "codeSystem": "2.16.840.1.113883.5.4"}, nil)
	assertXPath(t, xrn, "//entry/observation/effectiveTime/low", map[string]string{"value": "201510130800+0000"}, nil)
	assertXPath(t, xrn, "//entry/observation/effectiveTime/high", map[string]string{"value": "201510130800+0000"}, nil)
}

func TestProcedureIntoleranceTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.62"
	dataCriteriaName := "procedure_intolerance"
	entryName := "procedure_intolerance"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Procedure{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/observation", map[string]string{"classCode": "OBS", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/observation/templateId[@root='2.16.840.1.113883.10.20.24.3.62']", nil, nil)
	assertXPath(t, xrn, "//entry/observation/id", map[string]string{"root": "1.3.6.1.4.1.115"}, nil)
	assertXPath(t, xrn, "//entry/observation/code", map[string]string{"code": "ASSERTION", "codeSystem": "2.16.840.1.113883.5.4", "codeSystemName": "ActCode", "displayName": "Assertion"}, nil)
	assertContent(t, xrn, "//entry/observation/entryRelationship/procedure/text", "Procedure, Intolerance: Influenza Vaccination")
	assertXPath(t, xrn, "//entry/observation/effectiveTime/low", map[string]string{"value": "201411101000+0000"}, nil)
	assertXPath(t, xrn, "//entry/observation/entryRelationship/procedure/effectiveTime/high", map[string]string{"value": "201411101000+0000"}, nil)
}

func TestProcedureOrderTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.63"
	dataCriteriaName := "procedure_order"
	entryName := "procedure_order"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Procedure{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/procedure", map[string]string{"classCode": "PROC", "moodCode": "RQO"}, nil)
	assertXPath(t, xrn, "//entry/procedure/templateId[@root='2.16.840.1.113883.10.20.22.4.41']", nil, nil)
	assertXPath(t, xrn, "//entry/procedure/id", map[string]string{"root": "1.3.6.1.4.1.115"}, nil)
	assertContent(t, xrn, "//entry/procedure/text", "Procedure, Order: BH Counseling for Depression")
	assertXPath(t, xrn, "//entry/procedure/author/time", map[string]string{"value": "201501220845+0000"}, nil)
}

func TestProcedurePerformedTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.64"
	dataCriteriaName := "procedure_performed"
	entryName := "procedure_performed"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Procedure{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/procedure", map[string]string{"classCode": "PROC", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/procedure/templateId[@root='2.16.840.1.113883.10.20.24.3.64']", nil, nil)
	assertXPath(t, xrn, "//entry/procedure/id", map[string]string{"root": "1.3.6.1.4.1.115"}, nil)
	assertContent(t, xrn, "//entry/procedure/text", "Procedure, Performed: General or Neuraxial Anesthesia")
	assertXPath(t, xrn, "//entry/procedure/effectiveTime/low", map[string]string{"value": "201504070830+0000"}, nil)
	assertXPath(t, xrn, "//entry/procedure/entryRelationship/procedure/effectiveTime", map[string]string{"value": "201504070840+0000"}, nil)
}

func TestInterventionOrderTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.31"
	dataCriteriaName := "intervention_order"
	entryName := "intervention_order"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Procedure{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/act", map[string]string{"classCode": "ACT", "moodCode": "RQO"}, nil)
	assertXPath(t, xrn, "//entry/act/templateId[@root='2.16.840.1.113883.10.20.22.4.39']", nil, nil)
	assertXPath(t, xrn, "//entry/act/id", map[string]string{"root": "1.3.6.1.4.1.115"}, nil)
	assertContent(t, xrn, "//entry/act/text", "Intervention, Order: Comfort Measures")
	assertXPath(t, xrn, "//entry/act/effectiveTime/low", map[string]string{"value": "201504201945+0000"}, nil)
}

func TestInterventionPerformedTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.32"
	dataCriteriaName := "intervention_performed"
	entryName := "intervention_performed"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Procedure{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/act", map[string]string{"classCode": "ACT", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/act/templateId[@root='2.16.840.1.113883.10.20.22.4.12']", nil, nil)
	assertXPath(t, xrn, "//entry/act/id", map[string]string{"root": "1.3.6.1.4.1.115"}, nil)
	assertContent(t, xrn, "//entry/act/text", "Intervention, Performed: Chronic Wound Care")
	assertXPath(t, xrn, "//entry/act/effectiveTime/low", map[string]string{"value": "201505121000+0000"}, nil)
}

func TestDiagnosisActiveTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.11"
	dataCriteriaName := "diagnosis_active"
	entryName := "diagnosis_active"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Condition{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/act", map[string]string{"classCode": "ACT", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/act/templateId[@root='2.16.840.1.113883.10.20.22.4.3']", nil, nil)
	assertXPath(t, xrn, "//entry/act/effectiveTime/low", map[string]string{"value": "201503120830+0000"}, nil)
	assertXPath(t, xrn, "//entry/act/entryRelationship/observation/entryRelationship[@typeCode='REFR']/observation/value", map[string]string{"displayName": "Moderate or Severe"}, nil)
}

func TestFamilyHistoryTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.12"
	dataCriteriaName := "diagnosis_family_history"
	entryName := "diagnosis_family_history"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Condition{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/organizer", map[string]string{"classCode": "CLUSTER", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/organizer/templateId[@root='2.16.840.1.113883.10.20.24.3.12']", nil, nil)
	assertXPath(t, xrn, "//entry/organizer/component/observation/effectiveTime/low", map[string]string{"value": "201503120830+0000"}, nil)
}

func TestDiagnosisInactiveTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.13"
	dataCriteriaName := "diagnosis_inactive"
	entryName := "diagnosis_inactive"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Condition{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/act", map[string]string{"classCode": "ACT", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/act/templateId[@root='2.16.840.1.113883.10.20.22.4.3']", nil, nil)
	assertXPath(t, xrn, "//entry/act/effectiveTime/high", map[string]string{"value": "201505100759+0000"}, nil)
}

func TestDiagnosisResolvedTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.14"
	dataCriteriaName := "diagnosis_resolved"
	entryName := "diagnosis_resolved"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Condition{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/act", map[string]string{"classCode": "ACT", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/act/templateId[@root='2.16.840.1.113883.10.20.22.4.3']", nil, nil)
	assertXPath(t, xrn, "//entry/act/effectiveTime/low", map[string]string{"value": "201405100801+0000"}, nil)
}

func TestMedicationActiveTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.41"
	dataCriteriaName := "medication_active"
	entryName := "medication_active"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Medication{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/substanceAdministration", map[string]string{"classCode": "SBADM", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/substanceAdministration/templateId[@root='2.16.840.1.113883.10.20.24.3.41']", nil, nil)
	assertXPath(t, xrn, "//entry/substanceAdministration/effectiveTime/low", map[string]string{"value": "201501062200+0000"}, nil)
	assertXPath(t, xrn, "//entry/substanceAdministration/effectiveTime[@operator='A']/period", map[string]string{"value": "1", "unit": "d"}, nil)
}

func TestMedicationAdministeredTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.42"
	dataCriteriaName := "medication_administered"
	entryName := "medication_administered"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Medication{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/act", map[string]string{"classCode": "ACT", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/act/templateId[@root='2.16.840.1.113883.10.20.24.3.42']", nil, nil)
	assertXPath(t, xrn, "//entry/act/entryRelationship[@typeCode='COMP']/substanceAdministration/effectiveTime/low", map[string]string{"value": "201505151000+0000"}, nil)
}

func TestMedicationAdverseEffectTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.43"
	dataCriteriaName := "medication_adverse_effect"
	entryName := "medication_adverse_effect"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Medication{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/observation", map[string]string{"classCode": "OBS", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/observation/templateId[@root='2.16.840.1.113883.10.20.22.4.7']", nil, nil)
	assertXPath(t, xrn, "//entry/observation/effectiveTime/low", map[string]string{"value": "201501062200+0000"}, nil)
}

func TestMedicationAllergyTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.44"
	dataCriteriaName := "medication_allergy"
	entryName := "medication_allergy"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Medication{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/observation", map[string]string{"classCode": "OBS", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/observation/templateId[@root='2.16.840.1.113883.10.20.24.3.44']", nil, nil)
	assertXPath(t, xrn, "//entry/observation/effectiveTime/low", map[string]string{"value": "201505190800+0000"}, nil)
}

func TestMedicationDispensedTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.45"
	dataCriteriaName := "medication_dispensed"
	entryName := "medication_dispensed"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Medication{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/supply", map[string]string{"classCode": "SPLY", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/supply/templateId[@root='2.16.840.1.113883.10.20.22.4.18']", nil, nil)
	assertXPath(t, xrn, "//entry/supply/effectiveTime/high", map[string]string{"value": "201410190815+0000"}, nil)
}

func TestMedicationIntoleranceTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.46"
	dataCriteriaName := "medication_intolerance"
	entryName := "medication_intolerance"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Medication{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/observation", map[string]string{"classCode": "OBS", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/observation/templateId[@root='2.16.840.1.113883.10.20.22.4.7']", nil, nil)
	assertXPath(t, xrn, "//entry/observation/effectiveTime/low", map[string]string{"value": "201512070830+0000"}, nil)
}

func TestMedicationOrderTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.47"
	dataCriteriaName := "medication_order"
	entryName := "medication_order"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Medication{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/substanceAdministration", map[string]string{"classCode": "SBADM", "moodCode": "RQO"}, nil)
	assertXPath(t, xrn, "//entry/substanceAdministration/templateId[@root='2.16.840.1.113883.10.20.22.4.42']", nil, nil)
	assertXPath(t, xrn, "//entry/substanceAdministration/effectiveTime/low", map[string]string{"value": "201505140800+0000"}, nil)
	assertXPath(t, xrn, "//entry/substanceAdministration/effectiveTime[@operator='A']/period", map[string]string{"value": "12", "unit": "h"}, nil)
}

func TestMedicationDischargeTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.105"
	dataCriteriaName := "medication_discharge"
	entryName := "medication_discharge"

	ei := generateDataForTemplate(dataCriteriaName, entryName, &models.Medication{})

	xrn := xmlRootNodeForQrdaOidWithData(qrdaOid, ei)

	assertXPath(t, xrn, "//entry/act", map[string]string{"classCode": "ACT", "moodCode": "EVN"}, nil)
	assertXPath(t, xrn, "//entry/act/templateId[@root='2.16.840.1.113883.10.20.24.3.105']", nil, nil)
	assertXPath(t, xrn, "//entry/act/effectiveTime/low", map[string]string{"value": "201504150800+0000"}, nil)
	assertXPath(t, xrn, "//entry/act/entryRelationship[@typeCode='SUBJ']/substanceAdministration/effectiveTime[@operator='A']/period", map[string]string{"value": "1", "unit": "d"}, nil)
	assertXPath(t, xrn, "//entry/act/entryRelationship[@typeCode='SUBJ']/substanceAdministration/routeCode", map[string]string{"code": "ROUTE_CODE_1", "codeSystem": "2.16.840.1.113883.6.1"}, nil)
	assertXPath(t, xrn, "//entry/act/entryRelationship[@typeCode='SUBJ']/substanceAdministration/doseQuantity", map[string]string{"value": "1"}, nil)
}

// - - - - - - - - //
//   H E L P E R   //
// - - - - - - - - //

// Given the name of an "entry" file, a "dataCriteria" file, and a pointer to an entry object, return the required entryInfo struct for the template
func generateDataForTemplate(dataCriteriaName string, entryName string, entry models.HasEntry) models.EntryInfo {
	hds := models.NewHds()
	dc, err := ioutil.ReadFile(fmt.Sprintf("../fixtures/data_criteria/%s.json", dataCriteriaName))
	if err != nil {
		log.Fatalln(err)
	}

	ent, err := ioutil.ReadFile(fmt.Sprintf("../fixtures/entries/%s.json", entryName))
	if err != nil {
		log.Fatalln(err)
	}

	var dataCriteria models.DataCriteria
	json.Unmarshal(dc, &dataCriteria)

	json.Unmarshal(ent, &entry)

	udc := models.UniqueDataCriteria([]models.DataCriteria{dataCriteria})
	hds.SetCodeDisplaysForEntry(entry.GetEntry(), udc[0], "r3")

	ei := models.EntryInfo{
		EntrySection:    entry,
		MapDataCriteria: udc[0],
	}

	return ei
}

// asserts the xml path exists in xml string
// asserts that each expected attribute is on the tag
// asserts that each unexpected attribute is not on the tag
func assertXPath(t *testing.T, elem *xml.ElementNode, pathString string, expectedAttributes map[string]string, unexpectedAttributes []string) {
	path := xpath.Compile(pathString)
	nodes, err := elem.Search(path)
	if err != nil {
		log.Fatalln(err)
	}
	assert.NotEqual(t, len(nodes), 0)
	for _, node := range nodes {
		for attr, val := range expectedAttributes {
			if attrVal := node.Attribute(attr); attrVal != nil {
				assert.Equal(t, val, attrVal.String())
			} else {
				assert.Fail(t, fmt.Sprintf("expected xml attribute %s was not found at xml path \"%s\"", attr, pathString))
			}
		}
		for _, attr := range unexpectedAttributes {
			assert.Nil(t, node.Attribute(attr))
		}
	}
}

// assert the xml path does not exist in the xml string
func assertNoXPath(t *testing.T, elem *xml.ElementNode, pathString string) {
	path := xpath.Compile(pathString)
	res, err := elem.Search(path)
	if err != nil {
		log.Fatalln(err)
	}
	assert.Nil(t, res)
}

// assert all xml tags at the xml contain the content
func assertContent(t *testing.T, elem *xml.ElementNode, pathString string, content string) {
	path := xpath.Compile(pathString)
	nodes, err := elem.Search(path)
	if err != nil {
		log.Fatalln(err)
	}
	for _, node := range nodes {
		assert.Equal(t, content, node.Content())
	}
}

// just for debugging. remove later
func printXmlString(xmlString string) {
	fmt.Printf("\n====================\n")
	fmt.Printf(xmlString)
	fmt.Printf("\n====================\n\n")
}

// - - - - - - - - - - - - - - - - - - - //
//   G E N E R A T E   T E M P L A T E   //
// - - - - - - - - - - - - - - - - - - - //

func xmlRootNodeForQrdaOid(qrdaOid string) *xml.ElementNode {
	fileName := "_" + qrdaOid + ".xml"
	printXmlString(generateXML(fileName, getDataForQrdaOid(qrdaOid)))
	return xmlRootNode(generateXML(fileName, getDataForQrdaOid(qrdaOid)))
}

// same as xmlRootNodeForQrdaOid() function but allows custom input data (should be an EntryInfo struct)
func xmlRootNodeForQrdaOidWithData(qrdaOid string, data interface{}) *xml.ElementNode {
	fileName := "_" + qrdaOid + ".xml"
	printXmlString(generateXML(fileName, data))
	return xmlRootNode(generateXML(fileName, data))
}

func getDataForQrdaOid(qrdaOid string) models.EntryInfo {
	var p models.Record
	var m []models.Measure
	var vs []models.ValueSet
	setPatientMeasuresAndValueSets(&p, &m, &vs)
	ei, err := getEntryInfo(p, m, qrdaOid, vs) // ei stands for entry info
	if err != nil {
		log.Fatalln(err)
	}
	return ei
}

func xmlRootNode(xmlString string) *xml.ElementNode {
	doc, err := xml.Parse([]byte(xmlString), nil, nil, xml.DefaultParseOption, xml.DefaultEncodingBytes)
	if err != nil {
		log.Fatalln(err)
	}
	return doc.Root()
}

func generateXML(fileName string, templateData interface{}) string {
	var p models.Record
	var m []models.Measure
	var vs []models.ValueSet
	setPatientMeasuresAndValueSets(&p, &m, &vs)
	vsMap := models.NewValueSetMap(vs)
	return generateTemplateForFile(makeTemplate("r3", vsMap), fileName, templateData)
}

func setPatientMeasuresAndValueSets(patient *models.Record, measures *[]models.Measure, valueSets *[]models.ValueSet) {
	measureData := append([]byte("["), append(append(fixtures.Cms9v4a, append([]byte(","), fixtures.Cms26v3...)...), []byte("]")...)...)
	json.Unmarshal(fixtures.TestPatientDataAmi, patient)
	json.Unmarshal(measureData, measures)
	json.Unmarshal(fixtures.Cms9_26, valueSets)
}

func makeTemplate(qrdaVersion string, vsMap models.ValueSetMap) *template.Template {
	if qrdaVersion == "" {
		qrdaVersion = "r3"
	}
	temp := template.New("cat1")
	temp.Funcs(exporterFuncMap(temp, vsMap))
	fileNames, err := AssetDir("templates/cat1/" + qrdaVersion)
	if err != nil {
		log.Fatalln(err)
	}
	for _, fileName := range fileNames {
		asset, err := Asset("templates/cat1/" + qrdaVersion + "/" + fileName)
		if err != nil {
			log.Fatalln(err)
		}
		template.Must(temp.New(fileName).Parse(string(asset)))
	}
	return temp
}

func generateTemplateForFile(temp *template.Template, fileName string, templateData interface{}) string {
	var buf bytes.Buffer
	if err := temp.ExecuteTemplate(&buf, fileName, templateData); err != nil {
		if err != nil {
			log.Fatalln(err)
		}
	}
	return buf.String()
}

func getEntryInfo(patient models.Record, measures []models.Measure, qrdaOid string, vs []models.ValueSet) (models.EntryInfo, error) {
	hds := models.NewHds()
	entryInfos := patient.EntryInfosForPatient(measures, models.NewValueSetMap(vs), "r3")
	for _, ei := range entryInfos {
		if qrdaOid == hds.HqmfToQrdaOid(ei.EntrySection.GetEntry().Oid, ei.MapDataCriteria.DcKey.ValueSetOid) {
			return ei, nil
		}
	}
	if len(entryInfos) == 0 {
		return models.EntryInfo{}, errors.New("no entry infos found for patient and measures")
	}
	return models.EntryInfo{}, errors.New(fmt.Sprintf("no entry info found with qrda oid \"%s\"", qrdaOid))
}
