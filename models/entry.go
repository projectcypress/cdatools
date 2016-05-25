package models

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
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
	NegationInd    bool                `json:"negationInd,omitempty"`
	NegationReason CodedConcept        `json:"negationReason,omitempty"`
	Values         []ResultValue       `bson:"values,omitempty"`
	StatusCode     map[string][]string `json:"status_code,omitempty"`
	Reason         CodedConcept        `json:"reason,omitempty"`
	TransferTo     Transfer            `json:"transferTo,omitempty"`
	TransferFrom   Transfer            `json:"transferFrom,omitempty"`
	References     []Reference         `json:"references,omitempty"`
	CodeDisplays   []CodeDisplay       `json:"code_displays,omitempty"`
}

// used by exporter template to display a code. ex. (if TagName is priorityCode) <priorityCode code="1234"></priorityCode>
type CodeDisplay struct {
	CodeType          string        `json:"code_type,omitempty"`
	TagName           string        `json:"tag_name,omitempty"`
	Attribute         string        `json:"attribute,omitempty"`
	ExcludeNullFlavor bool          `json:"exclude_null_flavor,omitempty"`
	ExtraContent      string        `json:"extra_content,omitempty"`
	PreferredCodeSets []string      `json:"preferred_code_sets,omitempty"`
	PreferredCode     PreferredCode `json:"preferred_code,omitempty"`
}

type PreferredCode struct {
	Code    string
	CodeSet string
}

func ExtractEntry(entry *interface{}) *Entry {

	switch t := (*entry).(type) {
	case Encounter:
		return t.Entry
	case LabResult:
		return &t.Entry
	case InsuranceProvider:
		return &t.Entry
	case Procedure:
		return &t.Entry
	case Allergy:
		return &t.Entry
	case Medication:
		return &t.Entry
	case Communication:
		return &t.Entry
	case Condition:
		return &t.Entry
	case ProviderPerformance:
		return &t.Entry
	case Entry:
		return &t
	default:
		spew.Dump(reflect.TypeOf(entry))
		panic("We don't know how to extract an Entry from this type")
	}
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

func (e *Entry) AddScalarValue(value int64, units string) {
	val := ResultValue{}
	val.Scalar = value
	val.Units = units
	e.Values = append(e.Values, val)
}

func (e *Entry) AddStringValue(value string, units string) {
	val := ResultValue{}
	val.Value = value
	val.Units = units
	e.Values = append(e.Values, val)
}

func (e *Entry) PreferredCode(preferredCodeSets []string) Concept {
	codeTypes := make([]string, len(e.Coded.Codes))
	i := 0
	for k := range e.Coded.Codes {
		codeTypes[i] = k
		i++
	}
	codes := computeIntersection(preferredCodeSets, codeTypes)
	if len(codes) > 0 {
		return Concept{CodeSystem: codes[0], Code: e.Coded.Codes[codes[0]][0]}
	}
	return Concept{}
}
