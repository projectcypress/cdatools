package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOidForCodeForField(t *testing.T) {
	valueSets, err := ioutil.ReadFile("../fixtures/value_sets/cms9_26.json")
	if err != nil {
		log.Fatalln(err)
	}
	vs := []ValueSet{}
	json.Unmarshal(valueSets, &vs)
	vsMap := NewValueSetMap(vs)

	coded := CodedConcept{Code: "3950001", CodeSystem: "2.16.840.1.113883.6.96"}
	vsoids := []string{"2.16.840.1.113883.3.117.1.7.1.70", "2.16.840.1.113883.3.117.1.7.1.27", "2.16.840.1.113883.3.117.1.7.1.26", "2.16.840.1.113883.3.117.1.7.1.25"}
	fieldOids := make(map[string][]string)
	fieldOids["REASON"] = vsoids
	assert.Equal(t, "2.16.840.1.113883.3.117.1.7.1.70", vsMap.OidForCodeForField(coded, fieldOids, "REASON"))

	fieldOidsNoReason := make(map[string][]string)
	assert.Equal(t, "", vsMap.OidForCodeForField(coded, fieldOidsNoReason, "REASON"))

	fieldOidsNoOids := make(map[string][]string)
	fieldOidsNoOids["Reason"] = []string{}
	assert.Equal(t, "", vsMap.OidForCodeForField(coded, fieldOidsNoOids, "REASON"))
}

func TestOidForCode(t *testing.T) {
	valueSets, err := ioutil.ReadFile("../fixtures/value_sets/cms9_26.json")
	if err != nil {
		log.Fatalln(err)
	}
	vs := []ValueSet{}
	json.Unmarshal(valueSets, &vs)
	vsMap := NewValueSetMap(vs)
	coded := CodedConcept{Code: "3950001", CodeSystem: "2.16.840.1.113883.6.96"}
	coded2 := CodedConcept{Code: "3950001222", CodeSystem: "2.16.840.1.113883.6.96"}
	vsoids := []string{"2.16.840.1.113883.3.117.1.7.1.70", "2.16.840.1.113883.3.117.1.7.1.27", "2.16.840.1.113883.3.117.1.7.1.26", "2.16.840.1.113883.3.117.1.7.1.25"}

	assert.Equal(t, vsMap.OidForCode(coded, vsoids), "2.16.840.1.113883.3.117.1.7.1.70")
	assert.Equal(t, vsMap.OidForCode(coded2, vsoids), "")
}
