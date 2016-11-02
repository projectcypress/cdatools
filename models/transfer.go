package models

type Transfer struct {
	Coded
	CodedConcept
	Time      int64 `json:"time,omitempty"`
	StartTime int64 `json:"start_time,omitempty"`
}
