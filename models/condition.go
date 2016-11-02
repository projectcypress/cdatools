package models

// Condition encompasses anything that is wrong with a patient that's not a Diagnosis
type Condition struct {
	Entry              `bson:",inline"`
	Type               string       `json:"type,omitempty"`
	CauseOfDeath       bool         `json:"cause_of_death,omitempty"`
	TimeOfDeath        int64        `json:"time_of_death,omitempty"`
	Priority           int64        `json:"priority,omitempty"`
	Name               string       `json:"name,omitempty"`
	Ordinality         Ordinality   `json:"ordinality,omitempty"`
	Severity           CodedConcept `json:"severity,omitempty"`
	Laterality         CodedConcept `json:"laterality,omitempty"`
	AnatomicalLocation CodedConcept `json:"anatomical_location,omitempty"`
}

type Ordinality struct {
	CodedConcept
	CodeList []string `json:"SNOMED-CT,omitempty"`
}

type Laterality struct {
	CodedConcept `bson:",inline"`
	Title        string `json:"title,omitempty"`
}

func (con *Condition) GetEntry() *Entry {
	return &con.Entry
}
