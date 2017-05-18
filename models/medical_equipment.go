package models

type MedicalEquipment struct {
	Entry               `bson:",inline"`
	Manufacturer        string       `json:"manufacturer,omitempty"`
	AnatomicalStructure CodedConcept `json:"anatomical_target,omitempty"`
	RemovalTime         *int64       `json:"removal_time,omitempty"`
}
