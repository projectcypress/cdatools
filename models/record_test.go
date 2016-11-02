package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReasonInCodesTrue(t *testing.T) {
	code := CodeSet{}
	reason := CodedConcept{}
	code.Values = []Concept{Concept{Code: "test", CodeSystem: "codeSystem"}}
	reason.Code = "test"
	reason.CodeSystem = "codeSystem"
	assert.True(t, reasonInCodes(code, reason))
}

func TestReasonInCodesFalse(t *testing.T) {
	code := CodeSet{}
	reason := CodedConcept{}
	code.Values = []Concept{Concept{Code: "not code", CodeSystem: "codeSystem"}}
	reason.Code = "test"
	reason.CodeSystem = "codeSystem"
	assert.False(t, reasonInCodes(code, reason))
}
