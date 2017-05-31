package models

var oidMap = map[string]string{
	"2.16.840.1.113883.6.1":           "LOINC",
	"2.16.840.1.113883.6.96":          "SNOMED-CT",
	"2.16.840.1.113883.6.12":          "CPT",
	"2.16.840.1.113883.6.88":          "RxNorm",
	"2.16.840.1.113883.6.103":         "ICD-9-CM",
	"2.16.840.1.113883.6.104":         "ICD-9-PCS",
	"2.16.840.1.113883.6.4":           "ICD-10-PCS",
	"2.16.840.1.113883.6.90":          "ICD-10-CM",
	"2.16.840.1.113883.6.14":          "HCP",
	"2.16.840.1.113883.6.285":         "HCPCS",
	"2.16.840.1.113883.5.2":           "HL7 Marital Status",
	"2.16.840.1.113883.12.292":        "CVX",
	"2.16.840.1.113883.5.83":          "HITSP C80 Observation Status",
	"2.16.840.1.113883.3.26.1.1":      "NCI Thesaurus",
	"2.16.840.1.113883.3.88.12.80.20": "FDA",
	"2.16.840.1.113883.4.9":           "UNII",
	"2.16.840.1.113883.6.69":          "NDC",
	"2.16.840.1.113883.5.14":          "HL7 ActStatus",
	"2.16.840.1.113883.6.259":         "HL7 Healthcare Service Location",
	"2.16.840.1.113883.12.112":        "DischargeDisposition",
	"2.16.840.1.113883.5.4":           "ActCode",
	"2.16.840.1.113883.1.11.18877":    "HL7 Relationship Code",
	"2.16.840.1.113883.6.238":         "CDC Race",
	"2.16.840.1.113883.6.177":         "NLM MeSH",
	"2.16.840.1.113883.5.1076":        "Religious Affiliation",
	"2.16.840.1.113883.1.11.19717":    "HL7 ActNoImmunicationReason",
	"2.16.840.1.113883.3.88.12.80.33": "NUBC",
	"2.16.840.1.113883.1.11.78":       "HL7 Observation Interpretation",
	"2.16.840.1.113883.3.221.5":       "Source of Payment Typology",
	"2.16.840.1.113883.6.13":          "CDT",
	"2.16.840.1.113883.18.2":          "AdministrativeSex"}

var codeSystemAliases = map[string]string{
	"FDA SPL": "NCI Thesaurus",
	"HSLOC":   "HL7 Healthcare Service Location",
	"SOP":     "Source of Payment Typology"}

var oidAliases = map[string]string{
	"2.16.840.1.113883.6.59": "2.16.840.1.113883.12.292"}

func CodeSystemFor(oid string) string {
	var oidString = oid
	var alias, ok = oidAliases[oid]
	if ok {
		oidString = alias
	}
	return oidMap[oidString]
}

func invertMap(toInvert map[string]string) map[string]string {
	var inverted = map[string]string{}
	for k, v := range toInvert {
		inverted[v] = k
	}
	return inverted
}

func OidForCodeSystem(codeSystem string) string {
	var cs = codeSystem
	var alias, ok = codeSystemAliases[codeSystem]
	if ok {
		cs = alias
	}
	var inverted = invertMap(oidMap)
	return inverted[cs]
}

func CodeSystemNames() []string {
	keys := make([]string, 0, len(oidMap)+len(codeSystemAliases))
	for _, v := range oidMap {
		keys = append(keys, v)
	}
	for _, v := range codeSystemAliases {
		keys = append(keys, v)
	}
	return keys
}

func CodeSystems() map[string]string {
	return oidMap
}
