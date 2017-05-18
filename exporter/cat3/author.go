package cat3

import (
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
		Ids:          CDAIdentifiers(a.Ids),
		Addresses:    Addresses(a.Addresses),
		Telecoms:     Telecoms(a.Telecoms),
		Organization: NewOrganization(a.Organization),
	}
}

type Authors []Author

func NewAuthors(a []models.Author) Authors {
	var authors []Author
	for _, value := range a {
		authors = append(authors, NewAuthor(value))
	}
	return Authors(authors)
}

func (a Authors) Template() string {
	t := `<!-- SHALL have 1..* author. MAY be device or person. 
    The author of the CDA document in this example is a device at a data submission vendor/registry. -->
{{range . -}}
<author>
	<time value="{{.Time}}" />
	<assignedAuthor>
		{{Print .Ids.Template .Ids}}
		{{Print .Addresses.Template .Addresses}}
		{{Print .Telecoms.Template .Telecoms}}
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
		{{Print .Organization.Template .Organization}} <!--TagName "representedOrganization"-->
	</assignedAuthor>
</author>
{{end -}}`

	return t
}
