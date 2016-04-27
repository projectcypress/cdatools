package models

// Reference is a link from one entry to another, used in "fulfills" among others
type Reference struct {
	Type           string `json:"type,omitempty"`
	ReferencedType string `json:"referenced_type,omitempty"`
	ReferencedID   string `json:"referenced_id,omitempty"`
}
