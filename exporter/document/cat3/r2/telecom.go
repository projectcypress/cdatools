package document

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/projectcypress/cdatools/exporter/document"
	"github.com/projectcypress/cdatools/models"
)

// Telecoms most often come in groups, so it should be useful to print them as a group,
// rather than one at a time
type Telecoms struct {
	Telecoms []models.Telecom
}

func NewTelecoms(t []models.Telecom) Telecoms {
	return Telecoms{Telecoms: t}
}

func (t Telecoms) Print() string {
	tmpl := template.New("")
	tmpl, err := tmpl.Funcs(document.ExporterFuncMapCat3(tmpl)).Parse(t.cat3Template())
	if err != nil {
		fmt.Println("error making template:")
		fmt.Println(err)
		return ""
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, t)
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
func (te Telecoms) cat3Template() string {
	t := `{{range .Telecoms -}}
<telecom use="{{escape .Use}}" value="tel:+{{escape .Value}}" />
{{end -}}`
	return t
}
