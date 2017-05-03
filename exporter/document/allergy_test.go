package document_test

import (
	"fmt"
	"testing"

	"github.com/projectcypress/cdatools/exporter/document"
	"github.com/projectcypress/cdatools/models"
)

func TestPrint(t *testing.T) {
	startTime := int64(1)
	var tests = []struct {
		n        document.Allergy
		expected string
	}{
		{
			document.NewAllergy(models.Allergy{
				Entry: models.Entry{StartTime: &startTime, Description: "A test allergy", ID: models.CDAIdentifier{Extension: "extension"}}}),
			fmt.Sprintf(allergyCat1TestTemplate, "extension", 1, "A test allergy"),
		},
	}

	for _, tt := range tests {
		actual := tt.n.Print()
		if actual != tt.expected {
			t.Errorf("Allergy.Print(): expected %s, actual %s", tt.expected, actual)
		}
	}
}

var allergyCat1TestTemplate = `
<entry>
  <observation classCode="OBS" moodCode="EVN">
    <!-- consolidation CDA Allergy Observation template -->
    <templateId root="2.16.840.1.113883.10.20.22.4.7" extension="2014-06-09"/>
    <templateId root="2.16.840.1.113883.10.20.24.3.46" extension="2014-12-01"/>
    <id root="1.3.6.1.4.1.115" extension="%s"/>
    <code code="ASSERTION" 
          displayName="Assertion" 
          codeSystem="2.16.840.1.113883.5.4" 
          codeSystemName="ActCode"/>
    <statusCode code="completed"/>
    <effectiveTime>
      <!--<low valueOrNullFlavor .StartTime/>-->
      <low %d/>
    </effectiveTime>
    <value xsi:type="CD" 
           code="59037007" 
           displayName="Drug intolerance"
           codeSystem="2.16.840.1.113883.6.96" 
           codeSystemName="SNOMED CT"/>
    <participant typeCode="CSM">
      <participantRole classCode="MANU">
        <playingEntity classCode="MMAT">
          <name>%s</name>
        </playingEntity>
      </participantRole>
    </participant>
  </observation>
</entry>`
