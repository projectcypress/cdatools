package models

type ProviderPerformance struct {
	Entry     `bson:",inline"`
	StartDate *int64   `json:"start_date,omitempty"`
	EndDate   *int64   `json:"end_date,omitempty"`
	Provider  Provider `json:"provider,omitempty"`
}

func (pp *ProviderPerformance) GetEntry() *Entry {
	return &pp.Entry
}

// GetProviders returns the Provider objects from a slice of ProviderPerformances
func GetProviders(perfs []ProviderPerformance) []Provider {
	provs := []Provider{}
	for _, pp := range perfs {
		provs = append(provs, pp.Provider)
	}
	return provs
}
