package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

// AllergyExtractor extracts allergy/intolerance/adverse event-specific data, such as the specific reaction & the severity
func AllergyExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	allergy := models.Allergy{}
	allergy.Entry = *entry

	codeXPath := xpath.Compile("./cda:participant/cda:participantRole/cda:playingEntity/cda:code")
	ExtractCodes(&allergy.Entry.Coded, entryElement, codeXPath)

	typeXPath := xpath.Compile("./cda:code")
	allergy.Type.Codes = map[string][]string{}
	ExtractCodes(&allergy.Type, entryElement, typeXPath)

	reactionXPath := xpath.Compile("./cda:entryRelationship[@typeCode='MFST']/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.24.3.85']/cda:value")
	allergy.Reaction.Codes = map[string][]string{}
	ExtractCodes(&allergy.Reaction, entryElement, reactionXPath)

	severityXPath := xpath.Compile("./cda:entryRelationship[@typeCode='SUBJ']/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.22.4.8']/cda:value")
	allergy.Severity.Codes = map[string][]string{}
	ExtractCodes(&allergy.Severity, entryElement, severityXPath)

	return allergy
}

// ProcedureIntoleranceExtractor extracts the intolerance of a patient to a procedure
func ProcedureIntoleranceExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	allergy := models.Allergy{}
	allergy.Entry = *entry

	codeXPath := xpath.Compile("./cda:code")
	ExtractCodes(&allergy.Entry.Coded, entryElement, codeXPath)

	valueXPath := xpath.Compile("../../cda:value")
	allergy.Entry.Values = make([]models.ResultValue, 0)
	ExtractValues(&allergy.Entry, entryElement, valueXPath)

	return allergy
}
