package cat3_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/projectcypress/cdatools/exporter/cat3"
)

func TestMeasureDataTemplate(t *testing.T) {
	var tests = []struct {
		n        cat3.MeasureData
		expected string
	}{
		{
			cat3.MeasureData{
				Population: cat3.Population{
					ID:              "test id",
					Type:            "MSRPOPL",
					Value:           1,
					Stratifications: []cat3.Stratification{cat3.Stratification{ID: "strat id", Value: 2}},
				},
				Observ: cat3.Population{
					ID:              "observ test id",
					Type:            "observ test type",
					Value:           3,
					Stratifications: []cat3.Stratification{cat3.Stratification{ID: "strat id", Value: 4}},
				},
			},
			fmt.Sprintf(measureDataTestTemplate, "MSRPOPL", "test id", "MSRPOPL", 1, "MSRPOPL", "test id", "strat id", 2, 4, "strat id", "strat id", 3, "observ test id", "test id"),
		},
	}

	for _, tt := range tests {
		actual := cat3.Print(tt.n.Template(), tt.n)
		if strings.TrimSpace(actual) != strings.TrimSpace(tt.expected) {
			t.Errorf("MeasureData.Template(): expected===\n%s\n===actual===\n%s\n===", tt.expected, actual)
		}
	}
}

var measureDataTestTemplate = `<!--	MEASURE DATA REPORTING FOR	%s %s	-->
<component>
<observation classCode="OBS" moodCode="EVN">
	<!--	Measure Data template	-->
	<templateId root="2.16.840.1.113883.10.20.27.3.5" extension="2016-09-01"/>
	<code code="ASSERTION"
				codeSystem="2.16.840.1.113883.5.4"
				displayName="Assertion"
				codeSystemName="ActCode"/>
	<statusCode code="completed"/>
	<value xsi:type="CD" code="%s"
				 codeSystem="2.16.840.1.113883.5.4"
				 codeSystemName="ActCode"/>
	<!-- Aggregate Count -->
	<entryRelationship typeCode="SUBJ" inversionInd="true">
		<observation classCode="OBS" moodCode="EVN">
			<templateId root="2.16.840.1.113883.10.20.27.3.3"/>
			<code code="MSRAGG"
				displayName="rate aggregation"
				codeSystem="2.16.840.1.113883.5.4"
				codeSystemName="ActCode"/>
			<value xsi:type="INT" value="%d"/>
			<methodCode code="COUNT"
				displayName="Count"
				codeSystem="2.16.840.1.113883.5.84"
				codeSystemName="ObservationMethod"/>
		</observation>
	</entryRelationship>
	
	<!--	Stratification Reporting Template for %s %s	Stratification %s	-->

	<entryRelationship typeCode="COMP">
		<observation classCode="OBS" moodCode="EVN">
			<templateId root="2.16.840.1.113883.10.20.27.3.4"/>
			<code code="ASSERTION"
						codeSystem="2.16.840.1.113883.5.4"
						displayName="Assertion"
						codeSystemName="ActCode"/>
			<statusCode code="completed"/>
			<value xsi:type="CD" nullFlavor="OTH">
			 <originalText>Stratum</originalText>
			</value>
			<entryRelationship typeCode="SUBJ" inversionInd="true">
				<observation classCode="OBS" moodCode="EVN">
					<templateId root="2.16.840.1.113883.10.20.27.3.3"/>
					<code code="MSRAGG"
								displayName="rate aggregation"
								codeSystem="2.16.840.1.113883.5.4"
								codeSystemName="ActCode"/>
					<value xsi:type="INT" value="%d"/>
					<methodCode code="COUNT"
											displayName="Count"
											codeSystem="2.16.840.1.113883.5.84"
											codeSystemName="ObservationMethod"/>
				</observation>
			</entryRelationship>
			<entryRelationship typeCode="COMP">
				<observation classCode="OBS" moodCode="EVN">
					<templateId root="2.16.840.1.113883.10.20.27.3.2"/>
					<code nullFlavor="OTH">
						<originalText>Time Difference</originalText>
					</code>
					<statusCode code="completed"/>
					<value xsi:type="PQ" value="%d" unit="min"/>
					<methodCode code="MEDIAN"
											displayName="Median"
											codeSystem="2.16.840.1.113883.5.84"
											codeSystemName="ObservationMethod"/>
					<reference typeCode="REFR">
						<!-- reference to the relevant measure observation in the eMeasure -->
						<externalObservation classCode="OBS" moodCode="EVN">
							<id root="%s"/>
						</externalObservation>
					</reference>
				</observation>
			</entryRelationship>
			<reference typeCode="REFR">
				<externalObservation classCode="OBS" moodCode="EVN">
					<id root="%s"/>
				</externalObservation>
			</reference>
		</observation>
	</entryRelationship>
	<entryRelationship typeCode="COMP">
		<observation classCode="OBS" moodCode="EVN">
			<templateId root="2.16.840.1.113883.10.20.27.3.2"/>
			<code nullFlavor="OTH">
				<originalText>Time Difference</originalText>
			</code>
			<statusCode code="completed"/>
			<value xsi:type="PQ" value="%d" unit="min"/>
			<methodCode code="MEDIAN"
									displayName="Median"
									codeSystem="2.16.840.1.113883.5.84"
									codeSystemName="ObservationMethod"/>
			<reference typeCode="REFR">
				<!-- reference to the relevant measure observation in the eMeasure -->
				<externalObservation classCode="OBS" moodCode="EVN">
					<id root="%s"/>
				</externalObservation>
			</reference>
		</observation>
	</entryRelationship>
	<reference typeCode="REFR">
		 <externalObservation classCode="OBS" moodCode="EVN">
				<id root="%s"/>
		 </externalObservation>
	</reference>
</observation></component>`
