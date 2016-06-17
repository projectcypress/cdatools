package models

type Record struct {
	Person                `json:",inline"`
	MedicalRecordNumber   string                `json:"medical_record_number,omitempty"`
	MedicalRecordAssigner string                `json:"medical_record_assigner,omitempty"`
	DeathDate             int64                 `json:"deathdate,omitempty"`
	Expired               bool                  `json:"expired,omitempty"`
	Encounters            []Encounter           `json:"encounters,omitempty"`
	LabResults            []LabResult           `json:"results,omitempty"`
	Languages             []Language            `json:"languages,omitempty"`
	ProviderPerformances  []ProviderPerformance `json:"provider_performances,omitempty"`
	InsuranceProviders    []InsuranceProvider   `json:"insurance_providers,omitempty"`
	Procedures            []Procedure           `json:"procedures,omitempty"`
	Medications           []Medication          `json:"medications,omitempty"`
	Allergies             []Allergy             `json:"allergies,omitempty"`
	Conditions            []Condition           `json:"conditions,omitempty"`
	Communications        []Communication       `json:"communications,omitempty"`
	MedicalEquipment      []MedicalEquipment    `json:"medical_equipment,omitempty"`
	CareGoals             []CareGoal            `json:"care_goals,omitempty"`
}

// Entries returns all the entries from the Encounters, Diagnoses, and LabResults for a Record
func (r *Record) Entries() []HasEntry {
	var entries []HasEntry

	// This whole "for loop for each of these things" is unavoidable, because elements must be copied individually to a []HasEntry
	for i, _ := range r.Encounters {
		entries = append(entries, &r.Encounters[i])
	}

	for i, _ := range r.LabResults {
		entries = append(entries, &r.LabResults[i])
	}

	for i, _ := range r.InsuranceProviders {
		entries = append(entries, &r.InsuranceProviders[i])
	}

	for i, _ := range r.ProviderPerformances {
		entries = append(entries, &r.ProviderPerformances[i])
	}

	for i, _ := range r.Procedures {
		entries = append(entries, &r.Procedures[i])
	}

	for i, _ := range r.Medications {
		entries = append(entries, &r.Medications[i])
	}

	for i, _ := range r.Allergies {
		entries = append(entries, &r.Allergies[i])
	}

	for i, _ := range r.Conditions {
		entries = append(entries, &r.Conditions[i])
	}

	return entries
}

// EntriesForOid returns all the entries which include the OID
func (r *Record) EntriesForOid(oid string) []HasEntry {
	var matchedEntries []HasEntry
	for _, entry := range r.Entries() {
		if entry.GetEntry().Oid == oid {
			matchedEntries = append(matchedEntries, entry)
		}
	}
	return matchedEntries
}
