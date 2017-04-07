package document_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/projectcypress/cdatools/exporter/document/cat3/r2"
	"github.com/projectcypress/cdatools/models"
)

func TestTelecomsPrint(t *testing.T) {
	var tests = []struct {
		n        document.Telecoms
		expected string
	}{
		{
			document.NewTelecoms([]models.Telecom{models.Telecom{
				Use: "use", Value: "value"}}),
			fmt.Sprintf(telecomsCat3TestTemplate, "use", "value"),
		},
	}

	for _, tt := range tests {
		actual := tt.n.Print()
		if strings.TrimSpace(actual) != strings.TrimSpace(tt.expected) {
			t.Errorf("Telecoms.Print(): expected \n%s\n, actual \n%s", tt.expected, actual)
		}
	}
}

var telecomsCat3TestTemplate = `<telecom use="%s" value="tel:+%s" />`
