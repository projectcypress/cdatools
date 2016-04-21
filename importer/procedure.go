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

func ProcedurePerformedExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	procedurePerformed := models.Procedure{}
	procedurePerformed.Entry = *entry

	var codePath = xpath.Compile("cda:code")
	ExtractCodes(&procedurePerformed.Entry.Coded, entryElement, codePath)
	extractOrdinality(&procedurePerformed.Ordinality, entryElement)
	extractPerformer(&procedurePerformed.Performer, entryElement)
	extractAnatomicalTarget(&procedurePerformed.AnatomicalTarget, entryElement)
	ExtractReasonOrNegation(&procedurePerformed.Entry, entryElement)
	// extractProcedureScalar(&procedurePerformed.Entry, entryElement)
	scalarPath := xpath.Compile("cda:value")
	ExtractValues(&procedurePerformed.Entry, entryElement, scalarPath)
	extractIncisionTime(&procedurePerformed, entryElement)

	return procedurePerformed
}

// private

func extractOrdinality(ordinality *models.Coded, entryElement xml.Node) {
	ordinality.Codes = map[string][]string{} // create code map
	var codePath = xpath.Compile("cda:priorityCode")
	ExtractCodes(ordinality, entryElement, codePath)
}

func extractPerformer(performer *models.Performer, entryElement xml.Node) {

	// extract performer here

}

func extractAnatomicalTarget(anitomicalTarget *models.CodedConcept, entryElement xml.Node) {
	codePath := xpath.Compile("cda:targetSiteCode")
	ExtractCodedConcept(anitomicalTarget, entryElement, codePath)
}

func extractIncisionTime(procedure *models.Procedure, entryElement xml.Node) {
	incisionTimeXPath := xpath.Compile("cda:entryRelationship/cda:procedure[cda:templateId[@root='2.16.840.1.113883.10.20.24.3.89']]/cda:effectiveTime/@value")
	procedure.IncisionTime = GetTimestamp(incisionTimeXPath, entryElement)
}
