package models

type ProviderPerformance struct {
	Entry     `bson:",inline"`
	StartDate int64 `json:"startDate,omitempty"`
	EndDate   int64 `json:"endDate,omitempty"`
}

func (pp *ProviderPerformance) GetEntry() *Entry {
	return &pp.Entry
}
