package exporter

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/projectcypress/cdatools/models"
)

var hqmfMap map[string]models.DataCriteria
var idMap map[string]string
var hqmfMapInit sync.Once

func initializeMap() {
	hqmfMapInit.Do(func() {
		importHQMFTemplateJSON()
	})
}

func importHQMFTemplateJSON() {
	data, err := Asset("hqmfr2_template_oid_map.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(data, &hqmfMap)
	idMap = map[string]string{}
	for id, data := range hqmfMap {
		idMap[makeDefinitionKey(data.Definition, data.Status, data.Negation)] = id
	}
}

func makeDefinitionKey(definition string, status string, negation bool) string {
	return fmt.Sprintf("%s-%s-%t", definition, status, negation)
}

func GetTemplateDefinition(id string) models.DataCriteria {
	initializeMap()
	return hqmfMap[id]
}

func GetID(data models.DataCriteria) string {
	initializeMap()
	return idMap[makeDefinitionKey(data.Definition, data.Status, data.Negation)]
}

func GetMap() map[string]models.DataCriteria {
	initializeMap()
	return hqmfMap
}
