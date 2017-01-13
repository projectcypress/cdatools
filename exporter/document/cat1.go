package document

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/projectcypress/cdatools/models"
)

type Cat1 struct {
	//EntryInfos []models.EntryInfo
	Record models.Record
	Header models.Header
	//Measures   []models.Measure
	//ValueSets  []models.ValueSet
	//StartDate  int64
	//EndDate    int64
}

func NewCat1() Cat1 {
	return *new(Cat1)
}

func (c Cat1) Print() string {
	tmpl, err := template.New("").Parse(c.template())
	if err != nil {
		fmt.Println("error making template:")
		fmt.Println(err)
		return ""
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, c)
	if err != nil {
		fmt.Println("error executing template:")
		fmt.Println(err)
		return ""
	}
	return b.String()
}

func (c Cat1) template() string {
	template := `<?xml version="1.0" encoding="utf-8"?>
<ClinicalDocument xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
 xmlns="urn:hl7-org:v3"
 xmlns:voc="urn:hl7-org:v3/voc"
 xmlns:sdtc="urn:hl7-org:sdtc">
  <!-- QRDA Header -->
  <realmCode code="US"/>
  <typeId root="2.16.840.1.113883.1.3" extension="POCD_HD000040"/>
  <!-- US Realm Header Template Id -->
  <templateId root="2.16.840.1.113883.10.20.22.1.1" extension="2014-06-09" />
  <!-- QRDA templateId -->
  <templateId root="2.16.840.1.113883.10.20.24.1.1" extension="2014-12-01" />
  <!-- QDM-based QRDA templateId -->
  <templateId root="2.16.840.1.113883.10.20.24.1.2" />
  <!-- This is the globally unique identifier for this QRDA document -->
  <id root="{{ newRandom }}"/>
  <!-- QRDA document type code -->
  <code code="55182-0" codeSystem="2.16.840.1.113883.6.1" codeSystemName="LOINC" displayName="Quality Measure Report"/>
  <title>QRDA Incidence Report</title>
  <!-- This is the document creation time -->
  <effectiveTime value="{{ timeNow }}"/>
  <confidentialityCode code="N" codeSystem="2.16.840.1.113883.5.25"/>
  <languageCode code="en"/>
  <!-- reported patient -->
    {{template "_record_target.xml" .Record}}
  {{ .Header.Cat1View() }}
  {{template "_providers.xml" .Record}}
  <component>
    <structuredBody>
      {{template "_measures.xml" .Measures}}
      {{template "_reporting_parameters.xml" .}}
      {{ .PatientData.Cat1View() }}
    </structuredBody>
  </component>

</ClinicalDocument>
`

	return template
}
