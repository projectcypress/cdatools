package models

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
)

type Entry struct {
	Coded
	StartTime        *int64              `json:"start_time,omitempty"`
	BSONID           string              `json:"bson_id,omitempty"`
	EndTime          *int64              `json:"end_time,omitempty"`
	Time             *int64              `json:"time,omitempty"`
	ID               CDAIdentifier       `json:"cda_identifier,omitempty"`
	Oid              string              `json:"oid,omitempty"`
	ObjectIdentifier ObjectIdentifier    `json:"_id,omitempty"`
	Description      string              `json:"description,omitempty"`
	NegationInd      *bool               `json:"negationInd,omitempty"`
	NegationReason   CodedConcept        `json:"negationReason,omitempty"`
	Values           []ResultValue       `json:"values,omitempty"`
	StatusCode       map[string][]string `json:"status_code,omitempty"`
	Reason           CodedConcept        `json:"reason,omitempty"`
	References       []Reference         `json:"references,omitempty"`
	CodeDisplays     []CodeDisplay       `json:"code_displays,omitempty"`
}

// Reference is a link from one entry to another, used in "fulfills" among others
type Reference struct {
	Type           string `json:"type,omitempty"`
	ReferencedType string `json:"referenced_type,omitempty"`
	ReferencedID   string `json:"referenced_id,omitempty"`
	ExportedRef    string `json:"exported_ref,omitempty"`
}

// used by exporter template to display a code. ex. (if TagName is priorityCode) <priorityCode code="1234"></priorityCode>
type CodeDisplay struct {
	CodeType               string     `json:"code_type,omitempty"`
	TagName                string     `json:"tag_name,omitempty"`
	Attribute              string     `json:"attribute,omitempty"`
	ExcludeNullFlavor      bool       `json:"exclude_null_flavor,omitempty"`
	ExtraContent           string     `json:"extra_content,omitempty"`
	PreferredCodeSets      []string   `json:"preferred_code_sets,omitempty"`
	CodeSetRequired        bool       `json:"code_set_required,omitempty"`
	ValueSetPreferred      bool       `json:"value_set_preferred,omitempty"`
	PreferredCode          Concept    `json:"preferred_code,omitempty"`
	EntryFieldNameForCoded string     `json:"entry_field_name_for_coded"`
	Description            string     `json:"description"`
	Laterality             Laterality `json:"laterality,omitempty"`
	Translations           []Concept
	MapDataCriteria        Mdc
}

// Used to uniquely identify an entry
type ObjectIdentifier struct {
	ID string `json:"$oid,omitempty"`
}

type HasEntry interface {
	GetEntry() *Entry
}

func (entry *Entry) GetEntry() *Entry {
	return entry
}

// GetCodeDisplay returns codeDisplay. also returns true if code display was found and false if not found
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

func (e *Entry) AsPointInTime() *int64 {
	if e.Time != nil {
		return e.Time
	} else if e.StartTime != nil {
		return e.StartTime
	} else {
		return e.EndTime
	}
}

func (e *Entry) IsValuesEmpty() bool {
	return len(e.Values) == 0
}

func (e *Entry) HasReason() bool {
	return e.NegationReason != (CodedConcept{}) || e.Reason != (CodedConcept{})
}

func (e *Entry) WrapResultValue(val ResultValue, MapDataCriteria Mdc) ResultValueWrap {
	return ResultValueWrap{ResultValueEntry: e, ResultValueMdc: MapDataCriteria, Values: []ResultValue{val}}
}

func (e *Entry) WrapResultValues(vals []ResultValue, MapDataCriteria Mdc) ResultValueWrap {
	return ResultValueWrap{ResultValueEntry: e, ResultValueMdc: MapDataCriteria, Values: vals}
}

func (e *Entry) NonEmptyResultValueArray() []ResultValue {
	if len(e.Values) > 0 {
		return e.Values
	} else {
		return []ResultValue{ResultValue{}}
	}
}

func (c CodeDisplay) RenderExtraContent() string {
	var tmplOut bytes.Buffer
	tmpl, err := template.New("renderExtraContent").Parse(c.ExtraContent)
	if err != nil {
		log.Fatalln(err)
	}
	err = tmpl.Execute(&tmplOut, c)
	if err != nil {
		log.Fatalln(err)
	}
	return tmplOut.String()
}
