package models

import (
	"regexp"

	"github.com/pborman/uuid"
)

type Record struct {
	Person
	RecordGroup
	MedicalRecordNumber   string                `json:"medical_record_number,omitempty"`
	MedicalRecordAssigner string                `json:"medical_record_assigner,omitempty"`
	BirthDate             *int64                `json:"birthdate,omitempty"`
	DeathDate             *int64                `json:"deathdate,omitempty"`
	Expired               bool                  `json:"expired,omitempty"`
	ProviderPerformances  []ProviderPerformance `json:"provider_performances,omitempty"`
	// Private values to handle querying of Entries

}

type RecordGroup struct {
	Encounters           EncounterGroup           `json:"encounters,omitempty"`
	LabResults           LabResultsGroup          `json:"results,omitempty"`
	ProviderPerformances ProviderPerformanceGroup `json:"provider_performances,omitempty"`
	InsuranceProviders   InsuranceProviderGroup   `json:"insurance_providers,omitempty"`
	Procedures           ProcedureGroup           `json:"procedures,omitempty"`
	Medications          MedicationGroup          `json:"medications, omitempty"`
	Allergies            AllergyGroup             `json:"allergies,omitempty"`
	Conditions           ConditionGroup           `json:"conditions,omitempty"`
	VitalSigns           VitalSignGroup           `json:"vital_signs,omitempty"`
	Communications   	 CommunicationGroup       `json:"communications,omitempty"`
	MedicalEquipment 	 MedicalEquipmentGroup    `json:"medical_equipment,omitempty"`

	// These weren't in the Entries() method.
	Languages        LanguagesGroup        `json:"languages,omitempty"`
	CareGoals        EntryGroup            `json:"care_goals,omitempty"`
}

type Language struct {
	Coded
}

type EntryGroup []Entry
type EncounterGroup []Encounter
type LabResultsGroup []LabResult
type LanguagesGroup []Language
type ProviderPerformanceGroup []ProviderPerformance
type InsuranceProviderGroup []InsuranceProvider
type ProcedureGroup []Procedure
type MedicationGroup []Medication
type AllergyGroup []Allergy
type ConditionGroup []Condition
type CommunicationGroup []Communication
type MedicalEquipmentGroup []MedicalEquipment
type VitalSignGroup []VitalSign

type EntryService interface {
	EntriesForDataCriteria(DataCriteria) EntryGroup
	EntriesForOid(oid string) EntryGroup
}

// Entries returns all the entries from the Encounters, Diagnoses, and LabResults for a Record
func (r *Record) Entries() []HasEntry {
	var entries []HasEntry

	// This whole "for loop for each of these things" is unavoidable, because elements must be copied individually to a []HasEntry
	for i := range r.Encounters {
		entries = append(entries, &r.Encounters[i])
	}

	for i := range r.LabResults {
		entries = append(entries, &r.LabResults[i])
	}

	for i := range r.InsuranceProviders {
		entries = append(entries, &r.InsuranceProviders[i])
	}

	for i := range r.ProviderPerformances {
		entries = append(entries, &r.ProviderPerformances[i])
	}

	for i := range r.Procedures {
		entries = append(entries, &r.Procedures[i])
	}

	for i := range r.Medications {
		entries = append(entries, &r.Medications[i])
	}

	for i := range r.Allergies {
		entries = append(entries, &r.Allergies[i])
	}

	for i := range r.Conditions {
		entries = append(entries, &r.Conditions[i])
	}
	
	for i := range r.Communications {
		entries = append(entries, &r.Communications[i])
	}
	
	for i := range r.MedicalEquipment {
		entries = append(entries, &r.MedicalEquipment[i])
	}

	for i := range r.VitalSigns {
		entries = append(entries, &r.VitalSigns[i])
	}

	return entries
}

var negationRegexp *regexp.Regexp = regexp.MustCompile(`2\.16\.840\.1\.113883\.3\.526\.3\.100[7-9]`)

// GetEntriesForOids returns all the entries which include the list of OIDs given
func (r *Record) GetEntriesForOids(dataCriteria DataCriteria, codes []CodeSet, oids ...string) []HasEntry {
	var entries []HasEntry
	for _, entry := range r.Entries() {
		for _, oid := range oids {
			if entry.GetEntry().Oid == oid {
				entryData := entry.GetEntry()
				dataCriteriaOid := dataCriteria.HQMFOid
				if negationRegexp.FindStringIndex(dataCriteria.CodeListID) != nil {
					// Add the entry to FilteredEntries if Entry.negationReason is in codes
					if reasonInCodes(codes[0], entryData.NegationReason) {
						entries = append(entries, entry)
					}
				} else if dataCriteriaOid == "2.16.840.1.113883.3.560.1.71" {
					if transferFrom := &entry.(*Encounter).TransferFrom; transferFrom != nil {
						if transferFrom.Codes == nil {
							transferFrom.Codes = make(map[string][]string)
						}
						transferFrom.Codes[transferFrom.CodeSystem] = []string{transferFrom.Code}
						if (len(codes) > 0 && codeSetContainsCode(codes,transferFrom.CodedConcept)) {
							newId := ObjectIdentifier{ID: uuid.New()}
							newEntry := Entry{Oid: "2.16.840.1.113883.3.560.1.71", ObjectIdentifier: newId}
							newEntry.StartTime = transferFrom.Time
							newEntry.EndTime = transferFrom.Time
							newTransCode := CodedConcept{Code: transferFrom.Code, CodeSystem: transferFrom.CodeSystem}
							newTrans := Transfer{CodedConcept: newTransCode, Time: transferFrom.Time}
							var newTransElm HasEntry = &Encounter{Entry: newEntry, TransferFrom: newTrans}
							entries = append(entries, newTransElm)
						}
						
					}
				} else if dataCriteriaOid == "2.16.840.1.113883.3.560.1.72" {
					if transferTo := &entry.(*Encounter).TransferTo; transferTo != nil {
						if transferTo.Codes == nil {
							transferTo.Codes = make(map[string][]string)
						}
						transferTo.Codes[transferTo.CodeSystem] = []string{transferTo.Code}
						if len(codes) > 0 && codeSetContainsCode(codes,transferTo.CodedConcept) {
							newId := ObjectIdentifier{ID: uuid.New()}
							newEntry := Entry{Oid: "2.16.840.1.113883.3.560.1.72", ObjectIdentifier: newId}
							newEntry.StartTime = transferTo.Time
							newEntry.EndTime = transferTo.Time
							newTransCode := CodedConcept{Code: transferTo.Code, CodeSystem: transferTo.CodeSystem}
							newTrans := Transfer{CodedConcept: newTransCode, Time: transferTo.Time}
							var newTransElm HasEntry = &Encounter{Entry: newEntry, TransferTo: newTrans}
							entries = append(entries, newTransElm)
						}
					}
				} else {
					if entryData.IsInCodeSet(codes) && entryData.NegationInd != nil {
						if *entryData.NegationInd == dataCriteria.Negation {
							entries = append(entries, entry)
						}
					} else if entryData.IsInCodeSet(codes) && entryData.NegationInd == nil && !dataCriteria.Negation {
						entries = append(entries, entry)
					}
				}
			}
		}
	}
	return entries
}

func (r *Record) EntriesForDataCriteria(dataCriteria DataCriteria, vsMap map[string][]CodeSet) []HasEntry {
	dataCriteriaOid := dataCriteria.HQMFOid
	var entries []HasEntry

	switch dataCriteriaOid {
	case "2.16.840.1.113883.3.560.1.404":
		entries = r.handlePatientExpired()
	case "2.16.840.1.113883.3.560.1.405":
		entries = r.handlePayerInformation()
	default:
		var codes []CodeSet
		codes = vsMap[dataCriteria.CodeListID]

		switch dataCriteriaOid {
		case "2.16.840.1.113883.3.560.1.5", "2.16.840.1.113883.3.560.1.12":
			// If Lab Test: Performed, look for Lab Test: Result too
			entries = r.GetEntriesForOids(dataCriteria, codes, "2.16.840.1.113883.3.560.1.5", "2.16.840.1.113883.3.560.1.12")
		case "2.16.840.1.113883.3.560.1.6", "2.16.840.1.113883.3.560.1.63":
			entries = r.GetEntriesForOids(dataCriteria, codes, "2.16.840.1.113883.3.560.1.6", "2.16.840.1.113883.3.560.1.63")
		case "2.16.840.1.113883.3.560.1.3", "2.16.840.1.113883.3.560.1.11":
			entries = r.GetEntriesForOids(dataCriteria, codes, "2.16.840.1.113883.3.560.1.3", "2.16.840.1.113883.3.560.1.11")
		case "2.16.840.1.113883.3.560.1.71", "2.16.840.1.113883.3.560.1.72":
			// Transfers (either from or to)
			if dataCriteria.FieldValues != nil {
				if dataCriteria.FieldValues["TRANSFER_FROM"].CodeListID != "" {
					codes = vsMap[dataCriteria.FieldValues["TRANSFER_FROM"].CodeListID]
					entries = r.GetEntriesForOids(dataCriteria, codes, dataCriteriaOid, "2.16.840.1.113883.3.560.1.79")
				} else if dataCriteria.FieldValues["TRANSFER_TO"].CodeListID != "" {
					codes = vsMap[dataCriteria.FieldValues["TRANSFER_TO"].CodeListID]
					entries = r.GetEntriesForOids(dataCriteria, codes, dataCriteriaOid, "2.16.840.1.113883.3.560.1.79")
				}
			}
		default:
			entries = r.GetEntriesForOids(dataCriteria, codes, dataCriteriaOid)
		}

		// Gonna have to do for now. First time I've ever made Go panic and I have no clue how it happened.
		// Get a slice containing only unique entries
		//		ids := make(map[string]struct{})
		//		uniqueEntries := make([]HasEntry, len(entries))
		//		for _, entry := range entries {
		//			if _, ok := ids[entry.GetEntry().BSONID]; ok {
		//				continue
		//			}
		//			uniqueEntries = append(uniqueEntries, entry)
		//			ids[entry.GetEntry().BSONID] = *new(struct{})
		//		}

	}
	return entries
}

func (r *Record) handlePatientExpired() []HasEntry {
	if r.Expired {
		exp := make([]HasEntry, 0)
		return append(exp, &Entry{StartTime: r.DeathDate, Oid: "2.16.840.1.113883.3.560.1.404"})
	}
	return nil
}

// create entryInfos for each entry. entryInfos have mapped data criteria (mdc) recieved from the uniqueDataCriteria() function
// also adds code displays struct to each entry
func (r *Record) EntryInfosForPatient(measures []Measure, vsMap map[string][]CodeSet, qrdaVersion string) []EntryInfo {
	mappedDataCriterias := UniqueDataCriteria(allDataCriteria(measures))
	var entryInfos []EntryInfo
	for _, mappedDataCriteria := range mappedDataCriterias {
		var entrySections []HasEntry = r.EntriesForDataCriteria(mappedDataCriteria.DataCriteria, vsMap)
		// add code displays struct to each entry
		for i, entrySection := range entrySections {
			if entrySection != nil {
				entry := entrySections[i].GetEntry()
				hds.SetCodeDisplaysForEntry(entry, mappedDataCriteria, qrdaVersion)
			}
		}
		entryInfos = AppendEntryInfos(entryInfos, entrySections, mappedDataCriteria)
	}
	return entryInfos
}

// ResolveReference takes a Reference object, and finds the Entry that it refers to
func (r *Record) ResolveReference(ref Reference) HasEntry {
	switch ref.ReferencedType {
	case "Conditions":
		for _, entry := range r.Conditions {
			if entry.GetEntry().ID.Extension == ref.ReferencedID {
				return &entry
			}
		}
	case "Allergies":
		for _, entry := range r.Allergies {
			if entry.GetEntry().ID.Extension == ref.ReferencedID {
				return &entry
			}
		}
	case "Medications":
		for _, entry := range r.Medications {
			if entry.GetEntry().ID.Extension == ref.ReferencedID {
				return &entry
			}
		}
	case "Procedures":
		for _, entry := range r.Procedures {
			if entry.GetEntry().ID.Extension == ref.ReferencedID {
				return &entry
			}
		}
	case "ProviderPerformances":
		for _, entry := range r.ProviderPerformances {
			if entry.GetEntry().ID.Extension == ref.ReferencedID {
				return &entry
			}
		}
	case "InsuranceProviders":
		for _, entry := range r.InsuranceProviders {
			if entry.GetEntry().ID.Extension == ref.ReferencedID {
				return &entry
			}
		}
	case "LabResults":
		for _, entry := range r.LabResults {
			if entry.GetEntry().ID.Extension == ref.ReferencedID {
				return &entry
			}
		}
	case "Encounters":
		for _, entry := range r.Encounters {
			if entry.GetEntry().ID.Extension == ref.ReferencedID {
				return &entry
			}
		}
	}
	return nil
}

// TODO: This probably needs to be on MeasureGroup.
func allDataCriteria(measures []Measure) []DataCriteria {
	var dc []DataCriteria
	for _, measure := range measures {
		for _, crit := range measure.HQMFDocument.DataCriteria {
			dc = append(dc, crit)
		}
	}
	return dc
}

// TODO: most likely belongs on a `type InsuranceProvidersGroup []InsuranceProviders`
func (r *Record) handlePayerInformation() []HasEntry {
	providers := make([]HasEntry, len(r.InsuranceProviders))
	for _, prov := range r.InsuranceProviders {
		prov.Oid = "2.16.840.1.113883.3.560.1.405"
		providers = append(providers, &prov)
	}
	return providers
}

// TODO: this belongs on CodeSet.
func reasonInCodes(code CodeSet, reason CodedConcept) bool {
	for _, value := range code.Values {
		if reason.Code == value.Code && reason.CodeSystem == value.CodeSystem {
			return true
		}
	}
	return false
}
