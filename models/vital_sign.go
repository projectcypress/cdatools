package models

type VitalSign struct {
	LabResult `bson:",inline"`
}
