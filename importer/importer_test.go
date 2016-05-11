package importer

import (
	"io/ioutil"
	"testing"

	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
	. "gopkg.in/check.v1"
	"strconv"
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
	c.Assert(labResult.Entry.Values[0].Scalar, Equals, "positive")
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
	c.Assert(gestationalAge.Values[0].Scalar, Equals, strconv.Itoa(int(36)))
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

func (i *ImporterSuite) TestDiagnosisResolved(c *C) {
	var diagonsisResolvedXPath = xpath.Compile("//cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.14']")
	rawDiagnosesResolved := ExtractSection(i.patientElement, diagonsisResolvedXPath, ConditionExtractor, "2.16.840.1.113883.3.560.1.24")
	i.patient.Conditions = make([]models.Condition, len(rawDiagnosesResolved))
	for j := range rawDiagnosesResolved {
		i.patient.Conditions[j] = rawDiagnosesResolved[j].(models.Condition)
	}
	diagnosisResolved := i.patient.Conditions[0]
	c.Assert(diagnosisResolved.ID.Root, Equals, "50f84c187042f98775000089")
	c.Assert(diagnosisResolved.Oid, Equals, "2.16.840.1.113883.3.560.1.24")
	c.Assert(diagnosisResolved.Codes["SNOMED-CT"][0], Equals, "94643001")
	c.Assert(diagnosisResolved.Codes["ICD-10-CM"][0], Equals, "C21.8")
	c.Assert(diagnosisResolved.Codes["ICD-9-CM"][0], Equals, "197.5")
}

func (i *ImporterSuite) TestLabResultPerformed(c *C) {
	var labResultPerformedXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.38']")
	rawLabResults := ExtractSection(i.patientElement, labResultPerformedXPath, ResultExtractor, "2.16.840.1.113883.3.560.1.5")
	i.patient.LabResults = make([]models.LabResult, len(rawLabResults))
	for j := range rawLabResults {
		i.patient.LabResults[j] = rawLabResults[j].(models.LabResult)
	}
	labResult := i.patient.LabResults[0]
	c.Assert(labResult.ID.Root, Equals, "50f84c1d7042f98775000353")
	c.Assert(labResult.Oid, Equals, "2.16.840.1.113883.3.560.1.5")
	// These are commented out until the CodeSystems get added to the Codesystemmap
	// c.Assert(labResult.Interpretation.Code, Equals, "N")
	// c.Assert(labResult.Interpretation.CodeSystem, Equals, "HITSP C80 Observation Status")
	c.Assert(labResult.ReferenceRange, Equals, "M 13-18 g/dl; F 12-16 g/dl")
}

func (i *ImporterSuite) TestMedicalEquipmentApplied(c *C) {
	var medEquipAppliedXPath = xpath.Compile("//cda:procedure[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.7']")
	rawMedEquipApplied := ExtractSection(i.patientElement, medEquipAppliedXPath, MedicalEquipmentExtractor, "2.16.840.1.113883.3.560.1.110")
	i.patient.MedicalEquipment = make([]models.MedicalEquipment, len(rawMedEquipApplied))
	for j := range rawMedEquipApplied {
		i.patient.MedicalEquipment[j] = rawMedEquipApplied[j].(models.MedicalEquipment)
	}
	medEquipApplied := i.patient.MedicalEquipment[0]
	c.Assert(medEquipApplied.ID.Root, Equals, "510969b3944dfe9bd7000056")
	c.Assert(medEquipApplied.StartTime, Equals, int64(481091888))
	c.Assert(medEquipApplied.Codes["ICD-9-CM"][0], Equals, "37.98")
	c.Assert(medEquipApplied.AnatomicalStructure.Code, Equals, "thigh")
	c.Assert(medEquipApplied.AnatomicalStructure.CodeSystem, Equals, "2.16.840.1.113883.6.96")
	c.Assert(medEquipApplied.AnatomicalStructure.CodeSystemName, Equals, "SNOMED-CT")
}

func (i *ImporterSuite) TestExtractProcedurePerformed(c *C) {
	var procedurePerformedXPath = xpath.Compile("//cda:entry/cda:procedure[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.64']")
	rawProcedurePerformed := ExtractSection(i.patientElement, procedurePerformedXPath, ProcedurePerformedExtractor, "2.16.840.1.113883.3.560.1.6")
	i.patient.Procedures = make([]models.Procedure, len(rawProcedurePerformed))
	for j := range rawProcedurePerformed {
		i.patient.Procedures[j] = rawProcedurePerformed[j].(models.Procedure)
	}

	procedurePerformed := i.patient.Procedures[0]
	c.Assert(len(i.patient.Procedures), Equals, 2) // there are two procedure performed in cat1_good.xml
	c.Assert(procedurePerformed.ID.Root, Equals, "51083f0e944dfe9bd7000004")
	c.Assert(procedurePerformed.Oid, Equals, "2.16.840.1.113883.3.560.1.6") // hqmf oid
	c.Assert(procedurePerformed.Codes["SNOMED-CT"][0], Equals, "236211007")
	c.Assert(procedurePerformed.Ordinality.Codes["SNOMED-CT"][0], Equals, "63161005")
	c.Assert(procedurePerformed.StartTime, Equals, int64(506358845))
	c.Assert(procedurePerformed.EndTime, Equals, int64(506409573))
	c.Assert(procedurePerformed.IncisionTime, Equals, int64(506358905))
	c.Assert(procedurePerformed.NegationInd, Equals, true)
	c.Assert(procedurePerformed.NegationReason, Equals, models.CodedConcept{}) // no negation reason

	// tests not included in health data standards
	c.Assert(procedurePerformed.AnatomicalTarget.Code, Equals, "28273000")
	c.Assert(procedurePerformed.AnatomicalTarget.CodeSystem, Equals, "SNOMED-CT")
	c.Assert(procedurePerformed.AnatomicalTarget.CodeSystemOid, Equals, "2.16.840.1.113883.6.96")

	// find second procedurePerformed
	for j, procedure := range i.patient.Procedures {
		if procedure.ID.Root == "51083f0e944dfe9bd7001234" {
			procedurePerformed = i.patient.Procedures[j]
		}
	}
	// second procedure performed has negation reasons (not just negation indicator with no reason)
	c.Assert(procedurePerformed.ID.Root, Equals, "51083f0e944dfe9bd7001234")
	c.Assert(procedurePerformed.NegationInd, Equals, true)
	c.Assert(procedurePerformed.NegationReason.Code, Equals, "308292007")
	c.Assert(procedurePerformed.NegationReason.CodeSystem, Equals, "SNOMED-CT")

	// second procedure performed also has values tags with different formats
	c.Assert(procedurePerformed.Values[0].Scalar, Equals, "6")
	c.Assert(procedurePerformed.Values[0].Units, Equals, "m[IU]/L")
	c.Assert(procedurePerformed.Values[1].Scalar, Equals, "true")
	c.Assert(procedurePerformed.Values[1].Units, Equals, "")
	c.Assert(procedurePerformed.Values[2].Scalar, Equals, "my_string_value")
	c.Assert(procedurePerformed.Values[2].Units, Equals, "")
}

func (i *ImporterSuite) TestExtractPhysicalExamPerformed(c *C) {
	var physicalExamPerformedXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.59']")
	rawPhysicalExamPerformed := ExtractSection(i.patientElement, physicalExamPerformedXPath, ProcedureExtractor, "2.16.840.1.113883.3.560.1.57")
	i.patient.Procedures = make([]models.Procedure, len(rawPhysicalExamPerformed))
	for j := range rawPhysicalExamPerformed {
		i.patient.Procedures[j] = rawPhysicalExamPerformed[j].(models.Procedure)
	}

	physicalExamPerformed := i.patient.Procedures[0]
	c.Assert(len(i.patient.Procedures), Equals, 1)
	c.Assert(physicalExamPerformed.ID.Root, Equals, "5101a4f7944dfe3db4000006")
	c.Assert(physicalExamPerformed.Oid, Equals, "2.16.840.1.113883.3.560.1.57") // hqmf oid
	c.Assert(physicalExamPerformed.Codes["LOINC"][0], Equals, "8462-4")
	c.Assert(physicalExamPerformed.StartTime, Equals, int64(751003636))
	c.Assert(physicalExamPerformed.EndTime, Equals, int64(751060302))
	c.Assert(physicalExamPerformed.NegationInd, Equals, true)
	c.Assert(physicalExamPerformed.NegationReason, Equals, models.CodedConcept{})
}

func (i *ImporterSuite) TestExtractInterventionOrder(c *C) {
	var interventionOrderXPath = xpath.Compile("//cda:entry/cda:act[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.31']")
	rawInterventionOrder := ExtractSection(i.patientElement, interventionOrderXPath, ProcedureExtractor, "2.16.840.1.113883.3.560.1.45")
	i.patient.Procedures = make([]models.Procedure, len(rawInterventionOrder))
	for j := range rawInterventionOrder {
		i.patient.Procedures[j] = rawInterventionOrder[j].(models.Procedure)
	}

	interventionOrder := i.patient.Procedures[0]
	c.Assert(len(i.patient.Procedures), Equals, 1)
	c.Assert(interventionOrder.ID.Root, Equals, "510831719eae47faed000150")
	c.Assert(interventionOrder.Oid, Equals, "2.16.840.1.113883.3.560.1.45")
	c.Assert(interventionOrder.Codes["CPT"][0], Equals, "43644")
	c.Assert(interventionOrder.Codes["ICD-9-CM"][0], Equals, "V65.3")
	c.Assert(interventionOrder.Codes["ICD-10-CM"][0], Equals, "Z71.3")
	c.Assert(interventionOrder.Codes["SNOMED-CT"][0], Equals, "304549008")
	c.Assert(interventionOrder.StartTime, Equals, int64(1277424000))
}

func (i *ImporterSuite) TestExtractInterventionPerformed(c *C) {
	var interventionPerformedXPath = xpath.Compile("//cda:entry/cda:act[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.32']")
	rawInterventionPerformed := ExtractSection(i.patientElement, interventionPerformedXPath, ProcedureExtractor, "2.16.840.1.113883.3.560.1.46")
	i.patient.Procedures = make([]models.Procedure, len(rawInterventionPerformed))
	for j := range rawInterventionPerformed {
		i.patient.Procedures[j] = rawInterventionPerformed[j].(models.Procedure)
	}

	interventionPerformed := i.patient.Procedures[0]
	c.Assert(len(i.patient.Procedures), Equals, 1)
	c.Assert(interventionPerformed.ID.Root, Equals, "1.3.6.1.4.1.115")
	c.Assert(interventionPerformed.Oid, Equals, "2.16.840.1.113883.3.560.1.46")
	c.Assert(interventionPerformed.Codes["SNOMED-CT"][0], Equals, "103699006")
	c.Assert(interventionPerformed.StartTime, Equals, int64(1396189800))
	c.Assert(interventionPerformed.EndTime, Equals, int64(1396189800))
}

func (i *ImporterSuite) TestExtractProcedureInterventionResults(c *C) {
	var procedureInterventionResultXPath = xpath.Compile("//cda:entry/cda:act[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.34']")
	rawProcedureInterventionResults := ExtractSection(i.patientElement, procedureInterventionResultXPath, ProcedureExtractor, "2.16.840.1.113883.3.560.1.47")
	i.patient.Procedures = make([]models.Procedure, len(rawProcedureInterventionResults))
	for j := range rawProcedureInterventionResults {
		i.patient.Procedures[j] = rawProcedureInterventionResults[j].(models.Procedure)
	}

	interventionResults := i.patient.Procedures[0]
	c.Assert(len(i.patient.Procedures), Equals, 1)
	c.Assert(interventionResults.ID.Root, Equals, "50f84c1c7042f987750002d1")
	c.Assert(interventionResults.Oid, Equals, "2.16.840.1.113883.3.560.1.47")
	c.Assert(interventionResults.Codes["SNOMED-CT"][0], Equals, "428181000124104")
	c.Assert(interventionResults.StartTime, Equals, int64(1097940444))
	c.Assert(interventionResults.EndTime, Equals, int64(1097959712))
}

func (i *ImporterSuite) TestExtractProcedureOrder(c *C) {
	var procedureOrderXPath = xpath.Compile("//cda:entry/cda:procedure[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.63']")
	rawProcedureOrders := ExtractSection(i.patientElement, procedureOrderXPath, ProcedureOrderExtractor, "2.16.840.1.113883.3.560.1.62")
	i.patient.Procedures = make([]models.Procedure, len(rawProcedureOrders))
	for j := range rawProcedureOrders {
		i.patient.Procedures[j] = rawProcedureOrders[j].(models.Procedure)
	}

	procedureOrder := i.patient.Procedures[0]
	c.Assert(len(i.patient.Procedures), Equals, 1)
	c.Assert(procedureOrder.ID.Root, Equals, "5106e039944dfe4d2000000d")
	c.Assert(procedureOrder.Oid, Equals, "2.16.840.1.113883.3.560.1.62")
	c.Assert(procedureOrder.Codes["CPT"][0], Equals, "90870") // only code tested by health data standards
	c.Assert(procedureOrder.Codes["ICD-10-PCS"][0], Equals, "GZB4ZZZ")
	c.Assert(procedureOrder.Codes["SNOMED-CT"][0], Equals, "313020008")
	c.Assert(procedureOrder.Codes["ICD-9-CM"][0], Equals, "94.27")

	c.Assert(procedureOrder.Time, Equals, int64(1306230203))
	c.Assert(procedureOrder.StatusCode["HL7 ActStatus"][0], Equals, "ordered")
}

func (i *ImporterSuite) TestExtractProcedureResults(c *C) {
	var procedureResultsXPath = xpath.Compile("//cda:entry/cda:procedure[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.66']")
	rawProcedureResults := ExtractSection(i.patientElement, procedureResultsXPath, ProcedureExtractor, "2.16.840.1.113883.3.560.1.63")
	i.patient.Procedures = make([]models.Procedure, len(rawProcedureResults))
	for j := range rawProcedureResults {
		i.patient.Procedures[j] = rawProcedureResults[j].(models.Procedure)
	}

	procedureResult := i.patient.Procedures[0]
	c.Assert(len(i.patient.Procedures), Equals, 1)
	c.Assert(procedureResult.ID.Root, Equals, "51095fc3944dfe9bd7000012")
	c.Assert(procedureResult.Oid, Equals, "2.16.840.1.113883.3.560.1.63")
	c.Assert(procedureResult.Codes["SNOMED-CT"][0], Equals, "116783008")
	c.Assert(procedureResult.StartTime, Equals, int64(1007264866))
	c.Assert(procedureResult.EndTime, Equals, int64(1007316283))
}

func (i *ImporterSuite) TestExtractRiskCategoryAssessment(c *C) {
	var riskCategoryAssessmentXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.69']")
	rawRiskCategoryAssessments := ExtractSection(i.patientElement, riskCategoryAssessmentXPath, ProcedureExtractor, "2.16.840.1.113883.3.560.1.21")
	i.patient.Procedures = make([]models.Procedure, len(rawRiskCategoryAssessments))
	for j := range rawRiskCategoryAssessments {
		i.patient.Procedures[j] = rawRiskCategoryAssessments[j].(models.Procedure)
	}

	riskCategoryAssessment := i.patient.Procedures[0]
	c.Assert(len(i.patient.Procedures), Equals, 1)
	c.Assert(riskCategoryAssessment.ID.Root, Equals, "510963e9944dfe9bd7000047")
	c.Assert(riskCategoryAssessment.Oid, Equals, "2.16.840.1.113883.3.560.1.21")
	c.Assert(riskCategoryAssessment.Codes["LOINC"][0], Equals, "72136-5")
	c.Assert(riskCategoryAssessment.StartTime, Equals, int64(744555728))
}

func (i *ImporterSuite) TestExtractDiagnosticStudyNotPerformed(c *C) {
	var diagnosticStudyNotPerformedXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.18']")
	rawDiagnosticStudyNotPerformed := ExtractSection(i.patientElement, diagnosticStudyNotPerformedXPath, ProcedureExtractor, "2.16.840.1.113883.3.560.1.103")
	i.patient.Procedures = make([]models.Procedure, len(rawDiagnosticStudyNotPerformed))
	for j := range rawDiagnosticStudyNotPerformed {
		i.patient.Procedures[j] = rawDiagnosticStudyNotPerformed[j].(models.Procedure)
	}

	diagnosticStudyNotPerformed := i.patient.Procedures[0]
	c.Assert(len(i.patient.Procedures), Equals, 1)
	c.Assert(diagnosticStudyNotPerformed.ID.Root, Equals, "50f84dbb7042f9366f000143")
	c.Assert(diagnosticStudyNotPerformed.Oid, Equals, "2.16.840.1.113883.3.560.1.103")
	c.Assert(diagnosticStudyNotPerformed.Codes["LOINC"][0], Equals, "69399-4")
	c.Assert(diagnosticStudyNotPerformed.StartTime, Equals, int64(1225314966))
	c.Assert(diagnosticStudyNotPerformed.EndTime, Equals, int64(1225321540))
}

func (i *ImporterSuite) TestExtractDiagnosticStudyResult(c *C) {
	var diagnosticStudyResultXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.20']")
	rawDiagnosticStudyResults := ExtractSection(i.patientElement, diagnosticStudyResultXPath, ProcedureExtractor, "2.16.840.1.113883.3.560.1.11")
	i.patient.Procedures = make([]models.Procedure, len(rawDiagnosticStudyResults))
	for j := range rawDiagnosticStudyResults {
		i.patient.Procedures[j] = rawDiagnosticStudyResults[j].(models.Procedure)
	}

	diagnosticStudyResult := i.patient.Procedures[0]
	c.Assert(len(i.patient.Procedures), Equals, 1)
	c.Assert(diagnosticStudyResult.ID.Root, Equals, "50f84c1b7042f987750001e7")
	c.Assert(diagnosticStudyResult.Oid, Equals, "2.16.840.1.113883.3.560.1.11")
	c.Assert(diagnosticStudyResult.Codes["LOINC"][0], Equals, "71485-7")
	c.Assert(diagnosticStudyResult.StartTime, Equals, int64(622535563))
	c.Assert(diagnosticStudyResult.EndTime, Equals, int64(622548751))
	c.Assert(diagnosticStudyResult.NegationInd, Equals, true)
	c.Assert(diagnosticStudyResult.NegationReason, Equals, models.CodedConcept{Code: "79899007"})
}

func (i *ImporterSuite) TestExtractCareGoal(c *C) {
	var careGoalXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.1']")
	rawCareGoals := ExtractSection(i.patientElement, careGoalXPath, CareGoalExtractor, "2.16.840.1.113883.3.560.1.9")
	i.patient.CareGoals = make([]models.CareGoal, len(rawCareGoals))
	for j := range rawCareGoals {
		i.patient.CareGoals[j] = rawCareGoals[j].(models.CareGoal)
	}

	careGoal := i.patient.CareGoals[0]
	c.Assert(len(i.patient.CareGoals), Equals, 1)
	c.Assert(careGoal.Codes["SNOMED-CT"][0], Equals, "252465000")
	c.Assert(careGoal.Oid, Equals, "2.16.840.1.113883.3.560.1.9")
	c.Assert(careGoal.StartTime, Equals, int64(1293890400))
}

func (i *ImporterSuite) TestExtractPatientCharacteristicExpired(c *C) {

	var patientExpiredXPath = xpath.Compile("//cda:entry/cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.54']")
	rawPatientExpireds := ExtractSection(i.patientElement, patientExpiredXPath, ConditionExtractor, "2.16.840.1.113883.3.560.1.404")
	i.patient.Conditions = make([]models.Condition, len(rawPatientExpireds))
	for j := range rawPatientExpireds {
		i.patient.Conditions[j] = rawPatientExpireds[j].(models.Condition)
	}

	// before extracting patient characteristic expired, patient should not be dead
	c.Assert(false, Equals, i.patient.Expired)

	// set Expired and DeathDate if patient is dead
	set_patient_expired(i.patient, i.patientElement)

	// after extracting patient characteristic expired, patient should be dead
	c.Assert(true, Equals, i.patient.Expired)

	patientExpired := i.patient.Conditions[0]
	c.Assert(len(i.patient.Conditions), Equals, 1)
	c.Assert(patientExpired.ID.Root, Equals, "22aeb750-4308-0130-0ade-680688cbd736")
	c.Assert(patientExpired.Oid, Equals, "2.16.840.1.113883.3.560.1.404")
	c.Assert(i.patient.DeathDate, Equals, int64(1450141290))
}
