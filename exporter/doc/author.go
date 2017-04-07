package doc

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/projectcypress/cdatools/models"
)

type Author struct {
	models.Author
	Ids          CDAIdentifiers
	Addresses    Addresses
	Telecoms     Telecoms
	Organization Organization
}

func NewAuthor(a models.Author) Author {
	return Author{
		Author:       a,
		Ids:          CDAIdentifiers{CDAIdentifiers: a.Ids},
		Addresses:    Addresses{Addresses: a.Addresses},
		Telecoms:     NewTelecoms(a.Telecoms),
		Organization: NewOrganization(a.Organization),
	}
}

type Authors struct {
	Authors []Author
}

func NewAuthors(a []models.Author) Authors {
	var authors []Author
	for _, value := range a {
		authors = append(authors, NewAuthor(value))
	}
	return Authors{Authors: authors}
}

func (a Authors) Print() string {
	tmpl := template.New("")
	tmpl, err := tmpl.Funcs(ExporterFuncMapCat3(tmpl)).Parse(a.cat3Template())
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

func (a Authors) cat3Template() string {
	t := `<!-- SHALL have 1..* author. MAY be device or person. 
    The author of the CDA document in this example is a device at a data submission vendor/registry. -->
{{range .Authors -}}
<author>
	<time value="{{.Time}}" />
	<assignedAuthor>
		{{.Ids.Print}}
		{{.Addresses.Print}}
		{{.Telecoms.Print}}
		{{if .Person.First -}}
		<assignedPerson>
			<name>
				<given>{{escape .Person.First}}</given>
				<family>{{escape .Person.Last}}</family>
			</name>
		</assignedPerson>
		{{else if .Device.Model -}}
		<assignedAuthoringDevice>
			<manufacturerModelName>{{escape .Device.Model}}</manufacturerModelName>
			<softwareName>{{escape .Device.Name}}</softwareName>
		</assignedAuthoringDevice>
		{{end -}}
		{{.Organization.Print}} <!--TagName "representedOrganization"-->
	</assignedAuthor>
</author>
{{end -}}`

	return t
}
