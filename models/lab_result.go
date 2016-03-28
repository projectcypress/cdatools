package models

type LabResult struct {
	Entry          `bson:",inline"`
	ReferenceRange string `json:"referenceRange,omitempty"`
}
