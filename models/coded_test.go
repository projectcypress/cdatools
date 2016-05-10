package models

import "testing"

func TestIntersection(t *testing.T) {
	if len(computeIntersection([]string{"a", "b"}, []string{"a"})) != 1 {
		t.Error("Incorrect number of intersecting elements")
	}
	if len(computeIntersection([]string{"a", "b"}, []string{"a", "b"})) != 2 {
		t.Error("Incorrect number of intersecting elements")
	}
	if len(computeIntersection([]string{"a", "b"}, []string{"c", "d"})) != 0 {
		t.Error("Incorrect number of intersecting elements")
	}
}

func TestPreferredCode(t *testing.T) {
	codes := make(map[string][]string, 2)
	codes["a"] = []string{"aa", "ab"}
	codes["b"] = []string{"ba", "bb"}
	coded := Coded{Codes: codes}
	entry := Entry{Coded: coded}
	entry.PreferredCode([]string{"b", "a"})
}
