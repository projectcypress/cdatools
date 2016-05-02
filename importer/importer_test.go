package importer

import (
	"io/ioutil"
	"testing"

	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
	. "gopkg.in/check.v1"
)

type ImporterSuite struct {
	patientElement xml.Node
	patient        *models.Record
}

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&ImporterSuite{})

func (i *ImporterSuite) SetUpSuite(c *C) {
	data, err := ioutil.ReadFile("../fixtures/cat1_good.xml")
	util.CheckErr(err)

	doc, err := xml.Parse(data, nil, nil, 0, xml.DefaultEncodingBytes)
	util.CheckErr(err)

	xp := doc.DocXPathCtx()
	xp.RegisterNamespace("cda", "urn:hl7-org:v3")
	xp.RegisterNamespace("sdtc", "urn:hl7-org:sdtc")

	var patientXPath = xpath.Compile("/cda:ClinicalDocument/cda:recordTarget/cda:patientRole/cda:patient")
	patientElements, err := doc.Root().Search(patientXPath)
	util.CheckErr(err)
	i.patientElement = patientElements[0]
	i.patient = &models.Record{}
}

func (i *ImporterSuite) TestExtractDemograpics(c *C) {
	ExtractDemographics(i.patient, i.patientElement)
	c.Assert(i.patient.First, Equals, "Norman")
	c.Assert(i.patient.Last, Equals, "Flores")
	c.Assert(i.patient.Birthdate, Equals, int64(599646600))
	c.Assert(i.patient.Race.Code, Equals, "1002-5")
	c.Assert(i.patient.Race.CodeSet, Equals, "CDC Race and Ethnicity")
	c.Assert(i.patient.Ethnicity.Code, Equals, "2186-5")
	c.Assert(i.patient.Ethnicity.CodeSet, Equals, "CDC Race and Ethnicity")
}

func (i *ImporterSuite) TestExtractEncountersPerformed(c *C) {
	var encounterXPath = xpath.Compile("//cda:encounter[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.23']")
	rawEncounters := ExtractSection(i.patientElement, encounterXPath, EncounterPerformedExtractor, "2.16.840.1.113883.3.560.1.79")
	i.patient.Encounters = make([]models.Encounter, len(rawEncounters))
	for j := range rawEncounters {
		i.patient.Encounters[j] = rawEncounters[j].(models.Encounter)
	}

	c.Assert(len(i.patient.Encounters), Equals, 3)

	encounter := i.patient.Encounters[0]
	c.Assert(encounter.ID.Root, Equals, "1.3.6.1.4.1.115")
	c.Assert(encounter.ID.Extension, Equals, "50d3a288da5fe6e14000016c")
	c.Assert(encounter.Codes["CPT"][0], Equals, "99201")
	c.Assert(encounter.StartTime, Equals, int64(1288612800))
	c.Assert(encounter.EndTime, Equals, int64(1288616400))
}

func (i *ImporterSuite) TestExtractEncounterOrdered(c *C) {
	var encounterOrderXPath = xpath.Compile("//cda:encounter[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.22']")
	rawEncounterOrders := ExtractSection(i.patientElement, encounterOrderXPath, EncounterOrderExtractor, "")
	i.patient.Encounters = make([]models.Encounter, len(rawEncounterOrders))
	for j := range rawEncounterOrders {
		i.patient.Encounters[j] = rawEncounterOrders[j].(models.Encounter)
	}

	c.Assert(len(i.patient.Encounters), Equals, 1)

	encounter := i.patient.Encounters[0]
	c.Assert(encounter.ID.Root, Equals, "50f84c1b7042f9877500025e")
	c.Assert(encounter.Codes["SNOMED-CT"][0], Equals, "76168009")
	c.Assert(encounter.Codes["CPT"][0], Equals, "90815")
	c.Assert(encounter.Codes["ICD-9-CM"][0], Equals, "94.49")
	c.Assert(encounter.Codes["ICD-10-PCS"][0], Equals, "GZHZZZZ")
	c.Assert(encounter.StartTime, Equals, int64(1135608034))
	c.Assert(encounter.EndTime, Equals, int64(1135608034))
}

func (i *ImporterSuite) TestExtractDiagnosesActive(c *C) {
	var diagnosisXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.11']")
	rawDiagnoses := ExtractSection(i.patientElement, diagnosisXPath, ConditionExtractor, "2.16.840.1.113883.3.560.1.2")
	i.patient.Conditions = make([]models.Condition, len(rawDiagnoses))
	for j := range rawDiagnoses {
		i.patient.Conditions[j] = rawDiagnoses[j].(models.Condition)
	}

	c.Assert(len(i.patient.Conditions), Equals, 3)
	firstDiagnosis := i.patient.Conditions[0]
	c.Assert(firstDiagnosis.ID.Root, Equals, "1.3.6.1.4.1.115")
	c.Assert(firstDiagnosis.ID.Extension, Equals, "54c1142869702d2cd2520100")
	c.Assert(firstDiagnosis.Codes["SNOMED-CT"][0], Equals, "195080001")
	c.Assert(firstDiagnosis.Description, Equals, "Diagnosis, Active: Atrial Fibrillation/Flutter")
	c.Assert(firstDiagnosis.StartTime, Equals, int64(1332775800))
	c.Assert(firstDiagnosis.EndTime, Equals, int64(0))
	c.Assert(firstDiagnosis.Severity["SNOMED-CT"][0], Equals, "55561003")

	secondDiagnosis := i.patient.Conditions[1]
	c.Assert(secondDiagnosis.ID.Root, Equals, "1.3.6.1.4.1.115")
	c.Assert(secondDiagnosis.ID.Extension, Equals, "54c1142969702d2cd2cd0200")
	c.Assert(secondDiagnosis.Codes["SNOMED-CT"][0], Equals, "237244005")
	c.Assert(secondDiagnosis.Description, Equals, "Diagnosis, Active: Pregnancy Dx")
	c.Assert(secondDiagnosis.StartTime, Equals, int64(1362150000))
	c.Assert(secondDiagnosis.EndTime, Equals, int64(1382284800))

	thirdDiagnosis := i.patient.Conditions[2]
	c.Assert(thirdDiagnosis.ID.Root, Equals, "1.3.6.1.4.1.115")
	c.Assert(thirdDiagnosis.ID.Extension, Equals, "54c1142869702d2cd2760100")
	c.Assert(thirdDiagnosis.Codes["SNOMED-CT"][0], Equals, "46635009")
	c.Assert(thirdDiagnosis.Description, Equals, "Diagnosis, Active: Diabetes")
	c.Assert(thirdDiagnosis.StartTime, Equals, int64(1361890800))
	c.Assert(thirdDiagnosis.EndTime, Equals, int64(0))
}

func (i *ImporterSuite) TestExtractDiagnosesInactive(c *C) {
	var diagnosisInactiveXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.13']")
	rawDiagnosesInactive := ExtractSection(i.patientElement, diagnosisInactiveXPath, DiagnosisInactiveExtractor, "2.16.840.1.113883.3.560.1.2")
	i.patient.Conditions = make([]models.Condition, len(rawDiagnosesInactive))
	for j := range rawDiagnosesInactive {
		i.patient.Conditions[j] = rawDiagnosesInactive[j].(models.Condition)
	}

	diagnosis := i.patient.Conditions[0]
	c.Assert(len(i.patient.Conditions), Equals, 1)
	c.Assert(diagnosis.ID.Root, Equals, "50f84c1d7042f98775000352")
	c.Assert(diagnosis.Codes["SNOMED-CT"][0], Equals, "76795007")
	c.Assert(diagnosis.Codes["ICD-9-CM"][0], Equals, "V02.61")
	c.Assert(diagnosis.Codes["ICD-10-CM"][0], Equals, "Z22.51")
	c.Assert(diagnosis.StartTime, Equals, int64(1092658739))
	c.Assert(diagnosis.EndTime, Equals, int64(1092686969))
}

func (i *ImporterSuite) TestExtractLabResults(c *C) {
	var labResultXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.40']")
	rawLabResults := ExtractSection(i.patientElement, labResultXPath, LabResultExtractor, "")
	i.patient.LabResults = make([]models.LabResult, len(rawLabResults))
	for j := range rawLabResults {
		i.patient.LabResults[j] = rawLabResults[j].(models.LabResult)
	}

	labResult := i.patient.LabResults[0]
	c.Assert(len(i.patient.LabResults), Equals, 1)
	c.Assert(labResult.ID.Root, Equals, "1.3.6.1.4.1.115")
	c.Assert(labResult.ID.Extension, Equals, "50d3a288da5fe6e1400002a9")
	c.Assert(labResult.Codes["LOINC"][0], Equals, "11268-0")
	c.Assert(labResult.StartTime, Equals, int64(674670276))
	c.Assert(len(labResult.Entry.Values), Equals, 1)
	c.Assert(labResult.Entry.Values[0].Value, Equals, "positive")
}

func (i *ImporterSuite) TestExtractLabOrders(c *C) {
	var labOrderXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.37']")
	rawLabOrders := ExtractSection(i.patientElement, labOrderXPath, LabOrderExtractor, "")
	i.patient.LabResults = make([]models.LabResult, len(rawLabOrders))
	for j := range rawLabOrders {
		i.patient.LabResults[j] = rawLabOrders[j].(models.LabResult)
	}

	labOrder := i.patient.LabResults[0]
	c.Assert(len(i.patient.LabResults), Equals, 1)
	c.Assert(labOrder.ID.Root, Equals, "50f84c1d7042f9877500039e")
	c.Assert(labOrder.Codes["SNOMED-CT"][0], Equals, "8879006")
	c.Assert(labOrder.StartTime, Equals, int64(674670276))
}

func (i *ImporterSuite) TestExtractInsuranceProviders(c *C) {
	var insuranceProviderXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.55']")
	rawInsuranceProviders := ExtractSection(i.patientElement, insuranceProviderXPath, InsuranceProviderExtractor, "2.16.840.1.113883.3.560.1.405")
	i.patient.InsuranceProviders = make([]models.InsuranceProvider, len(rawInsuranceProviders))
	for j := range rawInsuranceProviders {
		i.patient.InsuranceProviders[j] = rawInsuranceProviders[j].(models.InsuranceProvider)
	}

	insuranceProvider := i.patient.InsuranceProviders[0]
	c.Assert(len(i.patient.InsuranceProviders), Equals, 1)
	c.Assert(insuranceProvider.ID.Root, Equals, "1.3.6.1.4.1.115")
	c.Assert(insuranceProvider.Codes["SOP"][0], Equals, "349")
	c.Assert(insuranceProvider.StartTime, Equals, int64(1111851000)) // March 26, 2005 @ 15:30:00 GMT
}

func (i *ImporterSuite) TestExtractDiagnosticStudyOrders(c *C) {
	var diagnosticStudyOrderXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.17']")
	rawDiagnosticStudyOrders := ExtractSection(i.patientElement, diagnosticStudyOrderXPath, DiagnosticStudyOrderExtractor, "2.16.840.1.113883.3.560.1.40")
	i.patient.Procedures = make([]models.Procedure, len(rawDiagnosticStudyOrders))
	for j := range rawDiagnosticStudyOrders {
		i.patient.Procedures[j] = rawDiagnosticStudyOrders[j].(models.Procedure)
	}

	diagnosticStudyOrder := i.patient.Procedures[0]
	c.Assert(len(i.patient.Procedures), Equals, 1)
	c.Assert(diagnosticStudyOrder.ID.Root, Equals, "50f84dbb7042f9366f00014c")
	c.Assert(diagnosticStudyOrder.Codes["LOINC"][0], Equals, "69399-4")
	c.Assert(diagnosticStudyOrder.StartTime, Equals, int64(629709860)) // start and end time should be equal for diagnostic study orders
	c.Assert(diagnosticStudyOrder.EndTime, Equals, int64(629709860))
}

func (i *ImporterSuite) TestExtractTransferFrom(c *C) {
	var transferFromXPath = xpath.Compile("//cda:encounter[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.81']")
	rawTransferFroms := ExtractSection(i.patientElement, transferFromXPath, TransferFromExtractor, "2.16.840.1.113883.3.560.1.71")
	i.patient.Encounters = make([]models.Encounter, len(rawTransferFroms))
	for j := range rawTransferFroms {
		i.patient.Encounters[j] = rawTransferFroms[j].(models.Encounter)
	}

	transferFromEncounter := i.patient.Encounters[0]
	c.Assert(len(i.patient.Encounters), Equals, 1)
	c.Assert(transferFromEncounter.ID.Root, Equals, "49d75f61-0dec-4972-9a51-e2490b18c772")
	c.Assert(transferFromEncounter.Codes["LOINC"][0], Equals, "77305-1")
	c.Assert(transferFromEncounter.StartTime, Equals, int64(1415097000))
	c.Assert(transferFromEncounter.TransferFrom.Time, Equals, int64(1415097000))
	c.Assert(transferFromEncounter.TransferFrom.Codes["SNOMED-CT"][0], Equals, "309911002")
}

func (i *ImporterSuite) TestExtractTransferTo(c *C) {
	var transferToXPath = xpath.Compile("//cda:encounter[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.82']")
	rawTransferTos := ExtractSection(i.patientElement, transferToXPath, TransferToExtractor, "2.16.840.1.113883.3.560.1.72")
	i.patient.Encounters = make([]models.Encounter, len(rawTransferTos))
	for j := range rawTransferTos {
		i.patient.Encounters[j] = rawTransferTos[j].(models.Encounter)
	}

	transferToEncounter := i.patient.Encounters[0]
	c.Assert(len(i.patient.Encounters), Equals, 1)
	c.Assert(transferToEncounter.ID.Root, Equals, "49d75f61-0dec-4972-9a51-e2490b18c772")
	c.Assert(transferToEncounter.Codes["LOINC"][0], Equals, "77306-9")
	c.Assert(transferToEncounter.StartTime, Equals, int64(1415097000))
	c.Assert(transferToEncounter.TransferTo.Time, Equals, int64(1415097000))
	c.Assert(transferToEncounter.TransferTo.Codes["SNOMED-CT"][0], Equals, "309911002")
}

func (i *ImporterSuite) TestMedicationActive(c *C) {
	var medicationActiveXPath = xpath.Compile("//cda:substanceAdministration[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.41']")
	rawMedicationActives := ExtractSection(i.patientElement, medicationActiveXPath, MedicationActiveExtractor, "2.16.840.1.113883.3.560.1.13")
	i.patient.Medications = make([]models.Medication, len(rawMedicationActives))
	for j := range rawMedicationActives {
		i.patient.Medications[j] = rawMedicationActives[j].(models.Medication)
	}

	medicationActive := i.patient.Medications[0]
	c.Assert(len(i.patient.Medications), Equals, 1)
	c.Assert(medicationActive.ID.Root, Equals, "c0ea7bf3-50e7-4e7a-83a3-e5a9ccbb8541")
	c.Assert(medicationActive.Codes["RxNorm"][0], Equals, "105152")
	c.Assert(medicationActive.AdministrationTiming.InstitutionSpecified, Equals, true)
	c.Assert(medicationActive.AdministrationTiming.Period.Unit, Equals, "h")
	c.Assert(medicationActive.AdministrationTiming.Period.Value, Equals, int64(6))
	c.Assert(medicationActive.StartTime, Equals, int64(1092658739))
	c.Assert(medicationActive.EndTime, Equals, int64(1092676026))
	c.Assert(medicationActive.Oid, Equals, "2.16.840.1.113883.3.560.1.13")
	c.Assert(medicationActive.Route.Codes["NCI Thesaurus"][0], Equals, "C38288")
	c.Assert(medicationActive.ProductForm.Codes["NCI Thesaurus"][0], Equals, "C42944")
	c.Assert(medicationActive.DoseRestriction.Numerator.Unit, Equals, "oz")
	c.Assert(medicationActive.DoseRestriction.Numerator.Value, Equals, int64(42))
	c.Assert(medicationActive.DoseRestriction.Denominator.Unit, Equals, "oz")
	c.Assert(medicationActive.DoseRestriction.Denominator.Value, Equals, int64(100))
	c.Assert(medicationActive.OrderInformation[0].OrderNumber, Equals, "12345")
	c.Assert(medicationActive.OrderInformation[0].Fills, Equals, int64(1))
	c.Assert(medicationActive.OrderInformation[0].QuantityOrdered.Value, Equals, int64(75))
	c.Assert(medicationActive.OrderInformation[0].OrderNumber, Equals, "12345")
	c.Assert(medicationActive.OrderInformation[0].OrderDate, Equals, int64(1092676026))

}

func (i *ImporterSuite) TestMedicationDispensed(c *C) {
	var medicationDispensedXPath = xpath.Compile("//cda:supply[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.45']")
	rawMedicationDispenseds := ExtractSection(i.patientElement, medicationDispensedXPath, MedicationDispensedExtractor, "2.16.840.1.113883.3.560.1.8")
	i.patient.Medications = make([]models.Medication, len(rawMedicationDispenseds))
	for j := range rawMedicationDispenseds {
		i.patient.Medications[j] = rawMedicationDispenseds[j].(models.Medication)
	}

	medicationDispensed := i.patient.Medications[0]
	c.Assert(len(i.patient.Medications), Equals, 1)
	c.Assert(medicationDispensed.ID.Root, Equals, "50f84c1b7042f9877500023e")
	c.Assert(medicationDispensed.Codes["RxNorm"][0], Equals, "977869")
	c.Assert(medicationDispensed.StartTime, Equals, int64(822072083))
	c.Assert(medicationDispensed.EndTime, Equals, int64(822089605))
}

func (i *ImporterSuite) TestAllergy(c *C) {
	var allergyXpath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.43']")
	rawAllergies := ExtractSection(i.patientElement, allergyXpath, AllergyExtractor, "2.16.840.1.113883.3.560.1.7")

	i.patient.Allergies = make([]models.Allergy, len(rawAllergies))
	for j := range rawAllergies {
		i.patient.Allergies[j] = rawAllergies[j].(models.Allergy)
	}

	medAllergy := i.patient.Allergies[0]
	c.Assert(len(i.patient.Allergies), Equals, 1)
	c.Assert(medAllergy.ID.Root, Equals, "50f84db97042f9366f00000e")
	c.Assert(medAllergy.Codes["RxNorm"][0], Equals, "996994")
	c.Assert(medAllergy.StartTime, Equals, int64(303055256))
	c.Assert(medAllergy.Type.Codes["ActCode"][0], Equals, "ASSERTION")
	c.Assert(medAllergy.Reaction.Codes["SNOMED-CT"][0], Equals, "422587007")
	c.Assert(medAllergy.Severity.Codes["SNOMED-CT"][0], Equals, "371924009")
}

func (i *ImporterSuite) TestProcedureIntolerance(c *C) {
	var procedureIntoleranceXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.62']/cda:entryRelationship/cda:procedure[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.64']")
	rawProcedureIntolerances := ExtractSection(i.patientElement, procedureIntoleranceXPath, ProcedureIntoleranceExtractor, "2.16.840.1.113883.3.560.1.61")

	i.patient.Allergies = make([]models.Allergy, len(rawProcedureIntolerances))
	for j := range rawProcedureIntolerances {
		i.patient.Allergies[j] = rawProcedureIntolerances[j].(models.Allergy)
	}
	procedureIntolerance := i.patient.Allergies[0]
	c.Assert(procedureIntolerance.ID.Root, Equals, "5102936b944dfe3db4000016")
	c.Assert(procedureIntolerance.Codes["CPT"][0], Equals, "90668")
	c.Assert(procedureIntolerance.Codes["SNOMED-CT"][0], Equals, "86198006")
	c.Assert(procedureIntolerance.StartTime, Equals, int64(1094992715))
	c.Assert(procedureIntolerance.EndTime, Equals, int64(1095042729))
	c.Assert(procedureIntolerance.Oid, Equals, "2.16.840.1.113883.3.560.1.61")
	c.Assert(procedureIntolerance.Values[0].Codes["SNOMED-CT"][0], Equals, "102460003")
}

func (i *ImporterSuite) TestGestationalAge(c *C) {
	var gestationalAgeXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.101']")
	rawGestationalAges := ExtractSection(i.patientElement, gestationalAgeXPath, GestationalAgeExtractor, "2.16.840.1.113883.3.560.1.1001")

	i.patient.Conditions = make([]models.Condition, len(rawGestationalAges))
	for j := range rawGestationalAges {
		i.patient.Conditions[j] = rawGestationalAges[j].(models.Condition)
	}
	gestationalAge := i.patient.Conditions[0]
	c.Assert(gestationalAge.ID.Root, Equals, "50f6c6da7042f9cdd0000233")
	c.Assert(gestationalAge.Oid, Equals, "2.16.840.1.113883.3.560.1.1001")
	c.Assert(gestationalAge.Codes["SNOMED-CT"][0], Equals, "931004")
	c.Assert(gestationalAge.Values[0].Scalar, Equals, int64(36))
	c.Assert(gestationalAge.Values[0].Units, Equals, "wk")
}

func (i *ImporterSuite) TestCommunication(c *C) {
	var communicationXPath = xpath.Compile("//cda:entry/cda:act[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.2']")
	rawCommunications := ExtractSection(i.patientElement, communicationXPath, CommunicationExtractor, "2.16.840.1.113883.3.560.1.30")

	i.patient.Communications = make([]models.Communication, len(rawCommunications))
	for j := range rawCommunications {
		i.patient.Communications[j] = rawCommunications[j].(models.Communication)
	}
	communication := i.patient.Communications[0]
	c.Assert(communication.ID.Root, Equals, "50f84c187042f987750000e5")
	c.Assert(communication.Oid, Equals, "2.16.840.1.113883.3.560.1.30")
	c.Assert(communication.Direction, Equals, "communication_from_patient_to_provider")
	c.Assert(communication.Codes["SNOMED-CT"][0], Equals, "315640000")
	c.Assert(communication.NegationInd, Equals, false)
	c.Assert(communication.Reason.Code, Equals, "105480006")
	c.Assert(communication.Reason.CodeSystem, Equals, "SNOMED-CT")
	reference := communication.References[0]
	c.Assert(reference.ReferencedID, Equals, "56c237ee02d40565bb00030e")
	c.Assert(reference.ReferencedType, Equals, "Procedure")
	c.Assert(reference.Type, Equals, "fulfills")
}

func (i *ImporterSuite) TestEcogStatus(c *C) {
	var ecogStatusXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.103']")
	rawEcogStatuses := ExtractSection(i.patientElement, ecogStatusXPath, ConditionExtractor, "2.16.840.1.113883.3.560.1.1001")
	i.patient.Conditions = make([]models.Condition, len(rawEcogStatuses))
	for j := range rawEcogStatuses {
		i.patient.Conditions[j] = rawEcogStatuses[j].(models.Condition)
	}
	ecogStatus := i.patient.Conditions[0]
	c.Assert(ecogStatus.ID.Root, Equals, "50f6c6067042f91c7c000272")
	c.Assert(ecogStatus.Oid, Equals, "2.16.840.1.113883.3.560.1.1001")
	c.Assert(ecogStatus.Codes["SNOMED-CT"][0], Equals, "423237006")
}

func (i *ImporterSuite) TestSymptomActive(c *C) {
	var symptomActiveXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.76']")
	rawActiveSymptoms := ExtractSection(i.patientElement, symptomActiveXPath, ConditionExtractor, "2.16.840.1.113883.3.560.1.69")
	i.patient.Conditions = make([]models.Condition, len(rawActiveSymptoms))
	for j := range rawActiveSymptoms {
		i.patient.Conditions[j] = rawActiveSymptoms[j].(models.Condition)
	}
	activeSymptom := i.patient.Conditions[0]
	c.Assert(activeSymptom.Codes["SNOMED-CT"][0], Equals, "95815000")
	c.Assert(activeSymptom.StartTime, Equals, int64(729814935))
	c.Assert(activeSymptom.EndTime, Equals, int64(729867188))
	c.Assert(activeSymptom.ID.Root, Equals, "50f84dbb7042f9366f0001ac")
	c.Assert(activeSymptom.Oid, Equals, "2.16.840.1.113883.3.560.1.69")
}
