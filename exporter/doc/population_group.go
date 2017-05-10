package doc

import (
	"bytes"
	"fmt"
	"text/template"
)

type PopulationGroup struct {
	Populations []Population `json:"populations,omitempty"`
}

// Find Finds the type of population given.
// Possible values for query:
// NUMER - The numerator
// DENOM - The denominator
// DENEXCEP - The denominator exceptions
// DENEX - The denominator exclusions
func (pg PopulationGroup) Find(query string) *Population {
	for _, pop := range pg.Populations {
		if pop.Type == query {
			return &pop
		}
	}
	return nil
}

func (pg PopulationGroup) Numer() int {
	pop := pg.Find("NUMER")
	if pop != nil {
		return pop.Value
	}
	return 0
}

func (pg PopulationGroup) Denom() int {
	pop := pg.Find("DENOM")
	if pop != nil {
		return pop.Value
	}
	return 0
}

func (pg PopulationGroup) Denexcep() int {
	pop := pg.Find("DENEXCEP")
	if pop != nil {
		return pop.Value
	}
	return 0
}

func (pg PopulationGroup) Denex() int {
	pop := pg.Find("DENEX")
	if pop != nil {
		return pop.Value
	}
	return 0
}

func (pg PopulationGroup) PerformanceRate() float64 {
	return (float64(pg.Numer()) / float64(pg.PerformanceRateDenom()))
}

func (pg PopulationGroup) PerformanceRateDenom() int {
	return (pg.Denom() - pg.Denexcep() - pg.Denex())
}

func (pg PopulationGroup) Print() string {
	tmpl := template.New("")
	tmpl, err := tmpl.Funcs(ExporterFuncMapCat3(tmpl)).Parse(pg.cat3Template())
	if err != nil {
		fmt.Println("error making template:")
		fmt.Println(err)
		return ""
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, pg)
	if err != nil {
		fmt.Println("error executing template:")
		fmt.Println(err)
		return ""
	}
	return b.String()
}

func (pg PopulationGroup) cat3Template() string {
	t := `<observation classCode="OBS" moodCode="EVN">
	<templateId root="2.16.840.1.113883.10.20.27.3.14" extension="2016-09-01"/>
	<templateId root="2.16.840.1.113883.10.20.27.3.30" extension="2016-09-01"/>
	<code code="72510-1" codeSystem="2.16.840.1.113883.6.1"
		displayName="Performance Rate"
		codeSystemName="2.16.840.1.113883.6.1"/>
	<statusCode code="completed"/>
	{{- if gt .PerformanceRateDenom 0 }}
	<value xsi:type="REAL" value="{{.PerformanceRate}}"/>
	{{- else }}
		<value xsi:type="REAL" nullFlavor="NA"/>
	{{- end }}
	<reference typeCode="REFR">
		<externalObservation classCode="OBS" moodCode="EVN">
			{{- $numerator := .Find "NUMER" }}
			<id root="{{ $numerator.ID }}"/>
			<code code="NUMER" displayName="Numerator" codeSystem="2.16.840.1.113883.5.4" codeSystemName="ObservationValue"/>
		</externalObservation>
	</reference>
</observation>`
	return t
}
