package doc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/projectcypress/cdatools/exporter/doc"
	"github.com/projectcypress/cdatools/models"
)

func TestAddressesPrint(t *testing.T) {
	var tests = []struct {
		n        doc.Addresses
		expected string
	}{
		{
			doc.NewAddresses([]models.Address{models.Address{
				Street: []string{"line 1", "line 2"}, City: "city", State: "state", Zip: "zip", Country: "country", Use: "use"}}),
			fmt.Sprintf(addressesCat3TestTemplate, "use", "line 1", "line 2", "city", "state", "zip", "country"),
		},
	}

	for _, tt := range tests {
		actual := tt.n.Print()
		if strings.TrimSpace(actual) != strings.TrimSpace(tt.expected) {
			t.Errorf("Addresses.Print(): expected \n%s\n, actual \n%s", tt.expected, actual)
		}
	}
}

var addressesCat3TestTemplate = `<addr use="%s">
	<streetAddressLine>%s</streetAddressLine>
	<streetAddressLine>%s</streetAddressLine>
	<city>%s</city>
	<state>%s</state>
	<postalCode>%s</postalCode>
	<country>%s</country>
</addr>`
