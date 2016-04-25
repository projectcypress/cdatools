package importer

import (
	"strconv"

	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
)

func MedicationActiveExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	medicationActive := extractMedication(entry, entryElement)

	codeXPath := xpath.Compile("./cda:consumable/cda:manufacturedProduct/cda:manufacturedMaterial/cda:code")
	ExtractCodes(&medicationActive.Entry.Coded, entryElement, codeXPath)

	return medicationActive
}

func MedicationDispensedExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	medicationDispensed := extractMedication(entry, entryElement)

	codeXPath := xpath.Compile("./cda:product/cda:manufacturedProduct/cda:manufacturedMaterial/cda:code")
	ExtractCodes(&medicationDispensed.Entry.Coded, entryElement, codeXPath)

	return medicationDispensed
}

func MedicationExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	medication := extractMedication(entry, entryElement)

	return medication
}

//private
func extractMedication(entry *models.Entry, entryElement xml.Node) models.Medication {
	var medication models.Medication
	medication.Entry = *entry

	extractAdministrationTiming(&medication, entryElement)

	routeXPath := xpath.Compile("./cda:routeCode")
	medication.Route.Codes = map[string][]string{}
	ExtractCodes(&medication.Route, entryElement, routeXPath)

	doseXPath := xpath.Compile("./cda:doseQuantity")
	ExtractScalar(&medication.Dose, entryElement, doseXPath)

	aaXPath := xpath.Compile("./cda:approachSiteCode")
	medication.AnatomicalApproach.Codes = map[string][]string{}
	ExtractCodes(&medication.AnatomicalApproach, entryElement, aaXPath)

	extractDoseRestriction(&medication, entryElement)

	pfXPath := xpath.Compile("./cda:administrationUnitCode")
	medication.ProductForm.Codes = map[string][]string{}
	ExtractCodes(&medication.ProductForm, entryElement, pfXPath)

	dmXPath := xpath.Compile("./cda:code")
	medication.DeliveryMethod.Codes = map[string][]string{}
	ExtractCodes(&medication.DeliveryMethod, entryElement, dmXPath)

	tomXPath := xpath.Compile("./cda:entryRelationship[@typeCode='SUBJ']/cda:observation[cda:templateId/@root='2.16.840.1.113883.3.88.11.83.8.1']/cda:code")
	medication.TypeOfMedication.Codes = map[string][]string{}
	ExtractCodes(&medication.TypeOfMedication, entryElement, tomXPath)

	indXPath := xpath.Compile("./cda:entryRelationship[@typeCode='RSON']/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.1.28']/cda:code")
	medication.Indication.Codes = map[string][]string{}
	ExtractCodes(&medication.Indication, entryElement, indXPath)

	vehicleXPath := xpath.Compile("cda:participant/cda:participantRole[cda:code/@code='412307009' and cda:code/@codeSystem='2.16.840.1.113883.6.96']/cda:playingEntity/cda:code")
	medication.Vehicle.Codes = map[string][]string{}
	ExtractCodes(&medication.Vehicle, entryElement, vehicleXPath)

	medication.OrderInformation = []models.OrderInformation{}
	extractOrderInformation(&medication, entryElement)

	medication.FulfillmentHistory = []models.FulfillmentHistory{}
	extractFulfillmentHistory(&medication, entryElement)

	ExtractReasonOrNegation(&medication.Entry, entryElement)
	return medication
}

func extractAdministrationTiming(medication *models.Medication, entryElement xml.Node) {
	adminTimingXPath := xpath.Compile("./cda:effectiveTime[2]")
	adminTimingElements, err := entryElement.Search(adminTimingXPath)
	util.CheckErr(err)

	for _, ATElement := range adminTimingElements {
		institutionSpecifiedAttr := ATElement.Attribute("institutionSpecified")
		if institutionSpecifiedAttr != nil {
			medication.AdministrationTiming.InstitutionSpecified, err = strconv.ParseBool(institutionSpecifiedAttr.String())
			util.CheckErr(err)
		}
		periodXPath := xpath.Compile("./cda:period")
		ExtractScalar(&medication.AdministrationTiming.Period, ATElement, periodXPath)
	}
}

func extractDoseRestriction(medication *models.Medication, entryElement xml.Node) {
	drXPath := xpath.Compile("./cda:maxDoseQuantity")
	drElements, err := entryElement.Search(drXPath)
	util.CheckErr(err)

	for _, drElement := range drElements {
		numXPath := xpath.Compile("./cda:numerator")
		denomXPath := xpath.Compile("./cda:denominator")
		ExtractScalar(&medication.DoseRestriction.Numerator, drElement, numXPath)
		ExtractScalar(&medication.DoseRestriction.Denominator, drElement, denomXPath)
	}
}

func extractOrderInformation(medication *models.Medication, entryElement xml.Node) {
	oiXPath := xpath.Compile("./cda:entryRelationship[@typeCode='REFR']/cda:supply[@moodCode='INT']")
	oiElements, err := entryElement.Search(oiXPath)
	util.CheckErr(err)

	for _, oiElement := range oiElements {
		oi := models.OrderInformation{}
		//provider information not captured, unsure if necessary
		oi.OrderNumber = FirstElementContent(xpath.Compile("./cda:id/@root"), oiElement)
		oi.Fills, err = strconv.ParseInt(FirstElementContent(xpath.Compile("./cda:repeatNumber/@value"), oiElement), 10, 64)
		oi.OrderDate = GetTimestamp(xpath.Compile("./cda:effectiveTime/cda:low/@value"), oiElement)
		util.CheckErr(err)

		qoXPath := xpath.Compile("./cda:quantity")
		ExtractScalar(&oi.QuantityOrdered, oiElement, qoXPath)

		medication.OrderInformation = append(medication.OrderInformation, oi)
	}
}

func extractFulfillmentHistory(medication *models.Medication, entryElement xml.Node) {
	fhXPath := xpath.Compile("./cda:entryRelationship/cda:supply[@moodCode='EVN']")
	fhElements, err := entryElement.Search(fhXPath)
	util.CheckErr(err)

	for _, fhElement := range fhElements {
		fh := models.FulfillmentHistory{}
		fh.PrescriptionNumber = FirstElementContent(xpath.Compile("./cda:id/@root"), fhElement)
		fh.DispenseDate = GetTimestamp(xpath.Compile("./cda:effectiveTime/@value"), fhElement)
		ExtractScalar(&fh.QuantityDispensed, fhElement, xpath.Compile("./cda:quantity"))
		fillNumber := FirstElementContent(xpath.Compile("./cda:entryRelationship[@typeCode='COMP']/cda:sequenceNumber/@value"), fhElement)
		if fillNumber != "" {
			fh.FillNumber, err = strconv.ParseInt(fillNumber, 10, 64)
			util.CheckErr(err)
		}

		medication.FulfillmentHistory = append(medication.FulfillmentHistory, fh)
	}
}
