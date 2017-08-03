package importer

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"

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
	if err != nil {
		panic(err.Error())
	}
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
	if err != nil {
		panic(err.Error())
	}

	for _, ref := range refs {
		reference := models.Reference{}
		switch ref.Attr("typeCode") {
		case "FLFS":
			reference.Type = "fulfills"
		}

		ar, err := ref.Search("./sdtc:actReference")
		if err != nil {
			panic(err.Error())
		}
		switch ar[0].Attr("classCode") {
		case "ACT":
			reference.ReferencedType = "Procedure"
		}
		idElem, err := ref.Search("./sdtc:actReference/sdtc:id")
		if err != nil {
			panic(err.Error())
		}
		reference.ReferencedID = idElem[0].Attr("extension")
		reference.ExportedRef = idElem[0].Attr("extension")
		entry.References = append(entry.References, reference)
	}
}
