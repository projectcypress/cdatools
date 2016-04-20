package importer

import (
  "github.com/moovweb/gokogiri/xml"
  "github.com/moovweb/gokogiri/xpath"
  "github.com/projectcypress/cdatools/models"
)

func DiagnosticStudyOrderExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
  diagnosticStudyOrder := models.Procedure{}
  diagnosticStudyOrder.Entry = *entry

  // extract codes
  var codePath = xpath.Compile("cda:code")
  ExtractCodes(&diagnosticStudyOrder.Entry.Coded, entryElement, codePath)

  // extract order specific dates
  var orderTimeXPath = xpath.Compile("cda:author/cda:time/@value")
  diagnosticStudyOrder.StartTime = GetTimestamp(orderTimeXPath, entryElement)
  diagnosticStudyOrder.EndTime = GetTimestamp(orderTimeXPath, entryElement)

  return diagnosticStudyOrder
}
