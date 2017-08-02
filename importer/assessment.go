package importer

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

func AssessmentPerformedExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	assessmentPerformed := models.Assessment{}
	assessmentPerformed.Entry = *entry
	
	var codePath = xpath.Compile("cda:code")
	ExtractCodes(&assessmentPerformed.Entry.Coded, entryElement, codePath)
	ExtractReasonOrNegation(&assessmentPerformed.Entry, entryElement)
	scalarPath := xpath.Compile(".//cda:value")
	ExtractValues(&assessmentPerformed.Entry, entryElement, scalarPath)

	// set Status Code
	assessmentPerformed.StatusCode = map[string][]string{}
	assessmentPerformed.StatusCode["HL7 ActStatus"] = []string{"performed"}

	return assessmentPerformed
}