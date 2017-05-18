package models

import "log"

// It is very important that NewValueSetMap gets called to get a populated ValueSetMap
type ValueSetMap map[string][]CodeSet

func NewValueSetMap(vs []ValueSet) ValueSetMap {
	vsMap := make(ValueSetMap)
	for _, valueSet := range vs {
		vsMap[valueSet.Oid] = valueSet.CodeSetMap()
	}
	return vsMap
}

// ValueSet is a set of concepts relating to a particular topic
type ValueSet struct {
	Oid         string    `json:"oid"`
	DisplayName string    `json:"display_name"`
	Version     string    `json:"version"`
	Concepts    []Concept `json:"concepts"`
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


func (v ValueSetMap) CodeDisplayWithPreferredCodeForField(entry *Entry, coded *Coded, MapDataCriteria Mdc, codeType string, field string) CodeDisplay {
	codeDisplay, err := entry.GetCodeDisplay(codeType)
	if err != nil {
		log.Fatalln(err)
	}
	codeDisplay.MapDataCriteria = MapDataCriteria
	for FieldOidInd := 0; FieldOidInd < len(MapDataCriteria.FieldOids[field]); FieldOidInd++ {
		codeDisplay.PreferredCode = coded.PreferredCode(codeDisplay.PreferredCodeSets, codeDisplay.CodeSetRequired, codeDisplay.ValueSetPreferred, v, MapDataCriteria.FieldOids[field][FieldOidInd])
		if codeDisplay.PreferredCode.Code != "" {
			break
		}
	}
	return codeDisplay
}

func (v ValueSetMap) CodeDisplayWithPreferredCode(entry *Entry, coded *Coded, MapDataCriteria Mdc, codeType string) CodeDisplay {
	codeDisplay, err := entry.GetCodeDisplay(codeType)
	if err != nil {
		log.Fatalln(err)
	}
	codeDisplay.MapDataCriteria = MapDataCriteria
	codeDisplay.PreferredCode = coded.PreferredCode(codeDisplay.PreferredCodeSets, codeDisplay.CodeSetRequired, codeDisplay.ValueSetPreferred, v, MapDataCriteria.ValueSetOid)
	return codeDisplay
}

func (v ValueSetMap) CodeDisplayWithPreferredCodeForResultValue(entry *Entry, coded *Coded, MapDataCriteria Mdc, codeType string) CodeDisplay {
	codeDisplay, err := entry.GetCodeDisplay(codeType)
	if err != nil {
		log.Fatalln(err)
	}
	codeDisplay.MapDataCriteria = MapDataCriteria
	for ResultOidInd := 0; ResultOidInd < len(MapDataCriteria.ResultOids); ResultOidInd++ {
		codeDisplay.PreferredCode = coded.PreferredCode(codeDisplay.PreferredCodeSets, codeDisplay.CodeSetRequired, codeDisplay.ValueSetPreferred, v, MapDataCriteria.ResultOids[ResultOidInd])
		if codeDisplay.PreferredCode.Code != "" {
			break
		}
	}
	return codeDisplay
}

func (v ValueSetMap) CodeDisplayWithPreferredCodeAndLaterality(entry *Entry, coded *Coded, codeType string, laterality Laterality, MapDataCriteria Mdc) CodeDisplay {
	codeDisplay, err := entry.GetCodeDisplay(codeType)
	if err != nil {
		log.Fatal(err)
	}
	codeDisplay.PreferredCode = coded.PreferredCode(codeDisplay.PreferredCodeSets, codeDisplay.CodeSetRequired, codeDisplay.ValueSetPreferred, v, MapDataCriteria.ValueSetOid)
	codeDisplay.Laterality = laterality
	codeDisplay.MapDataCriteria = MapDataCriteria
	return codeDisplay
}

func (v ValueSetMap) OidForCode(codedValue CodedConcept, valuesetOids []string) string {
	for _, vsoid := range valuesetOids {
		oidlist := v[vsoid]
		if codeSetContainsCode(oidlist, codedValue) {
			return vsoid
		}
	}
	return ""
}

func codeSetContainsCode(sets []CodeSet, codedValue CodedConcept) bool {
	for _, cs := range sets {
		for _, val := range cs.Values {
			if (val.CodeSystem == codedValue.CodeSystem ||
				val.CodeSystemName == codedValue.CodeSystemName ||
				val.CodeSystemName == codedValue.CodeSystem) &&
				val.Code == codedValue.Code {
				return true
			}
		}
	}
	return false
}
