package exporter

import (
	"bytes"
	"strings"
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
		"<id root=\"rootId\" extension=\"extensionId\" />",
	},
	{
		"_id.xml",
		"r3_1",
		models.CDAIdentifier{
			Root:      "rootId",
			Extension: "extensionId",
		},
		nil,
		"<id root=\"rootId\" extension=\"extensionId\" />",
	},
	{
		"_reason.xml",
		"r3_1",
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
		"<entryRelationship typeCode=\"RSON\">\n      <observation classCode=\"OBS\" moodCode=\"EVN\">\n        <templateId root=\"2.16.840.1.113883.10.20.24.3.88\" extension=\"2014-12-01\"/>\n        <id root=\"1.3.6.1.4.1.115\" extension=\"7EBA4DE48292A48FE6A35FAD579E74C5\"/>\n          <code code=\"77301-0\" codeSystem=\"2.16.840.1.113883.6.1\" displayName=\"reason\" codeSystemName=\"LOINC\"/>\n        <statusCode code=\"completed\"/>\n          <effectiveTime>\n            <start_time_value>1480704345</start_time_value>\n            <low value='201612021845+0000'/>\n          </effectiveTime>\n        <value xsi:type=\"CD\" code=\"reason_code\" codeSystem=\"reason_code_system\" sdtc:valueSet=\"reason_code_system\"/>\n      </observation>\n    </entryRelationship>",
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
		"<entryRelationship typeCode=\"RSON\">\n      <observation classCode=\"OBS\" moodCode=\"EVN\">\n        <templateId root=\"2.16.840.1.113883.10.20.24.3.88\" extension=\"2014-12-01\"/>\n        <id root=\"1.3.6.1.4.1.115\" extension=\"7EBA4DE48292A48FE6A35FAD579E74C5\"/>\n          <code code=\"77301-0\" codeSystem=\"2.16.840.1.113883.6.1\" displayName=\"reason\" codeSystemName=\"LOINC\"/>\n        <statusCode code=\"completed\"/>\n          <effectiveTime>\n            <start_time_value>1480704345</start_time_value>\n            <low value='201612021845+0000'/>\n          </effectiveTime>\n        <value xsi:type=\"CD\" code=\"reason_code\" codeSystem=\"reason_code_system\" sdtc:valueSet=\"reason_code_system\"/>\n      </observation>\n    </entryRelationship>",
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
		if cleanString(buf.String()) != cleanString(tt.expected) {
			t.Errorf("Template %s does not match", tt.templateName)
			diffs := dmp.DiffMain(buf.String(), tt.expected, false)
			t.Errorf(dmp.DiffPrettyText(diffs))
			spew.Dump(cleanString(buf.String()))
		}
	}
}

func cleanString(s string) string {
	return strings.Trim(s, "\n\t ")
}
