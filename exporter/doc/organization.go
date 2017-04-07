package doc

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/projectcypress/cdatools/models"
)

type Organization struct {
	models.Organization
	Ids CDAIdentifiers
}

func NewOrganization(i models.Organization) Organization {
	return Organization{
		Organization: i,
		Ids:          NewCDAIdentifiers(i.Ids),
	}
}

func (i Organization) Print() string {
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

func (i Organization) cat3Template() string {
	t := `<{{escape .TagName}}>
	<!-- Represents unique registry organization TIN -->
	{{.Ids.Print}}
	<!-- Contains name - specific registry not required-->
	<name>{{escape .Name}}</name>
</{{escape .TagName}}>`

	return t
}
