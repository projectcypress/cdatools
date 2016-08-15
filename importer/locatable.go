package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
)

func ImportAddress(addressElement xml.Node) models.Address {
	var address = models.Address{}
	address.Use = addressElement.Attr("use")
	streetNameElements, err := addressElement.Search(xpath.Compile("cda:streetAddressLine"))
	util.CheckErr(err)
	address.Street = make([]string, len(streetNameElements))
	for i, streetNameElement := range streetNameElements {
		address.Street[i] = streetNameElement.Content()
	}
	address.City = FirstElementContent(xpath.Compile("cda:city"), addressElement)
	address.State = FirstElementContent(xpath.Compile("cda:state"), addressElement)
	address.Zip = FirstElementContent(xpath.Compile("cda:postalCode"), addressElement)
	address.Country = FirstElementContent(xpath.Compile("cda:country"), addressElement)

	return address
}

func ImportTelecom(telecomEntry xml.Node) models.Telecom {
	var telecom = models.Telecom{}
	value := telecomEntry.Attr("value")
	telecom.Value = value
	use := telecomEntry.Attr("use")
	telecom.Use = use
	return telecom
}