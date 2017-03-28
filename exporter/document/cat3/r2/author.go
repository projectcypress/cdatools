package document

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/projectcypress/cdatools/exporter/document"
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
		Ids:          NewCDAIdentifiers(a.Ids),
		Addresses:    NewAddresses(a.Addresses),
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
