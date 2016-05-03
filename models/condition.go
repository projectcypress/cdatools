package models

// Condition encompasses anything that is wrong with a patient that's not a Diagnosis
type Condition struct {
	Entry        `bson:",inline"`
	Type         string              `json:"type,omitempty"`
	CauseOfDeath bool                `json:"cause_of_death,omitempty"`
	TimeOfDeath  int64               `json:"time_of_death,omitempty"`
	Priority     int64               `json:"priority,omitempty"`
	Name         string              `json:"name,omitempty"`
	Ordinality   Ordinality          `json:"ordinality,omitempty"`
	Severity     map[string][]string `json:"severity,omitempty"`
}

type Ordinality struct {
	CodedConcept `bson:",inline"`
	Title        string `json:"title,omitempty"`
}

type Laterality struct {
	CodedConcept `bson:",inline"`
	Title        string `json:"title,omitempty"`
}
