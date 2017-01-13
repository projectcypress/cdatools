package exporter

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/projectcypress/cdatools/models"
)

func escape(s string) string {
	b := new(bytes.Buffer)
	err := xml.EscapeText(b, []byte(s))
	if err != nil {
		log.Fatalln(err)
	}
	return b.String()
}

// TimeToFormat parses time from a seconds since Epoch value, and spits out a string in the supplied format
func timeToFormat(t int64, f string) string {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Fatalln(err, "Time Zone Location failed to load")
	}
	parsedTime := time.Unix(t, 0)
	return escape(parsedTime.In(utc).Format(f))
}

// IdentifierFor MD5s a byte slice, and returns a String
func identifierFor(b []byte) string {
	md := md5.Sum(b)
	return escape(strings.ToUpper(hex.EncodeToString(md[:])))
}

func identifierForInterface(objs ...interface{}) string {
	b := make([]byte, len(objs))
	for _, obj := range objs {
		switch otyped := obj.(type) {
		case int64:
			b = append(b, []byte(strconv.FormatInt(otyped, 10))...)
		case *int64:
			if otyped != nil {
				b = append(b, []byte(strconv.FormatInt(*otyped, 10))...)
			}
		case string:
			b = append(b, []byte(otyped)...)
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

// IdentifierForInt generates an MD5 representation of a set of *int64s
func identifierForIntp(objs ...*int64) string {
	b := make([]byte, len(objs))
	for _, val := range objs {
		if val != nil {
			b = append(b, []byte(strconv.FormatInt(*val, 10))...)
		}
	}
	return identifierFor(b)
}

// IdentifierForString generates an MD5 representation of a set of string objects
func identifierForString(objs ...string) string {
	b := strings.Join(objs, ",")
	return identifierFor([]byte(b))
}

// git blame schreiber
// returns a function for executing a template based on the qrda oid
//   this is done so we have access to cat1Template when calling this function from _patient_data.xml
func generateExecuteTemplateForEntry(cat1Template *template.Template) func(models.EntryInfo) string {
	hds := models.NewHds()
	return func(ei models.EntryInfo) string {
		entry := ei.EntrySection.GetEntry()
		qrdaOid := hds.HqmfToQrdaOid(entry.Oid, ei.MapDataCriteria.DcKey.ValueSetOid)

		templateName := fmt.Sprintf("_%v.xml", qrdaOid)
		var b bytes.Buffer
		if err := cat1Template.ExecuteTemplate(&b, templateName, ei); err != nil {
			log.Fatalln(err)
		}

		return b.String()
	}
}

func negationIndicator(entry models.Entry) string {
	if entry.NegationInd != nil && *entry.NegationInd {
		return "negationInd='true'"
	}
	return ""
}

func valueOrNullFlavor(i interface{}) string {
	var s string
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Fatalln(err, "Time Zone Location failed to load")
	}
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
	case *int64:
		if str != nil {
			var t = time.Unix(*str, 0)
			s = fmt.Sprintf("value='%s'", t.In(utc).Format("200601021504-0700"))
		} else {
			s = "nullFlavor='UNK'"
		}
	case *int:
		if str != nil {
			var t = time.Unix(int64(*str), 0)
			s = fmt.Sprintf("value='%s'", t.In(utc).Format("200601021504-0700"))
		} else {
			s = "nullFlavor='UNK'"
		}
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

// conditional assignment. returns the second value only if the first value is empty, zero, or nil
func condAssign(first interface{}, second interface{}) interface{} {
	result := second
	switch first := first.(type) {
	case string:
		if first != "" {
			result = first
		}
	case int64:
		if first != 0 {
			result = first
		}
	case int:
		if first != 0 {
			result = first
		}
	case *int64:
		if first != nil {
			result = first
		}
	case *int:
		if first != nil {
			result = first
		}
	}
	return result
}

func codeDisplayWithPreferredCode(entry *models.Entry, coded *models.Coded, codeType string) models.CodeDisplay {
	codeDisplay, err := entry.GetCodeDisplay(codeType)
	if err != nil {
		log.Fatalln(err)
	}
	codeDisplay.PreferredCode = coded.PreferredCode(codeDisplay.PreferredCodeSets)
	return codeDisplay
}

func codeDisplayWithPreferredCodeAndLaterality(entry *models.Entry, coded *models.Coded, codeType string, laterality models.Laterality, MapDataCriteria models.Mdc) models.CodeDisplay {
	codeDisplay, err := entry.GetCodeDisplay(codeType)
	if err != nil {
		log.Fatal(err)
	}
	codeDisplay.PreferredCode = coded.PreferredCode(codeDisplay.PreferredCodeSets)
	codeDisplay.Laterality = laterality
	codeDisplay.MapDataCriteria = MapDataCriteria
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

func hasPreferredCode(pc models.Concept) bool {
	return pc.Code != "" && pc.CodeSystem != ""
}

func hasLaterality(l models.Laterality) bool {
	return l != models.Laterality{}
}

func codeDisplayAttributeIsCodes(attribute string) bool {
	return attribute == "codes"
}

func isNil(i interface{}) bool {
	return i == nil
}

func derefBool(i *bool) bool {
	return *i
}
