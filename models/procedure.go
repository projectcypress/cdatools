package models

type Procedure struct {
	Entry            `bson:",inline"`
	Ordinality       Ordinality   `json:"ordinality,omitempty" bson:"ordinality,omitempty"`
	Performer        Performer    `json:"performer,omitempty" bson:"performer,omitempty"`
	AnatomicalTarget CodedConcept `json:"anatomical_target,omitempty" bson:"anatomical_target,omitempty"`
	IncisionTime     *int64       `json:"incisionTime,omitempty" bson:"incisionTime,omitempty"`
}

type Performer struct {
}

func (proc *Procedure) GetEntry() *Entry {
	return &proc.Entry
}

func (proc *Procedure) HasSetIncisionTime() bool {
	return proc.IncisionTime != nil
}

func (proc *Procedure) HasSetOrdinality() bool {
	return proc.Ordinality != Ordinality{}
}
