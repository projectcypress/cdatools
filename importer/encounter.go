package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

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
