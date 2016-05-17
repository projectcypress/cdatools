package exporter

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/projectcypress/cdatools/models"
	"github.com/stretchr/testify/assert"
)

func TestValueOrNullFlavor(t *testing.T) {
	assert.Equal(t, valueOrNullFlavor(nil), "nullFlavor='UNK'")
	assert.Equal(t, valueOrNullFlavor(0), "value='196912311900'")
	assert.Equal(t, valueOrNullFlavor(int64(0)), "value='196912311900'")
	assert.Equal(t, valueOrNullFlavor("0"), "value='196912311900'")
}

func TestEscape(t *testing.T) {
	assert.Equal(t, escape("&"), "&amp;")
	assert.Equal(t, escape(1), "1")
	assert.Equal(t, escape(nil), "")
}

func TestValueOrDefault(t *testing.T) {
	assert.Equal(t, valueOrDefault(nil, "hey"), "hey")
	assert.Equal(t, valueOrDefault("hey", "hey thar"), "hey")
}

func TestCodeDisplay(t *testing.T) {
	var entry = models.Entry{}
	var m = make(map[string]interface{})
	m["preferred_code_sets"] = []string{"*"}
	assert.Equal(t, codeDisplay(entry, m), "<code code='1234' codeSystem='2.16.840.1.113883.6.96' ><originalText></originalText> </code>")
}

func TestOidForCode(t *testing.T) {
	valueSets, _ := ioutil.ReadFile("../fixtures/value_sets/CMS9_26.json")
	vs := []models.ValueSet{}
	json.Unmarshal(valueSets, &vs)
	initializeVsMap(vs)
	coded := models.CodedConcept{Code: "3950001", CodeSystem: "2.16.840.1.113883.6.96"}
	coded2 := models.CodedConcept{Code: "3950001222", CodeSystem: "2.16.840.1.113883.6.96"}
	vsoids := []string{"2.16.840.1.113883.3.117.1.7.1.70", "2.16.840.1.113883.3.117.1.7.1.27", "2.16.840.1.113883.3.117.1.7.1.26", "2.16.840.1.113883.3.117.1.7.1.25"}

	assert.Equal(t, oidForCode(coded, vsoids), "2.16.840.1.113883.3.117.1.7.1.70")
	assert.Equal(t, oidForCode(coded2, vsoids), "")
}
