package doc

import (
	"bytes"
	"fmt"
	"text/template"

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
	tmpl, err := tmpl.Funcs(ExporterFuncMapCat3(tmpl)).Parse(t.cat3Template())
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

func (te Telecoms) cat3Template() string {
	t := `{{range .Telecoms -}}
<telecom use="{{escape .Use}}" value="tel:+{{escape .Value}}" />
{{end -}}`
	return t
}
