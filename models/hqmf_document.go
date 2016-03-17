package models

//HQMFDocument is a type that holds the Golang representation of an HQMF measure
type HQMFDocument struct {
	DataCriteria map[string]DataCriteria `json:"data_criteria"`
}

type DataCriteria struct {
	Title              string                `json:"title"`
	Description        string                `json:"description"`
	CodeListID         string                `json:"code_list_id"`
	Type               string                `json:"type"`
	Definition         string                `json:"definition"`
	Status             string                `json:"status"`
	HardStatus         bool                  `json:"hard_status"`
	Negation           bool                  `json:"negation"`
	SourceDataCriteria string                `json:"source_data_criteria"`
	Variable           bool                  `json:"variable"`
	FieldValues        map[string]FieldValue `json:"field_values"`
	Value              MetaValue             `json:"value"`
}

type Range struct {
	Low   Val `json:"low"`
	High  Val `json:"high"`
	Width Val `json:"width"`
}

type Coded struct {
	System     string `json:"system"`
	Code       string `json:"code"`
	CodeListID string `json:"code_list_id"`
}

type Val struct {
	Unit       string `json:"unit"`
	Value      string `json:"value"`
	Inclusive  bool   `json:"inclusive"`
	Derived    bool   `json:"derived"`
	Expression string `json:"expression"`
}

type MetaValue struct {
	Type string `json:"type"`
	Val
	Coded
	Range
}

type FieldValue struct {
	Type       string          `json:"type"`
	CodeListID string          `json:"code_list_id"`
	Title      string          `json:"title"`
	High       FieldValueValue `json:"high"`
	Low        FieldValueValue `json:"low"`
}

type FieldValueValue struct {
	Type      string `json:"type"`
	Unit      string `json:"unit"`
	Value     string `json:"value"`
	Inclusive bool   `json:"inclusive?"`
	Derived   bool   `json:"derived?"`
}
