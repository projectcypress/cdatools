package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	emptyTransfer := Transfer{}
	assert.Equal(t, true, emptyTransfer.IsEmpty())

	transfer := Transfer{Time: int64(4)}
	assert.Equal(t, false, transfer.IsEmpty())

	coded := Coded{Codes: map[string][]string{"my_key": []string{"my_first_val"}}}
	transfer = Transfer{Coded: coded}
	assert.Equal(t, false, transfer.IsEmpty())

	transfer = Transfer{CodedConcept: CodedConcept{Code: "my_code"}}
	assert.Equal(t, false, transfer.IsEmpty())

	transfer = Transfer{CodedConcept: CodedConcept{CodeSystem: "my_code_system"}}
	assert.Equal(t, false, transfer.IsEmpty())

	transfer = Transfer{CodedConcept: CodedConcept{CodeSystemName: "my_code_system_name"}}
	assert.Equal(t, false, transfer.IsEmpty())

	transfer = Transfer{CodedConcept: CodedConcept{CodeSystemOid: "my_code_system_oid"}}
	assert.Equal(t, false, transfer.IsEmpty())
}
