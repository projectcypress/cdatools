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
	parsedTime := time.Unix(t, 0)
	return escapeString(parsedTime.Format(f))
}

// IdentifierFor MD5s a byte slice, and returns a String
func identifierFor(b []byte) string {
	md := md5.Sum(b)
	return escapeString(strings.ToUpper(hex.EncodeToString(md[:])))
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

func entriesForPatient(patient models.Record, measures []models.Measure) []interface{} {
	udcs := uniqueDataCriteria(allDataCriteria(measures))
	var entries []interface{}
	for _, udc := range udcs {
		entries = append(entries, entriesForDataCriteria(udc.DataCriteria, patient)...)
	}
	var retEntries []interface{}
	for _, entry := range entries {
		if entry != nil {
			retEntries = append(retEntries, entry)
		}
	}
	return retEntries
}

// git blame schreiber
// returns a function for executing a template based on the qrda oid
//   this is done so we have access to cat1Template when calling this function from _patient_data.xml
func generateExecuteTemplateForEntry(cat1Template *template.Template) func(interface{}) string {
	return func(e interface{}) string {
		entry := models.ExtractEntry(e)
		qrdaOid := HqmfToQrdaOid(entry.Oid)

		templateName := fmt.Sprintf("_%v.xml", qrdaOid)
		var b bytes.Buffer
		if err := cat1Template.ExecuteTemplate(&b, templateName, entry); err != nil {
			panic(err)
		}

		return b.String()
	}
}

func negationIndicator(entry models.Entry) string {
	if entry.NegationInd {
		return "negationInd='true'"
	}
	return ""
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

// def oid_for_code(codedValue, valueset_oids,  valueset)
// 	return nil if codedValue.nil?
// 	valueset_oids ||=[]
// 	code = codedValue["code"]
// 	code_system = codedValue["code_set"] || codedValue["code_system"]
// 	vs_map = (value_set_map(bundle_id) || {})
// 	valueset_oids.each do |vs_oid|
// 		oid_list = (vs_map[vs_oid] || [])
// 		oid_map = Hash[oid_list.collect{|x| [x["set"],x["values"]]}]
// 		if (oid_map[code_system] || []).index code
// 			return vs_oid
// 		end
// 	end
// 	return nil
// end

func valueOrNullFlavor(i interface{}) string {
	var s string
	switch str := i.(type) {
	case string:
		ival, err := strconv.Atoi(str)
		if err == nil {
			var t = time.Unix(int64(ival), 0)
			s = fmt.Sprintf("value='%s'", t.Format("200601021504"))
		}
	case int64:
		var t = time.Unix(str, 0)
		s = fmt.Sprintf("value='%s'", t.Format("200601021504"))
	case int:
		var t = time.Unix(int64(str), 0)
		s = fmt.Sprintf("value='%s'", t.Format("200601021504"))
	default:
		s = "nullFlavor='UNK'"
	}
	return s
}

func valueOrDefault(val interface{}, def interface{}) interface{} {
	if val != nil {
		return val
	}
	return def
}

func codeDisplay(entry models.Entry, options map[string]interface{}) string {
	tagName := valueOrDefault(options["tag_name"], "code")
	attribute := valueOrDefault(options["attribute"], "codes")
	excludeNullFlavor := valueOrDefault(options["exclude_null_flavor"], false)
	extraContent := valueOrDefault(options["extra_content"], "")
	var codeString string
	var pcs []string
	if options["preferred_code_sets"] == "*" {
		pcs = models.CodeSystemNames()
	} else {
		pcs = options["preferred_code_sets"].([]string)
	}

	if pcs != nil {
	}
	// need to replace this with actual call to entry.perferredCode once implmented
	preferredCode := map[string]string{"code": "1234",
		"code_set": "SNOMED-CT"}
	if preferredCode != nil {
		oid := models.OidForCodeSystem(preferredCode["code_set"])
		codeString = fmt.Sprintf("<%s code='%s' codeSystem='%s' %s>", tagName, preferredCode["code"], oid, extraContent)
	} else {
		var buffer bytes.Buffer
		buffer.WriteString(fmt.Sprintf("<%s ", tagName))
		if !excludeNullFlavor.(bool) {
			buffer.WriteString(" nullFlavor='UNK' ")
		}
		buffer.WriteString(extraContent.(string))
		codeString = buffer.String()
	}
	if attribute == "codes" {
		codeString += fmt.Sprintf("<originalText>%s</originalText>", entry.Description)
		// add a bunch of translation codes if they exist
	}

	//       if options["attribute"] == :codes && entry.respond_to?(:translation_codes)
	//         code_string += "<originalText>#{ERB::Util.html_escape entry.description}</originalText>" if entry.respond_to?(:description)
	//         entry.translation_codes(options['preferred_code_sets'], options['value_set_map']).each do |translation|
	//           code_string += "<translation code=\"#{translation['code']}\" codeSystem=\"#{HealthDataStandards::Util::CodeSystemHelper.oid_for_code_system(translation['code_set'])}\"/>\n"
	//         end
	//       end

	return fmt.Sprintf("%s </%s>", codeString, tagName)
}

func isR2Compatable(i interface{}) bool {
	return false
}
