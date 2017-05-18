package cat3_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/projectcypress/cdatools/exporter/cat3"
	"github.com/projectcypress/cdatools/models"
)

func TestAddressesTemplate(t *testing.T) {
	var tests = []struct {
		n        cat3.Addresses
		expected string
	}{
		{
			cat3.Addresses([]models.Address{models.Address{
				Street: []string{"line 1", "line 2"}, City: "city", State: "state", Zip: "zip", Country: "country", Use: "use"}}),
			fmt.Sprintf(addressesTestTemplate, "use", "line 1", "line 2", "city", "state", "zip", "country"),
		},
	}

	for _, tt := range tests {
		actual := cat3.Print(tt.n.Template(), tt.n)
		if strings.TrimSpace(actual) != strings.TrimSpace(tt.expected) {
			t.Errorf("Addresses.Template(): expected \n%s\n, actual \n%s", tt.expected, actual)
		}
	}
}

var addressesTestTemplate = `<addr use="%s">
	<streetAddressLine>%s</streetAddressLine>
	<streetAddressLine>%s</streetAddressLine>
	<city>%s</city>
	<state>%s</state>
	<postalCode>%s</postalCode>
	<country>%s</country>
</addr>`
