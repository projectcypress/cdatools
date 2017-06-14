package importer

import (
	"strings"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
)

func MedicalEquipmentExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	medicalEquipment := models.MedicalEquipment{}
	medicalEquipment.Entry = *entry

	codeXPath := xpath.Compile("./cda:participant/cda:participantRole/cda:playingDevice/cda:code")
	ExtractCodes(&medicalEquipment.Entry.Coded, entryElement, codeXPath)

	extractManufacturer(&medicalEquipment, entryElement)
	extractAnatomicalStructure(&medicalEquipment, entryElement)
	extractRemovalTime(&medicalEquipment, entryElement)

	if entryElement.Parent().Name() == "entryRelationship" {
		ExtractReasonOrNegation(&medicalEquipment.Entry, entryElement.Parent().Parent())
	} else {
		ExtractReasonOrNegation(&medicalEquipment.Entry, entryElement)
	}
	return medicalEquipment
}

func extractManufacturer(medicalEquipment *models.MedicalEquipment, entryElement xml.Node) {
	mfXPath := xpath.Compile("./cda:participant/cda:participantRole/cda:scopingEntity/cda:desc")
	manufacturerElements, err := entryElement.Search(mfXPath)
	util.CheckErr(err)

	for _, MFElement := range manufacturerElements {
		manufacturerAttr := MFElement.Attribute("inner_text")
		if manufacturerAttr != nil {
			medicalEquipment.Manufacturer = strings.TrimSpace(manufacturerAttr.String())
		}
	}
}

func extractAnatomicalStructure(medicalEquipment *models.MedicalEquipment, entryElement xml.Node) {
	asXPath := xpath.Compile("./cda:targetSiteCode")
	asElements, err := entryElement.Search(asXPath)
	util.CheckErr(err)

	for _, ASElement := range asElements {
		codeElement := ASElement.Attribute("code")
		codeSystemElement := ASElement.Attribute("codeSystem")
		if codeElement != nil && codeSystemElement != nil {
			medicalEquipment.AnatomicalStructure.Code = codeElement.String()
			medicalEquipment.AnatomicalStructure.CodeSystem = codeSystemElement.String()
			medicalEquipment.AnatomicalStructure.CodeSystemName = models.CodeSystemFor(codeSystemElement.String())
		}
	}
}

func extractRemovalTime(medicalEquipment *models.MedicalEquipment, entryElement xml.Node) {
	rtXPath := xpath.Compile("cda:effectiveTime/cda:high/@value")
	medicalEquipment.RemovalTime = GetTimestamp(rtXPath, entryElement)
}
