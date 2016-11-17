package models

import (
	"regexp"
	"sync"
	"fmt"
	"encoding/json"
	"io/ioutil"
	
	"github.com/pebbe/util"
)

type Record struct {
	Person
	RecordGroup
	MedicalRecordNumber   string                `json:"medical_record_number,omitempty"`
	MedicalRecordAssigner string                `json:"medical_record_assigner,omitempty"`
	BirthDate             int64                 `json:"birthdate,omitempty"`
	DeathDate             int64                 `json:"deathdate,omitempty"`
	Expired               bool                  `json:"expired,omitempty"`
	// Private values to handle querying of Entries
	
}

// TODO: These maps need to be put in their own struct, in their own package, called Hds.
var idMap map[string]string
var idR2Map map[string]string
var hqmfR2Map map[string]DataCriteria
var hqmfMap map[string]DataCriteria
var hqmfQrdaMap map[string]map[string]string           // maps qrda oids to hqmf oids
var qrdaCodeDisplayMap map[string][]CodeDisplay // maps qrda oids to maps containing code display information
var vsMapInit sync.Once
var vsMap map[string][]CodeSet
var hqmfMapInit sync.Once

// TODO: This type is only used once on line :58 of this file.
type HqmfQrdaOidsWithCodeDisplays struct {
	HqmfName     string               `json:"hqmf_name,omitempty"`
	HqmfOid      string               `json:"hqmf_oid,omitempty"`
	QrdaName     string               `json:"qrda_name,omitempty"`
	QrdaOid      string               `json:"qrda_oid,omitempty"`
	CodeDisplays []CodeDisplay `json:"code_displays,omitempty"`
}

func initializeMap() {
	hqmfMapInit.Do(func() {
		importHQMFTemplateJSON()
		importHqmfQrdaJSON()
	})
}

func InitializeVsMap(vs []ValueSet) map[string][]CodeSet {
	vsMapInit.Do(func() {
		vsMap = map[string][]CodeSet{}
		for _, valueSet := range vs {
			vsMap[valueSet.Oid] = valueSet.CodeSetMap()
		}
	})
	return vsMap
}


func makeDefinitionKey(definition string, status string, negation bool) string {
	return fmt.Sprintf("%s-%s-%t", definition, status, negation)
}

func importHqmfQrdaJSON() {
	data, err := ioutil.ReadFile("../exporter/hqmf_qrda_oids.json")
	if err != nil {
		util.CheckErr(err)
	}

	// unmarshal from "hqmf_qrda_oids.json" to hqmfQrdaOids variable
	var hqmfQrdaOids []HqmfQrdaOidsWithCodeDisplays
	if err := json.Unmarshal(data, &hqmfQrdaOids); err != nil {
		util.CheckErr(err)
	}

	// create qrdaCodeDisplayMap
	qrdaCodeDisplayMap = make(map[string][]CodeDisplay)
	for _, oidsElem := range hqmfQrdaOids {
		qrdaCodeDisplayMap[oidsElem.QrdaOid] = oidsElem.CodeDisplays
	}

	// create hqmfQrdaMap (map) of hqmf oid to map[string]string containing "hqmf_name", "hqmf_oid", "qrda_name", and qrda_oid
	hqmfQrdaMap = map[string]map[string]string{}
	for _, oidsElem := range hqmfQrdaOids {
		hqmfQrdaMapElem := make(map[string]string)
		hqmfQrdaMapElem["hqmf_name"] = oidsElem.HqmfName
		hqmfQrdaMapElem["hqmf_oid"] = oidsElem.HqmfOid
		hqmfQrdaMapElem["qrda_name"] = oidsElem.QrdaName
		hqmfQrdaMapElem["qrda_oid"] = oidsElem.QrdaOid
		hqmfQrdaMap[oidsElem.HqmfOid] = hqmfQrdaMapElem
	}
}

// NOTE: Had to remove the use of the Asset function that exists in exporter/templates.go:1228
// Now function just reads the json file directly. This will be changed when the Hds refactor
// happens because the way maps are initialized will be changed.
func importHQMFTemplateJSON() {
	data, err := ioutil.ReadFile("../exporter/hqmf_template_oid_map.json")
	if err != nil {
		fmt.Println("Matt... something's wrong...")
		fmt.Println(err)
	}
	json.Unmarshal(data, &hqmfMap)
	idMap = map[string]string{}
	for id, data := range hqmfMap {
		idMap[makeDefinitionKey(data.Definition, data.Status, data.Negation)] = id
	}
	data, err = ioutil.ReadFile("../exporter/hqmfr2_template_oid_map.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(data, &hqmfR2Map)
	idR2Map = map[string]string{}
	for id, data := range hqmfR2Map {
		idR2Map[makeDefinitionKey(data.Definition, data.Status, data.Negation)] = id
	}
}

// TODO: This function is only used in a test...
func GetTemplateDefinition(id string, r2Compat bool) DataCriteria {
	initializeMap()
	if r2Compat {
		return hqmfR2Map[id]
	} else {
		return hqmfMap[id]
	}
}

func GetID(data DataCriteria, r2Compat bool) string {
	initializeMap()
	if r2Compat {
		return idR2Map[makeDefinitionKey(data.Definition, data.Status, data.Negation)]
	} else {
		return idMap[makeDefinitionKey(data.Definition, data.Status, data.Negation)]
	}
}

// TODO: This function isn't used anywhere in the codebase...
//func GetMap(r2Compat bool) map[string]models.DataCriteria {
//	initializeMap()
//	if r2Compat {
//		return hqmfR2Map
//	} else {
//		return hqmfMap
//	}
//}

func HqmfToQrdaOid(hqmfOid string) string {
	initializeMap()
	var qrdaOidToReturn string
	for curHqmfOid, hqmfQrdaMapVal := range hqmfQrdaMap {
		if hqmfOid == curHqmfOid {
			if qrdaOidToReturn != "" {
				panic("There should only be one QRDA oid for one HQMF oid. If this is hit, there is a flaw in the logic of this code.")
			}
			qrdaOidToReturn = hqmfQrdaMapVal["qrda_oid"]
		}
	}
	return qrdaOidToReturn
}

func codeDisplayForQrdaOid(oid string) []CodeDisplay {
	initializeMap()
	if codeDisplays, ok := qrdaCodeDisplayMap[oid]; ok {
		return codeDisplays
	}
	return []CodeDisplay{}
}
// TODO: Why does this function exist...
func ReasonValueSetOid(codedValue CodedConcept, fieldOids map[string][]string) string {
	return OidForCode(codedValue, fieldOids["REASON"])
}

func OidForCode(codedValue CodedConcept, valuesetOids []string) string {
	for _, vsoid := range valuesetOids {
		oidlist := vsMap[vsoid]
		if codeSetContainsCode(oidlist, codedValue) {
			return vsoid
		}
	}
	return ""
}

func codeSetContainsCode(sets []CodeSet, codedValue CodedConcept) bool {
	for _, cs := range sets {
		for _, val := range cs.Values {
			if val.CodeSystem == codedValue.CodeSystem && val.Code == codedValue.Code {
				return true
			}
		}
	}
	return false
}

func stringInSlice(str string, list []string) bool {
	for _, elem := range list {
		if elem == str {
			return true
		}
	}
	return false
}

// END of code that needs to be put in an Hds package

type dcKey struct {
	DataCriteriaOid string
	ValueSetOid     string
}

type Mdc struct {
	FieldOids    map[string][]string
	ResultOids   []string
	DataCriteria DataCriteria
	dcKey
}

// passed into each qrda oid (entry) template
// EntrySection should be a struct that includes entry attributes (ex. Procedure, Medication, ...)
type EntryInfo struct {
	EntrySection    HasEntry
	MapDataCriteria Mdc
}

type RecordGroup struct {
	Encounters            EncounterGroup           `json:"encounters,omitempty"`
	LabResults            LabResultsGroup           `json:"results,omitempty"`
	ProviderPerformances  ProviderPerformanceGroup `json:"provider_performances,omitempty"`
	InsuranceProviders    InsuranceProviderGroup   `json:"insurance_providers,omitempty"`
	Procedures            ProcedureGroup           `json:"procedures,omitempty"`
	Medications           MedicationGroup          `json:"medications, omitempty"`
	Allergies             AllergyGroup             `json:"allergies,omitempty"`
	Conditions            ConditionGroup           `json:"conditions,omitempty"`
	
	// These weren't in the Entries() method.
	Languages             LanguagesGroup            `json:"languages,omitempty"`
	Communications        CommunicationGroup       `json:"communications,omitempty"`
	MedicalEquipment      MedicalEquipmentGroup    `json:"medical_equipment,omitempty"`
	CareGoals             EntryGroup               `json:"care_goals,omitempty"`
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


type EntryService interface {
	EntriesForDataCriteria(DataCriteria) EntryGroup
	EntriesForOid(oid string) EntryGroup
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

// GetEntriesForOids returns all the entries which include the list of OIDs given
func (r *Record) GetEntriesForOids(dataCriteria DataCriteria, codes []CodeSet, oids ...string) []HasEntry {
	var entries []HasEntry
	for _, entry := range r.Entries() {
		for _, oid := range oids {
			if entry.GetEntry().Oid == oid {
				negationRegexp := regexp.MustCompile(`2\.16\.840\.1\.113883\.3\.526\.3\.100[7-9]`)
				entryData := entry.GetEntry()
				dataCriteriaOid := dataCriteria.HQMFOid
				if negationRegexp.FindStringIndex(dataCriteria.CodeListID) != nil {
					// Add the entry to FilteredEntries if Entry.negationReason is in codes
					if reasonInCodes(codes[0], entryData.NegationReason) {
						entries = append(entries, entry)
					}
				} else if dataCriteriaOid == "2.16.840.1.113883.3.560.1.71" {
					if transferFrom := &entry.(*Encounter).TransferFrom; transferFrom != nil {
						transferFrom.Codes[transferFrom.CodeSystem] = []string{transferFrom.Code}
						tfc := transferFrom.Coded.CodesInCodeSet(codes[0].Set)
						if len(tfc) > 0 {
							entries = append(entries, entry)
						}
					}
				} else if dataCriteriaOid == "2.16.840.1.113883.3.560.1.72" {
					if transferTo := &entry.(*Encounter).TransferTo; transferTo != nil {
						transferTo.Codes[transferTo.CodeSystem] = []string{transferTo.Code}
						if len(transferTo.Coded.CodesInCodeSet(codes[0].Set)) > 0 {
							entries = append(entries, entry)
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
				if dataCriteria.FieldValues["TRANSFER_FROM"].CodeListID == "" {
					codes = vsMap[dataCriteria.FieldValues["TRANSFER_TO"].CodeListID]
				}
			}
			entries = r.GetEntriesForOids(dataCriteria, codes, dataCriteriaOid, "2.16.840.1.113883.3.560.1.79")
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
		exp := make([]HasEntry, 1)
		return append(exp, &Entry{StartTime: r.DeathDate})
	}
	return nil
}

// create entryInfos for each entry. entryInfos have mapped data criteria (mdc) recieved from the uniqueDataCriteria() function
// also adds code displays struct to each entry
func (r *Record) EntryInfosForPatient(measures []Measure, vsMap map[string][]CodeSet) []EntryInfo {
	mappedDataCriterias := UniqueDataCriteria(allDataCriteria(measures))
	var entryInfos []EntryInfo
	for _, mappedDataCriteria := range mappedDataCriterias {
		var entrySections []HasEntry = r.EntriesForDataCriteria(mappedDataCriteria.DataCriteria, vsMap)
		// add code displays struct to each entry
		for i, entrySection := range entrySections {
			if entrySection != nil {
				entry := entrySections[i].GetEntry()
				SetCodeDisplaysForEntry(entry)
			}
		}
		entryInfos = AppendEntryInfos(entryInfos, entrySections, mappedDataCriteria)
	}
	return entryInfos
}

// TODO: NOTE: This code seems like it has a similar pattern to what EntriesForDataCriteria had...
func UniqueDataCriteria(allDataCriteria []DataCriteria) []Mdc {
	mappedDataCriteria := map[dcKey]Mdc{}
	for _, dataCriteria := range allDataCriteria {
		// Based on the data criteria, get the HQMF oid associated with it]
		oid := dataCriteria.HQMFOid
		if oid == "" {
			oid = GetID(dataCriteria, false)
			if oid == "" {
				oid = GetID(dataCriteria, true)
			}
			if oid != "" {
				dataCriteria.HQMFOid = oid
			}
		}
		vsOid := dataCriteria.CodeListID

		// Special cases for the valueSet OID, taken from Health Data Standards
		if oid == "2.16.840.1.113883.3.560.1.71" {
			vsOid = dataCriteria.FieldValues["TRANSFER_FROM"].CodeListID
		} else if oid == "2.16.840.1.113883.3.560.1.72" {
			vsOid = dataCriteria.FieldValues["TRANSFER_TO"].CodeListID
		}

		// Generate the key for the mappedDataCriteria
		dc := dcKey{DataCriteriaOid: oid, ValueSetOid: vsOid}

		var mappedDc = mappedDataCriteria[dc]
		if mappedDc.FieldOids == nil {
			mappedDc = Mdc{DataCriteria: dataCriteria, FieldOids: make(map[string][]string)}
		}

		// Add all the codedValues onto the list of field OIDs
		for field, descr := range dataCriteria.FieldValues {
			if descr.Type == "CD" {
				mappedDc.FieldOids[field] = append(mappedDc.FieldOids[field], descr.CodeListID)
			}
		}

		// If the data criteria has a negation, add the reason onto the returned FieldOids
		if dataCriteria.Negation {
			mappedDc.FieldOids["REASON"] = append(mappedDc.FieldOids["REASON"], dataCriteria.NegationCodeListID)
		}

		// If the data criteria has a value, and it's a "coded" type, added the CodeListId into the result OID set
		if dataCriteria.Value.Type == "CD" {
			mappedDc.ResultOids = append(mappedDc.ResultOids, dataCriteria.CodeListID)
		}

		if dc.DataCriteriaOid != "" {
			mappedDataCriteria[dc] = mappedDc
		}
	}

	// Add the key to the value to get what HDS would have returned
	var retDataCriteria []Mdc
	for key, value := range mappedDataCriteria {
		value.DataCriteriaOid = key.DataCriteriaOid
		value.ValueSetOid = key.ValueSetOid
		retDataCriteria = append(retDataCriteria, value)
	}
	return retDataCriteria
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

// START code that needs to be put on other structs in models

// adds all code system names to preferred code sets if "*" is present in the existant preferred code sets
func allPerferredCodeSetsIfNeeded(cds []CodeDisplay) {
	for i, _ := range cds {
		if stringInSlice("*", cds[i].PreferredCodeSets) {
			cds[i].PreferredCodeSets = CodeSystemNames()
		}
	}
}

// TODO: needs to be put on Entry
func SetCodeDisplaysForEntry(e *Entry) {
	codeDisplays := codeDisplayForQrdaOid(HqmfToQrdaOid(e.Oid))
	allPerferredCodeSetsIfNeeded(codeDisplays)
	for i, _ := range codeDisplays {
		codeDisplays[i].Description = e.Description
	}
	e.CodeDisplays = codeDisplays
}

// TODO: Needs to be on entry? Something?
// append an entryInfo to entryInfos for each entry
func AppendEntryInfos(entryInfos []EntryInfo, entries []HasEntry, mappedDataCriteria Mdc) []EntryInfo {
	for _, entry := range entries {
		if entry != nil {
			entryInfo := EntryInfo{EntrySection: entry, MapDataCriteria: mappedDataCriteria}
			entryInfos = append(entryInfos, entryInfo)
		}
	}
	return entryInfos
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

