package models

import "github.com/projectcypress/cdatools"

type Measure struct {
	ID                string              `json:"id"`
	SubID             string              `json:"sub_id"`
	CmsID             string              `json:"cms_id"`
	Name              string              `json:"name"`
	Description       string              `json:"description"`
	Subtitle          string              `json:"subtitle"`
	ShortSubtitle     string              `json:"short_subtitle"`
	HQMFID            string              `json:"hqmf_id"`
	HQMFSetID         string              `json:"hqmf_set_id"`
	HQMFVersionNumber int                 `json:"hqmf_version_number"`
	NQFID             string              `json:"nqf_id"`
	Type              string              `json:"type"`
	Category          string              `json:"category"`
	PopulationIDs     map[string]string   `json:"population_ids"`
	OIDs              []string            `json:"oids"`
	HQMFDocument      models.HQMFDocument `json:"hqmf_document"`
}
