package models

import "reflect"

type Transfer struct {
	Coded
	CodedConcept
	Time int64 `json:"time,omitempty"`
}

func (t *Transfer) IsEmpty() bool {
	return reflect.DeepEqual(*t, Transfer{})
}
