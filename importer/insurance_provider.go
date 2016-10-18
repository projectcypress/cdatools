package importer

import (
	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

func InsuranceProviderExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	insuranceProvider := models.InsuranceProvider{}
	insuranceProvider.Entry = *entry

	var codePath = xpath.Compile("cda:value")
	ExtractCodes(&insuranceProvider.Entry.Coded, entryElement, codePath)

	return insuranceProvider
}
