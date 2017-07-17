package models

import (
	"encoding/json"
	"testing"

	"github.com/projectcypress/cdatools/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestReasonInCodesTrue(t *testing.T) {
	code := CodeSet{}
	reason := CodedConcept{}
	code.Values = []Concept{Concept{Code: "test", CodeSystemName: "codeSystem"}}
	reason.Code = "test"
	reason.CodeSystem = "codeSystem"
	assert.True(t, reasonInCodes(code, reason))
}

func TestReasonInCodesFalse(t *testing.T) {
	code := CodeSet{}
	reason := CodedConcept{}
	code.Values = []Concept{Concept{Code: "not code", CodeSystemName: "codeSystem"}}
	reason.Code = "test"
	reason.CodeSystem = "codeSystem"
	assert.False(t, reasonInCodes(code, reason))
}

func TestAppendEntryInfos(t *testing.T) {
	// create entry sections
	entries := make([]Entry, 0)
	entries = append(entries, Entry{Description: "my description"})
	var entrySections []HasEntry
	for _, entry := range entries {
		entrySections = append(entrySections, &Encounter{Entry: entry})
	}
	entrySections = append(entrySections, nil) // appendEntryInfos() function should not include nil entry sections

	entryInfos := AppendEntryInfos([]EntryInfo{}, entrySections, Mdc{})
	assert.Equal(t, 1, len(entryInfos))
	assert.Equal(t, "my description", entryInfos[0].EntrySection.GetEntry().Description)
}

func TestEntriesForDataCriteria(t *testing.T) {
	patient := &Record{}
	measure := &Measure{}
	var vs []ValueSet
	json.Unmarshal(fixtures.Cms9_26, &vs)
	vsMap := NewValueSetMap(vs)

	json.Unmarshal(fixtures.TestPatientDataAmi, patient)
	json.Unmarshal(fixtures.Cms9v4a, measure)

	var entries []HasEntry
	for _, crit := range measure.HQMFDocument.DataCriteria {
		if crit.HQMFOid != "" {
			for _, entryForDataCriteria := range patient.EntriesForDataCriteria(crit, vsMap) {
				entries = append(entries, entryForDataCriteria)
			}
		}
	}
	// TODO: This test will have to change when we get a new export of CMS9v4a with all the HQMFOid fields filled.
	assert.Equal(t, len(entries), 1)
}

func TestGetIdentifiersForRoot(t *testing.T) {
	rec := Record{
		ProviderPerformances: []ProviderPerformance{
			ProviderPerformance{
				Provider: Provider{
					CDAIdentifiers: []CDAIdentifier{
						CDAIdentifier{
							Root:      "test",
							Extension: "value1",
						},
						CDAIdentifier{
							Root:      "notvalid",
							Extension: "notthisvalue",
						},
					},
				},
			},
			ProviderPerformance{
				Provider: Provider{
					CDAIdentifiers: []CDAIdentifier{
						CDAIdentifier{
							Root:      "test",
							Extension: "value2",
						},
						CDAIdentifier{
							Root:      "notvalid",
							Extension: "notthisvalueeither",
						},
					},
				},
			},
		},
	}

	ids := rec.ProviderIdentifiersForRoot("test", "default")
	assert.Equal(t, 2, len(ids))

	ids = rec.ProviderIdentifiersForRoot("nonexistent", "default")
	assert.Equal(t, 1, len(ids))
	assert.Contains(t, ids, "default")
}
