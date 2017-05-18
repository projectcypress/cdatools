package cat3_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/projectcypress/cdatools/exporter/cat3"
	"github.com/projectcypress/cdatools/models"
)

func TestOrganizationTemplate(t *testing.T) {
	var tests = []struct {
		n        cat3.Organization
		expected string
	}{
		{
			cat3.NewOrganization(models.Organization{
				Name: "name", TagName: "tag"}),
			fmt.Sprintf(organizationTestTemplate, "tag", "name", "tag"),
		},
	}

	for _, tt := range tests {
		actual := cat3.Print(tt.n.Template(), tt.n)
		if strings.TrimSpace(actual) != strings.TrimSpace(tt.expected) {
			t.Errorf("Organization.Template(): expected \n%s\n, actual \n%s", tt.expected, actual)
		}
	}
}

var organizationTestTemplate = `<%s>
	<!-- Represents unique registry organization TIN -->
	
	<!-- Contains name - specific registry not required-->
	<name>%s</name>
</%s>`
