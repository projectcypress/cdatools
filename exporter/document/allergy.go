package document

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/projectcypress/cdatools/models"
)

type Allergy struct {
	models.Allergy
}

func NewAllergy(a models.Allergy) Allergy {
	return Allergy{Allergy: a}
}

func (a Allergy) Print() string {
	tmpl, err := template.New("").Parse(a.cat1Template())
	if err != nil {
		fmt.Println("error making template:")
		fmt.Println(err)
		return ""
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, a)
	if err != nil {
		fmt.Println("error executing template:")
		fmt.Println(err)
		return ""
	}
	return b.String()
}

// NOTE: Need to add this into the template above .Description
// <!--<%== code_display(entry,'value_set_map'
// => filtered_vs_map, 'preferred_code_sets'
// => ['RxNorm', 'SNOMED-CT', 'CVX'], 'extra_content'
// => "sdtc:valueSet=\"#{value_set_oid}\"") %>-->
func (a Allergy) cat1Template() string {
	t := `
<entry>
  <observation classCode="OBS" moodCode="EVN">
    <!-- consolidation CDA Allergy Observation template -->
    <templateId root="2.16.840.1.113883.10.20.22.4.7" extension="2014-06-09"/>
    <templateId root="2.16.840.1.113883.10.20.24.3.46" extension="2014-12-01"/>
    <id root="1.3.6.1.4.1.115" extension="{{ .ID.Extension }}"/>
    <code code="ASSERTION" 
          displayName="Assertion" 
          codeSystem="2.16.840.1.113883.5.4" 
          codeSystemName="ActCode"/>
    <statusCode code="completed"/>
    <effectiveTime>
      <!--<low valueOrNullFlavor .StartTime/>-->
      <low {{ .StartTime }}/>
    </effectiveTime>
    <value xsi:type="CD" 
           code="59037007" 
           displayName="Drug intolerance"
           codeSystem="2.16.840.1.113883.6.96" 
           codeSystemName="SNOMED CT"/>
    <participant typeCode="CSM">
      <participantRole classCode="MANU">
        <playingEntity classCode="MMAT">
          <name>{{ .Description }}</name>
        </playingEntity>
      </participantRole>
    </participant>
  </observation>
</entry>`

	return t
}
