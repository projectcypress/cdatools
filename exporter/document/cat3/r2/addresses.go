package document

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/projectcypress/cdatools/exporter/document"
	"github.com/projectcypress/cdatools/models"
)

// Addresses come in groups, so it should be useful to print them as a group,
// rather than one at a time
type Addresses struct {
	Addresses []models.Address
}

func NewAddresses(a []models.Address) Addresses {
	return Addresses{Addresses: a}
}

func (a Addresses) Print() string {
	tmpl := template.New("")
	tmpl, err := tmpl.Funcs(document.ExporterFuncMapCat3(tmpl)).Parse(a.cat3Template())
	if err != nil {
		fmt.Println("error making template:")
		fmt.Println(err)
		return ""
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, a)
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
func (a Addresses) cat3Template() string {
	t := `{{range .Addresses -}}
<addr use="{{escape .Use}}">
	{{range .Street -}}
	<streetAddressLine>{{escape .}}</streetAddressLine>
	{{end -}}
	<city>{{escape .City}}</city>
	<state>{{escape .State}}</state>
	<postalCode>{{escape .Zip}}</postalCode>
	<country>{{escape .Country}}</country>
</addr>
{{end -}}`

	return t
}
