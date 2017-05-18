package cat3

type MeasureResults struct {
	MeasureID        string            `json:"measure_id,omitempty"`
	PopulationGroups []PopulationGroup `json:"population_groups,omitempty"`
	Populations      []Population      `json:"populations,omitempty"`
}
