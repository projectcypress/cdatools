package models

type Record struct {
	Person
	MedicalRecordNumber   string                `json:"medical_record_number,omitempty"`
	MedicalRecordAssigner string                `json:"medical_record_assigner,omitempty"`
	BirthDate             int64                 `json:"birthdate,omitempty"`
	DeathDate             int64                 `json:"deathdate,omitempty"`
	Expired               bool                  `json:"expired,omitempty"`
	Encounters            []Encounter           `json:"encounters,omitempty"`
	LabResults            []LabResult           `json:"results,omitempty"`
	Languages             []Language            `json:"languages,omitempty"`
	ProviderPerformances  []ProviderPerformance `json:"provider_performances,omitempty"`
	InsuranceProviders    []InsuranceProvider   `json:"insurance_providers,omitempty"`
	Procedures            []Procedure           `json:"procedures,omitempty"`
	Medications           []Medication          `json:"medications, omitempty"`
	Allergies             []Allergy             `json:"allergies,omitempty"`
	Conditions            []Condition           `json:"conditions,omitempty"`
	Communications        []Communication       `json:"communications,omitempty"`
	MedicalEquipment      []MedicalEquipment    `json:"medical_equipment,omitempty"`
	CareGoals             []Entry               `json:"care_goals,omitempty"`
}

type Language struct {
	Coded
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

// Entries returns all the entries from the Encounters, Diagnoses, and LabResults for a Record
func (r *Record) EntryMap() map[string][]HasEntry {
	var entries map[string][]HasEntry

	// This whole "for loop for each of these things" is unavoidable, because elements must be copied individually to a []HasEntry
	for i := range r.Encounters {
		entries["Encounters"] = append(entries["Encounters"], &r.Encounters[i])
	}

	for i := range r.LabResults {
		entries["LabResults"] = append(entries["LabResults"], &r.LabResults[i])
	}

	for i := range r.InsuranceProviders {
		entries["InsuranceProviders"] = append(entries["InsuranceProviders"], &r.InsuranceProviders[i])
	}

	for i := range r.ProviderPerformances {
		entries["ProviderPerformances"] = append(entries["ProviderPerformances"], &r.ProviderPerformances[i])
	}

	for i := range r.Procedures {
		entries["Procedures"] = append(entries["Procedures"], &r.Procedures[i])
	}

	for i := range r.Medications {
		entries["Medications"] = append(entries["Medications"], &r.Medications[i])
	}

	for i := range r.Allergies {
		entries["Allergies"] = append(entries["Allergies"], &r.Allergies[i])
	}

	for i := range r.Conditions {
		entries["Conditions"] = append(entries["Conditions"], &r.Conditions[i])
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

// ResolveReference takes a Reference object, and finds the Entry that it refers to
func (r *Record) ResolveReference(ref Reference) HasEntry {
	for _, entry := range r.EntryMap()[ref.ReferencedType] {
		if entry.GetEntry().ID.Extension == ref.ReferencedID {
			return entry
		}
	}
	return nil
}
