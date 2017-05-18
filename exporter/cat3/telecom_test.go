package cat3_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/projectcypress/cdatools/exporter/cat3"
	"github.com/projectcypress/cdatools/models"
)

func TestTelecomsTemplate(t *testing.T) {
	var tests = []struct {
		n        cat3.Telecoms
		expected string
	}{
		{
			cat3.Telecoms([]models.Telecom{models.Telecom{
				Use: "use", Value: "value"}}),
			fmt.Sprintf(telecomsTestTemplate, "use", "value"),
		},
	}

	for _, tt := range tests {
		actual := cat3.Print(tt.n.Template(), tt.n)
		if strings.TrimSpace(actual) != strings.TrimSpace(tt.expected) {
			t.Errorf("Telecoms.Template(): expected \n%s\n, actual \n%s", tt.expected, actual)
		}
	}
}

var telecomsTestTemplate = `<telecom use="%s" value="tel:+%s" />`
