package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
	"strconv"
)

func GestationalAgeExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	gestationalAge := models.Condition{}
	gestationalAge.Entry = *entry
	codeXPath := xpath.Compile("./cda:code")
	ExtractCodes(&gestationalAge.Entry.Coded, entryElement, codeXPath)

	valueXPath := xpath.Compile("./cda:value")
	entry.Values = make([]models.ResultValue, 0)
	ExtractValues(&gestationalAge.Entry, entryElement, valueXPath)

	gestationalAgeScalar, err := strconv.Atoi(gestationalAge.Entry.Values[0].Scalar)
	util.CheckErr(err)
	switch gestationalAgeScalar {
	case 39:
		entry.Codes["SNOMED-CT"] = []string{"80487005"}
	case 38:
		entry.Codes["SNOMED-CT"] = []string{"13798002"}
	case 37:
		entry.Codes["SNOMED-CT"] = []string{"43697006"}
	case 36:
		entry.Codes["SNOMED-CT"] = []string{"931004"}
	}

	return gestationalAge
}

func ConditionExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	condition := models.Condition{}
	condition.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:value")
	ExtractCodes(&condition.Entry.Coded, entryElement, codePath)

	//extract severity
	var severityXPath = xpath.Compile("cda:entryRelationship/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.22.4.8']/cda:value")
	ExtractSeverity(&condition, entryElement, severityXPath)

	//extract laterality
	var lateralityXPath = xpath.Compile("cda:targetSiteCode")
	ExtractLaterality(&condition, entryElement, lateralityXPath)

	return condition
}

func DiagnosisInactiveExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	diagnosisInactive := models.Condition{}
	diagnosisInactive.Entry = *entry

	//extract codes
	var codePath = xpath.Compile("cda:value")
	ExtractCodes(&diagnosisInactive.Entry.Coded, entryElement, codePath)

	return diagnosisInactive
}

func ExtractSeverity(diagnosis *models.Condition, entryElement xml.Node, severityXPath *xpath.Expression) {
	severityElement := FirstElement(severityXPath, entryElement)
	if severityElement != nil {
		diagnosis.Severity.AddCodeIfPresent(severityElement)
	}
}

func ExtractLaterality(diagnosis *models.Condition, entryElement xml.Node, lateralityXPath *xpath.Expression) {
	lateralityElement := FirstElement(lateralityXPath, entryElement)
	if lateralityElement != nil {
		diagnosis.Laterality.AddCodeIfPresent(lateralityElement)
		diagnosis.AnatomicalLocation.AddCodeIfPresent(lateralityElement)
	}
}
