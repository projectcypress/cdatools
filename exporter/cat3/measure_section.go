package cat3

import (
	"github.com/projectcypress/cdatools/models"
)

type MeasureSection struct {
	models.Measure
	Results         map[string]AggregateCount
	PerformanceRate PopulationGroup
	Data            MeasureData
}

type AggregateCount struct {
	Measure_HQMFID   string
	Populations      []Population
	PopulationGroups []PopulationGroup
}

func (ac AggregateCount) IsCV() bool {
	for _, pop := range ac.Populations {
		if pop.Type == "MSRPOPL" {
			return true
		}
	}
	return false
}

func (ac AggregateCount) FindObserv() *Population {
	for _, pop := range ac.Populations {
		if pop.Type == "OBSERV" {
			return &pop
		}
	}
	return nil
}

func (ms MeasureSection) NewMeasureData(pop Population, obs Population) MeasureData {
	return MeasureData{pop, obs}
}

func (ms MeasureSection) GetAggregateCount() AggregateCount {
	return ms.Results[ms.HQMFID]
}

// TODO: UUID.generate for when ID does not exist
func (ms MeasureSection) Template() string {
	t := `{{ $ms := . }}<component><structuredBody><component><section>
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
	<id extension="{{.ID}}"/>
	<statusCode code="completed"/>
	<reference typeCode="REFR">
	<externalDocument classCode="DOC" moodCode="EVN">
		<!-- SHALL: required Id but not restricted to the eMeasure Document/Id-->
		<!-- QualityMeasureDocument/id This is the version specific identifier for eMeasure -->
		<id root="2.16.840.1.113883.4.738" extension="{{.HQMFID}}"/>

		<!-- SHOULD This is the title of the eMeasure -->
		<text>{{.Name}}</text>
		<!-- SHOULD: setId is the eMeasure version neutral id	-->
		<setId root="{{.HQMFSetID}}"/>
		<!-- This is the sequential eMeasure Version number -->
		<versionNumber value="1"/>
	</externalDocument>
	</reference>
	{{- $aggregateCount := $ms.GetAggregateCount }}
	{{- if not $aggregateCount.IsCV }}
	{{- range $pg := $aggregateCount.PopulationGroups }}
	<component>
	{{- $pg.Print}}
	</component>
	{{- end }}
	{{- end }}
	{{- range $pop := $aggregateCount.Populations }}
	{{- if (ne $pop.Type "OBSERV") }}
	{{- $md := $ms.NewMeasureData $pop $pop }}
	{{- Print $md.Template $md }}
	{{- end }}
	{{- end }}
</organizer>
</entry></section>
	</component>
	</structuredBody>
	</component>`
	return t
}
