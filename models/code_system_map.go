package models

var oidMap = map[string]string{
	"2.16.840.1.113883.6.12":     "CPT",
	"2.16.840.1.113883.6.1":      "LOINC",
	"2.16.840.1.113883.6.96":     "SNOMED-CT",
	"2.16.840.1.113883.6.88":     "RxNorm",
	"2.16.840.1.113883.12.292":   "CVX",
	"2.16.840.1.113883.6.103":    "ICD-9-CM",
	"2.16.840.1.113883.6.104":    "ICD-9-PCS",
	"2.16.840.1.113883.6.4":      "ICD-10-PCS",
	"2.16.840.1.113883.6.90":     "ICD-10-CM",
	"2.16.840.1.113883.3.221.5":  "SOP",
	"2.16.840.1.113883.3.26.1.1": "NCI Thesaurus",
	"2.16.840.1.113883.5.4":      "ActCode",
}

// old oids still exist in old data. this maps retired oid values to new oid values
var oidAliases = map[string]string{
	"2.16.840.1.113883.6.59": "2.16.840.1.113883.12.292", // CVX
}

func CodeSystemFor(oid string) string {
	if newOid, ok := oidAliases[oid]; ok {
		oid = newOid
	}
	return oidMap[oid]
}
