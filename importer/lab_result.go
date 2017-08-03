package importer

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

func ResultExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	result := models.LabResult{}
	result.Entry = *entry

	var codePath = xpath.Compile("cda:code")
	ExtractCodes(&result.Entry.Coded, entryElement, codePath)

	//extract values
	var valuePath = xpath.Compile("./cda:value | ./cda:entryRelationship/cda:observation[./cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.87']/cda:value")
	ExtractValues(&result.Entry, entryElement, valuePath)

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
	ExtractReasonOrNegation(&labResult.Entry, entryElement)

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
	interpretationElement := FirstElement(interpretationXPath, entryElement)
	if interpretationElement != nil {
		code := interpretationElement.Attr("code")
		codeSystem := models.CodeSystemFor(interpretationElement.Attr("codeSystem"))
		result.Interpretation.Code = code
		result.Interpretation.CodeSystem = codeSystem
	}
}

func extractReferenceRange(result *models.LabResult, entryElement xml.Node) {
	var referenceRangeXPath = xpath.Compile("./cda:referenceRange/cda:observationRange/cda:text")
	referenceElement := FirstElement(referenceRangeXPath, entryElement)
	if referenceElement != nil {
		result.ReferenceRange = referenceElement.Content()
	}
}
