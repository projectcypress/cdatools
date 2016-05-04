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
	xp.RegisterNamespace("sdtc", "urn:hl7-org:sdtc")

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
	rawDiagnosesActive := ExtractSection(patientElement, diagnosisActiveXPath, ConditionExtractor, "2.16.840.1.113883.3.560.1.2")
	patient.Conditions = make([]models.Condition, len(rawDiagnosesActive))
	for i := range rawDiagnosesActive {
		patient.Conditions[i] = rawDiagnosesActive[i].(models.Condition)
	}

	//diagnosis inactive
	var diagnosisInactiveXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.13']")
	rawDiagnosesInactive := ExtractSection(patientElement, diagnosisInactiveXPath, DiagnosisInactiveExtractor, "2.16.840.1.113883.3.560.1.2")
	for i := range rawDiagnosesInactive {
		patient.Conditions = append(patient.Conditions, rawDiagnosesInactive[i].(models.Condition))
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

	// transfer to
	var transferToXPath = xpath.Compile("//cda:encounter[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.82']")
	rawTransferTos := ExtractSection(patientElement, transferToXPath, TransferToExtractor, "2.16.840.1.113883.3.560.1.72")
	for i := range rawTransferTos {
		patient.Encounters = append(patient.Encounters, rawTransferTos[i].(models.Encounter))
	}

	//medication active
	var medicationActiveXPath = xpath.Compile("./cda:entry/cda:substanceAdministration[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.41']")
	rawMedicationActives := ExtractSection(patientElement, medicationActiveXPath, MedicationActiveExtractor, "2.16.840.1.113883.3.560.1.13")
	patient.Medications = make([]models.Medication, len(rawMedicationActives))
	for i := range rawMedicationActives {
		patient.Medications[i] = rawMedicationActives[i].(models.Medication)
	}

	//medication dispensed
	var medicationDispensedXPath = xpath.Compile("./cda:entry/cda:supply[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.45']")
	rawMedicationDispenseds := ExtractSection(patientElement, medicationDispensedXPath, MedicationDispensedExtractor, "2.16.840.1.113883.3.560.1.8")
	for i := range rawMedicationDispenseds {
		patient.Medications = append(patient.Medications, rawMedicationDispenseds[i].(models.Medication))
	}

	//medication administered
	var medicationAdministeredXPath = xpath.Compile("./cda:entry/cda:act[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.42']/cda:entryRelationship/cda:substanceAdministration[cda:templateId/@root='2.16.840.1.113883.10.20.22.4.16']")
	rawMedicationAdministereds := ExtractSection(patientElement, medicationAdministeredXPath, MedicationExtractor, "2.16.840.1.113883.3.560.1.14")
	for i := range rawMedicationAdministereds {
		patient.Medications = append(patient.Medications, rawMedicationAdministereds[i].(models.Medication))
	}

	//medication ordered
	var medicationOrderedXPath = xpath.Compile("./cda:entry/cda:substanceAdministration[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.47']")
	rawMedicationOrdereds := ExtractSection(patientElement, medicationOrderedXPath, MedicationExtractor, "2.16.840.1.113883.3.560.1.17")
	for i := range rawMedicationOrdereds {
		patient.Medications = append(patient.Medications, rawMedicationOrdereds[i].(models.Medication))
	}

	//discharge medication active
	var medicationDischargeActiveXPath = xpath.Compile("./cda:entry/cda:act[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.105']/cda:entryRelationship/cda:substanceAdministration[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.41']")
	rawMedicationDischargeActives := ExtractSection(patientElement, medicationDischargeActiveXPath, MedicationExtractor, "2.16.840.1.113883.3.560.1.199")
	for i := range rawMedicationDischargeActives {
		patient.Medications = append(patient.Medications, rawMedicationDischargeActives[i].(models.Medication))
	}

	// medication intolerance
	var medicationIntoleranceXPath = xpath.Compile("./cda:entry/cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.46']")
	rawMedicationIntolerances := ExtractSection(patientElement, medicationIntoleranceXPath, AllergyExtractor, "2.16.840.1.113883.3.560.1.67")
	for i := range rawMedicationIntolerances {
		patient.Allergies = append(patient.Allergies, rawMedicationIntolerances[i].(models.Allergy))
	}

	// medication adverse event
	var medicationAdverseEventXPath = xpath.Compile("./cda:entry/cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.43']")
	rawMedicationAdverseEvents := ExtractSection(patientElement, medicationAdverseEventXPath, AllergyExtractor, "2.16.840.1.113883.3.560.1.7")
	for i := range rawMedicationAdverseEvents {
		patient.Allergies = append(patient.Allergies, rawMedicationAdverseEvents[i].(models.Allergy))
	}

	// medication allergy
	var medicationAllergyXPath = xpath.Compile("./cda:entry/cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.44']")
	rawMedicationAllergies := ExtractSection(patientElement, medicationAllergyXPath, AllergyExtractor, "2.16.840.1.113883.3.560.1.1")
	for i := range rawMedicationAllergies {
		patient.Allergies = append(patient.Allergies, rawMedicationAllergies[i].(models.Allergy))
	}

	// procedure intolerance (such as flu shot intolerance)
	var procedureIntoleranceXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.62']/cda:entryRelationship/cda:procedure[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.64']")
	rawProcedureIntolerances := ExtractSection(patientElement, procedureIntoleranceXPath, ProcedureIntoleranceExtractor, "2.16.840.1.113883.3.560.1.61")
	for i := range rawProcedureIntolerances {
		patient.Allergies = append(patient.Allergies, rawProcedureIntolerances[i].(models.Allergy))
	}

	// Gestational Age (technically a condition)
	var gestationalAgeXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.101']")
	rawGestationalAges := ExtractSection(patientElement, gestationalAgeXPath, GestationalAgeExtractor, "2.16.840.1.113883.3.560.1.1001")
	for i := range rawGestationalAges {
		patient.Conditions = append(patient.Conditions, rawGestationalAges[i].(models.Condition))
	}

	// Communication: patient to provider
	var communicationPatientToProviderXPath = xpath.Compile("//cda:entry/cda:act[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.2']")
	rawCommunicationsPatientToProvider := ExtractSection(patientElement, communicationPatientToProviderXPath, CommunicationExtractor, "2.16.840.1.113883.3.560.1.30")
	for i := range rawCommunicationsPatientToProvider {
		patient.Communications = append(patient.Communications, rawCommunicationsPatientToProvider[i].(models.Communication))
	}

	// Communication: provider to provider
	var communicationProviderToProviderXPath = xpath.Compile("//cda:entry/cda:act[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.4']")
	rawCommunicationsProviderToProvider := ExtractSection(patientElement, communicationProviderToProviderXPath, CommunicationExtractor, "2.16.840.1.113883.3.560.1.129")
	for i := range rawCommunicationsProviderToProvider {
		patient.Communications = append(patient.Communications, rawCommunicationsProviderToProvider[i].(models.Communication))
	}

	// Communication: provider to patient: not done
	var communicationProviderToPatientXPath = xpath.Compile("//cda:entry/cda:act[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.3']")
	rawCommunicationsProviderToPatient := ExtractSection(patientElement, communicationProviderToPatientXPath, CommunicationExtractor, "2.16.840.1.113883.3.560.1.31")
	for i := range rawCommunicationsProviderToPatient {
		patient.Communications = append(patient.Communications, rawCommunicationsProviderToPatient[i].(models.Communication))
	}

	// ECOG Status
	var ecogStatusXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.103']")
	rawEcogStatuses := ExtractSection(patientElement, ecogStatusXPath, ConditionExtractor, "2.16.840.1.113883.3.560.1.1001")
	for i := range rawEcogStatuses {
		patient.Conditions = append(patient.Conditions, rawEcogStatuses[i].(models.Condition))
	}

	// Symptom, active
	var symptomActiveXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.76']")
	rawActiveSymptoms := ExtractSection(patientElement, symptomActiveXPath, ConditionExtractor, "2.16.840.1.113883.3.560.1.69")
	for i := range rawActiveSymptoms {
		patient.Conditions = append(patient.Conditions, rawActiveSymptoms[i].(models.Condition))
	}

	// Diagnosis, Resolved
	var diagonsisResolvedXPath = xpath.Compile("//cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.14']")
	rawDiagnosesResolved := ExtractSection(patientElement, diagonsisResolvedXPath, ConditionExtractor, "2.16.840.1.113883.3.560.1.24")
	for i := range rawDiagnosesResolved {
		patient.Conditions = append(patient.Conditions, rawDiagnosesResolved[i].(models.Condition))
	}

	//Medical Equipment Applied
	var medEquipAppliedXPath = xpath.Compile("./cda:entry/cda:procedure[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.7']")
	rawMedEquipApplied := ExtractSection(patientElement, medEquipAppliedXPath, MedicalEquipmentExtractor, "2.16.840.1.113883.3.560.1.110")
	patient.MedicalEquipment = make([]models.MedicalEquipment, len(rawMedEquipApplied))
	for i := range rawMedEquipApplied {
		patient.MedicalEquipment[i] = rawMedEquipApplied[i].(models.MedicalEquipment)
	}

	//Medical Equipment
	var medEquipXPath = xpath.Compile("./cda:entry/cda:act[cda:code/@code = 'SPLY']")
	rawMedEquip := ExtractSection(patientElement, medEquipXPath, MedicalEquipmentExtractor, "2.16.840.1.113883.3.560.1.137")
	for i := range rawMedEquip {
		patient.MedicalEquipment = append(patient.MedicalEquipment, rawMedEquipApplied[i].(models.MedicalEquipment))
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

func ExtractScalar(scalar *models.Scalar, entryElement xml.Node, scalarPath *xpath.Expression) {
	scalarElements, err := entryElement.Search(scalarPath)
	util.CheckErr(err)

	for _, scalarElement := range scalarElements {
		unitAttr := scalarElement.Attribute("unit")
		valueAttr := scalarElement.Attribute("value")

		if valueAttr != nil {
			if unitAttr != nil {
				scalar.Unit = unitAttr.String()
			}
			scalar.Value, err = strconv.ParseInt(valueAttr.String(), 10, 64)
			util.CheckErr(err)
		}
	}
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
				val.Codes = map[string][]string{}
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

func ExtractReasonOrNegation(entry *models.Entry, entryElement xml.Node) {
	reasonXPath := xpath.Compile("./cda:entryRelationship[@typeCode='RSON']/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.88']/cda:value | ./cda:entryRelationship[@typeCode='RSON']/cda:act[cda:templateId/@root='2.16.840.1.113883.10.20.1.27']/cda:code")
	reasonElements, err := entryElement.Search(reasonXPath)
	util.CheckErr(err)

	for _, reasonElement := range reasonElements {
		codeSystemOidAttr := reasonElement.Attribute("codeSystem")
		codeAttr := reasonElement.Attribute("code")
		if codeSystemOidAttr != nil && codeAttr != nil {
			codeSystem := models.CodeSystemFor(codeSystemOidAttr.String())
			code := codeAttr.String()
			negationAttr := entryElement.Attribute("negationInd")
			if negationAttr != nil {
				negationInd := negationAttr.String()
				if negationInd == "true" {
					entry.NegationInd = true
					entry.NegationReason.Code = code
					entry.NegationReason.CodeSystem = codeSystem
					return
				}
			}
			entry.Reason.Code = code
			entry.Reason.CodeSystem = codeSystem
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
