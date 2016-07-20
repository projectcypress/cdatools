package importer

import (
	"strconv"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
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

	codeXPath := xpath.Compile("./cda:consumable/cda:manufacturedProduct/cda:manufacturedMaterial/cda:code")
	ExtractCodes(&medication.Entry.Coded, entryElement, codeXPath)

	routeXPath := xpath.Compile("./cda:routeCode")
	medication.Route = &models.CodedConcept{}
	medication.Route.AddCodeIfPresent(FirstElement(routeXPath, entryElement))

	doseXPath := xpath.Compile("./cda:doseQuantity")
	ExtractScalar(&medication.Dose, entryElement, doseXPath)

	aaXPath := xpath.Compile("./cda:approachSiteCode")
	medication.AnatomicalApproach = &models.CodedConcept{}
	medication.AnatomicalApproach.AddCodeIfPresent(FirstElement(aaXPath, entryElement))

	extractDoseRestriction(&medication, entryElement)

	pfXPath := xpath.Compile("./cda:administrationUnitCode")
	medication.ProductForm = &models.CodedConcept{}
	medication.ProductForm.AddCodeIfPresent(FirstElement(pfXPath, entryElement))

	dmXPath := xpath.Compile("./cda:code")
	medication.DeliveryMethod = &models.CodedConcept{}
	medication.DeliveryMethod.AddCodeIfPresent(FirstElement(dmXPath, entryElement))

	tomXPath := xpath.Compile("./cda:entryRelationship[@typeCode='SUBJ']/cda:observation[cda:templateId/@root='2.16.840.1.113883.3.88.11.83.8.1']/cda:code")
	medication.TypeOfMedication = &models.CodedConcept{}
	medication.TypeOfMedication.AddCodeIfPresent(FirstElement(tomXPath, entryElement))

	indXPath := xpath.Compile("./cda:entryRelationship[@typeCode='RSON']/cda:observation[cda:templateId/@root='2.16.840.1.113883.10.20.1.28']/cda:code")
	medication.Indication = &models.CodedConcept{}
	medication.Indication.AddCodeIfPresent(FirstElement(indXPath, entryElement))

	vehicleXPath := xpath.Compile("cda:participant/cda:participantRole[cda:code/@code='412307009' and cda:code/@codeSystem='2.16.840.1.113883.6.96']/cda:playingEntity/cda:code")
	medication.Vehicle = &models.CodedConcept{}
	medication.Vehicle.AddCodeIfPresent(FirstElement(vehicleXPath, entryElement))

	medication.OrderInformation = []models.OrderInformation{}
	extractSupplyInformation(&medication, entryElement)
	extractOrderInformation(&medication, entryElement)

	medication.FulfillmentHistory = []models.FulfillmentHistory{}
	extractFulfillmentHistory(&medication, entryElement)

	if entryElement.Parent().Name() == "entryRelationship" {
		ExtractReasonOrNegation(&medication.Entry, entryElement.Parent().Parent())
	} else {
		ExtractReasonOrNegation(&medication.Entry, entryElement)
	}
	return medication
}

func extractAdministrationTiming(medication *models.Medication, entryElement xml.Node) {
	adminTimingXPath := xpath.Compile("./cda:effectiveTime[2]")
	adminTimingElement := FirstElement(adminTimingXPath, entryElement)

	if adminTimingElement != nil {
		institutionSpecifiedAttr := adminTimingElement.Attribute("institutionSpecified")
		if institutionSpecifiedAttr != nil {
			institutionSpecified, err := strconv.ParseBool(institutionSpecifiedAttr.String())
			util.CheckErr(err)
			medication.AdministrationTiming.InstitutionSpecified = institutionSpecified
		}
		periodXPath := xpath.Compile("./cda:period")
		ExtractScalar(&medication.AdministrationTiming.Period, adminTimingElement, periodXPath)
	}
}

func extractDoseRestriction(medication *models.Medication, entryElement xml.Node) {
	drXPath := xpath.Compile("./cda:maxDoseQuantity")
	drElement := FirstElement(drXPath, entryElement)

	if drElement != nil {
		numXPath := xpath.Compile("./cda:numerator")
		denomXPath := xpath.Compile("./cda:denominator")
		ExtractScalar(&medication.DoseRestriction.Numerator, drElement, numXPath)
		ExtractScalar(&medication.DoseRestriction.Denominator, drElement, denomXPath)
	}
}

func extractOrderInformation(medication *models.Medication, entryElement xml.Node) {
	oiXPath := xpath.Compile("./cda:repeatNumber")
	oiElement := FirstElement(oiXPath, entryElement)

	if oiElement != nil {
		fills, err := strconv.ParseInt(FirstElementContent(xpath.Compile("./@value"), oiElement), 10, 64)
		util.CheckErr(err)
		medication.AllowedAdministrations = &fills
	}
}

func extractSupplyInformation(medication *models.Medication, entryElement xml.Node) {
	oiXPath := xpath.Compile("./cda:entryRelationship[@typeCode='REFR']/cda:supply[@moodCode='INT']")
	oiElement := FirstElement(oiXPath, entryElement)

	if oiElement != nil {
		oi := models.OrderInformation{}
		//provider information not captured, unsure if necessary
		oi.OrderNumber = FirstElementContent(xpath.Compile("./cda:id/@root"), oiElement)
		fills, err := strconv.ParseInt(FirstElementContent(xpath.Compile("./cda:repeatNumber/@value"), oiElement), 10, 64)
		util.CheckErr(err)
		oi.Fills = fills
		oi.OrderDate = GetTimestamp(xpath.Compile("./cda:effectiveTime/cda:low/@value"), oiElement)

		qoXPath := xpath.Compile("./cda:quantity")
		ExtractScalar(&oi.QuantityOrdered, oiElement, qoXPath)

		medication.AllowedAdministrations = &fills
		medication.OrderInformation = append(medication.OrderInformation, oi)
	}
}

func extractFulfillmentHistory(medication *models.Medication, entryElement xml.Node) {
	fhXPath := xpath.Compile("./cda:entryRelationship/cda:supply[@moodCode='EVN']")
	fulfillmentElements, err := entryElement.Search(fhXPath)
	util.CheckErr(err)
	if len(fulfillmentElements) > 0 {
		for _, fhElement := range fulfillmentElements {
			if fhElement != nil {
				fh := models.FulfillmentHistory{}
				fh.PrescriptionNumber = FirstElementContent(xpath.Compile("./cda:id/@root"), fhElement)
				fh.DispenseDate = GetTimestamp(xpath.Compile("./cda:effectiveTime/@value"), fhElement)
				ExtractScalar(&fh.QuantityDispensed, fhElement, xpath.Compile("./cda:quantity"))
				fillNumber := FirstElementContent(xpath.Compile("./cda:entryRelationship[@typeCode='COMP']/cda:sequenceNumber/@value"), fhElement)
				if fillNumber != "" {
					fillnumber, err := strconv.ParseInt(fillNumber, 10, 64)
					util.CheckErr(err)
					fh.FillNumber = fillnumber
				}

				medication.FulfillmentHistory = append(medication.FulfillmentHistory, fh)
			}
		}
	}
}
