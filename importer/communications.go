package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
)

func CommunicationExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	communication := models.Communication{}
	communication.Entry = *entry

	codeXPath := xpath.Compile("./cda:code")
	ExtractCodes(&communication.Entry.Coded, entryElement, codeXPath)

	ExtractReasonOrNegation(&communication.Entry, entryElement)

	communication.Direction = findCommunicationDirection(entryElement)

	extractReferences(&communication.Entry, entryElement)

	return communication
}

func findCommunicationDirection(entryElement xml.Node) string {
	node, err := entryElement.Search("./cda:templateId")
	util.CheckErr(err)
	switch node[0].Attr("root") {
	case "2.16.840.1.113883.10.20.24.3.3":
		return "communication_from_provider_to_patient"
	case "2.16.840.1.113883.10.20.24.3.2":
		return "communication_from_patient_to_provider"
	case "2.16.840.1.113883.10.20.24.3.4":
		return "communication_from_provider_to_provider"
	default:
		return ""
	}
}

func extractReferences(entry *models.Entry, entryElement xml.Node) {
	refs, err := entryElement.Search("./sdtc:inFulfillmentOf1")
	util.CheckErr(err)

	for _, ref := range refs {
		reference := models.Reference{}
		switch ref.Attr("typeCode") {
		case "FLFS":
			reference.Type = "fulfills"
		}

		ar, err := ref.Search("./sdtc:actReference")
		util.CheckErr(err)
		switch ar[0].Attr("classCode") {
		case "ACT":
			reference.ReferencedType = "Procedure"
		}
		idElem, err := ref.Search("./sdtc:actReference/sdtc:id")
		util.CheckErr(err)
		reference.ReferencedID = idElem[0].Attr("extension")
		entry.References = append(entry.References, reference)
	}
}
