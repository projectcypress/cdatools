package exporter

import "github.com/projectcypress/cdatools/models"

func handlePatientExpired(patient models.Record) []models.Entry {
	if patient.Expired {
		return []models.Entry{models.Entry{StartTime: patient.DeathDate}}
	}
	return nil
}

func handlePayerInformation(patient models.Record) []models.Entry {
	providers := make([]models.Entry, len(patient.InsuranceProviders))
	for _, prov := range patient.InsuranceProviders {
		// NOTE this redefines the InsuranceProviders as Entries, so we can add them into the entries array
		// This dumps a lot of data out of it, but it doesn't look like the cat1 needs/wants it.
		entry := models.Entry{ID: prov.ID, StartTime: prov.StartTime, EndTime: prov.EndTime, Codes: prov.Codes}
		providers = append(providers, entry)
	}
	return providers
}

func entriesForDataCriteria(dataCriteria models.DataCriteria, patient models.Record) {
	dataCriteriaOid := GetID(dataCriteria)
	var entries []models.Entry
	var filteredEntries []models.Entry
	switch dataCriteriaOid {
	case "2.16.840.1.113883.3.560.1.404":
		filteredEntries = handlePatientExpired(patient)
	case "2.16.840.1.113883.3.560.1.405":
		filteredEntries = handlePayerInformation(patient)
	default:
		entries = append(entries, patient.EntriesForOid(dataCriteriaOid)...)
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
				}
				// TODO finish this method
			}
		}
	}
}
