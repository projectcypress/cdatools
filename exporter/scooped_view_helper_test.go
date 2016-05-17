package exporter

import (
	"testing"

	"github.com/projectcypress/cdatools/models"
	"github.com/stretchr/testify/assert"
)

func TestReasonInCodesTrue(t *testing.T) {
	code := models.CodeSet{}
	reason := models.CodedConcept{}
	code.Values = []models.Concept{models.Concept{Code: "test", CodeSystem: "codeSystem"}}
	reason.Code = "test"
	reason.CodeSystem = "codeSystem"
	assert.True(t, reasonInCodes(code, reason))

}

func TestReasonInCodesFalse(t *testing.T) {
	code := models.CodeSet{}
	reason := models.CodedConcept{}
	code.Values = []models.Concept{models.Concept{Code: "not code", CodeSystem: "codeSystem"}}
	reason.Code = "test"
	reason.CodeSystem = "codeSystem"
	assert.False(t, reasonInCodes(code, reason))

}
