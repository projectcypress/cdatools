package doc

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/projectcypress/cdatools/models"
)

type CDAIdentifier struct {
	models.CDAIdentifier
}

func (i CDAIdentifier) Print() string {
	tmpl := template.New("")
	tmpl, err := tmpl.Funcs(ExporterFuncMapCat3(tmpl)).Parse(i.cat3Template())
	if err != nil {
		fmt.Println("error making template:")
		fmt.Println(err)
		return ""
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, i)
	if err != nil {
		fmt.Println("error executing template:")
		fmt.Println(err)
		return ""
	}
	return b.String()
}

func (i CDAIdentifier) cat3Template() string {
	t := `
<id {{if .Root}}root="{{escape .Root}}"{{end}} extension="{{escape .Extension}}" />`

	return t
}

// CDAIdentifiers most often come in groups, so it should be useful to print them as a group,
// rather than one at a time
type CDAIdentifiers struct {
	CDAIdentifiers []models.CDAIdentifier
}

func (i CDAIdentifiers) Print() string {
	tmpl := template.New("")
	tmpl, err := tmpl.Funcs(ExporterFuncMapCat3(tmpl)).Parse(i.cat3Template())
	if err != nil {
		fmt.Println("error making template:")
		fmt.Println(err)
		return ""
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, i)
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
func (i CDAIdentifiers) cat3Template() string {
	t := `{{range .CDAIdentifiers -}}
	<id {{if .Root}}root="{{escape .Root}}"{{end}} extension="{{escape .Extension}}" />
{{end -}}`
	return t
}
