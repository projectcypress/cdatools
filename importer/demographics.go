package importer

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

func ExtractDemographics(patient *models.Record, patientElement xml.Node) {
	var firstNameXPath = xpath.Compile("cda:name/cda:given")
	patient.First = FirstElementContent(firstNameXPath, patientElement)
	var lastNameXPath = xpath.Compile("cda:name/cda:family")
	patient.Last = FirstElementContent(lastNameXPath, patientElement)

	var genderXPath = xpath.Compile("cda:administrativeGenderCode/@code")
	genderResults, err := patientElement.Search(genderXPath)
	if err != nil {
		panic(err.Error())
	}
	if len(genderResults) == 0 {
		genderResults, err = patientElement.Search(xpath.Compile("cda:administrativeGenderCode/@nullFlavor"))
		if err != nil {
			panic(err.Error())
		}
	}
	if len(genderResults) == 0 {
		patient.Gender = "NULL"
	} else {
		firstNode := genderResults[0]
		patient.Gender = firstNode.Content()
	}
	var birthTimeXPath = xpath.Compile("cda:birthTime/@value")
	patient.BirthDate = GetTimestamp(birthTimeXPath, patientElement)

	patient.Race = &models.CodedConcept{}
	var raceXPath = xpath.Compile("cda:raceCode/@code")
	patient.Race.Code = FirstElementContent(raceXPath, patientElement)
	var raceCodeSetXPath = xpath.Compile("cda:raceCode/@codeSystemName")
	patient.Race.CodeSystem = FirstElementContent(raceCodeSetXPath, patientElement)
	var raceDisplayNameXPath = xpath.Compile("cda:raceCode/@displayName")
	patient.Race.DisplayName = FirstElementContent(raceDisplayNameXPath, patientElement)

	patient.Ethnicity = &models.CodedConcept{}
	var ethnicityXPath = xpath.Compile("cda:ethnicGroupCode/@code")
	patient.Ethnicity.Code = FirstElementContent(ethnicityXPath, patientElement)
	var ethnicityCodeSetXPath = xpath.Compile("cda:ethnicGroupCode/@codeSystemName")
	patient.Ethnicity.CodeSystem = FirstElementContent(ethnicityCodeSetXPath, patientElement)
	var ethnicityDisplayNameXPath = xpath.Compile("cda:ethnicityCode/@displayName")
	patient.Ethnicity.DisplayName = FirstElementContent(ethnicityDisplayNameXPath, patientElement)
}
