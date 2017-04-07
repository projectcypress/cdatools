package doc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/projectcypress/cdatools/exporter/doc"
	"github.com/projectcypress/cdatools/models"
)

func TestAuthorsPrint(t *testing.T) {
	aTime := int64(1)
	var tests = []struct {
		n        doc.Authors
		expected string
	}{
		{
			doc.NewAuthors([]models.Author{models.Author{
				Time: &aTime, Person: models.Person{First: "first", Last: "last"}},
				models.Author{Time: &aTime, Device: models.Device{Model: "model", Name: "name"}}}),
			fmt.Sprintf(authorsCat3TestTemplate1, 1, "first", "last", 1, "model", "name"),
		},
	}

	for _, tt := range tests {
		actual := tt.n.Print()
		if strings.TrimSpace(actual) != strings.TrimSpace(tt.expected) {
			t.Errorf("Author.Print(): expected \n%s\n, actual \n%s", tt.expected, actual)
		}
	}
}

var authorsCat3TestTemplate1 = `<!-- SHALL have 1..* author. MAY be device or person. 
    The author of the CDA document in this example is a device at a data submission vendor/registry. -->
<author>
	<time value="%d" />
	<assignedAuthor>
		
		
		
		<assignedPerson>
			<name>
				<given>%s</given>
				<family>%s</family>
			</name>
		</assignedPerson>
		<>
	<!-- Represents unique registry organization TIN -->
	
	<!-- Contains name - specific registry not required-->
	<name></name>
</> <!--TagName "representedOrganization"-->
	</assignedAuthor>
</author>
<author>
	<time value="%d" />
	<assignedAuthor>
		
		
		
		<assignedAuthoringDevice>
			<manufacturerModelName>%s</manufacturerModelName>
			<softwareName>%s</softwareName>
		</assignedAuthoringDevice>
		<>
	<!-- Represents unique registry organization TIN -->
	
	<!-- Contains name - specific registry not required-->
	<name></name>
</> <!--TagName "representedOrganization"-->
	</assignedAuthor>
</author>`
