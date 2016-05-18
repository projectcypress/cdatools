package exporter

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
)

var hqmfR2Map map[string]models.DataCriteria
var hqmfMap map[string]models.DataCriteria
var idMap map[string]string
var idR2Map map[string]string
var hqmfQrdaMap map[string]map[string]string
var hqmfMapInit sync.Once

func initializeMap() {
	hqmfMapInit.Do(func() {
		importHQMFTemplateJSON()
		importHqmfQrdaJSON()
	})
}

func importHQMFTemplateJSON() {
	data, err := Asset("hqmf_template_oid_map.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(data, &hqmfMap)
	idMap = map[string]string{}
	for id, data := range hqmfMap {
		idMap[makeDefinitionKey(data.Definition, data.Status, data.Negation)] = id
	}
	data, err = Asset("hqmfr2_template_oid_map.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(data, &hqmfR2Map)
	idR2Map = map[string]string{}
	for id, data := range hqmfR2Map {
		idR2Map[makeDefinitionKey(data.Definition, data.Status, data.Negation)] = id
	}
}

func makeDefinitionKey(definition string, status string, negation bool) string {
	return fmt.Sprintf("%s-%s-%t", definition, status, negation)
}

func importHqmfQrdaJSON() {
	data, err := Asset("hqmf_qrda_oids.json")
	if err != nil {
		util.CheckErr(err)
	}

	// unmarshal from "hqmf_qrda_oids.json" to hqmfQrdaOids variable
	var hqmfQrdaOids []map[string]string
	if err := json.Unmarshal(data, &hqmfQrdaOids); err != nil {
		util.CheckErr(err)
	}

	// create hqmfQrdaMap from hqmfQrdaOids variable
	hqmfQrdaMap = map[string]map[string]string{}
	for _, hqmfQrdaOidElement := range hqmfQrdaOids {
		hqmfQrdaMap[hqmfQrdaOidElement["hqmf_oid"]] = hqmfQrdaOidElement
	}
}

func GetTemplateDefinition(id string, r2Compat bool) models.DataCriteria {
	initializeMap()
	if r2Compat {
		return hqmfR2Map[id]
	} else {
		return hqmfMap[id]
	}
}

func GetID(data models.DataCriteria, r2Compat bool) string {
	initializeMap()
	if r2Compat {
		return idR2Map[makeDefinitionKey(data.Definition, data.Status, data.Negation)]
	} else {
		return idMap[makeDefinitionKey(data.Definition, data.Status, data.Negation)]
	}
}

func GetMap(r2Compat bool) map[string]models.DataCriteria {
	initializeMap()
	if r2Compat {
		return hqmfR2Map
	} else {
		return hqmfMap
	}
}

func HqmfToQrdaOid(hqmfOid string) string {
	initializeMap()
	var qrdaOidToReturn string
	for curHqmfOid, hqmfQrdaMapVal := range hqmfQrdaMap {
		if hqmfOid == curHqmfOid {
			if qrdaOidToReturn != "" {
				panic("There should only be one QRDA oid for one HQMF oid. If this is hit, there is a flaw in the logic of this code.")
			}
			qrdaOidToReturn = hqmfQrdaMapVal["qrda_oid"]
		}
	}
	return qrdaOidToReturn
}

func IsR2Compatible(i interface{}) bool {
	initializeMap()
	entry := models.ExtractEntry(i)
	return hqmfQrdaMap[entry.Oid] != nil
}
