package cat3_test

import (
	"fmt"
	"testing"

	"github.com/projectcypress/cdatools/exporter/cat3"
)

func TestPopulationTemplate(t *testing.T) {
	var testSupData = cat3.SupplementalData{Sex: map[string]int{"sex code": 1}, Ethnicity: map[string]int{"ethnicity code": 2}, Race: map[string]int{"race code": 3}, Payer: map[string]int{"payer code": 4}}
	var tests = []struct {
		n        cat3.Population
		expected string
	}{
		{
			cat3.Population{ID: "test id", Type: "test type", Value: 2, Stratifications: []cat3.Stratification{cat3.Stratification{ID: "strat id", Value: 4}}, SupplementalData: testSupData},
			fmt.Sprintf(populationTestTemplate, "test type", "test id", "sex code", 1, "test type", "test id", "ethnicity code", 2, "test type", "test id", "race code", 3, "test type", "test id", "payer code", 4),
		},
	}

	for _, tt := range tests {
		actual := cat3.Print(tt.n.Template(), tt.n)
		if actual != tt.expected {
			t.Errorf("Population.Template(): expected \n%s\n, actual \n%s", tt.expected, actual)
		}
	}
}

var populationTestTemplate = `<!--	SEX Supplemental Data Reporting for %s	%s	-->

<!--	Supplemental Data Template	-->

<entryRelationship typeCode="COMP">
	<observation classCode="OBS" moodCode="EVN">
		<!--	Sex Supplemental Data	-->
		<templateId root="2.16.840.1.113883.10.20.27.3.6" extension="2016-09-01"/>
		<id nullFlavor="NA"/>
		<code code="76689-9" codeSystem="2.16.840.1.113883.6.1"/>
		<statusCode code="completed"/>
		<value xsi:type="CD" code="%s" codeSystem="2.16.840.1.113883.5.1"/>
		<entryRelationship typeCode="SUBJ" inversionInd="true">
			<!--	Aggregate Count template	-->
			<observation classCode="OBS" moodCode="EVN">
				<templateId root="2.16.840.1.113883.10.20.27.3.3"/>
				<code code="MSRAGG" displayName="rate aggregation" codeSystem="2.16.840.1.113883.5.4" codeSystemName="ActCode"/>
				<value xsi:type="INT" value="%d"/>
				<methodCode code="COUNT" displayName="Count" codeSystem="2.16.840.1.113883.5.84" codeSystemName="ObservationMethod"/>
			</observation>
		</entryRelationship>
	</observation>
</entryRelationship>

<!--	ETHNICITY Supplemental Data Reporting for %s	%s	-->

<!--	Supplemental Data Template	-->

<entryRelationship typeCode="COMP">
	<observation classCode="OBS" moodCode="EVN">
		<!--	Ethnicity Supplemental Data	-->
		<templateId root="2.16.840.1.113883.10.20.27.3.7" extension="2016-09-01"/>
		<id nullFlavor="NA"/>
		<code code="69490-1" codeSystem="2.16.840.1.113883.6.1"/>
		<statusCode code="completed"/>
		<value xsi:type="CD" code="%s" codeSystem="2.16.840.1.113883.6.238"/>
		<entryRelationship typeCode="SUBJ" inversionInd="true">
			<!--	Aggregate Count template	-->
			<observation classCode="OBS" moodCode="EVN">
				<templateId root="2.16.840.1.113883.10.20.27.3.3"/>
				<code code="MSRAGG" displayName="rate aggregation" codeSystem="2.16.840.1.113883.5.4" codeSystemName="ActCode"/>
				<value xsi:type="INT" value="%d"/>
				<methodCode code="COUNT" displayName="Count" codeSystem="2.16.840.1.113883.5.84" codeSystemName="ObservationMethod"/>
			</observation>
		</entryRelationship>
	</observation>
</entryRelationship>

<!--	RACE Supplemental Data Reporting for %s	%s	-->

<!--	Supplemental Data Template	-->

<entryRelationship typeCode="COMP">
	<observation classCode="OBS" moodCode="EVN">
		<!--	Race Supplemental Data	-->
		<templateId root="2.16.840.1.113883.10.20.27.3.8" extension="2016-09-01"/>
		<id nullFlavor="NA"/>
		<code code="72826-1" codeSystem="2.16.840.1.113883.6.1"/>
		<statusCode code="completed"/>
		<value xsi:type="CD" code="%s" codeSystem="2.16.840.1.113883.6.238"/>
		<entryRelationship typeCode="SUBJ" inversionInd="true">
			<!--	Aggregate Count template	-->
			<observation classCode="OBS" moodCode="EVN">
				<templateId root="2.16.840.1.113883.10.20.27.3.3"/>
				<code code="MSRAGG" displayName="rate aggregation" codeSystem="2.16.840.1.113883.5.4" codeSystemName="ActCode"/>
				<value xsi:type="INT" value="%d"/>
				<methodCode code="COUNT" displayName="Count" codeSystem="2.16.840.1.113883.5.84" codeSystemName="ObservationMethod"/>
			</observation>
		</entryRelationship>
	</observation>
</entryRelationship>

<!--	PAYER Supplemental Data Reporting for %s	%s	-->

<!--	Supplemental Data Template	-->

<entryRelationship typeCode="COMP">
	<observation classCode="OBS" moodCode="EVN">
		<!--	Payer Supplemental Data	-->
		<templateId root="2.16.840.1.113883.10.20.27.3.9" extension="2016-02-01"/>
		<id nullFlavor="NA"/>
		<code code="48768-6" codeSystem="2.16.840.1.113883.6.1"/>
		<statusCode code="completed"/>
		<value xsi:type="CD" code="%s" codeSystem="2.16.840.1.113883.3.221.5"/>
		<entryRelationship typeCode="SUBJ" inversionInd="true">
			<!--	Aggregate Count template	-->
			<observation classCode="OBS" moodCode="EVN">
				<templateId root="2.16.840.1.113883.10.20.27.3.3"/>
				<code code="MSRAGG" displayName="rate aggregation" codeSystem="2.16.840.1.113883.5.4" codeSystemName="ActCode"/>
				<value xsi:type="INT" value="%d"/>
				<methodCode code="COUNT" displayName="Count" codeSystem="2.16.840.1.113883.5.84" codeSystemName="ObservationMethod"/>
			</observation>
		</entryRelationship>
	</observation>
</entryRelationship>`
