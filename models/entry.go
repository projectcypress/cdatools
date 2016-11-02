package models

import (
	"errors"
	"fmt"
)

type Entry struct {
	Coded
	StartTime      int64               `json:"start_time,omitempty"`
	BSONID         string              `json:"bson_id,omitempty"`
	EndTime        int64               `json:"end_time,omitempty"`
	Time           int64               `json:"time,omitempty"`
	ID             CDAIdentifier       `json:"cda_identifier,omitempty"`
	Oid            string              `json:"oid,omitempty"`
	Description    string              `json:"description,omitempty"`
	NegationInd    *bool               `json:"negationInd,omitempty"`
	NegationReason CodedConcept        `json:"negationReason,omitempty"`
	Values         []ResultValue       `json:"values,omitempty"`
	StatusCode     map[string][]string `json:"status_code,omitempty"`
	Reason         CodedConcept        `json:"reason,omitempty"`
	References     []Reference         `json:"references,omitempty"`
	CodeDisplays   []CodeDisplay       `json:"code_displays,omitempty"`
}

// used by exporter template to display a code. ex. (if TagName is priorityCode) <priorityCode code="1234"></priorityCode>
type CodeDisplay struct {
	CodeType               string   `json:"code_type,omitempty"`
	TagName                string   `json:"tag_name,omitempty"`
	Attribute              string   `json:"attribute,omitempty"`
	ExcludeNullFlavor      bool     `json:"exclude_null_flavor,omitempty"`
	ExtraContent           string   `json:"extra_content,omitempty"`
	PreferredCodeSets      []string `json:"preferred_code_sets,omitempty"`
	PreferredCode          Concept  `json:"preferred_code,omitempty"`
	EntryFieldNameForCoded string   `json:"entry_field_name_for_coded"`
	Description            string   `json:"description"`
}

type HasEntry interface {
	GetEntry() *Entry
}

func (entry *Entry) GetEntry() *Entry {
	return entry
}

// returns codeDisplay. also returns true if code display was found and false if not found
func (e *Entry) GetCodeDisplay(codeType string) (CodeDisplay, error) {
	for _, codeDisplay := range e.CodeDisplays {
		if codeDisplay.CodeType == codeType {
			return codeDisplay, nil
		}
	}
	var returnableCodeDisplay CodeDisplay
	return returnableCodeDisplay, errors.New(fmt.Sprintf("code display was not found for code type \"%s\"", codeType))
}

func (e *Entry) AddStringValue(value string, units string) {
	val := ResultValue{}
	val.Scalar = value
	val.Units = units
	e.Values = append(e.Values, val)
}

func (e *Entry) NegationReasonOrReason() CodedConcept {
	if e.NegationReason != (CodedConcept{}) {
		return e.NegationReason
	}
	return e.Reason
}

// In current implementation, this may give unexpected value if Time or StartTime
// are set to zero, and not just as a default
func (e *Entry) AsPointInTime() int64 {
	if e.Time != 0 {
		return e.Time
	} else if e.StartTime != 0 {
		return e.StartTime
	} else {
		return e.EndTime
	}
}

func (e *Entry) IsValuesEmpty() bool {
	return len(e.Values) == 0
}
