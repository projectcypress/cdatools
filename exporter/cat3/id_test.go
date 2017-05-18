package cat3_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/projectcypress/cdatools/exporter/cat3"
	"github.com/projectcypress/cdatools/models"
)

func TestIdTemplate(t *testing.T) {
	var tests = []struct {
		n        cat3.CDAIdentifiers
		expected string
	}{
		{
			cat3.CDAIdentifiers([]models.CDAIdentifier{models.CDAIdentifier{
				Root: "root", Extension: "extension"},
				models.CDAIdentifier{Extension: "extension"}}),
			fmt.Sprintf(cdaIdentifiersTestTemplate, "root", "extension", "extension"),
		},
	}

	for _, tt := range tests {
		actual := cat3.Print(tt.n.Template(), tt.n)
		if strings.TrimSpace(actual) != strings.TrimSpace(tt.expected) {
			t.Errorf("CDAIdentifiers.Template(): expected \n%s\n, actual \n%s", tt.expected, actual)
		}
	}
}

var cdaIdentifiersTestTemplate = `<id root="%s" extension="%s" />
<id  extension="%s" />`
