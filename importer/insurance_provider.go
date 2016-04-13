package importer

import (
  "github.com/moovweb/gokogiri/xml"
  "github.com/moovweb/gokogiri/xpath"
  "github.com/projectcypress/cdatools/models"
)

func InsuranceProviderExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
  insuranceProvider := models.InsuranceProvider{}
  insuranceProvider.Entry = *entry

  var codePath = xpath.Compile("cda:value")
  ExtractCodes(&insuranceProvider.Entry, entryElement, codePath)

  return insuranceProvider
}
