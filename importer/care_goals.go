package importer

import (
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

func CareGoalExtractor(entry *models.Entry, entryElement xml.Node) interface{} {
	careGoal := models.CareGoal{}
	careGoal.Entry = *entry

	// extract codes
	codeXPath := xpath.Compile("cda:code")
	ExtractCodes(&careGoal.Entry.Coded, entryElement, codeXPath)

	return careGoal
}
