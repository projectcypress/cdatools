package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
)

func EncounterOrderExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	encounterOrder := models.Encounter{}
	encounterOrder.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:code")
	ExtractCodes(&encounterOrder.Entry.Coded, entryElement, codePath)

	//extract order specific dates
	var orderTimeXPath = xpath.Compile("cda:author/cda:time/@value")
	encounterOrder.StartTime = GetTimestamp(orderTimeXPath, entryElement)
	encounterOrder.EndTime = GetTimestamp(orderTimeXPath, entryElement)

	return encounterOrder
}

func EncounterPerformedExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	encounter := models.Encounter{}
	encounter.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:code")
	ExtractCodes(&encounter.Entry.Coded, entryElement, codePath)

	//set admit/discharge times
	encounter.AdmitTime = encounter.Entry.StartTime
	encounter.DischargeTime = encounter.Entry.EndTime

	//extract reason
	extractReason(&encounter, entryElement)

	//extract diagnoses
	var pdXPath = xpath.Compile("cda:entryRelationship/cda:observation[cda:code/@code='8319008']")
	var diagXPath = xpath.Compile("cda:entryRelationship/cda:act/cda:entryRelationship/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.22.4.4']")
	ExtractCodes(&encounter.PrincipalDiagnosis, entryElement, pdXPath)
	ExtractCodes(&encounter.Diagnosis, entryElement, diagXPath)

	//extract facility
	extractFacility(&encounter, entryElement)

	//extract discharge disposition
	var dischargeDispositionCodeXPath = xpath.Compile("sdtc:dischargeDispositionCode/@code")
	var dischargeDispositionCodeSystemXPath = xpath.Compile("sdtc:dischargeDispositionCode/@codeSystem")
	dischargeDispositionCode := FirstElementContent(dischargeDispositionCodeXPath, entryElement)
	dischargeDispositionCodeSystemOid := FirstElementContent(dischargeDispositionCodeSystemXPath, entryElement)
	dischargeDispositionCodeSystem := models.CodeSystemFor(dischargeDispositionCodeSystemOid)
	if dischargeDispositionCode != "" {
		encounter.DischargeDisposition = map[string]string{
			"code":          dischargeDispositionCode,
			"codeSystem":    dischargeDispositionCodeSystem,
			"codeSystemOid": dischargeDispositionCodeSystemOid,
		}
	}

	return encounter
}

func extractFacility(encounter *models.Encounter, entryElement xml.Node) {
	var participantXPath = xpath.Compile("cda:participant[@typeCode='LOC']/cda:participantRole[@classCode='SDLOC']")
	participantElement := FirstElement(participantXPath, entryElement)

	if participantElement != nil {
		var facility = models.Facility{}

		var nameXPath = xpath.Compile("cda:playingEntity/cda:name")
		facility.Name = FirstElementContent(nameXPath, participantElement)

		addressXPath := xpath.Compile("cda:addr")
		addressElements, err := participantElement.Search(addressXPath)
		util.CheckErr(err)
		facility.Addresses = make([]models.Address, len(addressElements))
		for i, addressElement := range addressElements {
			facility.Addresses[i] = ImportAddress(addressElement)
		}

		telecomXPath := xpath.Compile("cda:telecom")
		telecomElements, err := participantElement.Search(telecomXPath)
		util.CheckErr(err)
		facility.Telecoms = make([]models.Telecom, len(telecomElements))
		for i, telecomElement := range telecomElements {
			facility.Telecoms[i] = ImportTelecom(telecomElement)
		}

		facility.Code = &models.CodedConcept{}
		ExtractCodedConcept(facility.Code, participantElement, xpath.Compile("cda:code"))

		var timeLowXPath = xpath.Compile("cda:effectiveTime/cda:low/@value")
		var timeHighXPath = xpath.Compile("cda:effectiveTime/cda:high/@value")
		facility.StartTime = GetTimestamp(timeLowXPath, entryElement)
		facility.EndTime = GetTimestamp(timeHighXPath, entryElement)

		encounter.Facility = facility
	}
}

func extractReason(encounter *models.Encounter, entryElement xml.Node) {
	var reasonXPath = xpath.Compile("cda:entryRelationship[@typeCode='RSON']/cda:observation")
	reasonElement := FirstElement(reasonXPath, entryElement)

	if reasonElement != nil {
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
