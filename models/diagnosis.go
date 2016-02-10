package models

type Diagnosis struct {
	Entry    `bson:",inline"`
	Severity map[string][]string `json:"severity"`
}
