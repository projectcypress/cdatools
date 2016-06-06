package models

type LabResult struct {
	Entry          `bson:",inline"`
	ReferenceRange string       `json:"referenceRange,omitempty"`
	Interpretation CodedConcept `json:"interpretation,omitempty"`
}

func (lr *LabResult) GetEntry() *Entry {
	return &lr.Entry
}
