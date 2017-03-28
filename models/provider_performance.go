package models

type ProviderPerformance struct {
	Provider
	Entry     `bson:",inline"`
	StartDate *int64   `json:"start_date,omitempty"`
	EndDate   *int64   `json:"end_date,omitempty"`
	Provider  Provider `json:"provider,omitempty"`
}

type Provider struct {
	Ids []CDAIdentifier
}

func (pp *ProviderPerformance) GetEntry() *Entry {
	return &pp.Entry
}
