package doc_test

import (
	"fmt"
	"testing"

	"github.com/projectcypress/cdatools/exporter/doc"
	"github.com/projectcypress/cdatools/models"
)

func TestMeasureSectionPrint(t *testing.T) {
	var tests = []struct {
		n        doc.MeasureSection
		expected string
	}{
		{
			doc.MeasureSection{
				Measure: models.Measure{
					ID:        "measure test id",
					HQMFID:    "measure hqmf id",
					Name:      "measure name",
					HQMFSetID: "measure hqmf set id",
				},
			},
			fmt.Sprintf(measureSectionCat3TestTemplate, "measure test id", "measure hqmf id", "measure name", "measure hqmf set id"),
		},
	}

	for _, tt := range tests {
		actual := tt.n.Print()
		if actual != tt.expected {
			t.Errorf("MeasureSection.Print(): expected =%s=actual=%s=", tt.expected, actual)
		}
	}
}

var measureSectionCat3TestTemplate = `<component><structuredBody><component><section>
	<!-- Implied template Measure Section templateId -->
	<templateId root="2.16.840.1.113883.10.20.24.2.2"/>
	<!-- In this case the query is using an eMeasure -->
	<!-- QRDA Category III Measure Section template -->
	<templateId extension="2016-09-01" root="2.16.840.1.113883.10.20.27.2.1"/>
	<code code="55186-1" codeSystem="2.16.840.1.113883.6.1"/>
	<title>Measure Section</title>
	<text/>
	<entry>
<organizer classCode="CLUSTER" moodCode="EVN">
	<!-- Implied template Measure Reference templateId -->
	<templateId root="2.16.840.1.113883.10.20.24.3.98"/>
	<!-- SHALL 1..* (one for each referenced measure) Measure Reference and Results template -->
	<templateId root="2.16.840.1.113883.10.20.27.3.1" extension="2016-09-01"/>
	<id extension="%s"/>
	<statusCode code="completed"/>
	<reference typeCode="REFR">
	<externalDocument classCode="DOC" moodCode="EVN">
		<!-- SHALL: required Id but not restricted to the eMeasure Document/Id-->
		<!-- QualityMeasureDocument/id This is the version specific identifier for eMeasure -->
		<id root="2.16.840.1.113883.4.738" extension="%s"/>

		<!-- SHOULD This is the title of the eMeasure -->
		<text>%s</text>
		<!-- SHOULD: setId is the eMeasure version neutral id	-->
		<setId root="%s"/>
		<!-- This is the sequential eMeasure Version number -->
		<versionNumber value="1"/>
	</externalDocument>
	</reference>
</organizer>
</entry></section>
	</component>
	</structuredBody>
	</component>`
