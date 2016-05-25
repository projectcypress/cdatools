package exporter

import (
	"regexp"
	"sync"

	"github.com/pborman/uuid"
	"github.com/projectcypress/cdatools/models"
)

var vsMapInit sync.Once
var vsMap map[string][]models.CodeSet

func initializeVsMap(vs []models.ValueSet) {
	vsMapInit.Do(func() {
		vsMap = map[string][]models.CodeSet{}
		for _, valueSet := range vs {
			vsMap[valueSet.Oid] = valueSet.CodeSetMap()
		}
	})
}

func valueSetMap(vs []models.ValueSet) map[string][]models.CodeSet {
	initializeVsMap(vs)
	return vsMap
}

func handlePatientExpired(patient models.Record) []interface{} {
	if patient.Expired {
		exp := make([]interface{}, 1)
		return append(exp, models.Entry{StartTime: patient.DeathDate})
	}
	return nil
}

func handlePayerInformation(patient models.Record) []interface{} {
	providers := make([]interface{}, len(patient.InsuranceProviders))
	for _, prov := range patient.InsuranceProviders {
		providers = append(providers, prov)
	}
	return providers
}

func entriesForDataCriteria(dataCriteria models.DataCriteria, patient models.Record) []interface{} {
	dataCriteriaOid := dataCriteria.HQMFOid //GetID(dataCriteria, false)
	if dataCriteriaOid == "" {
		dataCriteriaOid = dataCriteria.HQMFOid //GetID(dataCriteria, true)
	}

	var entries []interface{}
	var filteredEntries []interface{}
	switch dataCriteriaOid {
	case "2.16.840.1.113883.3.560.1.404":
		filteredEntries = handlePatientExpired(patient)
	case "2.16.840.1.113883.3.560.1.405":
		filteredEntries = handlePayerInformation(patient)
	default:
		entries = patient.EntriesForOid(dataCriteriaOid)

		var codes []models.CodeSet
		switch dataCriteriaOid {
		case "2.16.840.1.113883.3.560.1.5":
			// If Lab Test: Performed, look for Lab Test: Result too
			entries = append(entries, patient.EntriesForOid("2.16.840.1.113883.3.560.1.12")...)
		case "2.16.840.1.113883.3.560.1.12":
			entries = append(entries, patient.EntriesForOid("2.16.840.1.113883.3.560.1.5")...)
		case "2.16.840.1.113883.3.560.1.6":
			entries = append(entries, patient.EntriesForOid("2.16.840.1.113883.3.560.1.63")...)
		case "2.16.840.1.113883.3.560.1.63":
			entries = append(entries, patient.EntriesForOid("2.16.840.1.113883.3.560.1.6")...)
		case "2.16.840.1.113883.3.560.1.3":
			entries = append(entries, patient.EntriesForOid("2.16.840.1.113883.3.560.1.11")...)
		case "2.16.840.1.113883.3.560.1.11":
			entries = append(entries, patient.EntriesForOid("2.16.840.1.113883.3.560.1.3")...)
		case "2.16.840.1.113883.3.560.1.71", "2.16.840.1.113883.3.560.1.72":
			// Transfers (either from or to)
			entries = append(entries, patient.EntriesForOid("2.16.840.1.113883.3.560.1.79")...)
			if dataCriteria.FieldValues != nil {
				codeListID := dataCriteria.FieldValues["TRANSFER_FROM"].CodeListID
				if codeListID == "" {
					codeListID = dataCriteria.FieldValues["TRANSFER_TO"].CodeListID
					codes = vsMap[codeListID]
				}
			}
		}

		if codes == nil {
			codes = vsMap[dataCriteria.CodeListID]
		}

		// Get a slice containing only unique entries, by adding them to a map, then iterating over that
		// NOTE: this makes me hate myself
		uniqueEntries := make(map[string]interface{})
		for _, entry := range entries {
			uniqueEntries[string(uuid.NewRandom())] = entry
		}
		var negationRegexp = regexp.MustCompile(`2\.16\.840\.1\.113883\.3\.526\.3\.100[7-9]`)
		for _, entry := range uniqueEntries {

			entryData := models.ExtractEntry(&entry)

			if negationRegexp.FindStringIndex(dataCriteria.CodeListID) != nil {
				// Add the entry to FilteredEntries if entry.negation_reason['code'] is in codes
				if reasonInCodes(codes[0], entryData.NegationReason) {
					filteredEntries = append(filteredEntries, entry)
				}
			} else if dataCriteriaOid == "2.16.840.1.113883.3.560.1.71" && &entryData.TransferFrom != nil {
				entryData.TransferFrom.Codes[entryData.TransferFrom.CodeSystem] = []string{entryData.TransferFrom.Code}
				tfc := entryData.TransferFrom.Coded.CodesInCodeSet(codes[0].Set)
				if len(tfc) > 0 {
					filteredEntries = append(filteredEntries, entry)
				}
			} else if dataCriteriaOid == "2.16.840.1.113883.3.560.1.72" && &entryData.TransferTo != nil {
				entryData.TransferTo.Codes[entryData.TransferTo.CodeSystem] = []string{entryData.TransferTo.Code}
				if len(entryData.TransferTo.Coded.CodesInCodeSet(codes[0].Set)) > 0 {
					filteredEntries = append(filteredEntries, entry)
				}
			} else {
				if entryData.IsInCodeSet(codes) && entryData.NegationInd == dataCriteria.Negation {
					filteredEntries = append(filteredEntries, entry)
				}
			}
		}
	}

	return filteredEntries
}

func reasonInCodes(code models.CodeSet, reason models.CodedConcept) bool {
	for _, value := range code.Values {
		if reason.Code == value.Code && reason.CodeSystem == value.CodeSystem {
			return true
		}
	}
	return false
}
