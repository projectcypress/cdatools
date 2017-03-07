package models

type Person struct {
	Entity
	First     string        `json:"first,omitempty"`
	Last      string        `json:"last,omitempty"`
	Gender    string        `json:"gender,omitempty"`
	Race      *CodedConcept `json:"race,omitempty"`
	Ethnicity *CodedConcept `json:"ethnicity,omitempty"`
}
