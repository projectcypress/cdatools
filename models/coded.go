package models

import "github.com/moovweb/gokogiri/xml"

// Coded is a meta-struct that adds Code fields to an object
type Coded struct {
	Codes map[string][]string `json:"codes,omitempty"`
}

// AddCode adds a code string to the given codeSystem
func (c *Coded) AddCode(code string, codeSystem string) {
	if _, ok := c.Codes[codeSystem]; ok {
		c.Codes[codeSystem] = append(c.Codes[codeSystem], code)
	} else {
		c.Codes[codeSystem] = []string{code}
	}
}

// CodesInCodeSet returns all the codes for a particular CodeSet (ICD-10, SNOMED-CT, etc.)
func (c *Coded) CodesInCodeSet(codeSet string) []string {
	return c.Codes[codeSet]
}

// AddCodeIfPresent adds a code to a given codeSet within the Coded, if the Code exists
func (c *Coded) AddCodeIfPresent(codeElement xml.Node) {
	var code string
	var codeSystem string

	//extract code from attribute if it exists
	codeAttribute := codeElement.Attribute("code")
	if codeAttribute != nil {
		code = codeAttribute.String()
	}

	//extract codeSystem from attribute if it exists
	codeSystemAttribute := codeElement.Attribute("codeSystem")
	if codeSystemAttribute != nil {
		codeSystem = CodeSystemFor(codeElement.Attribute("codeSystem").String())
	}

	if code != "" && codeSystem != "" {
		c.AddCode(code, codeSystem)
	}
}

// IsInCodeSet checks if a code is in the list of possible codes
func (c *Coded) IsInCodeSet(codeSet []CodeSet) bool {
	for codeSystem, _ := range c.Codes {
		for _, set := range codeSet {
			if set.Set == codeSystem {
				if doSetsIntersect(set.Values, c.Codes[codeSystem]) {
					return true
				}
			}
		}
	}
	return false
}

func doSetsIntersect(a []Concept, b []string) bool {
	var m = make(map[string]int, len(a)+len(b))
	for _, con := range a {
		m[con.Code]++
	}
	for _, str := range b {
		if m[str] == 1 {
			return true
		}
	}
	return false
}

func computeIntersection(a []string, b []string) []string {
	var intersect = make([]string, 0)
	// We start with making a map of one of the lists
	var m = make(map[string]int, len(a))
	for _, el := range a {
		m[el] = 1
	}
	// Now iterate over B to check each element is in the map
	for _, checkEl := range b {
		if m[checkEl] == 1 {
			intersect = append(intersect, checkEl)
		}
	}
	return intersect
}
