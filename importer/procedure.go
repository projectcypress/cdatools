package importer

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

func DiagnosticStudyOrderExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	diagnosticStudyOrder := models.Procedure{}
	diagnosticStudyOrder.Entry = *entry

	extractBaseProcedure(&diagnosticStudyOrder, entryElement)
	// extract order specific dates
	var orderTimeXPath = xpath.Compile("cda:author/cda:time/@value")
	diagnosticStudyOrder.StartTime = GetTimestamp(orderTimeXPath, entryElement)
	diagnosticStudyOrder.EndTime = GetTimestamp(orderTimeXPath, entryElement)

	return diagnosticStudyOrder
}

func ProcedureExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	procedurePerformed := models.Procedure{}
	procedurePerformed.Entry = *entry

	extractBaseProcedure(&procedurePerformed, entryElement)

	return procedurePerformed
}

func ProcedurePerformedExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	procedurePerformed := models.Procedure{}
	procedurePerformed.Entry = *entry

	extractBaseProcedure(&procedurePerformed, entryElement)
	extractIncisionTime(&procedurePerformed, entryElement)

	scalarPath := xpath.Compile("cda:entryRelationship/cda:observation[./cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.87']/cda:value")
	ExtractValues(&procedurePerformed.Entry, entryElement, scalarPath)
	return procedurePerformed
}

func ProcedureOrderExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	procedureOrder := models.Procedure{}
	procedureOrder.Entry = *entry

	extractBaseProcedure(&procedureOrder, entryElement)

	// set Status Code
	procedureOrder.StatusCode = map[string][]string{}
	procedureOrder.StatusCode["HL7 ActStatus"] = []string{"ordered"}

	// extract Time
	var orderTimeXPath = xpath.Compile("cda:author/cda:time/@value")
	procedureOrder.Time = GetTimestamp(orderTimeXPath, entryElement)

	return procedureOrder
}

// private

func extractBaseProcedure(procedure *models.Procedure, entryElement xml.Node) {
	var codePath = xpath.Compile("cda:code")
	ExtractCodes(&procedure.Entry.Coded, entryElement, codePath)
	var ordinalityXPathR3 = xpath.Compile("cda:priorityCode")
	var ordinalityXPathR31 = xpath.Compile("cda:entryRelationship/cda:observation[./cda:code[@code='260870009']]/cda:value")
	ExtractCodedConcept(&procedure.Ordinality.CodedConcept, entryElement, ordinalityXPathR3)
	ExtractCodedConcept(&procedure.Ordinality.CodedConcept, entryElement, ordinalityXPathR31)
	extractPerformer(&procedure.Performer, entryElement)
	extractAnatomicalTarget(&procedure.AnatomicalTarget, entryElement)
	ExtractReasonOrNegation(&procedure.Entry, entryElement)
	scalarPath := xpath.Compile("./cda:value | ./cda:entryRelationship/cda:observation[./cda:templateId/@root = '2.16.840.1.113883.10.20.24.3.87']/cda:value")
	ExtractValues(&procedure.Entry, entryElement, scalarPath)
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
