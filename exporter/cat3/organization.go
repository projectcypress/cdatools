package cat3

import (
	"github.com/projectcypress/cdatools/models"
)

type Organization struct {
	models.Organization
	Ids CDAIdentifiers
}

func NewOrganization(org models.Organization) Organization {
	return Organization{
		Organization: org,
		Ids:          CDAIdentifiers(org.Ids),
	}
}

func (org Organization) Template() string {
	t := `<{{escape .TagName}}>
	<!-- Represents unique registry organization TIN -->
	{{Print .Ids.Template .Ids}}
	<!-- Contains name - specific registry not required-->
	<name>{{escape .Name}}</name>
</{{escape .TagName}}>`

	return t
}
