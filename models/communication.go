package models

type Communication struct {
	Entry     `bson:",inline"`
	Direction string `json:"direction,omitempty"`
}
