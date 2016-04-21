package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
)

func MedicationActiveExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	medicationActive := extractMedication(entry, entryElement)
}

//private
func extractMedication(entry *models.Entry, entryElement xml.Node) *models.Medication {
	medication := models.Medication{}
	medication.Entry = *entry

	extractAdministrationTiming(medication, entryElement)

  routeXPath := xpath.Compile("./cda:routeCode")
  ExtractCodes(medication.Route, entryElement, routeXPath)

  doseXPath := xpath.Compile("./cda:doseQuantity")
  ExtractScalar(medication.Dose, entryElement, doseXPath)

  aaXPath := xpath.Compilse("./cda:approachSiteCode")
  ExtractCodes(medication.AnatomicalApproach, entryElement, aaXPath)

  extractDoseRestriction(medication, entryElement)

  pfXPath := xpath.Compile("./cda:administrationUnitCode")
  ExtractCodes(medication.ProductForm, entryElement, pfXPath)

  dmXPath := xpath.Compile("./cda:code")
  ExtractCodes(medication.DeliveryMethod, entryElement, dmXPath)

  tomXPath := xpath.Compile("./cda:entryRelationship[@typeCode='SUBJ']/cda:observation[cda:templateId/@root='2.16.840.1.113883.3.88.11.83.8.1']/cda:code")
  ExtractCodes(medication.TypeOfMedication, entryElement, tomXPath)

  indXPath := xpath.Compile("./cda:entryRelationship[@typeCode='RSON']/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.1.28']/cda:code")
  ExtractCodes(medication.Indication, entryElement, indXPath)

  vehicleXPath := xpath.Compile("cda:participant/cda:participantRole[cda:code/@code='412307009' and cda:code/@codeSystem='2.16.840.1.113883.6.96']/cda:playingEntity/cda:code")
  ExtractCodes(medication.Vehicle, entryElement, vehicleXPath)

  extractOrderInformation(medication, entryElement)
  extractFulfillmentHistory(medicaiton, entryElement)
  extractReasonOrNegation(medication, entryElement)
}

func extractAdministrationTiming(medication *models.Medication, entryElement xml.Node) {
	adminTimingXPath := xpath.Compile("./cda:effectiveTime[2]")
	adminTimingElements, err := entryElement.Search(adminTimingXPath)
	util.CheckErr(err)

	for _, ATElement := range adminTimingElements {
		institutionSpecifiedAttr := ATElement.Attribute("institutionSpecified")
		if institutionSpecifiedAttr != nil {
      medication.AdministrationTiming.InstitutionSpecified, err := strconv.ParseBool(institutionSpecifiedAttr.String())
      util.CheckErr(err)
		}
    periodXPath := xpath.Compile("./cda:period")
    ExtractScalar(medication.AdministrationTiming.Period, periodXPath)
	}
}

func extractDoseRestriction(medication *models.Medication, entryElement xml.Node) {
  drXPath := xpath.Compile("./cda:maxDoseQuantity")
  drElements, err := entryElement.Search(drXPath)

  for _, drElement := range drElements {
    numXPath := xpath.Compile("./cda:numerator")
    denomXPath := xpath.Compile("./cda:denominator")
    ExtractScalar(medication.DoseRestriction.Numerator, drElement, numXPath)
    ExtractScalar(medication.DoseRestriction.Denominator, drElement, denomXPath)
  }
}
