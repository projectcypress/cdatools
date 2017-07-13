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

// IdentifierForRoot takes a Provider and a string identifying the root OID, and returns the extension for that OID
func (prov *Provider) IdentifierForRoot(root string) string {
	for _, identifier := range prov.CDAIdentifiers {
		if identifier.Root == root {
			return identifier.Extension
		}
	}
	return ""
}

// IdentifiersForRoot applies IdentifierForRoot to a passed-in slice of Providers
func IdentifiersForRoot(provs []Provider, root string) []string {
	ids := []string{}
	for _, provider := range provs {
		if id := provider.IdentifierForRoot(root); id != "" {
			ids = append(ids, id)
		}
	}
	return ids
}
