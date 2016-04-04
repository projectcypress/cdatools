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
