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
		address.Street[i] = streetNameElement.Attr("text")
	}
	city := FirstElement(xpath.Compile("cda:city"), addressElement).Attr("text")
	address.City = city
	state := FirstElement(xpath.Compile("cda:state"), addressElement).Attr("text")
	address.State = state
	zip := FirstElement(xpath.Compile("cda:postalCode"), addressElement).Attr("text")
	address.Zip = zip
	country := FirstElement(xpath.Compile("cda:country"), addressElement).Attr("text")
	address.Country = country

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
