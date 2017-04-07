package doc_test

import (
	"fmt"
	"testing"

	"github.com/projectcypress/cdatools/exporter/doc"
)

func TestPopulationGroupPrint(t *testing.T) {
	var tests = []struct {
		n        doc.PopulationGroup
		expected string
	}{
		{
			doc.PopulationGroup{
				Populations: []doc.Population{
					doc.Population{
						ID:    "test num ID",
						Type:  "NUMER",
						Value: 2,
					},
					doc.Population{
						ID:    "test ID",
						Type:  "DENOM",
						Value: 1,
					},
					doc.Population{
						ID:    "test ID",
						Type:  "DENEXCEP",
						Value: 0,
					},
					doc.Population{
						ID:    "test ID",
						Type:  "DENEX",
						Value: 0,
					},
				},
			},
			fmt.Sprintf(populationGroupCat3TestTemplate, 2, "test num ID"),
		},
	}

	for _, tt := range tests {
		actual := tt.n.Print()
		if actual != tt.expected {
			t.Errorf("PopulationGroup.Print(): expected =%s=actual=%s=", tt.expected, actual)
		}
	}
}

var populationGroupCat3TestTemplate = `<observation classCode="OBS" moodCode="EVN">
	<templateId root="2.16.840.1.113883.10.20.27.3.14" extension="2016-09-01"/>
	<templateId root="2.16.840.1.113883.10.20.27.3.30" extension="2016-09-01"/>
	<code code="72510-1" codeSystem="2.16.840.1.113883.6.1"
		displayName="Performance Rate"
		codeSystemName="2.16.840.1.113883.6.1"/>
	<statusCode code="completed"/>
	<value xsi:type="REAL" value="%d"/>
	<reference typeCode="REFR">
		<externalObservation classCode="OBS" moodCode="EVN">
			<id root="%s"/>
			<code code="NUMER" displayName="Numerator" codeSystem="2.16.840.1.113883.5.4" codeSystemName="ObservationValue"/>
		</externalObservation>
	</reference>
</observation>`
