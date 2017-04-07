package doc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/projectcypress/cdatools/exporter/doc"
)

func TestMeasureDataPrint(t *testing.T) {
	var tests = []struct {
		n        doc.MeasureData
		expected string
	}{
		{
			doc.MeasureData{
				Population: doc.Population{
					ID:              "test id",
					Type:            "MSRPOPL",
					Value:           1,
					Stratifications: []doc.Stratification{doc.Stratification{ID: "strat id", Value: 2}},
				},
				Observ: doc.Population{
					ID:              "observ test id",
					Type:            "observ test type",
					Value:           3,
					Stratifications: []doc.Stratification{doc.Stratification{ID: "strat id", Value: 4}},
				},
			},
			fmt.Sprintf(measureDataCat3TestTemplate, "MSRPOPL", "test id", "MSRPOPL", 1, "MSRPOPL", "test id", "strat id", 2, 4, "strat id", "strat id", 3, "observ test id", "test id"),
		},
	}

	for _, tt := range tests {
		actual := tt.n.Print()
		if strings.TrimSpace(actual) != strings.TrimSpace(tt.expected) {
			t.Errorf("MeasureData.Print(): expected===\n%s\n===actual===\n%s\n===", tt.expected, actual)
		}
	}
}

var measureDataCat3TestTemplate = `<!--	MEASURE DATA REPORTING FOR	%s %s	-->
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
</observation>`
