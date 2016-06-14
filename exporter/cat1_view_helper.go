package exporter

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/pebbe/util"

	"github.com/projectcypress/cdatools/models"
)

func escape(i interface{}) string {
	switch str := i.(type) {
	case string:
		return escapeString(str)
	case int64:
		return escapeString(strconv.FormatInt(str, 10))
	case int:
		return escapeString(strconv.Itoa(str))
	}
	return ""
}

func escapeString(s string) string {
	b := new(bytes.Buffer)
	xml.Escape(b, []byte(s))
	return b.String()
}

// TimeToFormat parses time from a seconds since Epoch value, and spits out a string in the supplied format
func timeToFormat(t int64, f string) string {
	utc, err := time.LoadLocation("UTC")
	util.CheckErr(err, "Time Zone Location failed to load")
	parsedTime := time.Unix(t, 0)
	return escapeString(parsedTime.In(utc).Format(f))
}

// IdentifierFor MD5s a byte slice, and returns a String
func identifierFor(b []byte) string {
	md := md5.Sum(b)
	return escapeString(strings.ToUpper(hex.EncodeToString(md[:])))
}

func identifierForInterface(objs ...interface{}) string {
	b := make([]byte, len(objs))
	for _, obj := range objs {
		switch obj.(type) {
		case int64:
			b = append(b, []byte(strconv.FormatInt(obj.(int64), 10))...)
		case string:
			b = append(b, []byte(obj.(string))...)
		}
	}
	return identifierFor(b)
}

// IdentifierForInt generates an MD5 representation of a set of int64s
func identifierForInt(objs ...int64) string {
	b := make([]byte, len(objs))
	for _, val := range objs {
		b = append(b, []byte(strconv.FormatInt(val, 10))...)
	}
	return identifierFor(b)
}

// IdentifierForString generates an MD5 representation of a set of string objects
func identifierForString(objs ...string) string {
	b := strings.Join(objs, ",")
	return identifierFor([]byte(b))
}

// create entryInfos for each entry. entryInfos have mapped data criteria (mdc) recieved from the uniqueDataCriteria() function
// also adds code displays struct to each entry
func entryInfosForPatient(patient models.Record, measures []models.Measure) []entryInfo {
	mappedDataCriterias := uniqueDataCriteria(allDataCriteria(measures))
	var entryInfos []entryInfo
	for _, mappedDataCriteria := range mappedDataCriterias {
		var entrySections []models.HasEntry = entriesForDataCriteria(mappedDataCriteria.DataCriteria, patient)
		// add code displays struct to each entry
		for i, entrySection := range entrySections {
			if entrySection != nil {
				entry := entrySections[i].GetEntry()
				SetCodeDisplaysForEntry(entry)
			}
		}
		entryInfos = appendEntryInfos(entryInfos, entrySections, mappedDataCriteria)
	}
	return entryInfos
}

func SetCodeDisplaysForEntry(e *models.Entry) {
	codeDisplays := codeDisplayForQrdaOid(HqmfToQrdaOid(e.Oid))
	allPerferredCodeSetsIfNeeded(codeDisplays)
	for i, _ := range codeDisplays {
		codeDisplays[i].Description = e.Description
	}
	e.CodeDisplays = codeDisplays
}

// adds all code system names to preferred code sets if "*" is present in the existant preferred code sets
func allPerferredCodeSetsIfNeeded(cds []models.CodeDisplay) {
	for i, _ := range cds {
		if stringInSlice("*", cds[i].PreferredCodeSets) {
			cds[i].PreferredCodeSets = models.CodeSystemNames()
		}
	}
}

// append an entryInfo to entryInfos for each entry
func appendEntryInfos(entryInfos []entryInfo, entries []models.HasEntry, mappedDataCriteria mdc) []entryInfo {
	for _, entry := range entries {
		if entry != nil {
			entryInfo := entryInfo{EntrySection: entry, MapDataCriteria: mappedDataCriteria}
			entryInfos = append(entryInfos, entryInfo)
		}
	}
	return entryInfos
}

// git blame schreiber
// returns a function for executing a template based on the qrda oid
//   this is done so we have access to cat1Template when calling this function from _patient_data.xml
func generateExecuteTemplateForEntry(cat1Template *template.Template) func(entryInfo) string {
	return func(ei entryInfo) string {
		entry := ei.EntrySection.GetEntry()
		qrdaOid := HqmfToQrdaOid(entry.Oid)

		templateName := fmt.Sprintf("_%v.xml", qrdaOid)
		var b bytes.Buffer
		if err := cat1Template.ExecuteTemplate(&b, templateName, ei); err != nil {
			panic(err)
		}

		return b.String()
	}
}

func negationIndicator(entry models.Entry) string {
	if *entry.NegationInd {
		return "negationInd='true'"
	}
	return ""
}

func reasonValueSetOid(codedValue models.CodedConcept, fieldOids map[string][]string) string {
	return oidForCode(codedValue, fieldOids["REASON"])
}

func oidForCode(codedValue models.CodedConcept, valuesetOids []string) string {

	for _, vsoid := range valuesetOids {
		oidlist := vsMap[vsoid]
		if codeSetContainsCode(oidlist, codedValue) {
			return vsoid
		}

	}
	return ""
}

func codeSetContainsCode(sets []models.CodeSet, codedValue models.CodedConcept) bool {
	for _, cs := range sets {
		for _, val := range cs.Values {
			if val.CodeSystem == codedValue.CodeSystem && val.Code == codedValue.Code {
				return true
			}
		}
	}
	return false
}

func valueOrNullFlavor(i interface{}) string {
	var s string
	utc, err := time.LoadLocation("UTC")
	util.CheckErr(err, "Time Zone Location failed to load")
	switch str := i.(type) {
	case string:
		ival, err := strconv.Atoi(str)
		if err == nil {
			var t = time.Unix(int64(ival), 0)
			s = fmt.Sprintf("value='%s'", t.In(utc).Format("200601021504-0700"))
		}
	case int64:
		var t = time.Unix(str, 0)
		s = fmt.Sprintf("value='%s'", t.In(utc).Format("200601021504-0700"))
	case int:
		var t = time.Unix(int64(str), 0)
		s = fmt.Sprintf("value='%s'", t.In(utc).Format("200601021504-0700"))
	default:
		s = "nullFlavor='UNK'"
	}
	return s
}

func valueOrDefault(val interface{}, def interface{}) interface{} {
	switch val.(type) {
	case string:
		if val != "" {
			return val
		}
		return def
	}
	if val != nil {
		return val
	}
	return def
}

// conditional assignment. returns the second value only if the first value is zero
// TODO: make arguments and return type interface{}. add "value is empty, zero, or nil" to description
func condAssign(first int64, second int64) int64 {
	if first != 0 {
		return first
	}
	return second
}

func codeToDisplay(entrySection models.HasEntry, codeType string) (models.CodeDisplay, error) {
	entry := entrySection.GetEntry()
	return entry.GetCodeDisplay(codeType)
}

func codeDisplayWithPreferredCode(entry *models.Entry, coded *models.Coded, codeType string) models.CodeDisplay {
	codeDisplay, err := entry.GetCodeDisplay(codeType)
	util.CheckErr(err)
	codeDisplay.PreferredCode = coded.PreferredCode(codeDisplay.PreferredCodeSets)
	return codeDisplay
}

// dd stands for discharge disposition
func dischargeDispositionDisplay(dd map[string]string) string {
	// set code system
	codeSystem := models.OidForCodeSystem(dd["code_system"])
	if codeSystem == "" {
		codeSystem = dd["code_system"]
	}
	return fmt.Sprintf("<sdtc:dischargeDispositionCode code=\"%s\" codeSystem=\"%s\"/>", dd["code"], codeSystem)
}

func sdtcValueSetAttribute(oid string) string {
	if oid == "" {
		return ""
	}
	return "sdtc:valueSet=\"" + oid + "\""
}

func getTransferOid(dc models.DataCriteria, key string) string {
	if fieldValue := dc.FieldValues[key]; fieldValue != (models.FieldValue{}) {
		return fieldValue.CodeListID
	}
	return ""
}

func oidForCodeSystem(codeSystem string) string {
	oid := models.OidForCodeSystem(codeSystem)
	if oid != "" {
		return oid
	}
	return codeSystem
}

func hasReason(entry models.Entry) bool {
	if entry.NegationReason != (models.CodedConcept{}) || entry.Reason != (models.CodedConcept{}) {
		return true
	}
	return false
}

func hasPreferredCode(pc models.Concept) bool {
	return pc.Code != "" && pc.CodeSystem != ""
}

func codeDisplayAttributeIsCodes(attribute string) bool {
	return attribute == "codes"
}

func stringInSlice(str string, list []string) bool {
	for _, elem := range list {
		if elem == str {
			return true
		}
	}
	return false
}
