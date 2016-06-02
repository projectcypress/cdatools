package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

func EncounterOrderExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	encounterOrder := models.Encounter{}
	encounterOrder.Entry = entry

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
	encounter.Entry = entry

	//extract codes
	var codePath = xpath.Compile("cda:code")
	ExtractCodes(&encounter.Entry.Coded, entryElement, codePath)

	//set admit/discharge times
	encounter.AdmitTime = encounter.Entry.StartTime
	encounter.DischargeTime = encounter.Entry.EndTime

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
