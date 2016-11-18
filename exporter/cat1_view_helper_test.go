package exporter

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
	"github.com/stretchr/testify/assert"
)

var fieldOids = map[string][]string{"REASON": []string{"1.2.3.4.5.6.7.8.9.11"},
	"ORDINAL":  []string{"1.2.3.4.5.6.7.8.9.10"},
	"SEVERITY": []string{"1.2.3.4.5.6.7.8.9.13"},
	"ROUTE":    []string{"1.2.3.4.5.6.7.8.9.12"}}
var resultOids = []string{"1.2.3.4.5.6.7.8.9.14"}
var vsOid = "1.2.3.4.5.6.7.8.9"

func TestValueOrNullFlavor(t *testing.T) {
	assert.Equal(t, valueOrNullFlavor(nil), "nullFlavor='UNK'")
	assert.Equal(t, valueOrNullFlavor(0), "value='197001010000+0000'")
	assert.Equal(t, valueOrNullFlavor(int64(0)), "value='197001010000+0000'")
	assert.Equal(t, valueOrNullFlavor("0"), "value='197001010000+0000'")
}

func TestEscape(t *testing.T) {
	assert.Equal(t, escape("&"), "&amp;")
	assert.Equal(t, escape(1), "1")
	assert.Equal(t, "42", escape(int64(42)))
	assert.Equal(t, escape(nil), "")
}

func TestValueOrDefault(t *testing.T) {
	assert.Equal(t, valueOrDefault(nil, "hey"), "hey")
	assert.Equal(t, valueOrDefault("hey", "hey thar"), "hey")
}

func TestOidForCode(t *testing.T) {
	valueSets, _ := ioutil.ReadFile("../fixtures/value_sets/cms9_26.json")
	vs := []models.ValueSet{}
	json.Unmarshal(valueSets, &vs)
	initializeVsMap(vs)
	coded := models.CodedConcept{Code: "3950001", CodeSystem: "2.16.840.1.113883.6.96"}
	coded2 := models.CodedConcept{Code: "3950001222", CodeSystem: "2.16.840.1.113883.6.96"}
	vsoids := []string{"2.16.840.1.113883.3.117.1.7.1.70", "2.16.840.1.113883.3.117.1.7.1.27", "2.16.840.1.113883.3.117.1.7.1.26", "2.16.840.1.113883.3.117.1.7.1.25"}

	assert.Equal(t, oidForCode(coded, vsoids), "2.16.840.1.113883.3.117.1.7.1.70")
	assert.Equal(t, oidForCode(coded2, vsoids), "")
}

func TestIdentifierForString(t *testing.T) {
	assert.Equal(t, "ACBD18DB4CC2F85CEDEF654FCCC4A4D8", identifierForString("foo"))
}

func TestIdentifierForInt(t *testing.T) {
	assert.Equal(t, "0C34B280850AF1B31ED2973D71ED43DA", identifierForInt(42))
}

func TestTimeToFormat(t *testing.T) {
	assert.Equal(t, "20160101-700", timeToFormat(1451606400, "20060102-700"))
}

func TestNegationIndicator(t *testing.T) {
	fals := false
	assert.Equal(t, "", negationIndicator(models.Entry{NegationInd: &fals}))
	tru := true
	assert.Equal(t, "negationInd='true'", negationIndicator(models.Entry{NegationInd: &tru}))
}

func TestIdentifierForInterface(t *testing.T) {
	str1 := "first string"
	str2 := "second string"
	myInt := int64(5)
	assert.NotEqual(t, identifierForInterface(str1, myInt), identifierForInterface(str2, myInt), "identifiers should not be equal for unequal strings")
}

func TestAppendEntryInfos(t *testing.T) {
	// create entry sections
	entries := make([]models.Entry, 0)
	entries = append(entries, models.Entry{Description: "my description"})
	var entrySections []models.HasEntry
	for _, entry := range entries {
		entrySections = append(entrySections, &models.Encounter{Entry: entry})
	}
	entrySections = append(entrySections, nil) // appendEntryInfos() function should not include nil entry sections

	entryInfos := appendEntryInfos([]entryInfo{}, entrySections, mdc{})
	assert.Equal(t, 1, len(entryInfos))
	assert.Equal(t, "my description", entryInfos[0].EntrySection.GetEntry().Description)
}

func TestAllPreferredCodeSetsIfNeeded(t *testing.T) {
	// code sets without the "*" string
	preferredCodeSets := []string{"one", "two", "three"}
	codeDisplays := make([]models.CodeDisplay, 1)
	codeDisplays[0] = models.CodeDisplay{PreferredCodeSets: preferredCodeSets}
	allPerferredCodeSetsIfNeeded(codeDisplays)
	assert.Equal(t, preferredCodeSets, codeDisplays[0].PreferredCodeSets)

	// all code sets indicated by the "*" string
	preferredCodeSets = append(preferredCodeSets, "*")
	codeDisplays[0] = models.CodeDisplay{PreferredCodeSets: preferredCodeSets}
	allPerferredCodeSetsIfNeeded(codeDisplays)
	assert.Equal(t, true, unorderedStringSlicesEqual(models.CodeSystemNames(), codeDisplays[0].PreferredCodeSets), "preferred code sets should include all code system names")
}

func TestSetCodeDisplaysForEntry(t *testing.T) {
	entry := &models.Entry{Oid: "2.16.840.1.113883.3.560.1.79"} // encounter performed hqmf oid
	mapDataCriteria := mdc{}
	assert.Equal(t, 0, len(entry.CodeDisplays))
	SetCodeDisplaysForEntry(entry, mapDataCriteria)
	assert.Equal(t, 3, len(entry.CodeDisplays)) // three code displays the encounter performed entry
}

// This test needs to be fixed. entryInfosForPatient uses EntriesForDataCriteria to get entries.
// This test will always pass since it uses EntriesForDataCriteria to the expected result.
// The only reason why I haven't deleted it is so that it can be fixed.
//func TestEntryInfosForPatient(t *testing.T) {
//	// patient
//	record, _ := ioutil.ReadFile("../fixtures/records/1_n_n_ami.json")
//	patient := models.Record{}
//	json.Unmarshal(record, &patient)
//
//	// measures
//	measures := []models.Measure{}
//	setMeasures(&measures)
//
//	// calculate expected number of entry infos
//	numEntrySections := 0
//	for _, mdc := range uniqueDataCriteria(allDataCriteria(measures)) {
//		numEntrySections += numNonNil(patient.EntriesForDataCriteria(mdc.DataCriteria))
//	}

//	entryInfos := entryInfosForPatient(patient, measures)
//	assert.Equal(t, numEntrySections, len(entryInfos))
//}

func TestReasonValueSetOid(t *testing.T) {
	valueSets, _ := ioutil.ReadFile("../fixtures/value_sets/cms9_26.json")
	vs := []models.ValueSet{}
	json.Unmarshal(valueSets, &vs)
	initializeVsMap(vs)

	coded := models.CodedConcept{Code: "3950001", CodeSystem: "2.16.840.1.113883.6.96"}
	vsoids := []string{"2.16.840.1.113883.3.117.1.7.1.70", "2.16.840.1.113883.3.117.1.7.1.27", "2.16.840.1.113883.3.117.1.7.1.26", "2.16.840.1.113883.3.117.1.7.1.25"}
	fieldOids := make(map[string][]string)
	fieldOids["REASON"] = vsoids
	assert.Equal(t, "2.16.840.1.113883.3.117.1.7.1.70", reasonValueSetOid(coded, fieldOids))

	fieldOidsNoReason := make(map[string][]string)
	assert.Equal(t, "", reasonValueSetOid(coded, fieldOidsNoReason))

	fieldOidsNoOids := make(map[string][]string)
	fieldOidsNoOids["Reason"] = []string{}
	assert.Equal(t, "", reasonValueSetOid(coded, fieldOidsNoOids))
}

func TestCondAssign(t *testing.T) {
	var first, second int64
	second = 5
	assert.Equal(t, second, condAssign(first, second))
	first = 3
	assert.Equal(t, first, condAssign(first, second))
	assert.Equal(t, second, condAssign(second, first))
}

func TestCodeToDisplay(t *testing.T) {
	codeDisplays := []models.CodeDisplay{models.CodeDisplay{CodeType: "first code type"}, models.CodeDisplay{CodeType: "second code type"}}
	entry := models.Entry{CodeDisplays: codeDisplays}
	encounter := models.Encounter{Entry: entry}

	codeDisplay, err := codeToDisplay(&encounter, "first code type")
	assert.Nil(t, err)
	assert.Equal(t, models.CodeDisplay{CodeType: "first code type"}, codeDisplay)

	codeDisplay, err = codeToDisplay(&encounter, "not a code type")
	assert.NotNil(t, err)
}

func TestCodeDisplayWithPreferredCode(t *testing.T) {
	codeType := "my code type"
	expectedCodeDisplay := models.CodeDisplay{CodeType: codeType, PreferredCodeSets: []string{"codeSetB"}}
	entry := models.Entry{CodeDisplays: []models.CodeDisplay{expectedCodeDisplay}}
	codes := make(map[string][]string)
	codes["codeSetA"] = []string{"third", "fourth"}
	codes["codeSetB"] = []string{"first", "second"}
	coded := models.Coded{Codes: codes}

	actualCodeDisplay := codeDisplayWithPreferredCode(&entry, &coded, codeType)
	expectedCodeDisplay.PreferredCode = models.Concept{Code: "first", CodeSystem: "codeSetB"}
	assert.Equal(t, expectedCodeDisplay, actualCodeDisplay)
}

func TestDischargeDispositionDisplay(t *testing.T) {
	code := "my_code"
	codeSystem := "my_code_system"
	dischargeDisposition := map[string]string{"code": code, "code_system": codeSystem}
	expected := "<sdtc:dischargeDispositionCode code=\"" + code + "\" codeSystem=\"" + codeSystem + "\"/>"
	assert.Equal(t, expected, dischargeDispositionDisplay(dischargeDisposition))

	// code system with a code system oid
	codeSystem = "ICD-10-CM"
	expectedCodeSystemOid := "2.16.840.1.113883.6.90" // oid for code system ICD-10-CM
	dischargeDisposition = map[string]string{"code": code, "code_system": codeSystem}
	expected = "<sdtc:dischargeDispositionCode code=\"" + code + "\" codeSystem=\"" + expectedCodeSystemOid + "\"/>"
	assert.Equal(t, expected, dischargeDispositionDisplay(dischargeDisposition))

	// empty discharge disposition
	dischargeDisposition = map[string]string{}
	expected = "<sdtc:dischargeDispositionCode code=\"\" codeSystem=\"\"/>"
	assert.Equal(t, expected, dischargeDispositionDisplay(dischargeDisposition))
}

func TestSdtcValueSetAttribute(t *testing.T) {
	assert.Equal(t, "", sdtcValueSetAttribute(""))
	assert.Equal(t, "sdtc:valueSet=\"my_value_set\"", sdtcValueSetAttribute("my_value_set"))
}

func TestGetTransferOid(t *testing.T) {
	codeListId := "my code list id"
	fieldValue := models.FieldValue{CodeListID: codeListId}
	dc := models.DataCriteria{FieldValues: map[string]models.FieldValue{"my key": fieldValue}}
	assert.Equal(t, codeListId, getTransferOid(dc, "my key"))

	fieldValue = models.FieldValue{}
	dc = models.DataCriteria{FieldValues: map[string]models.FieldValue{"my key": fieldValue}}
	assert.Equal(t, "", getTransferOid(dc, "my key"))
}

func TestOidForCodeSystem(t *testing.T) {
	codeSystemNames := []string{"HL7 ActStatus", "ICD-10-CM", "Religious Affiliation", "not a code system name"}
	expectedCodeSystemOids := []string{"2.16.840.1.113883.5.14", "2.16.840.1.113883.6.90", "2.16.840.1.113883.5.1076", "not a code system name"}
	for i, codeSystemName := range codeSystemNames {
		assert.Equal(t, expectedCodeSystemOids[i], oidForCodeSystem(codeSystemName))
	}
}

func TestHasPreferredCode(t *testing.T) {
	assert.Equal(t, false, hasPreferredCode(models.Concept{}))
	assert.Equal(t, false, hasPreferredCode(models.Concept{Code: "my code"}))
	assert.Equal(t, false, hasPreferredCode(models.Concept{CodeSystem: "my code system"}))
	assert.Equal(t, true, hasPreferredCode(models.Concept{Code: "my code", CodeSystem: "my code system"}))
}

func TestStringInSlice(t *testing.T) {
	stringSlice := []string{"one", "two", "four"}
	assert.Equal(t, true, stringInSlice("one", stringSlice))
	assert.Equal(t, true, stringInSlice("two", stringSlice))
	assert.Equal(t, false, stringInSlice("three", stringSlice))
}

// - - - - - - - - //
//   H E L P E R   //
// - - - - - - - - //

func setMeasures(measures *[]models.Measure) {
	measureData, err := ioutil.ReadFile("../fixtures/measures/CMS9v4a.json")
	util.CheckErr(err)
	measureData2, err := ioutil.ReadFile("../fixtures/measures/CMS26v3.json")
	util.CheckErr(err)
	measureData = append([]byte("["), append(append(measureData, append([]byte(","), measureData2...)...), []byte("]")...)...)
	err = json.Unmarshal(measureData, measures)
	util.CheckErr(err)
}

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

func numNonNil(objs []models.HasEntry) int {
	var sum int
	for _, obj := range objs {
		if obj != nil {
			sum += 1
		}
	}
	return sum
}
