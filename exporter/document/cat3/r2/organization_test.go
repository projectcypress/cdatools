package document_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/projectcypress/cdatools/exporter/document/cat3/r2"
	"github.com/projectcypress/cdatools/models"
)

func TestOrganizationPrint(t *testing.T) {
	var tests = []struct {
		n        document.Organization
		expected string
	}{
		{
			document.NewOrganization(models.Organization{
				Name: "name", TagName: "tag"}),
			fmt.Sprintf(organizationCat3TestTemplate, "tag", "name", "tag"),
		},
	}

	for _, tt := range tests {
		actual := tt.n.Print()
		if strings.TrimSpace(actual) != strings.TrimSpace(tt.expected) {
			t.Errorf("Organization.Print(): expected \n%s\n, actual \n%s", tt.expected, actual)
		}
	}
}

var organizationCat3TestTemplate = `<%s>
	<!-- Represents unique registry organization TIN -->
	
	<!-- Contains name - specific registry not required-->
	<name>%s</name>
</%s>`
