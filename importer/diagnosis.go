package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

func DiagnosisActiveExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	diagnosisActive := models.Diagnosis{}
	diagnosisActive.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:value/@code")
	var codeSetPath = xpath.Compile("cda:value/@codeSystem")
	ExtractCodes(&diagnosisActive.Entry, entryElement, codePath, codeSetPath)

	//extract severity
	var severityCodeXPath = xpath.Compile("cda:entryRelationship/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.22.4.8']/cda:value/@code")
	var severityCodeSetXPath = xpath.Compile("cda:entryRelationship/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.22.4.8']/cda:value/@codeSystem")
	ExtractSeverity(&diagnosisActive, entryElement, severityXPath)

	return diagnosisActive
}

func DiagnosisInactiveExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	diagnosisInactive := models.Diagnosis{}
	diagnosisInactive.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:value/@code")
	var codeSetPath = xpath.Compile("cda:value/@codeSystem")
	ExtractCodes(&diagnosisActive.Entry, entryElement, codePath, codeSetPath)

	return diagnosisInactive
}

func ExtractSeverity(diagnosis *models.Diagnosis, entryElement xml.Node, severityCodeXPath *xpath.Expression, severityCodeSetXPath *xpath.Expression) {
	severityCode := FirstElementContent(severityCodeXPath, entryElement)
	severityCodeSystem := models.CodeSystemFor(FirstElementContent(severityCodeSetXPath, entryElement))
	diagnosis.Severity = map[string][]string{
		severityCodeSystem: []string{severityCode},
	}
}
