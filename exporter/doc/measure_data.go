package doc

import (
	"bytes"
	"fmt"
	"text/template"
)

type MeasureData struct {
	Population Population
	Observ     Population
}

func (md MeasureData) Print() string {
	tmpl := template.New("")
	tmpl, err := tmpl.Funcs(ExporterFuncMapCat3(tmpl)).Parse(md.cat3Template())
	if err != nil {
		fmt.Println("error making template:")
		fmt.Println(err)
		return ""
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, md)
	if err != nil {
		fmt.Println("error executing template:")
		fmt.Println(err)
		return ""
	}
	return b.String()
}

func (md MeasureData) cat3Template() string {
	t := `<!--	MEASURE DATA REPORTING FOR	{{.Population.Type}} {{.Population.ID}}	-->
<observation classCode="OBS" moodCode="EVN">
	<!--	Measure Data template	-->
	<templateId root="2.16.840.1.113883.10.20.27.3.5" extension="2016-09-01"/>
	<code code="ASSERTION"
				codeSystem="2.16.840.1.113883.5.4"
				displayName="Assertion"
				codeSystemName="ActCode"/>
	<statusCode code="completed"/>
	<value xsi:type="CD" code="{{ if (eq .Population.Type "IPP") -}} IPOP {{- else -}} {{.Population.Type}} {{- end}}"
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
			<value xsi:type="INT" value="{{.Population.Value}}"/>
			<methodCode code="COUNT"
				displayName="Count"
				codeSystem="2.16.840.1.113883.5.84"
				codeSystemName="ObservationMethod"/>
		</observation>
	</entryRelationship>
	{{- $md := .}}
	{{ range $strat := .Population.Stratifications }}
	<!--	Stratification Reporting Template for {{$md.Population.Type}} {{$md.Population.ID}}	Stratification {{$strat.ID}}	-->

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
					<value xsi:type="INT" value="{{$strat.Value}}"/>
					<methodCode code="COUNT"
											displayName="Count"
											codeSystem="2.16.840.1.113883.5.84"
											codeSystemName="ObservationMethod"/>
				</observation>
			</entryRelationship>
			{{- if (eq $md.Population.Type "MSRPOPL") }}
			{{- $obs_strat := $md.Observ.FindStratification $strat.ID }}
			{{- if $obs_strat }}
			<entryRelationship typeCode="COMP">
				<observation classCode="OBS" moodCode="EVN">
					<templateId root="2.16.840.1.113883.10.20.27.3.2"/>
					<code nullFlavor="OTH">
						<originalText>Time Difference</originalText>
					</code>
					<statusCode code="completed"/>
					<value xsi:type="PQ" value="{{$obs_strat.Value}}" unit="min"/>
					<methodCode code="MEDIAN"
											displayName="Median"
											codeSystem="2.16.840.1.113883.5.84"
											codeSystemName="ObservationMethod"/>
					<reference typeCode="REFR">
						<!-- reference to the relevant measure observation in the eMeasure -->
						<externalObservation classCode="OBS" moodCode="EVN">
							<id root="{{$obs_strat.ID}}"/>
						</externalObservation>
					</reference>
				</observation>
			</entryRelationship>
			{{- end }}
			{{- end }}
			<reference typeCode="REFR">
				<externalObservation classCode="OBS" moodCode="EVN">
					<id root="{{$strat.ID}}"/>
				</externalObservation>
			</reference>
		</observation>
	</entryRelationship>
	{{- end }}
	{{- if $md.Population.SupplementalData }}
	{{- $md.Population.Print }}
	{{- end }}
	{{- if (eq $md.Population.Type "MSRPOPL") }}
	<entryRelationship typeCode="COMP">
		<observation classCode="OBS" moodCode="EVN">
			<templateId root="2.16.840.1.113883.10.20.27.3.2"/>
			<code nullFlavor="OTH">
				<originalText>Time Difference</originalText>
			</code>
			<statusCode code="completed"/>
			<value xsi:type="PQ" value="{{$md.Observ.Value}}" unit="min"/>
			<methodCode code="MEDIAN"
									displayName="Median"
									codeSystem="2.16.840.1.113883.5.84"
									codeSystemName="ObservationMethod"/>
			<reference typeCode="REFR">
				<!-- reference to the relevant measure observation in the eMeasure -->
				<externalObservation classCode="OBS" moodCode="EVN">
					<id root="{{$md.Observ.ID}}"/>
				</externalObservation>
			</reference>
		</observation>
	</entryRelationship>
	{{- end }}
	<reference typeCode="REFR">
		 <externalObservation classCode="OBS" moodCode="EVN">
				<id root="{{$md.Population.ID}}"/>
		 </externalObservation>
	</reference>
</observation>`

	return t
}
