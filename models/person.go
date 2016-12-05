package models

type Person struct {
	Entity
	First     string        `json:"first,omitempty"`
	Last      string        `json:"last,omitempty"`
	Birthdate *int64        `json:"birthdate,omitempty"`
	Gender    string        `json:"gender,omitempty"`
	Race      *CodedConcept `json:"race,omitempty"`
	Ethnicity *CodedConcept `json:"ethnicity,omitempty"`
}
