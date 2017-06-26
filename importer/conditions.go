package importer

import (
	"strconv"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
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

	//extract ordinality
	var ordinalityXPath = xpath.Compile("cda:priorityCode")
	ExtractOrdinality(&condition, entryElement, ordinalityXPath)

	//extract severity
	var severityXPath = xpath.Compile("cda:entryRelationship/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.22.4.8']/cda:value")
	ExtractSeverity(&condition, entryElement, severityXPath)

	//extract laterality
	var lateralityXPath = xpath.Compile("cda:value/cda:qualifier/cda:value")
	ExtractLaterality(&condition, entryElement, lateralityXPath)

	//extract anatomical Location
	var anatomicalLocationXPath = xpath.Compile("cda:targetSiteCode")
	ExtractAnatomicalLocation(&condition, entryElement, anatomicalLocationXPath)

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

func ExtractOrdinality(diagnosis *models.Condition, entryElement xml.Node, ordinalityXPath *xpath.Expression) {
	ordinalityElement := FirstElement(ordinalityXPath, entryElement)
	if ordinalityElement != nil {
		diagnosis.Ordinality.AddCodeIfPresent(ordinalityElement)
	}
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

func ExtractAnatomicalLocation(diagnosis *models.Condition, entryElement xml.Node, anatomicalLocationXPath *xpath.Expression) {
	anatomicalLocationElement := FirstElement(anatomicalLocationXPath, entryElement)
	if anatomicalLocationElement != nil {
		diagnosis.AnatomicalLocation.AddCodeIfPresent(anatomicalLocationElement)
	}
}