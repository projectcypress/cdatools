package models

import (
	"encoding/json"
	"testing"

	"github.com/projectcypress/cdatools/fixtures"
	"github.com/stretchr/testify/assert"
)

/*func TestCodeDisplayWithPreferredCode(t *testing.T) {
	vs := []ValueSet{}
	json.Unmarshal(fixtures.Cms9_26, &vs)
	vsMap := NewValueSetMap(vs)
	codeType := "my code type"
	mapDataCriteria := Mdc{DcKey: DcKey{ValueSetOid: "2.16.840.1.113883.3.117.1.7.1.279"}}
	expectedCodeDisplay := CodeDisplay{CodeType: codeType, PreferredCodeSets: []string{"codeSetB"}, MapDataCriteria: mapDataCriteria}
	entry := Entry{CodeDisplays: []CodeDisplay{expectedCodeDisplay}}
	codes := make(map[string][]string)
	codes["codeSetA"] = []string{"third", "fourth"}
	codes["codeSetB"] = []string{"first", "second"}
	coded := Coded{Codes: codes}

	actualCodeDisplay := vsMap.CodeDisplayWithPreferredCode(&entry, &coded, mapDataCriteria, codeType)
	expectedCodeDisplay.PreferredCode = Concept{Code: "first", CodeSystem: "codeSetB"}
	assert.Equal(t, expectedCodeDisplay, actualCodeDisplay)
}*/

func TestGenerateCodeDisplay(t *testing.T) {
	vs := []ValueSet{}
	json.Unmarshal(fixtures.Cms9_26, &vs)
	vsMap := NewValueSetMap(vs)
	codeType := "my code type"
	mapDataCriteria := Mdc{DcKey: DcKey{ValueSetOid: "2.16.840.1.113883.3.117.1.7.1.279"}}
	expectedCodeDisplay := CodeDisplay{CodeType: codeType, PreferredCodeSets: []string{"codeSetB"}, MapDataCriteria: mapDataCriteria}
	entry := Entry{CodeDisplays: []CodeDisplay{expectedCodeDisplay}}
	codes := make(map[string][]string)
	codes["codeSetA"] = []string{"third", "fourth"}
	codes["codeSetB"] = []string{"first", "second"}
	coded := Coded{Codes: codes}

	actualCodeDisplay := vsMap.GenerateCodeDisplay(&entry, &coded, mapDataCriteria, codeType)
	expectedCodeDisplay.PreferredCode = Concept{Code: "first", CodeSystem: "codeSetB"}
	assert.Equal(t, expectedCodeDisplay, actualCodeDisplay)
}

func TestOidForCode(t *testing.T) {
	vs := []ValueSet{}
	json.Unmarshal(fixtures.Cms9_26, &vs)
	vsMap := NewValueSetMap(vs)
	coded := CodedConcept{Code: "3950001", CodeSystem: "2.16.840.1.113883.6.96"}
	coded2 := CodedConcept{Code: "3950001222", CodeSystem: "2.16.840.1.113883.6.96"}
	vsoids := []string{"2.16.840.1.113883.3.117.1.7.1.70", "2.16.840.1.113883.3.117.1.7.1.27", "2.16.840.1.113883.3.117.1.7.1.26", "2.16.840.1.113883.3.117.1.7.1.25"}

	assert.Equal(t, vsMap.OidForCode(coded, vsoids), "2.16.840.1.113883.3.117.1.7.1.70")
	assert.Equal(t, vsMap.OidForCode(coded2, vsoids), "")
}
