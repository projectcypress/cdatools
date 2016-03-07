package models

type HQMFDocument struct {
	DataCriteria map[string]DataCriteria `json:"data_criteria"`
}

type DataCriteria struct {
	Title              string `json:"title"`
	Description        string `json:"description"`
	CodeListID         string `json:"code_list_id"`
	Type               string `json:"type"`
	Definition         string `json:"definition"`
	Status             string `json:"status"`
	HardStatus         bool   `json:"hard_status"`
	Negation           bool   `json:"negation"`
	SourceDataCriteria string `json:"source_data_criteria"`
	Variable           bool   `json:"variable"`
}
