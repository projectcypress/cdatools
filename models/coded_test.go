package models

import (
	"encoding/json"
	"testing"

	"github.com/projectcypress/cdatools/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	if len(computeIntersection([]string{"a", "b"}, []string{"a"})) != 1 {
		t.Error("Incorrect number of intersecting elements")
	}
	if len(computeIntersection([]string{"a", "b"}, []string{"a", "b"})) != 2 {
		t.Error("Incorrect number of intersecting elements")
	}
	if len(computeIntersection([]string{"a", "b"}, []string{"c", "d"})) != 0 {
		t.Error("Incorrect number of intersecting elements")
	}
}

func TestPreferredCodes(t *testing.T) {
	mdcOid := "2.16.840.1.113883.3.117.1.7.1.279"
	codes := make(map[string][]string, 2)
	codes["a"] = []string{"aa", "ab"}
	codes["b"] = []string{"ba", "bb"}
	codes["2.16.840.1.113883.6.96"] = []string{"3950001", "3950001222"}
	codes2 := make(map[string][]string, 2)
	codes2["a"] = []string{"ba", "ab"}
	codes2["b"] = []string{"ba", "bb"}
	coded := Coded{Codes: codes}
	prefCode := coded.PreferredCodes([]string{"b"}, true, false, nil, mdcOid)
	assert.Equal(t, len(prefCode), 0)
	prefCode = coded.PreferredCodes([]string{"b"}, false, false, nil, mdcOid)
	assert.Equal(t, len(prefCode), 0)

	vs := []ValueSet{}
	json.Unmarshal(fixtures.Cms9_26, &vs)
	vsMap := NewValueSetMap(vs)
	prefCode = coded.PreferredCodes([]string{"b"}, true, false, vsMap, mdcOid)
	assert.Equal(t, len(prefCode), 0)
	prefCode = coded.PreferredCodes([]string{"2.16.840.1.113883.3.117.1.7.1.70"}, true, true, vsMap, mdcOid)
	assert.Equal(t, prefCode[0].Code, "3950001")
	prefCode = coded.PreferredCodes([]string{"2.16.840.1.113883.3.117.1.7.1.70"}, false, true, vsMap, mdcOid)
	assert.Equal(t, prefCode[0].Code, "3950001")
	assert.Equal(t, prefCode[1].Code, "3950001222")
	prefCode = coded.PreferredCodes([]string{"2.16.840.1.113883.3.117.1.7.1.26"}, true, true, vsMap, mdcOid)
	assert.Equal(t, len(prefCode), 0)
	coded = Coded{Codes: codes2}
	prefCode = coded.PreferredCodes([]string{"2.16.840.1.113883.3.117.1.7.1.26"}, false, true, vsMap, mdcOid)
	assert.Equal(t, len(prefCode), 0)
}
