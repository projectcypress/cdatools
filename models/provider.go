package models

type Provider struct {
	Title          string          `json:"title,omitempty"`
	GivenName      string          `json:"given_name,omitempty"`
	FamilyName     string          `json:"family_name,omitempty"`
	Organization   Organization    `json:"organization,omitempty"`
	Specialty      string          `json:"specialty,omitempty"`
	Start          *int64          `json:"start,omitempty"`
	End            *int64          `json:"end,omitempty"`
	Npi            string          `json:"npi,omitempty"`
	Addresses      []Address       `json:"addresses,omitempty"`
	Telecoms       []Telecom       `json:"telecoms,omitempty"`
	CDAIdentifiers []CDAIdentifier `json:"cda_identifiers,omitempty"`
}
