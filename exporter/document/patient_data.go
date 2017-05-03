package document

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/projectcypress/cdatools/models"
)

type PatientData struct {
	models.Record
}

func NewPatientData() PatientData {
	return *new(PatientData)
}

func (p PatientData) Print() string {
	tmpl, err := template.New("").Parse(p.cat1Template())
	if err != nil {
		fmt.Println("error making template:")
		fmt.Println(err)
		return ""
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, p)
	if err != nil {
		fmt.Println("error executing template:")
		fmt.Println(err)
		return ""
	}
	return b.String()
}

func (p PatientData) cat1Template() string {
	t := `<component>
<section>
    <!-- This is the templateId for Patient Data section -->
    <templateId root="2.16.840.1.113883.10.20.17.2.4"/>
    <!-- This is the templateId for Patient Data QDM section -->
    <templateId root="2.16.840.1.113883.10.20.24.2.1" extension="2014-12-01" />
    <code code="55188-7" codeSystem="2.16.840.1.113883.6.1"/>
    <title>Patient Data</title>
    <text>
    </text>
    {{range .EntryInfos}}
      {{executeTemplateForEntry .}}
    {{end}}
</section>
</component>`

	return t
}
