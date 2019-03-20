package importer

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

// TransferFromExtractor returns Encounter with TransferFrom field
func TransferFromExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	transferFromEncounter := models.Encounter{}
	transferFromEncounter.Entry = *entry

	extractCodes(&transferFromEncounter.Entry.Coded, entryElement)
	extractTimes(&transferFromEncounter, &transferFromEncounter.TransferFrom, entryElement)

	var locationCodePath = xpath.Compile("cda:participant[@typeCode='ORG']/cda:participantRole[@classCode='LOCE']/cda:code")
	extractLocation(&transferFromEncounter.TransferFrom, entryElement, locationCodePath)
	ExtractCodedConcept(&transferFromEncounter.TransferFrom.CodedConcept, entryElement, locationCodePath)

	return transferFromEncounter
}

// TransferToExtractor returns Encounter with TransferTo field
func TransferToExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	transferToEncounter := models.Encounter{}
	transferToEncounter.Entry = *entry

	extractCodes(&transferToEncounter.Entry.Coded, entryElement)
	extractTimes(&transferToEncounter, &transferToEncounter.TransferTo, entryElement)

	var locationCodePath = xpath.Compile("cda:participant[@typeCode='DST']/cda:participantRole[@classCode='LOCE']/cda:code")
	extractLocation(&transferToEncounter.TransferTo, entryElement, locationCodePath)
	ExtractCodedConcept(&transferToEncounter.TransferTo.CodedConcept, entryElement, locationCodePath)

	return transferToEncounter
}

// private

func extractCodes(coded *models.Coded, entryElement xml.Node) {
	var codePath = xpath.Compile("cda:code")
	ExtractCodes(coded, entryElement, codePath)
}

func extractTimes(encounter *models.Encounter, transfer *models.Transfer, entryElement xml.Node) {
	var lowTimeXPath = xpath.Compile("cda:participant/cda:time/cda:low/@value")
	var timeStamp = GetTimestamp(lowTimeXPath, entryElement)
	encounter.StartTime = timeStamp // set start time on encounter from extracted low time
	transfer.Time = timeStamp       // set time on transfer attribute from extracted low time
}

// code path is xpath to location code
func extractLocation(transfer *models.Transfer, entryElement xml.Node, codePath *xpath.Expression) {
	transfer.Codes = map[string][]string{} // create code map
	ExtractCodes(&transfer.Coded, entryElement, codePath)
}
