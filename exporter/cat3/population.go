package cat3

type Population struct {
	ID               string           `json:"id,omitempty"`
	Type             string           `json:"type,omitempty"`
	Value            int              `json:"value,omitempty"`
	Stratifications  []Stratification `json:"stratifications,omitempty"`
	SupplementalData SupplementalData `json:"supplemental_data,omitempty"`
}

func (p Population) FindStratification(ID string) *Stratification {
	for _, strat := range p.Stratifications {
		if strat.ID == ID {
			return &strat
		}
	}
	return nil
}

type Stratification struct {
	ID    string `json:"id,omitempty"`
	Value int    `json:"value,omitempty"`
}

type SupplementalData struct {
	Sex       map[string]int `json:"SEX,omitempty"`
	Ethnicity map[string]int `json:"ETHNICITY,omitempty"`
	Race      map[string]int `json:"RACE,omitempty"`
	Payer     map[string]int `json:"PAYER,omitempty"`
}

func (p Population) Template() string {
	t := `{{- $pop := . }}
{{- if $pop.SupplementalData.Sex }}
{{- range $sex, $count := .SupplementalData.Sex -}}
<!--	SEX Supplemental Data Reporting for {{$pop.Type}}	{{$pop.ID}}	-->

<!--	Supplemental Data Template	-->

<entryRelationship typeCode="COMP">
	<observation classCode="OBS" moodCode="EVN">
		<!--	Sex Supplemental Data	-->
		<templateId root="2.16.840.1.113883.10.20.27.3.6" extension="2016-09-01"/>
		<id nullFlavor="NA"/>
		<code code="76689-9" codeSystem="2.16.840.1.113883.6.1"/>
		<statusCode code="completed"/>
		{{- if (or (eq $sex "") (eq $sex "UNK")) }}
		<value xsi:type="CD" nullFlavor="UNK" />
		{{- else }}
		<value xsi:type="CD" code="{{ $sex }}" codeSystem="2.16.840.1.113883.5.1"/>
		{{- end }}
		<entryRelationship typeCode="SUBJ" inversionInd="true">
			<!--	Aggregate Count template	-->
			<observation classCode="OBS" moodCode="EVN">
				<templateId root="2.16.840.1.113883.10.20.27.3.3"/>
				<code code="MSRAGG" displayName="rate aggregation" codeSystem="2.16.840.1.113883.5.4" codeSystemName="ActCode"/>
				<value xsi:type="INT" value="{{ $count }}"/>
				<methodCode code="COUNT" displayName="Count" codeSystem="2.16.840.1.113883.5.84" codeSystemName="ObservationMethod"/>
			</observation>
		</entryRelationship>
	</observation>
</entryRelationship>
{{- end }}
{{- end }}
{{- if $pop.SupplementalData.Ethnicity }}
{{ range $ethnicity, $count := .SupplementalData.Ethnicity }}
<!--	ETHNICITY Supplemental Data Reporting for {{ $pop.Type }}	{{ $pop.ID }}	-->

<!--	Supplemental Data Template	-->

<entryRelationship typeCode="COMP">
	<observation classCode="OBS" moodCode="EVN">
		<!--	Ethnicity Supplemental Data	-->
		<templateId root="2.16.840.1.113883.10.20.27.3.7" extension="2016-09-01"/>
		<id nullFlavor="NA"/>
		<code code="69490-1" codeSystem="2.16.840.1.113883.6.1"/>
		<statusCode code="completed"/>
		{{- if (or (eq $ethnicity "") (eq $ethnicity "UNK")) }}
		<value xsi:type="CD" nullFlavor="UNK" />
		{{- else }}
		<value xsi:type="CD" code="{{ $ethnicity }}" codeSystem="2.16.840.1.113883.6.238"/>
		{{- end }}
		<entryRelationship typeCode="SUBJ" inversionInd="true">
			<!--	Aggregate Count template	-->
			<observation classCode="OBS" moodCode="EVN">
				<templateId root="2.16.840.1.113883.10.20.27.3.3"/>
				<code code="MSRAGG" displayName="rate aggregation" codeSystem="2.16.840.1.113883.5.4" codeSystemName="ActCode"/>
				<value xsi:type="INT" value="{{ $count }}"/>
				<methodCode code="COUNT" displayName="Count" codeSystem="2.16.840.1.113883.5.84" codeSystemName="ObservationMethod"/>
			</observation>
		</entryRelationship>
	</observation>
</entryRelationship>
{{- end }}
{{- end }}
{{- if $pop.SupplementalData.Race }}
{{ range $race, $count := .SupplementalData.Race }}
<!--	RACE Supplemental Data Reporting for {{ $pop.Type }}	{{ $pop.ID }}	-->

<!--	Supplemental Data Template	-->

<entryRelationship typeCode="COMP">
	<observation classCode="OBS" moodCode="EVN">
		<!--	Race Supplemental Data	-->
		<templateId root="2.16.840.1.113883.10.20.27.3.8" extension="2016-09-01"/>
		<id nullFlavor="NA"/>
		<code code="72826-1" codeSystem="2.16.840.1.113883.6.1"/>
		<statusCode code="completed"/>
		{{- if (or (eq $race "") (eq $race "UNK")) }}
		<value xsi:type="CD" nullFlavor="UNK" />
		{{- else }}
		<value xsi:type="CD" code="{{ $race }}" codeSystem="2.16.840.1.113883.6.238"/>
		{{- end }}
		<entryRelationship typeCode="SUBJ" inversionInd="true">
			<!--	Aggregate Count template	-->
			<observation classCode="OBS" moodCode="EVN">
				<templateId root="2.16.840.1.113883.10.20.27.3.3"/>
				<code code="MSRAGG" displayName="rate aggregation" codeSystem="2.16.840.1.113883.5.4" codeSystemName="ActCode"/>
				<value xsi:type="INT" value="{{ $count }}"/>
				<methodCode code="COUNT" displayName="Count" codeSystem="2.16.840.1.113883.5.84" codeSystemName="ObservationMethod"/>
			</observation>
		</entryRelationship>
	</observation>
</entryRelationship>
{{- end }}
{{- end }}
{{- if $pop.SupplementalData.Payer }}
{{ range $payer, $count := .SupplementalData.Payer }}
<!--	PAYER Supplemental Data Reporting for {{ $pop.Type }}	{{ $pop.ID }}	-->

<!--	Supplemental Data Template	-->

<entryRelationship typeCode="COMP">
	<observation classCode="OBS" moodCode="EVN">
		<!--	Payer Supplemental Data	-->
		<templateId root="2.16.840.1.113883.10.20.27.3.9" extension="2016-02-01"/>
		<id nullFlavor="NA"/>
		<code code="48768-6" codeSystem="2.16.840.1.113883.6.1"/>
		<statusCode code="completed"/>
		{{- if (or (eq $payer "") (eq $payer "UNK")) }}
		<value xsi:type="CD" nullFlavor="UNK" />
		{{- else }}
		<value xsi:type="CD" code="{{ $payer }}" codeSystem="2.16.840.1.113883.3.221.5"/>
		{{- end }}
		<entryRelationship typeCode="SUBJ" inversionInd="true">
			<!--	Aggregate Count template	-->
			<observation classCode="OBS" moodCode="EVN">
				<templateId root="2.16.840.1.113883.10.20.27.3.3"/>
				<code code="MSRAGG" displayName="rate aggregation" codeSystem="2.16.840.1.113883.5.4" codeSystemName="ActCode"/>
				<value xsi:type="INT" value="{{ $count }}"/>
				<methodCode code="COUNT" displayName="Count" codeSystem="2.16.840.1.113883.5.84" codeSystemName="ObservationMethod"/>
			</observation>
		</entryRelationship>
	</observation>
</entryRelationship>
{{- end }}
{{- end }}`

	return t
}
