package models

import (	
	"encoding/json"
	"io/ioutil"
	"testing"
	
	"github.com/pebbe/util"
	"github.com/stretchr/testify/assert"
)

func TestReasonInCodesTrue(t *testing.T) {
	code := CodeSet{}
	reason := CodedConcept{}
	code.Values = []Concept{Concept{Code: "test", CodeSystem: "codeSystem"}}
	reason.Code = "test"
	reason.CodeSystem = "codeSystem"
	assert.True(t, reasonInCodes(code, reason))
}

func TestReasonInCodesFalse(t *testing.T) {
	code := CodeSet{}
	reason := CodedConcept{}
	code.Values = []Concept{Concept{Code: "not code", CodeSystem: "codeSystem"}}
	reason.Code = "test"
	reason.CodeSystem = "codeSystem"
	assert.False(t, reasonInCodes(code, reason))
}

func TestGetAllDataCriteriaForOneMeasure(t *testing.T) {
	mes := make([]Measure, 1)
	measureData, err := ioutil.ReadFile("../fixtures/measures/CMS9v4a.json")
	util.CheckErr(err)
	measureData = append([]byte("["), append(measureData, []byte("]")...)...)
	json.Unmarshal(measureData, &mes)
	assert.Equal(t, len(allDataCriteria(mes)), 27)
}

func TestGetallDatacriteriaForMultipleMeasures(t *testing.T) {
	mes := make([]Measure, 2)
	measureData, err := ioutil.ReadFile("../fixtures/measures/CMS9v4a.json")
	measureData2, err := ioutil.ReadFile("../fixtures/measures/CMS26v3.json")
	util.CheckErr(err)
	measureData = append([]byte("["), append(append(measureData, append([]byte(","), measureData2...)...), []byte("]")...)...)
	json.Unmarshal(measureData, &mes)

	assert.Equal(t, len(allDataCriteria(mes)), 47)
}

func TestGetUniqueDataCriteriaForOneMeasure(t *testing.T) {
	mes := make([]Measure, 1)
	measureData, err := ioutil.ReadFile("../fixtures/measures/CMS9v4a.json")
	util.CheckErr(err)
	measureData = append([]byte("["), append(measureData, []byte("]")...)...)
	json.Unmarshal(measureData, &mes)
	assert.Equal(t, len(UniqueDataCriteria(allDataCriteria(mes))), 14)
}

func TestAppendEntryInfos(t *testing.T) {
	// create entry sections
	entries := make([]Entry, 0)
	entries = append(entries, Entry{Description: "my description"})
	var entrySections []HasEntry
	for _, entry := range entries {
		entrySections = append(entrySections, &Encounter{Entry: entry})
	}
	entrySections = append(entrySections, nil) // appendEntryInfos() function should not include nil entry sections

	entryInfos := AppendEntryInfos([]EntryInfo{}, entrySections, Mdc{})
	assert.Equal(t, 1, len(entryInfos))
	assert.Equal(t, "my description", entryInfos[0].EntrySection.GetEntry().Description)
}

func TestSetCodeDisplaysForEntry(t *testing.T) {
	entry := &Entry{Oid: "2.16.840.1.113883.3.560.1.79"} // encounter performed hqmf oid
	assert.Equal(t, 0, len(entry.CodeDisplays))
	SetCodeDisplaysForEntry(entry)
	assert.Equal(t, 3, len(entry.CodeDisplays)) // three code displays the encounter performed entry
}

func TestEntriesForDataCriteria(t *testing.T) {
	patientData, err := ioutil.ReadFile("../fixtures/records/1_n_n_ami.json")
	util.CheckErr(err)

	measureData, err := ioutil.ReadFile("../fixtures/measures/CMS9v4a.json")
	util.CheckErr(err)

	patient := &Record{}
	measure := &Measure{}
	valueSetData, err := ioutil.ReadFile("../fixtures/value_sets/cms9_26.json")
	var vs []ValueSet
	json.Unmarshal(valueSetData, &vs)
	vsMap := InitializeVsMap(vs)

	json.Unmarshal(patientData, patient)
	json.Unmarshal(measureData, measure)

	var entries []HasEntry
	for _, crit := range measure.HQMFDocument.DataCriteria {
		if crit.HQMFOid != "" {
			for _, entryForDataCriteria := range patient.EntriesForDataCriteria(crit, vsMap) {
				entries = append(entries, entryForDataCriteria)
			}
		}
	}
	// TODO: This test will have to change when we get a new export of CMS9v4a with all the HQMFOid fields filled.
	assert.Equal(t, len(entries), 1)
}

// TODO: This needs to be broken into two tests or changed into a table test.
func TestAllPreferredCodeSetsIfNeeded(t *testing.T) {
	// code sets without the "*" string
	preferredCodeSets := []string{"one", "two", "three"}
	codeDisplays := make([]CodeDisplay, 1)
	codeDisplays[0] = CodeDisplay{PreferredCodeSets: preferredCodeSets}
	allPerferredCodeSetsIfNeeded(codeDisplays)
	assert.Equal(t, preferredCodeSets, codeDisplays[0].PreferredCodeSets)

	// all code sets indicated by the "*" string
	preferredCodeSets = append(preferredCodeSets, "*")
	codeDisplays[0] = CodeDisplay{PreferredCodeSets: preferredCodeSets}
	allPerferredCodeSetsIfNeeded(codeDisplays)
	assert.Equal(t, true, unorderedStringSlicesEqual(CodeSystemNames(), codeDisplays[0].PreferredCodeSets), "preferred code sets should include all code system names")
}

// TODO: NOTE: Start of test code that needs to be in its own Hds package
func TestImportHQMFTemplateJSON(t *testing.T) {
	var origID = "2.16.840.1.113883.10.20.28.3.19"
	var def = GetTemplateDefinition(origID, true)
	assert.Equal(t, def.Definition, "diagnosis")
	assert.Equal(t, def.Status, "resolved")
	var id = GetID(def, true)
	assert.Equal(t, id, origID)
}

func TestHqmfToQrdaOid(t *testing.T) {
	// ["Device, Applied", "Encounter, Performed", "Diagnostic Study, Intolerance"]
	hqmfOids := []string{"2.16.840.1.113883.3.560.1.10", "2.16.840.1.113883.3.560.1.79", "2.16.840.1.113883.3.560.1.39"}
	qrdaOids := []string{"2.16.840.1.113883.10.20.24.3.7", "2.16.840.1.113883.10.20.24.3.23", "2.16.840.1.113883.10.20.24.3.16"}
	for i, hqmfOid := range hqmfOids {
		assert.Equal(t, qrdaOids[i], HqmfToQrdaOid(hqmfOid))
	}
}


func TestCodeDisplayForQrdaOid(t *testing.T) {
	// invalid qrda oid
	codeDisplays := codeDisplayForQrdaOid("not a valid qrda oid")
	assert.Equal(t, 0, len(codeDisplays))

	// qrda oid with multiple code displays
	codeDisplays = codeDisplayForQrdaOid("2.16.840.1.113883.10.20.24.3.23")
	assert.Equal(t, 3, len(codeDisplays))
	assert.Equal(t, "entryCode", codeDisplays[0].CodeType)
	assert.Equal(t, "code", codeDisplays[0].TagName)
	assert.Equal(t, false, codeDisplays[0].ExcludeNullFlavor)
	assert.Equal(t, []string{"SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "CPT"}, codeDisplays[0].PreferredCodeSets)
}

func TestReasonValueSetOid(t *testing.T) {
	valueSets, _ := ioutil.ReadFile("../fixtures/value_sets/cms9_26.json")
	vs := []ValueSet{}
	json.Unmarshal(valueSets, &vs)
	InitializeVsMap(vs)

	coded := CodedConcept{Code: "3950001", CodeSystem: "2.16.840.1.113883.6.96"}
	vsoids := []string{"2.16.840.1.113883.3.117.1.7.1.70", "2.16.840.1.113883.3.117.1.7.1.27", "2.16.840.1.113883.3.117.1.7.1.26", "2.16.840.1.113883.3.117.1.7.1.25"}
	fieldOids := make(map[string][]string)
	fieldOids["REASON"] = vsoids
	assert.Equal(t, "2.16.840.1.113883.3.117.1.7.1.70", ReasonValueSetOid(coded, fieldOids))

	fieldOidsNoReason := make(map[string][]string)
	assert.Equal(t, "", ReasonValueSetOid(coded, fieldOidsNoReason))

	fieldOidsNoOids := make(map[string][]string)
	fieldOidsNoOids["Reason"] = []string{}
	assert.Equal(t, "", ReasonValueSetOid(coded, fieldOidsNoOids))
}

func TestOidForCode(t *testing.T) {
	valueSets, _ := ioutil.ReadFile("../fixtures/value_sets/cms9_26.json")
	vs := []ValueSet{}
	json.Unmarshal(valueSets, &vs)
	InitializeVsMap(vs)
	coded := CodedConcept{Code: "3950001", CodeSystem: "2.16.840.1.113883.6.96"}
	coded2 := CodedConcept{Code: "3950001222", CodeSystem: "2.16.840.1.113883.6.96"}
	vsoids := []string{"2.16.840.1.113883.3.117.1.7.1.70", "2.16.840.1.113883.3.117.1.7.1.27", "2.16.840.1.113883.3.117.1.7.1.26", "2.16.840.1.113883.3.117.1.7.1.25"}

	assert.Equal(t, OidForCode(coded, vsoids), "2.16.840.1.113883.3.117.1.7.1.70")
	assert.Equal(t, OidForCode(coded2, vsoids), "")
}

func TestStringInSlice(t *testing.T) {
	stringSlice := []string{"one", "two", "four"}
	assert.Equal(t, true, stringInSlice("one", stringSlice))
	assert.Equal(t, true, stringInSlice("two", stringSlice))
	assert.Equal(t, false, stringInSlice("three", stringSlice))
}

// END code that needs to be in its own Hds package

// NOTE: This is only used for a test.
func unorderedStringSlicesEqual(a, b []string) bool {
	if a == nil || b == nil || len(a) != len(b) {
		return false
	}
	for _, elemA := range a {
		if !stringInSlice(elemA, b) {
			return false
		}
	}
	return true
}