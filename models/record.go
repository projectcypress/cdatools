package models

import (
	"regexp"

	"github.com/pborman/uuid"
)

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

//type EntryGroup []Entry
//type EncounterGroup []Encounter
//type LabResults []LabResult

//type EntryService interface {
//	EntriesForDataCriteria(DataCriteria) []EntryGroup
//	EntriesForOid(oid string) []EntryGroup
//}

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

// GetEntriesForOids returns all the entries which include the list of OIDs given
func (r *Record) GetEntriesForOids(oids ...string) []HasEntry {
	var matchedEntries []HasEntry
	for _, entry := range r.Entries() {
		for _, oid := range oids {
			if entry.GetEntry().Oid == oid {
				matchedEntries = append(matchedEntries, entry)
			}
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

func (r *Record) EntriesForDataCriteria(dataCriteria DataCriteria, vsMap map[string][]CodeSet) []HasEntry {
	dataCriteriaOid := dataCriteria.HQMFOid

	var entries []HasEntry
	var filteredEntries []HasEntry
	switch dataCriteriaOid {
	case "2.16.840.1.113883.3.560.1.404":
		filteredEntries = r.handlePatientExpired()
	case "2.16.840.1.113883.3.560.1.405":
		filteredEntries = r.handlePayerInformation()
	default:

		var codes []CodeSet
		switch dataCriteriaOid {
		case "2.16.840.1.113883.3.560.1.5", "2.16.840.1.113883.3.560.1.12":
			// If Lab Test: Performed, look for Lab Test: Result too
			entries = r.GetEntriesForOids("2.16.840.1.113883.3.560.1.5", "2.16.840.1.113883.3.560.1.12")
		case "2.16.840.1.113883.3.560.1.6", "2.16.840.1.113883.3.560.1.63":
			entries = r.GetEntriesForOids("2.16.840.1.113883.3.560.1.6", "2.16.840.1.113883.3.560.1.63")
		case "2.16.840.1.113883.3.560.1.3", "2.16.840.1.113883.3.560.1.11":
			entries = r.GetEntriesForOids("2.16.840.1.113883.3.560.1.3", "2.16.840.1.113883.3.560.1.11")
		case "2.16.840.1.113883.3.560.1.71", "2.16.840.1.113883.3.560.1.72":
			// Transfers (either from or to)
			entries = r.GetEntriesForOids(dataCriteriaOid, "2.16.840.1.113883.3.560.1.79")
			if dataCriteria.FieldValues != nil {
				codeListID := dataCriteria.FieldValues["TRANSFER_FROM"].CodeListID
				if codeListID == "" {
					codeListID = dataCriteria.FieldValues["TRANSFER_TO"].CodeListID
					codes = vsMap[codeListID]
				}
			}
		default:
			entries = r.GetEntriesForOids(dataCriteriaOid)
		}

		if codes == nil {
			codes = vsMap[dataCriteria.CodeListID]
		}

		// Get a slice containing only unique entries, by adding them to a map, then iterating over that
		// NOTE: this makes me hate myself
		uniqueEntries := make(map[string]HasEntry)
		for _, entry := range entries {
			uniqueEntries[string(uuid.NewRandom())] = entry
		}
		var negationRegexp = regexp.MustCompile(`2\.16\.840\.1\.113883\.3\.526\.3\.100[7-9]`)
		for _, entry := range uniqueEntries {

			entryData := entry.GetEntry()

			if negationRegexp.FindStringIndex(dataCriteria.CodeListID) != nil {
				// Add the entry to FilteredEntries if entry.negation_reason['code'] is in codes
				if reasonInCodes(codes[0], entryData.NegationReason) {
					filteredEntries = append(filteredEntries, entry)
				}
			} else if dataCriteriaOid == "2.16.840.1.113883.3.560.1.71" {
				if transferFrom := &entry.(*Encounter).TransferFrom; transferFrom != nil {
					transferFrom.Codes[transferFrom.CodeSystem] = []string{transferFrom.Code}
					tfc := transferFrom.Coded.CodesInCodeSet(codes[0].Set)
					if len(tfc) > 0 {
						filteredEntries = append(filteredEntries, entry)
					}
				}
			} else if dataCriteriaOid == "2.16.840.1.113883.3.560.1.72" {
				if transferTo := &entry.(*Encounter).TransferTo; transferTo != nil {
					transferTo.Codes[transferTo.CodeSystem] = []string{transferTo.Code}
					if len(transferTo.Coded.CodesInCodeSet(codes[0].Set)) > 0 {
						filteredEntries = append(filteredEntries, entry)
					}
				}
			} else {
				if entryData.IsInCodeSet(codes) && entryData.NegationInd != nil {
					if *entryData.NegationInd == dataCriteria.Negation {
						filteredEntries = append(filteredEntries, entry)
					}
				} else if entryData.IsInCodeSet(codes) && entryData.NegationInd == nil && !dataCriteria.Negation {
					filteredEntries = append(filteredEntries, entry)
				}
			}
		}
	}

	return filteredEntries
}


func (r *Record) handlePatientExpired() []HasEntry {
	if r.Expired {
		exp := make([]HasEntry, 1)
		return append(exp, &Entry{StartTime: r.DeathDate})
	}
	return nil
}

func (r *Record) handlePayerInformation() []HasEntry {
	providers := make([]HasEntry, len(r.InsuranceProviders))
	for _, prov := range r.InsuranceProviders {
		providers = append(providers, &prov)
	}
	return providers
}

func reasonInCodes(code CodeSet, reason CodedConcept) bool {
	for _, value := range code.Values {
		if reason.Code == value.Code && reason.CodeSystem == value.CodeSystem {
			return true
		}
	}
	return false
}
