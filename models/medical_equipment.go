package models

type MedicalEquipment struct {
	Entry               `bson:",inline"`
	Manufacturer        string       `json:"manufacturer,omitempty"`
	AnatomicalStructure CodedConcept `json:"anatomicalStructure,omitempty"`
	RemovalTime         *int64       `json:"removal_time,omitempty"`
}
