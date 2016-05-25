package models

import "testing"

func TestPreferredCode(t *testing.T) {
	codes := make(map[string][]string, 2)
	codes["a"] = []string{"aa", "ab"}
	codes["b"] = []string{"ba", "bb"}
	coded := Coded{Codes: codes}
	entry := Entry{Coded: coded}
	prefCode := entry.PreferredCode([]string{"b"})
	if prefCode.Code != "ba" {
		t.Error("Returned incorrect code, expected", "ba", "got", prefCode.Code)
	}
}
