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
	NegationCodeListID string                `json:"negation_code_list_id"`
	SourceDataCriteria string                `json:"source_data_criteria"`
	Variable           bool                  `json:"variable"`
	FieldValues        map[string]FieldValue `json:"field_values"`
	Value              MetaValue             `json:"value"`
	HQMFOid            string                `json:"hqmf_oid"`
}

type Range struct {
	Low   Val `json:"low"`
	High  Val `json:"high"`
	Width Val `json:"width"`
}

type Val struct {
	Unit       string `json:"unit"`
	Value      string `json:"value"`
	Inclusive  bool   `json:"inclusive"`
	Derived    bool   `json:"derived"`
	Expression string `json:"expression"`
}

type MetaValue struct {
	Type   string `json:"type,omitempty"`
	System string `json:"system,omitempty"`
	Code   string `json:"code,omitempty"`
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

type DcKey struct {
	DataCriteriaOid string
	ValueSetOid     string
}

type Mdc struct {
	FieldOids    map[string][]string
	ResultOids   []string
	DataCriteria DataCriteria
	DcKey
}

// passed into each qrda oid (entry) template
// EntrySection should be a struct that includes entry attributes (ex. Procedure, Medication, ...)
type EntryInfo struct {
	EntrySection    HasEntry
	MapDataCriteria Mdc
}

func UniqueDataCriteria(allDataCriteria []DataCriteria) []Mdc {
	mappedDataCriteria := map[DcKey]Mdc{}
	for _, dataCriteria := range allDataCriteria {
		// Based on the data criteria, get the HQMF oid associated with it]
		oid := dataCriteria.HQMFOid
		if oid == "" {
			oid = hds.GetID(dataCriteria)
			if oid != "" {
				dataCriteria.HQMFOid = oid
			}
		}
		vsOid := dataCriteria.CodeListID

		// Special cases for the valueSet OID, taken from Health Data Standards
		if oid == "2.16.840.1.113883.3.560.1.71" {
			vsOid = dataCriteria.FieldValues["TRANSFER_FROM"].CodeListID
		} else if oid == "2.16.840.1.113883.3.560.1.72" {
			vsOid = dataCriteria.FieldValues["TRANSFER_TO"].CodeListID
		}

		// Generate the key for the mappedDataCriteria
		dc := DcKey{DataCriteriaOid: oid, ValueSetOid: vsOid}

		var mappedDc = mappedDataCriteria[dc]
		if mappedDc.FieldOids == nil {
			mappedDc = Mdc{DataCriteria: dataCriteria, FieldOids: make(map[string][]string)}
		}

		// Add all the codedValues onto the list of field OIDs
		for field, descr := range dataCriteria.FieldValues {
			if descr.Type == "CD" {
				mappedDc.FieldOids[field] = append(mappedDc.FieldOids[field], descr.CodeListID)
			}
		}

		// If the data criteria has a negation, add the reason onto the returned FieldOids
		if dataCriteria.Negation {
			mappedDc.FieldOids["REASON"] = append(mappedDc.FieldOids["REASON"], dataCriteria.NegationCodeListID)
		}

		// If the data criteria has a value, and it's a "coded" type, added the CodeListId into the result OID set
		if dataCriteria.Value.Type == "CD" {
			mappedDc.ResultOids = append(mappedDc.ResultOids, dataCriteria.CodeListID)
		}

		if dc.DataCriteriaOid != "" {
			mappedDataCriteria[dc] = mappedDc
		}
	}

	// Add the key to the value to get what HDS would have returned
	var retDataCriteria []Mdc
	for key, value := range mappedDataCriteria {
		value.DataCriteriaOid = key.DataCriteriaOid
		value.ValueSetOid = key.ValueSetOid
		retDataCriteria = append(retDataCriteria, value)
	}
	return retDataCriteria
}

// TODO: Needs to be on entry? Something like an entryInfoGroup?
// append an entryInfo to entryInfos for each entry
func AppendEntryInfos(entryInfos []EntryInfo, entries []HasEntry, mappedDataCriteria Mdc) []EntryInfo {
	for _, entry := range entries {
		if entry != nil {
			entryInfo := EntryInfo{EntrySection: entry, MapDataCriteria: mappedDataCriteria}
			entryInfos = append(entryInfos, entryInfo)
		}
	}
	return entryInfos
}
