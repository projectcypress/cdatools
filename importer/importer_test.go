package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"testing"
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
	xp.RegisterNamespace("stdc", "urn:hl7-org:sdtc")

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
	c.Assert(i.patient.Birthdate, Equals, int64(599616000))
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
	c.Assert(encounter.StartTime, Equals, int64(1288569600))
	c.Assert(encounter.EndTime, Equals, int64(1288569600))
}

func (i *ImporterSuite) TestExtractEncounterOrdered(c *C) {
	var encounterOrderXPath = xpath.Compile("//cda:encounter[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.22']")
	rawEncounterOrders := ExtractSection(i.patientElement, encounterOrderXPath, EncounterOrderExtractor, "")
	i.patient.Encounters = make([]models.Encounter, len(rawEncounterOrders))
	for j := range rawEncounterOrders {
		i.patient.Encounters = append(i.patient.Encounters, rawEncounterOrders[j].(models.Encounter))
	}

	c.Assert(len(i.patient.Encounters), Equals, 2)
}

func (i *ImporterSuite) TestExtractDiagnosesActive(c *C) {
	var diagnosisXPath = xpath.Compile("//cda:observation[cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.11']")
	rawDiagnoses := ExtractSection(i.patientElement, diagnosisXPath, DiagnosisActiveExtractor, "2.16.840.1.113883.3.560.1.2")
	i.patient.Diagnoses = make([]models.Diagnosis, len(rawDiagnoses))
	for j := range rawDiagnoses {
		i.patient.Diagnoses[j] = rawDiagnoses[j].(models.Diagnosis)
	}

	c.Assert(len(i.patient.Diagnoses), Equals, 3)
	firstDiagnosis := i.patient.Diagnoses[0]
	c.Assert(firstDiagnosis.ID.Root, Equals, "1.3.6.1.4.1.115")
	c.Assert(firstDiagnosis.ID.Extension, Equals, "54c1142869702d2cd2520100")
	c.Assert(firstDiagnosis.Codes["SNOMED-CT"][0], Equals, "195080001")
	c.Assert(firstDiagnosis.Description, Equals, "Diagnosis, Active: Atrial Fibrillation/Flutter")
	c.Assert(firstDiagnosis.StartTime, Equals, int64(1332720000))
	c.Assert(firstDiagnosis.EndTime, Equals, int64(0))

	secondDiagnosis := i.patient.Diagnoses[1]
	c.Assert(secondDiagnosis.ID.Root, Equals, "1.3.6.1.4.1.115")
	c.Assert(secondDiagnosis.ID.Extension, Equals, "54c1142969702d2cd2cd0200")
	c.Assert(secondDiagnosis.Codes["SNOMED-CT"][0], Equals, "237244005")
	c.Assert(secondDiagnosis.Description, Equals, "Diagnosis, Active: Pregnancy Dx")
	c.Assert(secondDiagnosis.StartTime, Equals, int64(1362096000))
	c.Assert(secondDiagnosis.EndTime, Equals, int64(1382227200))

	thirdDiagnosis := i.patient.Diagnoses[2]
	c.Assert(thirdDiagnosis.ID.Root, Equals, "1.3.6.1.4.1.115")
	c.Assert(thirdDiagnosis.ID.Extension, Equals, "54c1142869702d2cd2760100")
	c.Assert(thirdDiagnosis.Codes["SNOMED-CT"][0], Equals, "46635009")
	c.Assert(thirdDiagnosis.Description, Equals, "Diagnosis, Active: Diabetes")
	c.Assert(thirdDiagnosis.StartTime, Equals, int64(1361836800))
	c.Assert(thirdDiagnosis.EndTime, Equals, int64(0))
}

func (i *ImporterSuite) TestExtractDiagnosesOrdered(c *C) {
	//TODO:write this
}
