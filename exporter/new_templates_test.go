package exporter

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/davecgh/go-spew/spew"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var templateTests = []struct {
	templateName string
	qrdaVersion  string
	data         interface{}
	vsMap        models.ValueSetMap
	expected     string
}{
	{
		"_id.xml",
		"r3",
		models.CDAIdentifier{
			Root:      "rootId",
			Extension: "extensionId",
		},
		nil,
		"<id root=\"rootId\" extension=\"extensionId\" />\n",
	},
	{
		"_id.xml",
		"r3_1",
		models.CDAIdentifier{
			Root:      "rootId",
			Extension: "extensionId",
		},
		nil,
		"<id root=\"rootId\" extension=\"extensionId\" />\n",
	},
	{
		"_reason.xml",
		"r3",
		models.EntryInfo{
			EntrySection: &models.Encounter{
				Entry: models.Entry{
					StartTime: 1480704345,
					NegationReason: models.CodedConcept{
						Code:       "reason_code",
						CodeSystem: "reason_code_system",
					},
				},
			},
			MapDataCriteria: models.Mdc{
				FieldOids: map[string][]string{
					"REASON": {
						"reason_code",
						"reason_code_system",
					},
				},
			},
		},
		models.ValueSetMap{
			"reason_code_system": []models.CodeSet{
				models.CodeSet{
					Set: "reason_code_system",
					Values: []models.Concept{
						models.Concept{
							Code:       "reason_code",
							CodeSystem: "reason_code_system",
						},
					},
				},
			},
		},
		"<id root=\"rootId\" extension=\"extensionId\" />\n",
	},
}

func TestTemplates(t *testing.T) {
	dmp := diffmatchpatch.New()
	for _, tt := range templateTests {
		temp := template.New("cat1")
		temp.Funcs(exporterFuncMap(temp, tt.vsMap))
		asset, err := Asset("templates/cat1/" + tt.qrdaVersion + "/" + tt.templateName)
		util.CheckErr(err)
		template.Must(temp.New(tt.templateName).Parse(string(asset)))
		var buf bytes.Buffer
		if err := temp.ExecuteTemplate(&buf, tt.templateName, tt.data); err != nil {
			util.CheckErr(err)
		}
		if buf.String() != tt.expected {
			t.Errorf("Template %s does not match", tt.templateName)
			diffs := dmp.DiffMain(buf.String(), tt.expected, false)
			t.Errorf(dmp.DiffPrettyText(diffs))
			spew.Dump(buf.String())
		}
	}
}
