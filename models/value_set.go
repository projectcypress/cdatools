package models

// ValueSet is a set of concepts relating to a particular topic
type ValueSet struct {
	Oid         string    `json:"oid"`
	DisplayName string    `json:"display_name"`
	Version     string    `json:"version"`
	Concepts    []Concept `json:"concepts"`
}

// Concept is a mapping between code/system/version and a display name
type Concept struct {
	Code              string `json:"code"`
	CodeSystem        string `json:"code_system"`
	CodeSystemName    string `json:"code_system_name"`
	CodeSystemVersion string `json:"code_system_version"`
	DisplayName       string `json:"display_name"`
}

// CodeSet is a struct grouping Concepts by CodeSystem
type CodeSet struct {
	Set    string
	Values []Concept
}

// CodeSetMap returns a slice of CodeSets, each containing Concepts for a different code system
func (vs *ValueSet) CodeSetMap() []CodeSet {
	var cs = make(map[string][]Concept)
	for _, concept := range vs.Concepts {
		cs[concept.CodeSystemName] = append(cs[concept.CodeSystemName], concept)
	}
	var retVal = []CodeSet{}
	for key, val := range cs {
		retVal = append(retVal, CodeSet{Set: key, Values: val})
	}
	return retVal
}
