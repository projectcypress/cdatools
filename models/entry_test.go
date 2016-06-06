package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCodeDisplay(t *testing.T) {
	entry := Entry{CodeDisplays: []CodeDisplay{CodeDisplay{CodeType: "first code type"}, CodeDisplay{CodeType: "second code type"}}}

	codeDisplay, err := entry.GetCodeDisplay("first code type")
	assert.Nil(t, err)
	assert.Equal(t, CodeDisplay{CodeType: "first code type"}, codeDisplay)

	codeDisplay, err = entry.GetCodeDisplay("not a code type")
	assert.NotNil(t, err)
}

func TestNegationReasonOrReason(t *testing.T) {
	reason1 := CodedConcept{Code: "first code", CodeSystem: "first code system", CodeSystemName: "first code system name"}
	reason2 := CodedConcept{Code: "second code", CodeSystem: "second code system", CodeSystemName: "second code system name"}

	// return negation reason over reason
	entry := Entry{Reason: reason1, NegationReason: reason2}
	assert.Equal(t, reason2, entry.NegationReasonOrReason())

	// only negation reason
	entry = Entry{NegationReason: reason1}
	assert.Equal(t, reason1, entry.NegationReasonOrReason())

	// only reason
	entry = Entry{Reason: reason1}
	assert.Equal(t, reason1, entry.NegationReasonOrReason())

	// no negation reason or reason
	entry = Entry{}
	assert.Equal(t, CodedConcept{}, entry.NegationReasonOrReason())
}

func TestExtractEntryFromEncounter(t *testing.T) {
	entry := Entry{Description: "my entry's description"}
	var encounter HasEntry = &Encounter{Entry: entry}
	extractedEntry := encounter.GetEntry()
	assert.Equal(t, entry, *extractedEntry)
}

func TestExtractedEntryCanBeEdited(t *testing.T) {
	var encounter HasEntry = &Encounter{Entry: Entry{Description: "my entry's description"}}
	firstExtractedEntry := encounter.GetEntry()
	firstExtractedEntry.Description = "different description from before"
	secondExtractedEntry := encounter.GetEntry()
	assert.Equal(t, *firstExtractedEntry, *secondExtractedEntry)
}
