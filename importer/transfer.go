package importer

import (
  "github.com/moovweb/gokogiri/xml"
  "github.com/moovweb/gokogiri/xpath"
  "github.com/projectcypress/cdatools/models"
)

// returns Encounter with TransferFrom field
func TransferFromExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
  transferFromEncounter := models.Encounter{}
  transferFromEncounter.Entry = *entry

  // extract codes
  var codePath = xpath.Compile("cda:code")
  ExtractCodes(&transferFromEncounter.Entry.Coded, entryElement, codePath)

  // extract (start time for transfer from entry) and (time for transfer from field on transfer from entry)
  var lowTimeXPath = xpath.Compile("cda:participant/cda:time/cda:low/@value")
  var timeStamp = GetTimestamp(lowTimeXPath, entryElement)
  transferFromEncounter.StartTime = timeStamp // set start time from extracted low time
  transferFromEncounter.TransferFrom.Time = timeStamp // set time on TransferFrom attribute from extracted low time

  // extract transfer from location
  var locationCodePath = xpath.Compile("cda:participant/cda:participantRole[@classCode='LOCE']/cda:code")
  transferFromEncounter.TransferFrom.Codes = map[string][]string{} // create code map
  ExtractCodes(&transferFromEncounter.TransferFrom.Coded, entryElement, locationCodePath)

  return transferFromEncounter
}
