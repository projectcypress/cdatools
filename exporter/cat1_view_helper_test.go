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

func numNonNil(objs []models.HasEntry) int {
	var sum int
	for _, obj := range objs {
		if obj != nil {
			sum += 1
		}
	}
	return sum
}
