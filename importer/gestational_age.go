package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

func GestationalAgeExtractor(entry *models.Entry, entryElement xml.Node) interface{} {

	codeXPath := xpath.Compile("./cda:code")
	ExtractCodes(&entry.Coded, entryElement, codeXPath)

	valueXPath := xpath.Compile("./cda:value")
	entry.Values = make([]models.ResultValue, 0)
	ExtractValues(entry, entryElement, valueXPath)

	switch entry.Values[0].Scalar {
	case 39:
		entry.Codes["SNOMED-CT"] = []string{"80487005"}
	case 38:
		entry.Codes["SNOMED-CT"] = []string{"13798002"}
	case 37:
		entry.Codes["SNOMED-CT"] = []string{"43697006"}
	case 36:
		entry.Codes["SNOMED-CT"] = []string{"931004"}
	}

	return *entry
}
