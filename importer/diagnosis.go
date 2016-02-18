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
	var codePath = xpath.Compile("cda:value")
	ExtractCodes(&diagnosisActive.Entry, entryElement, codePath)

	//extract severity
	var severityCodeXPath = xpath.Compile("cda:entryRelationship/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.22.4.8']/cda:value/@code")
	var severityCodeSetXPath = xpath.Compile("cda:entryRelationship/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.22.4.8']/cda:value/@codeSystem")
	ExtractSeverity(&diagnosisActive, entryElement, severityCodeXPath, severityCodeSetXPath)

	return diagnosisActive
}

func DiagnosisInactiveExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	diagnosisInactive := models.Diagnosis{}
	diagnosisInactive.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:value")
	ExtractCodes(&diagnosisInactive.Entry, entryElement, codePath)

	return diagnosisInactive
}

func ExtractSeverity(diagnosis *models.Diagnosis, entryElement xml.Node, severityCodeXPath *xpath.Expression, severityCodeSetXPath *xpath.Expression) {
	severityCode := FirstElementContent(severityCodeXPath, entryElement)
	severityCodeSystem := models.CodeSystemFor(FirstElementContent(severityCodeSetXPath, entryElement))
	diagnosis.Severity = map[string][]string{
		severityCodeSystem: []string{severityCode},
	}
}
