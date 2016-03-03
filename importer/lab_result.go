package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

func LabResultExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	labResult := models.LabResult{}
	labResult.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:code")
	ExtractCodes(&labResult.Entry, entryElement, codePath)

	return labResult
}

func LabOrderExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	labOrder := models.LabResult{}
	labOrder.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:code")
	ExtractCodes(&labOrder.Entry, entryElement, codePath)

	//extract order specific dates
	var orderTimeXPath = xpath.Compile("cda:author/cda:time/@value")
	labOrder.StartTime = GetTimestamp(orderTimeXPath, entryElement)
	labOrder.EndTime = GetTimestamp(orderTimeXPath, entryElement)

	return labOrder
}
