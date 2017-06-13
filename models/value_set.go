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
	for i, oid := range MapDataCriteria.FieldOids[field] {
		preferredCodes := coded.PreferredCodes(codeDisplay.PreferredCodeSets, codeDisplay.CodeSetRequired, codeDisplay.ValueSetPreferred, v, MapDataCriteria.FieldOids[field][i])
		codeDisplay.setCodesFromPreferred(preferredCodes)
		if codeDisplay.PreferredCode.Code != "" {
			//Put the relevant oid in the 0 index for export
			oldoid := MapDataCriteria.FieldOids[field][0]
			MapDataCriteria.FieldOids[field][0] = oid
			MapDataCriteria.FieldOids[field][i] = oldoid
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
	preferredCodes := coded.PreferredCodes(codeDisplay.PreferredCodeSets, codeDisplay.CodeSetRequired, codeDisplay.ValueSetPreferred, v, MapDataCriteria.ValueSetOid)
	codeDisplay.setCodesFromPreferred(preferredCodes)
	return codeDisplay
}

func (v ValueSetMap) CodeDisplayWithPreferredCodeForResultValue(entry *Entry, coded *Coded, MapDataCriteria Mdc, codeType string) CodeDisplay {
	codeDisplay, err := entry.GetCodeDisplay(codeType)
	if err != nil {
		log.Fatalln(err)
	}
	codeDisplay.MapDataCriteria = MapDataCriteria
	for i, oid := range MapDataCriteria.ResultOids {
		preferredCodes := coded.PreferredCodes(codeDisplay.PreferredCodeSets, codeDisplay.CodeSetRequired, codeDisplay.ValueSetPreferred, v, oid)
		codeDisplay.setCodesFromPreferred(preferredCodes)
		if codeDisplay.PreferredCode.Code != "" {
			oldOid := MapDataCriteria.ResultOids[0]
			MapDataCriteria.ResultOids[0] = oid
			MapDataCriteria.ResultOids[i] = oldOid
			return codeDisplay
		}
	}
	for codeSystem, codes := range coded.Codes {
		codeDisplay.PreferredCode = Concept{CodeSystem: codeSystem, Code: codes[0]}
		break
	}
	return codeDisplay
}

func (v ValueSetMap) CodeDisplayWithPreferredCodeAndLaterality(entry *Entry, coded *Coded, codeType string, laterality Laterality, MapDataCriteria Mdc) CodeDisplay {
	codeDisplay, err := entry.GetCodeDisplay(codeType)
	if err != nil {
		log.Fatal(err)
	}
	preferredCodes := coded.PreferredCodes(codeDisplay.PreferredCodeSets, codeDisplay.CodeSetRequired, codeDisplay.ValueSetPreferred, v, MapDataCriteria.ValueSetOid)
	codeDisplay.setCodesFromPreferred(preferredCodes)
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

func (c *CodeDisplay) setCodesFromPreferred(preferredCodes []Concept) {
	if len(preferredCodes) > 0 {
		c.PreferredCode = preferredCodes[0]
		if len(preferredCodes) > 1 {
			c.Translations = preferredCodes[1:]
		}
	}
}

func codeSetContainsCode(sets []CodeSet, codedValue CodedConcept) bool {
	for _, cs := range sets {
		for _, val := range cs.Values {
			if ((val.CodeSystem == codedValue.CodeSystem ||
				val.CodeSystemName == codedValue.CodeSystemName ||
				val.CodeSystemName == codedValue.CodeSystem) ||
				(val.CodeSystem == codeSystemAliases[codedValue.CodeSystem] ||
					val.CodeSystemName == codeSystemAliases[codedValue.CodeSystemName] ||
					val.CodeSystemName == codeSystemAliases[codedValue.CodeSystem])) &&
				val.Code == codedValue.Code {
				return true
			}
		}
	}
	return false
}
