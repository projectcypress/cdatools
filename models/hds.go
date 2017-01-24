package models

import (
	"encoding/json"
	"fmt"
	"log"
)

// This type should not be created without
// the maps being populated first. This cannot
// be avoided because these maps rely on structs
// defined in models, and therefore cannot be moved
// to another package without a cyclic dependency.
type HdsMaps struct {
	IdMap     map[string]string
	IdR2Map   map[string]string
	HqmfR2Map map[string]DataCriteria
	HqmfMap   map[string]DataCriteria
	// maps qrda oids to hqmf oids
	HqmfQrdaMap map[string]map[string]string
	// maps qrda oids to maps containing code display information
	QrdaCodeDisplayMap map[string]map[string][]CodeDisplay
}

type HqmfQrdaOidsWithCodeDisplays struct {
	HqmfName     string        `json:"hqmf_name,omitempty"`
	HqmfOid      string        `json:"hqmf_oid,omitempty"`
	QrdaName     string        `json:"qrda_name,omitempty"`
	QrdaOid      string        `json:"qrda_oid,omitempty"`
	CodeDisplays []CodeDisplay `json:"code_displays,omitempty"`
}

// NOTE: Had to remove the use of the Asset function that exists in exporter/templates.go:1228
// Now function just reads the json file directly.
func (h *HdsMaps) importHqmfQrdaJSON() {
	var hqmfQrdaOids_r3 []HqmfQrdaOidsWithCodeDisplays
	err := json.Unmarshal(hqmf_qrda_oids, &hqmfQrdaOids_r3)
	if err != nil {
		log.Fatalln(err)
	}
	var hqmfQrdaOids_r3_1 []HqmfQrdaOidsWithCodeDisplays
	err = json.Unmarshal(hqmf_qrda_oids_r3_1, &hqmfQrdaOids_r3_1)
	if err != nil {
		log.Fatalln(err)
	}
	h.QrdaCodeDisplayMap["r3"] = make(map[string][]CodeDisplay)
	for _, oidsElem := range hqmfQrdaOids_r3 {
		h.QrdaCodeDisplayMap["r3"][oidsElem.QrdaOid] = oidsElem.CodeDisplays
	}
	h.QrdaCodeDisplayMap["r3_1"] = make(map[string][]CodeDisplay)
	for _, oidsElem := range hqmfQrdaOids_r3_1 {
		h.QrdaCodeDisplayMap["r3_1"][oidsElem.QrdaOid] = oidsElem.CodeDisplays
	}

	// create hqmfQrdaMap (map) of hqmf oid to map[string]string
	// containing "hqmf_name", "hqmf_oid", "qrda_name", and qrda_oid
	for _, oidsElem := range hqmfQrdaOids_r3 {
		hqmfQrdaMapElem := make(map[string]string)
		hqmfQrdaMapElem["hqmf_name"] = oidsElem.HqmfName
		hqmfQrdaMapElem["hqmf_oid"] = oidsElem.HqmfOid
		hqmfQrdaMapElem["qrda_name"] = oidsElem.QrdaName
		hqmfQrdaMapElem["qrda_oid"] = oidsElem.QrdaOid
		h.HqmfQrdaMap[oidsElem.HqmfOid] = hqmfQrdaMapElem
	}

}

// NOTE: Had to remove the use of the Asset function that exists in exporter/templates.go:1228
// Now function just reads the json file directly.
func (h *HdsMaps) importHQMFTemplateJSON() {
	err := json.Unmarshal(hqmf_template_oid_map, &h.HqmfMap)
	if err != nil {
		log.Fatalln(err)
	}
	for id, data := range h.HqmfMap {
		h.IdMap[makeDefinitionKey(data.Definition, data.Status, data.Negation)] = id
	}
}

func (h *HdsMaps) GetTemplateDefinition(id string) DataCriteria {
	return h.HqmfMap[id]
}

func (h *HdsMaps) GetID(data DataCriteria) string {
	return h.IdMap[makeDefinitionKey(data.Definition, data.Status, data.Negation)]
}

func makeDefinitionKey(definition string, status string, negation bool) string {
	return fmt.Sprintf("%s-%s-%t", definition, status, negation)
}

func (h *HdsMaps) HqmfToQrdaOid(hqmfOid string, vsOid string) string {
	var qrdaOidToReturn string
	if vsOid != "" && hqmfOid == "2.16.840.1.113883.3.560.1.1001" {
		return VsToQrdaOidPatientCharacteristic(vsOid)
	}
	for curHqmfOid, hqmfQrdaMapVal := range h.HqmfQrdaMap {
		if hqmfOid == curHqmfOid {
			if qrdaOidToReturn != "" {
				log.Fatalln("There should only be one QRDA oid for one HQMF oid. If this is hit, there is a flaw in the logic of this code.")
			}
			qrdaOidToReturn = hqmfQrdaMapVal["qrda_oid"]
		}
	}
	return qrdaOidToReturn
}

func (h *HdsMaps) CodeDisplayForQrdaOid(oid string, version string) []CodeDisplay {
	if codeDisplays, ok := h.QrdaCodeDisplayMap[version][oid]; ok {
		return codeDisplays
	} else if codeDisplays, ok := h.QrdaCodeDisplayMap["r3"][oid]; ok {
		return codeDisplays
	}
	return nil
}

// TODO: Add version stuff to here
func (h *HdsMaps) SetCodeDisplaysForEntry(e *Entry, mapDataCriteria Mdc, qrdaVersion string) {
	codeDisplays := h.CodeDisplayForQrdaOid(h.HqmfToQrdaOid(e.Oid, mapDataCriteria.DcKey.ValueSetOid), qrdaVersion)
	allPerferredCodeSetsIfNeeded(codeDisplays)
	for i := range codeDisplays {
		codeDisplays[i].Description = e.Description
	}
	e.CodeDisplays = codeDisplays
}

// adds all code system names to preferred code sets if "*" is present in the existent preferred code sets
func allPerferredCodeSetsIfNeeded(cds []CodeDisplay) {
	for i := range cds {
		if stringInSlice("*", cds[i].PreferredCodeSets) {
			cds[i].PreferredCodeSets = CodeSystemNames()
		}
	}
}

func VsToQrdaOidPatientCharacteristic(vsOid string) string {
	switch vsOid {
	case "2.16.840.1.113883.3.117.1.7.1.402", "2.16.840.1.113883.3.117.1.7.1.403",
		"2.16.840.1.113883.3.117.1.7.1.287", "2.16.840.1.113883.3.117.1.7.1.307":
		// Patient Charasteristic Gestational Age
		return "2.16.840.1.113883.10.20.24.3.101"
	default:
		// Patient Characteristic Observation Assertion template for
		// Patient Characteristic: ECOG Performance Status-Poor
		// Patient Characteristic Tobacco User/Non-User
		// return generic pc observation template for anything not specificly mapped to its own template
		return "2.16.840.1.113883.10.20.24.3.103"
	}
}

func stringInSlice(str string, list []string) bool {
	for _, elem := range list {
		if elem == str {
			return true
		}
	}
	return false
}

var hqmf_qrda_oids = []byte(`[
  {
    "hqmf_name": "Care Goal",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.9",
    "qrda_name": "Care Goal",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.1"
  },
  {
    "hqmf_name": "Communication: From Patient to Provider",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.30",
    "qrda_name": "Communication from Patient to Provider",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.2",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Communication: From Provider to Patient",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.31",
    "qrda_name": "Communication from Provider to Patient",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.3",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Communication: From Provider to Provider",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.29",
    "qrda_name": "Communication from Provider to Provider",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.4",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Device, Adverse Event",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.34",
    "qrda_name": "Device Adverse Event",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.5"
  },
  {
    "hqmf_name": "Device, Allergy",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.35",
    "qrda_name": "Device Allergy",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.6"
  },
  {
    "hqmf_name": "Device, Applied",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.10",
    "qrda_name": "Device Applied",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.7",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "CPT"]
      }
    ]
  },
  {
    "hqmf_name": "Device, Intolerance",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.36",
    "qrda_name": "Device Intolerance",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.8"
  },
  {
    "hqmf_name": "Device, Order",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.37",
    "qrda_name": "Device Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.9",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "CPT"]
      }
    ]
  },
  {
    "hqmf_name": "Device, Recommended",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.80",
    "qrda_name": "Device Recommended",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.10"
  },
  {
    "hqmf_name": "Diagnosis, Active",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.2",
    "qrda_name": "Diagnosis Active",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.11",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "ICD-9-CM", "ICD-10-CM"]
      }
    ]
  },
  {
    "hqmf_name": "Diagnosis, Family History",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.32",
    "qrda_name": "Diagnosis Family History",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.12",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"",
        "preferred_code_sets": ["SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Diagnosis, Inactive",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.23",
    "qrda_name": "Diagnosis Inactive",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.13",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "LOINC"]
      }
    ]
  },
  {
    "hqmf_name": "Diagnosis, Resolved",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.24",
    "qrda_name": "Diagnosis Resolved",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.14",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "CPT"]
      }
    ]
  },
  {
    "hqmf_name": "Diagnostic Study Adverse Event",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.38",
    "qrda_name": "Diagnostic Study Adverse Event",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.15"
  },
  {
    "hqmf_name": "Diagnostic Study, Intolerance",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.39",
    "qrda_name": "Diagnostic Study Intolerance",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.16"
  },
  {
    "hqmf_name": "Diagnostic Study, Order",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.40",
    "qrda_name": "Diagnostic Study Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.17",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC"]
      }
    ]
  },
  {
    "hqmf_name": "Diagnostic Study, Order not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.140",
    "qrda_name": "Diagnostic Study Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.17",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC"]
      }
    ]
  },
  {
    "hqmf_name": "Diagnostic Study, Performed",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.3",
    "qrda_name": "Diagnostic Study Performed",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.18",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "LOINC"]
      },
      {
        "code_type": "resultValue",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "LOINC"]
      }
    ]
  },
  {
    "hqmf_name": "Diagnostic Study, not Performed",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.103",
    "qrda_name": "Diagnostic Study Performed",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.18",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "LOINC"]
      },
      {
        "code_type": "resultValue",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "LOINC"]
      }
    ]
  },
  {
    "hqmf_name": "Diagnostic Study, Recommended",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.40",
    "qrda_name": "Diagnostic Study Recommended",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.19",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC"]
      }
    ]
  },
  {
    "hqmf_name": "Diagnostic Study, Result",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.11",
    "qrda_name": "Diagnostic Study Result",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.20"
  },
  {
    "hqmf_name": "Discharge Medication (Health Record Field Flow Attribute)",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.199",
    "qrda_name": "Discharge Medication - Active Medication",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.105",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm"]
      }
    ]
  },
  {
    "hqmf_name": "Encounter, Active",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.81",
    "qrda_name": "Encounter Active",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.21"
  },
  {
    "hqmf_name": "Encounter, Order",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.83",
    "qrda_name": "Encounter Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.22"
  },
  {
    "hqmf_name": "Encounter, Performed",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.79",
    "qrda_name": "Encounter Performed",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.23",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "CPT"]
      },
      {
        "code_type": "transferFrom",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "CPT"]
      },
      {
        "code_type": "transferTo",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "CPT"]
      }
    ]
  },
  {
    "hqmf_name": "Encounter, Recommended",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.84",
    "qrda_name": "Encounter Recommended",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.24"
  },
  {
    "hqmf_name": "Facility Location (attribute)",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1034",
    "qrda_name": "Facility Location",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.100"
  },
  {
    "hqmf_name": "Functional Status, Order",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.86",
    "qrda_name": "Functional Status Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.25"
  },
  {
    "hqmf_name": "Functional Status, Performed",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.85",
    "qrda_name": "Functional Status Performed",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.26"
  },
  {
    "hqmf_name": "Functional Status, Recommended",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.87",
    "qrda_name": "Functional Status Recommended",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.27"
  },
  {
    "hqmf_name": "Functional Status, Result",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.88",
    "qrda_name": "Functional Status Result",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.28"
  },
  {
    "hqmf_name": "Incision Datetime (attribute)",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1007",
    "qrda_name": "Incision Datetime",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.89"
  },
  {
    "hqmf_name": "Intervention, Adverse Event",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.43",
    "qrda_name": "Intervention Adverse Event",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.29"
  },
  {
    "hqmf_name": "Intervention, Intolerance",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.44",
    "qrda_name": "Intervention Intolerance",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.30"
  },
  {
    "hqmf_name": "Intervention, Order",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.45",
    "qrda_name": "Intervention Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.31",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "HCPCS"]
      }
    ]
  },
  {
    "hqmf_name": "Intervention, Order not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.145",
    "qrda_name": "Intervention Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.31",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "HCPCS"]
      }
    ]
  },
  {
    "hqmf_name": "Intervention, Performed",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.46",
    "qrda_name": "Intervention Performed",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.32",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT"]
      },
      {
        "code_type": "resultValue",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "LOINC"]
      }
    ]
  },
  {
    "hqmf_name": "Intervention, Recommended",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.89",
    "qrda_name": "Intervention Recommended",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.33"
  },
  {
    "hqmf_name": "Intervention, Result",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.47",
    "qrda_name": "Intervention Result",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.34"
  },
  {
    "hqmf_name": "Laboratory Test, Adverse Event",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.48",
    "qrda_name": "Laboratory Test Adverse Event",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.35"
  },
  {
    "hqmf_name": "Laboratory Test, Intolerance",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.49",
    "qrda_name": "Laboratory Test Intolerance",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.36"
  },
  {
    "hqmf_name": "Laboratory Test, Order",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.50",
    "qrda_name": "Laboratory Test Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.37",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC"]
      }
    ]
  },
  {
    "hqmf_name": "Laboratory Test, Order not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.150",
    "qrda_name": "Laboratory Test Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.37",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC"]
      }
    ]
  },
  {
    "hqmf_name": "Laboratory Test, Performed",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.5",
    "qrda_name": "Laboratory Test Performed",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.38",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT"]
      },
      {
        "code_type": "resultValue",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "ICD-9-CM", "ICD-10-CM"]
      }
    ]
  },
  {
    "hqmf_name": "Laboratory Test, Recommended",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.90",
    "qrda_name": "Laboratory Test Recommended",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.39"
  },
  {
    "hqmf_name": "Laboratory Test, Result",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.12",
    "qrda_name": "Laboratory Test Result",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.40"
  },
  {
    "hqmf_name": "Medication, Active",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.13",
    "qrda_name": "Medication Active",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.41",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "RxNorm"]
      },
      {
        "code_type": "medicationDispense",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm"]
      }
    ]
  },
  {
    "hqmf_name": "Medication, Administered",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.14",
    "qrda_name": "Medication Administered",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.42",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm", "CVX", "SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Medication, Adverse Effects",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.7",
    "qrda_name": "Medication Adverse Effect",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.43",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm", "SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Medication, Allergy",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1",
    "qrda_name": "Medication Allergy",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.44",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm", "SNOMED-CT", "CVX"]
      }
    ]
  },
  {
    "hqmf_name": "Medication, Dispensed",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.8",
    "qrda_name": "Medication Dispensed",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.45",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm"]
      }
    ]
  },
  {
    "hqmf_name": "Medication, Intolerance",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.15",
    "qrda_name": "Medication Intolerance",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.46",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm", "SNOMED-CT", "CVX"]
      }
    ]
  },
  {
    "hqmf_name": "Medication, Order",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.17",
    "qrda_name": "Medication Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.47",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm", "SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Patient Care Experience",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.96",
    "qrda_name": "Patient Care Experience",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.48"
  },
  {
    "hqmf_name": "Patient CharacteristicGestational Age",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1001",
    "qrda_name": "Patient Characteristic Gestational Age",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.101"
  },
  {
    "hqmf_name": "Patient Characteristic Clinical Trial Participant",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.401",
    "qrda_name": "Patient Characteristic Clinical Trial Participant",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.51",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "", 
        "preferred_code_sets": ["SNOMED-CT", "LOINC", "ICD-9-CM", "ICD-10-CM"]
      }
    ]
  },
  {
    "hqmf_name": "Patient Characteristic Expired",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.404",
    "qrda_name": "Patient Characteristic Expired",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.54",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "", 
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM"]
      }
    ]
  },
  {
    "hqmf_name": "Patient Characteristic Payer",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.405",
    "qrda_name": "Patient Characteristic Payer",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.55",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\" sdtc:valueSet=\"2.16.840.1.114222.4.11.3591\"", 
        "preferred_code_sets": ["SOP","Source of Payment Typology"]
      }
    ]
  },
  {
    "hqmf_name": "Patient Characteristic, ECOG Performance Status Poor",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1001",
    "qrda_name": "Patient Characteristic Observation Assertion",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.103",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"", 
        "preferred_code_sets": ["*"]
      }
    ]
  },
  {
    "hqmf_name": "Patient Charactersitic, Estimated Date of Conception",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1001",
    "qrda_name": "Patient Characteristic Estimated Date of Conception",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.102"
  },
  {
    "hqmf_name": "Patient Charactersitic, Tobacco user and Non-User",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1001",
    "qrda_name": "Tobacco Use",
    "qrda_oid": "2.16.840.1.113883.10.20.22.4.85",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"", 
        "preferred_code_sets": ["SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Patient Preference (attribute)",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1013",
    "qrda_name": "Patient Preference",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.83"
  },
  {
    "hqmf_name": "Physical Exam, Finding",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.18",
    "qrda_name": "Physical Exam Finding",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.57"
  },
  {
    "hqmf_name": "Physical Exam, Order",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.56",
    "qrda_name": "Physical Exam Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.58"
  },
  {
    "hqmf_name": "Physical Exam, Performed",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.57",
    "qrda_name": "Physical Exam Performed",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.59"
  },
  {
    "hqmf_name": "Physical Exam, Recommended",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.91",
    "qrda_name": "Physical Exam Recommended",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.60"
  },
  {
    "hqmf_name": "Procedure, Adverse Event",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.60",
    "qrda_name": "Procedure Adverse Event",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.61"
  },
  {
    "hqmf_name": "Procedure, Intolerance",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.61",
    "qrda_name": "Procedure Intolerance",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.62",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "", 
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "CPT"]
      }
    ]
  },
  {
    "hqmf_name": "Procedure, Order",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.62",
    "qrda_name": "Procedure Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.63",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "ICD-9-PCS", "ICD-10-PCS"]
      },
      {
        "code_type": "resultValue",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "LOINC"]
      }
    ]
  },
  {
    "hqmf_name": "Procedure, Performed",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.6",
    "qrda_name": "Procedure Performed",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.64",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "CPT", "ICD-9-PCS", "ICD-10-PCS"]
      },
      {
        "code_type": "resultValue",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "LOINC"]
      }
    ]
  },
  {
    "hqmf_name": "Procedure, Recommended",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.92",
    "qrda_name": "Procedure Recommended",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.65"
  },
  {
    "hqmf_name": "Procedure, Result",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.63",
    "qrda_name": "Procedure Result",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.66"
  },
  {
    "hqmf_name": "Provider Care Experience",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.28",
    "qrda_name": "Provider Care Experience",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.67"
  },
  {
    "hqmf_name": "Provider Preference (attribute)",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1014",
    "qrda_name": "Provider Preference",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.84"
  },
  {
    "hqmf_name": "Radiation Dosage (attribute)",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1015",
    "qrda_name": "Radiation Dosage and Duration",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.91"
  },
  {
    "hqmf_name": "Radiation Duration (attribute)",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.10005",
    "qrda_name": "Radiation Dosage and Duration",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.91"
  },
  {
    "hqmf_name": "Reaction (attribute)",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1016.1",
    "qrda_name": "Reaction",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.85"
  },
  {
    "hqmf_name": "Reason (attribute)",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1017.1",
    "qrda_name": "Reason",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.88"
  },
  {
    "hqmf_name": "Result (attribute)",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1019",
    "qrda_name": "Result",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.87"
  },
  {
    "hqmf_name": "Risk Category Assessment",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.21",
    "qrda_name": "Risk Category Assessment",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.69"
  },
  {
    "hqmf_name": "Severity (attribute)",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1021",
    "qrda_name": "Severity Observation",
    "qrda_oid": "2.16.840.1.113883.10.20.22.4.8"
  },
  {
    "hqmf_name": "Status (attribute)",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1022",
    "qrda_name": "Status",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.93"
  },
  {
    "hqmf_name": "Substance, Administered",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.64",
    "qrda_name": "Medication Administered",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.42",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm", "CVX", "SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Substance, Adverse Event",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.65",
    "qrda_name": "Medication Adverse Effect",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.43",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm", "SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Substance, Allergy",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.66",
    "qrda_name": "Medication Allergy",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.44",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm", "SNOMED-CT", "CVX"]
      }
    ]
  },
  {
    "hqmf_name": "Substance, Intolerance",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.67",
    "qrda_name": "Medication Intolerance",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.46",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm", "SNOMED-CT", "CVX"]
      }
    ]
  },
  {
    "hqmf_name": "Substance, Order",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.68",
    "qrda_name": "Medication Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.47",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm", "SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Substance, Recommended",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.93",
    "qrda_name": "Substance Recommended",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.75"
  },
  {
    "hqmf_name": "Symptom, Active",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.69",
    "qrda_name": "Symptom Active",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.76"
  },
  {
    "hqmf_name": "Symptom, Assessed",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.70",
    "qrda_name": "Symptom Assessed",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.77"
  },
  {
    "hqmf_name": "Symptom, Inactive",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.97",
    "qrda_name": "Symptom Inactive",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.78"
  },
  {
    "hqmf_name": "Symptom, Resolved",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.98",
    "qrda_name": "Symptom Resolved",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.79"
  },
  {
    "hqmf_name": "Transfer From (attribute)",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.71",
    "qrda_name": "Transfer From",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.81"
  },
  {
    "hqmf_name": "Transfer To (attribute)",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.72",
    "qrda_name": "Transfer To",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.82"
  },
  {
    "hqmf_name": "Laboratory Test, not Performed",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.105",
    "qrda_name": "Laboratory Test Performed",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.38",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT"]
      },
      {
        "code_type": "resultValue",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "ICD-9-CM", "ICD-10-CM"]
      }
    ]
  },
  {
    "hqmf_name": "Discharge Medication, not Performed",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.200",
    "qrda_name": "Discharge Medication - Active Medication",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.105",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm"]
      }
    ]
  },
  {
    "hqmf_name": "Procedure, not Performed",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.106",
    "qrda_name": "Procedure Performed",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.64",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "CPT", "ICD-9-PCS", "ICD-10-PCS"]
      },
      {
        "code_type": "resultValue",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "LOINC"]
      }
    ]
  },
  {
    "hqmf_name": "Medication, not Administered",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.114",
    "qrda_name": "Medication Administered",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.42",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm", "CVX", "SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Intervention, not Performed",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.146",
    "qrda_name": "Intervention Performed",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.32",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT"]
      },
      {
        "code_type": "resultValue",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "LOINC"]
      }
    ]
  },
  {
    "hqmf_name": "Medication, not Active",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.113",
    "qrda_name": "Medication Active",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.41",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "RxNorm"]
      },
      {
        "code_type": "medicationDispense",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm"]
      }
    ]
  },
  {
    "hqmf_name": "Medication, Order not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.78",
    "qrda_name": "Medication Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.47",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm", "SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Device, Applied not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.110",
    "qrda_name": "Device Applied",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.7",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "CPT"]
      }
    ]
  },
  {
    "hqmf_name": "Device, Order not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.137",
    "qrda_name": "Device Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.9",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "CPT"]
      }
    ]
  },
  {
    "hqmf_name": "Procedure, Order not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.162",
    "qrda_name": "Procedure Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.63",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["LOINC", "SNOMED-CT", "ICD-9-PCS", "ICD-10-PCS"]
      },
      {
        "code_type": "resultValue",
        "tag_name": "value",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "xsi:type=\"CD\"",
        "preferred_code_sets": ["SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "LOINC"]
      }
    ]
  },
  {
    "hqmf_name": "Substance, Administered not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.164",
    "qrda_name": "Medication Administered",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.42",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm", "CVX", "SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Communication: From Provider to Provider not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.129",
    "qrda_name": "Communication from Provider to Provider",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.4",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Communication: From Patient to Provider not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.130",
    "qrda_name": "Communication from Patient to Provider",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.2",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Risk Category Assessment not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.121",
    "qrda_name": "Risk Category Assessment",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.69"
  },
  {
    "hqmf_name": "Communication: From Provider to Patient not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.131",
    "qrda_name": "Communication from Provider to Patient",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.3",
     "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["SNOMED-CT"]
      }
    ]
  },
  {
    "hqmf_name": "Physical Exam, Performed not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.157",
    "qrda_name": "Physical Exam Performed",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.59"
  }
]
`)

var hqmf_qrda_oids_r3_1 = []byte(`[
    {
    "hqmf_name": "Medication, Administered",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.14",
    "qrda_name": "Medication Administered",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.42",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["*"]
      }
    ]
  },
  {
    "hqmf_name": "Substance, Administered",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.64",
    "qrda_name": "Medication Administered",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.42",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["*"]
      }
    ]
  },
  {
    "hqmf_name": "Medication, not Administered",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.114",
    "qrda_name": "Medication Administered",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.42",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["*"]
      }
    ]
  },
  {
    "hqmf_name": "Substance, Administered not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.164",
    "qrda_name": "Medication Administered",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.42",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["*"]
      }
    ]
  },
  {
    "hqmf_name": "Medication, Allergy",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.1",
    "qrda_name": "Medication Allergy",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.44",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["*"]
      }
    ]
  },
  {
    "hqmf_name": "Substance, Allergy",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.66",
    "qrda_name": "Medication Allergy",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.44",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["*"]
      }
    ]
  },
  {
    "hqmf_name": "Medication, Order",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.17",
    "qrda_name": "Medication Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.47",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm"]
      }
    ]
  },
  {
    "hqmf_name": "Substance, Order",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.68",
    "qrda_name": "Medication Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.47",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm"]
      }
    ]
  },
  {
    "hqmf_name": "Medication, Order not done",
    "hqmf_oid": "2.16.840.1.113883.3.560.1.78",
    "qrda_name": "Medication Order",
    "qrda_oid": "2.16.840.1.113883.10.20.24.3.47",
    "code_displays": [
      {
        "code_type": "entryCode",
        "tag_name": "code",
        "attribute": "",
        "exclude_null_flavor": false,
        "extra_content": "",
        "preferred_code_sets": ["RxNorm"]
      }
    ]
  }
]`)

var hqmf_template_oid_map = []byte(`{
  "2.16.840.1.113883.3.560.1.1001":{
      "definition":"patient_characteristic",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.25":{
      "definition":"patient_characteristic_birthdate",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.400":{
      "definition":"patient_characteristic_birthdate",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.401":{
      "definition":"patient_characteristic_clinical_trial_participant",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.402":{
      "definition":"patient_characteristic_gender",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.403":{
      "definition":"patient_characteristic_ethnicity",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.404":{
      "definition":"patient_characteristic_expired",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.405":{
      "definition":"patient_characteristic_payer",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.406":{
      "definition":"patient_characteristic_race",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.4":{
      "definition":"encounter",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.81":{
      "definition":"encounter",
      "status":"active",
      "negation":false},
  "2.16.840.1.113883.3.560.1.79":{
      "definition":"encounter",
      "status":"performed",
      "negation":false},
  "2.16.840.1.113883.3.560.1.83":{
      "definition":"encounter",
      "status":"ordered",
      "negation":false},
  "2.16.840.1.113883.3.560.1.84":{
      "definition":"encounter",
      "status":"recommended",
      "negation":false},
  "2.16.840.1.113883.3.560.1.179":{
      "definition":"encounter",
      "status":"performed",
      "negation":true},
  "2.16.840.1.113883.3.560.1.182":{
      "definition":"encounter",
      "status":"performed",
      "negation":true},
  "2.16.840.1.113883.3.560.1.181":{
      "definition":"encounter",
      "status":"active",
      "negation":true},
  "2.16.840.1.113883.3.560.1.183":{
      "definition":"encounter",
      "status":"ordered",
      "negation":true},
  "2.16.840.1.113883.3.560.1.184":{
      "definition":"encounter",
      "status":"recommended",
      "negation":true},
  "2.16.840.1.113883.3.560.1.104":{
      "definition":"encounter",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.6":{
      "definition":"procedure",
      "status":"performed",
      "negation":false},
  "2.16.840.1.113883.3.560.1.62":{
      "definition":"procedure",
      "status":"ordered",
      "negation":false},
  "2.16.840.1.113883.3.560.1.63":{
      "definition":"procedure_result",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.60":{
      "definition":"procedure_adverse_event",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.61":{
      "definition":"procedure_intolerance",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.92":{
      "definition":"procedure",
      "status":"recommended",
      "negation":false},
  "2.16.840.1.113883.3.560.1.162":{
      "definition":"procedure",
      "status":"ordered",
      "negation":true},
  "2.16.840.1.113883.3.560.1.163":{
      "definition":"procedure_result",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.160":{
      "definition":"procedure_adverse_event",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.161":{
      "definition":"procedure_intolerance",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.106":{
      "definition":"procedure",
      "status":"performed",
      "negation":true},
  "2.16.840.1.113883.3.560.1.192":{
      "definition":"procedure",
      "status":"recommended",
      "negation":true},
  "2.16.840.1.113883.3.560.1.2":{
      "definition":"diagnosis",
      "status":"active",
      "negation":false},
  "2.16.840.1.113883.3.560.1.24":{
      "definition":"diagnosis",
      "status":"resolved",
      "negation":false},
  "2.16.840.1.113883.3.560.1.32":{
      "definition":"diagnosis",
      "status":"family_history",
      "negation":false},
  "2.16.840.1.113883.3.560.1.23":{
      "definition":"diagnosis",
      "status":"inactive",
      "negation":false},
  "2.16.840.1.113883.3.560.1.33":{
      "definition":"diagnosis_risk_of",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.124":{
      "definition":"diagnosis",
      "status":"resolved",
      "negation":true},
  "2.16.840.1.113883.3.560.1.132":{
      "definition":"diagnosis",
      "status":"family_history",
      "negation":true},
  "2.16.840.1.113883.3.560.1.123":{
      "definition":"diagnosis",
      "status":"inactive",
      "negation":true},
  "2.16.840.1.113883.3.560.1.102":{
      "definition":"diagnosis",
      "status":"active",
      "negation":true},
  "2.16.840.1.113883.3.560.1.133":{
      "definition":"diagnosis_risk_of",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.3":{
      "definition":"diagnostic_study",
      "status":"performed",
      "negation":false},
  "2.16.840.1.113883.3.560.1.11":{
      "definition":"diagnostic_study_result",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.38":{
      "definition":"diagnostic_study_adverse_event",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.39":{
      "definition":"diagnostic_study_intolerance",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.40":{
      "definition":"diagnostic_study",
      "status":"ordered",
      "negation":false},
  "2.16.840.1.113883.3.560.1.103":{
      "definition":"diagnostic_study",
      "status":"performed",
      "negation":true},
  "2.16.840.1.113883.3.560.1.138":{
      "definition":"diagnostic_study_adverse_event",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.139":{
      "definition":"diagnostic_study_intolerance",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.140":{
      "definition":"diagnostic_study",
      "status":"ordered",
      "negation":true},
  "2.16.840.1.113883.3.560.1.111":{
      "definition":"diagnostic_study_result",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.8":{
      "definition":"medication",
      "status":"dispensed",
      "negation":false},
  "2.16.840.1.113883.3.560.1.17":{
      "definition":"medication",
      "status":"ordered",
      "negation":false},
  "2.16.840.1.113883.3.560.1.199":{
      "definition":"medication",
      "status":"discharge",
      "negation":false},
  "2.16.840.1.113883.3.560.1.200":{
      "definition":"medication",
      "status":"discharge",
      "negation":true},
  "2.16.840.1.113883.3.560.1.13":{
      "definition":"medication",
      "status":"active",
      "negation":false},
  "2.16.840.1.113883.3.560.1.14":{
      "definition":"medication",
      "status":"administered",
      "negation":false},
  "2.16.840.1.113883.3.560.1.7":{
      "definition":"medication_adverse_effects",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.1":{
      "definition":"medication_allergy",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.15":{
      "definition":"medication_intolerance",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.77":{
      "definition":"medication",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.78":{
      "definition":"medication",
      "status":"ordered",
      "negation":true},
  "2.16.840.1.113883.3.560.1.108":{
      "definition":"medication",
      "status":"dispensed",
      "negation":true},
  "2.16.840.1.113883.3.560.1.113":{
      "definition":"medication",
      "status":"active",
      "negation":true},
  "2.16.840.1.113883.3.560.1.107":{
      "definition":"medication_adverse_effects",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.101":{
      "definition":"medication_allergy",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.115":{
      "definition":"medication_intolerance",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.114":{
      "definition":"medication",
      "status":"administered",
      "negation":true},
  "2.16.840.1.113883.3.560.1.18":{
      "definition":"physical_exam",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.56":{
      "definition":"physical_exam",
      "status":"ordered",
      "negation":false},
  "2.16.840.1.113883.3.560.1.57":{
      "definition":"physical_exam",
      "status":"performed",
      "negation":false},
  "2.16.840.1.113883.3.560.1.91":{
      "definition":"physical_exam",
      "status":"recommended",
      "negation":false},
  "2.16.840.1.113883.3.560.1.156":{
      "definition":"physical_exam",
      "status":"ordered",
      "negation":true},
  "2.16.840.1.113883.3.560.1.157":{
      "definition":"physical_exam",
      "status":"performed",
      "negation":true},
  "2.16.840.1.113883.3.560.1.191":{
      "definition":"physical_exam",
      "status":"recommended",
      "negation":true},
  "2.16.840.1.113883.3.560.1.118":{
      "definition":"physical_exam",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.12":{
      "definition":"laboratory_test",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.5":{
      "definition":"laboratory_test",
      "status":"performed",
      "negation":false},
  "2.16.840.1.113883.3.560.1.48":{
      "definition":"laboratory_test_adverse_event",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.49":{
      "definition":"laboratory_test_intolerance",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.50":{
      "definition":"laboratory_test",
      "status":"ordered",
      "negation":false},
  "2.16.840.1.113883.3.560.1.90":{
      "definition":"laboratory_test",
      "status":"recommended",
      "negation":false},
  "2.16.840.1.113883.3.560.1.112":{
      "definition":"laboratory_test",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.148":{
      "definition":"laboratory_test_adverse_event",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.149":{
      "definition":"laboratory_test_intolerance",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.150":{
      "definition":"laboratory_test",
      "status":"ordered",
      "negation":true},
  "2.16.840.1.113883.3.560.1.190":{
      "definition":"laboratory_test",
      "status":"recommended",
      "negation":true},
  "2.16.840.1.113883.3.560.1.105":{
      "definition":"laboratory_test",
      "status":"performed",
      "negation":true},
  "2.16.840.1.113883.3.560.1.9":{
      "definition":"care_goal",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.109":{
      "definition":"care_goal",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.30":{
      "definition":"communication_from_patient_to_provider",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.31":{
      "definition":"communication_from_provider_to_patient",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.29":{
      "definition":"communication_from_provider_to_provider",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.130":{
      "definition":"communication_from_patient_to_provider",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.131":{
      "definition":"communication_from_provider_to_patient",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.129":{
      "definition":"communication_from_provider_to_provider",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.10":{
      "definition":"device",
      "status":"applied",
      "negation":false},
  "2.16.840.1.113883.3.560.1.34":{
      "definition":"device_adverse_event",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.35":{
      "definition":"device_allergy",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.36":{
      "definition":"device_intolerance",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.37":{
      "definition":"device",
      "status":"ordered",
      "negation":false},
  "2.16.840.1.113883.3.560.1.80":{
      "definition":"device",
      "status":"recommended",
      "negation":false},
  "2.16.840.1.113883.3.560.1.134":{
      "definition":"device_adverse_event",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.135":{
      "definition":"device_allergy",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.136":{
      "definition":"device_intolerance",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.137":{
      "definition":"device",
      "status":"ordered",
      "negation":true},
  "2.16.840.1.113883.3.560.1.180":{
      "definition":"device",
      "status":"recommended",
      "negation":true},
  "2.16.840.1.113883.3.560.1.110":{
      "definition":"device",
      "status":"applied",
      "negation":true},
  "2.16.840.1.113883.3.560.1.64":{
      "definition":"substance",
      "status":"administered",
      "negation":false},
  "2.16.840.1.113883.3.560.1.68":{
      "definition":"substance",
      "status":"ordered",
      "negation":false},
  "2.16.840.1.113883.3.560.1.65":{
      "definition":"substance_adverse_event",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.67":{
      "definition":"substance_intolerance",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.66":{
      "definition":"substance_allergy",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.93":{
      "definition":"substance",
      "status":"recommended",
      "negation":false},
  "2.16.840.1.113883.3.560.1.165":{
      "definition":"substance_adverse_event",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.167":{
      "definition":"substance_intolerance",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.164":{
      "definition":"substance",
      "status":"administered",
      "negation":true},
  "2.16.840.1.113883.3.560.1.168":{
      "definition":"substance",
      "status":"ordered",
      "negation":true},
  "2.16.840.1.113883.3.560.1.193":{
      "definition":"substance",
      "status":"recommended",
      "negation":true},
  "2.16.840.1.113883.3.560.1.166":{
      "definition":"substance_allergy",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.43":{
      "definition":"intervention_adverse_event",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.44":{
      "definition":"intervention_intolerance",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.45":{
      "definition":"intervention",
      "status":"ordered",
      "negation":false},
  "2.16.840.1.113883.3.560.1.46":{
      "definition":"intervention",
      "status":"performed",
      "negation":false},
  "2.16.840.1.113883.3.560.1.47":{
      "definition":"intervention_result",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.89":{
      "definition":"intervention",
      "status":"recommended",
      "negation":false},
  "2.16.840.1.113883.3.560.1.143":{
      "definition":"intervention_adverse_event",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.144":{
      "definition":"intervention_intolerance",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.145":{
      "definition":"intervention",
      "status":"ordered",
      "negation":true},
  "2.16.840.1.113883.3.560.1.146":{
      "definition":"intervention",
      "status":"performed",
      "negation":true},
  "2.16.840.1.113883.3.560.1.147":{
      "definition":"intervention_result",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.189":{
      "definition":"intervention",
      "status":"recommended",
      "negation":true},
  "2.16.840.1.113883.3.560.1.69":{
      "definition":"symptom",
      "status":"active",
      "negation":false},
  "2.16.840.1.113883.3.560.1.70":{
      "definition":"symptom",
      "status":"assessed",
      "negation":false},
  "2.16.840.1.113883.3.560.1.97":{
      "definition":"symptom",
      "status":"inactive",
      "negation":false},
  "2.16.840.1.113883.3.560.1.98":{
      "definition":"symptom",
      "status":"resolved",
      "negation":false},
  "2.16.840.1.113883.3.560.1.169":{
      "definition":"symptom",
      "status":"active",
      "negation":true},
  "2.16.840.1.113883.3.560.1.170":{
      "definition":"symptom",
      "status":"assessed",
      "negation":true},
  "2.16.840.1.113883.3.560.1.197":{
      "definition":"symptom",
      "status":"inactive",
      "negation":true},
  "2.16.840.1.113883.3.560.1.198":{
      "definition":"symptom",
      "status":"resolved",
      "negation":true},
  "2.16.840.1.113883.3.560.1.41":{
      "definition":"functional_status",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.85":{
      "definition":"functional_status",
      "status":"performed",
      "negation":false},
  "2.16.840.1.113883.3.560.1.86":{
      "definition":"functional_status",
      "status":"ordered",
      "negation":false},
  "2.16.840.1.113883.3.560.1.87":{
      "definition":"functional_status",
      "status":"recommended",
      "negation":false},
  "2.16.840.1.113883.3.560.1.88":{
      "definition":"functional_status_result",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.185":{
      "definition":"functional_status",
      "status":"performed",
      "negation":true},
  "2.16.840.1.113883.3.560.1.186":{
      "definition":"functional_status",
      "status":"ordered",
      "negation":true},
  "2.16.840.1.113883.3.560.1.187":{
      "definition":"functional_status",
      "status":"recommended",
      "negation":true},
  "2.16.840.1.113883.3.560.1.188":{
      "definition":"functional_status_result",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.141":{
      "definition":"functional_status",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.21":{
      "definition":"risk_category_assessment",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.121":{
      "definition":"risk_category_assessment",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.28":{
      "definition":"provider_care_experience",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.128":{
      "definition":"provider_care_experience",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.96":{
      "definition":"patient_care_experience",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.196":{
      "definition":"patient_care_experience",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.59":{
      "definition":"preference_provider",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.58":{
      "definition":"preference_patient",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.94":{
      "definition":"system_characteristic",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.194":{
      "definition":"system_characteristic",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.95":{
      "definition":"provider_characteristic",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.195":{
      "definition":"provider_characteristic",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.71":{
      "definition":"transfer_from",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.72":{
      "definition":"transfer_to",
      "status":"",
      "negation":false},
  "2.16.840.1.113883.3.560.1.171":{
      "definition":"transfer_from",
      "status":"",
      "negation":true},
  "2.16.840.1.113883.3.560.1.172":{
      "definition":"transfer_to",
      "status":"",
      "negation":true}
}
`)
