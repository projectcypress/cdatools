package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
)

func ResultExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	result := models.LabResult{}
	result.Entry = *entry

	var codePath = xpath.Compile("cda:code")
	ExtractCodes(&result.Entry.Coded, entryElement, codePath)

	extractInterpretation(&result, entryElement)
	extractReferenceRange(&result, entryElement)

	ExtractReasonOrNegation(&result.Entry, entryElement)

	return result
}

func LabResultExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	labResult := models.LabResult{}
	labResult.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:code")
	ExtractCodes(&labResult.Entry.Coded, entryElement, codePath)

	//extract values
	var valuePath = xpath.Compile("cda:value")
	ExtractValues(&labResult.Entry, entryElement, valuePath)

	return labResult
}

func LabOrderExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	labOrder := models.LabResult{}
	labOrder.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:code")
	ExtractCodes(&labOrder.Entry.Coded, entryElement, codePath)

	//extract order specific dates
	var orderTimeXPath = xpath.Compile("cda:author/cda:time/@value")
	labOrder.StartTime = GetTimestamp(orderTimeXPath, entryElement)
	labOrder.EndTime = GetTimestamp(orderTimeXPath, entryElement)

	return labOrder
}

func extractInterpretation(result *models.LabResult, entryElement xml.Node) {
	var interpretationXPath = xpath.Compile("cda:interpretationCode")
	interpretationElement, err := entryElement.Search(interpretationXPath)
	util.CheckErr(err)
	if len(interpretationElement) > 0 {
		code := interpretationElement[0].Attr("code")
		codeSystem := models.CodeSystemFor(interpretationElement[0].Attr("codeSystem"))
		result.Interpretation.Code = code
		result.Interpretation.CodeSystem = codeSystem
	}
}

func extractReferenceRange(result *models.LabResult, entryElement xml.Node) {
	var referenceRangeXPath = xpath.Compile("./cda:referenceRange/cda:observationRange/cda:text")
	referenceElement, err := entryElement.Search(referenceRangeXPath)
	util.CheckErr(err)
	if len(referenceElement) > 0 {
		result.ReferenceRange = referenceElement[0].Content()
	}
}
