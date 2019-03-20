package models

import (
	"log"
)

var hds *HdsMaps

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	hds = &HdsMaps{
		IdMap:              make(map[string]string),
		IdR2Map:            make(map[string]string),
		HqmfR2Map:          make(map[string]DataCriteria),
		HqmfMap:            make(map[string]DataCriteria),
		HqmfQrdaMap:        make(map[string]map[string]string),
		QrdaCodeDisplayMap: make(map[string]map[string][]CodeDisplay),
	}
	hds.importHqmfQrdaJSON()
	hds.importHQMFTemplateJSON()
}

// NewHds returns an Hds with properly populated maps.
func NewHds() *HdsMaps {
	return hds
}
