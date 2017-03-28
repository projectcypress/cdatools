package document

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/projectcypress/cdatools/exporter/document"
	"github.com/projectcypress/cdatools/models"
)

type ProviderPerformances struct {
	ProviderPerformances []models.ProviderPerformance
	Timestamp            int64
}

func NewProviderPerformances(p []models.ProviderPerformance) ProviderPerformances {
	timeNow := time.Now().UTC().Unix()
	return ProviderPerformances{ProviderPerformances: p, Timestamp: timeNow}
}

func (p ProviderPerformances) Print() string {
	tmpl := template.New("")
	tmpl, err := tmpl.Funcs(document.ExporterFuncMapCat3(tmpl)).Parse(p.cat3Template())
	if err != nil {
		fmt.Println("error making template:")
		fmt.Println(err)
		return ""
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, p)
	if err != nil {
		fmt.Println("error executing template:")
		fmt.Println(err)
		return ""
	}
	return b.String()
}

// NOTE: Need to add this into the template above .Description
// <!--<%== code_display(entry,'value_set_map'
// => filtered_vs_map, 'preferred_code_sets'
// => ['RxNorm', 'SNOMED-CT', 'CVX'], 'extra_content'
// => "sdtc:valueSet=\"#{value_set_oid}\"") %>-->
func (p ProviderPerformances) cat3Template() string {
	t := `<documentationOf typeCode="DOC">
	<serviceEvent classCode="PCPR"> <!-- care provision -->
		{{if .ProviderPerformances -}}
		{{range .ProviderPerformances -}}
		<effectiveTime>
			<low value="{{timeToFormat .StartDate "20060102"}}"/>
			<high value="{{timeToFormat .EndDate "20060102"}}"/>
		</effectiveTime>
		<!-- You can include multiple performers, each with an NPI, TIN, CCN. -->
		<performer typeCode="PRF">
			<time>
				<low value="{{timeToFormat .StartDate "20060102"}}"/>
				<high value="{{timeToFormat .EndDate "20060102"}}"/>
			</time>
			<assignedEntity>
				{{range .Provider.CDAIdentifiers -}}
				{{if ne .Root "2.16.840.1.113883.4.2" -}}
					<id root="{{ escape .Root}}" extension="{{ escape .Extension}}" />
				{{end -}}
				{{end -}}
				<representedOrganization>
					{{range .Provider.CDAIdentifiers -}}
					{{if eq .Root "2.16.840.1.113883.4.2" -}}
						<id root="2.16.840.1.113883.4.2" extension="{{ escape .Extension}}" />
					{{end -}}
					{{- end}}
				</representedOrganization>
			</assignedEntity>
		</performer>
		{{- end}}
		{{else -}}
		<!-- No provider data found in the patient record
			putting in a fake provider -->
		<effectiveTime>
			<low value="20020716"/>
			<high value="{{timeToFormat .Timestamp "20060102"}}"/>
		</effectiveTime>
		<!-- You can include multiple performers, each with an NPI, TIN, CCN. -->
		<performer typeCode="PRF">
			<time>
				<low value="20020716"/>
				<high value="{{timeToFormat .Timestamp "20060102"}}"/>
			</time>
			<assignedEntity>
				<!-- This is the provider NPI -->
				<id root="2.16.840.1.113883.4.6" extension="111111111" />
				<representedOrganization>
					<!-- This is the organization TIN -->
					<id root="2.16.840.1.113883.4.2" extension="1234567" />
					<!-- This is the organization CCN -->
					<id root="2.16.840.1.113883.4.336" extension="54321" />
				</representedOrganization>
			</assignedEntity>
		</performer>
		{{- end}}
	</serviceEvent>
</documentationOf>`

	return t
}
