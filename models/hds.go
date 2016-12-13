package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	QrdaCodeDisplayMap map[string][]CodeDisplay
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
	data, err := ioutil.ReadFile("../exporter/hqmf_qrda_oids.json")
	if err != nil {
		log.Fatalln(err)
	}

	var hqmfQrdaOids []HqmfQrdaOidsWithCodeDisplays
	if err := json.Unmarshal(data, &hqmfQrdaOids); err != nil {
		log.Fatalln(err)
	}

	for _, oidsElem := range hqmfQrdaOids {
		h.QrdaCodeDisplayMap[oidsElem.QrdaOid] = oidsElem.CodeDisplays
	}

	// create hqmfQrdaMap (map) of hqmf oid to map[string]string
	// containing "hqmf_name", "hqmf_oid", "qrda_name", and qrda_oid
	for _, oidsElem := range hqmfQrdaOids {
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
	data, err := ioutil.ReadFile("../exporter/hqmf_template_oid_map.json")
	if err != nil {
		log.Fatalf("Failed to read data file to make HdsMaps: %v \n", err)
	}
	json.Unmarshal(data, &h.HqmfMap)
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

func (h *HdsMaps) CodeDisplayForQrdaOid(oid string) []CodeDisplay {
	if codeDisplays, ok := h.QrdaCodeDisplayMap[oid]; ok {
		return codeDisplays
	}
	return nil
}

func (h *HdsMaps) SetCodeDisplaysForEntry(e *Entry, mapDataCriteria Mdc) {
	codeDisplays := h.CodeDisplayForQrdaOid(h.HqmfToQrdaOid(e.Oid, mapDataCriteria.DcKey.ValueSetOid))
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
