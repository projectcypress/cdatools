package models

type Allergy struct {
	Entry    `bson:",inline"`
	Type     Coded `json:"type,omitempty"`
	Reaction Coded `json:"reaction,omitempty"`
	Severity Coded `json:"severity,omitempty"`
}

func (al *Allergy) GetEntry() *Entry {
	return &al.Entry
}
