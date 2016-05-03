package models

type Transfer struct {
	Coded
	CodedConcept
	Time int64 `json:"time,omitempty"`
}
