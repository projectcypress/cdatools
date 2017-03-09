package models

import "github.com/jbowtie/gokogiri/xml"

type CodedConcept struct {
	Code           string `json:"code,omitempty"`
	CodeSystem     string `json:"code_system,omitempty"`
	CodeSystemName string `json:"code_system_name,omitempty"`
	CodeSystemOid  string `json:"code_system_oid,omitempty"`
	DisplayName    string `json:"display_name,omitempty"`
}

func (c *CodedConcept) AddCodeIfPresent(codeElement xml.Node) {
	if codeElement != nil {
		// get code from the code XML attribute
		if codeAttribute := codeElement.Attribute("code"); codeAttribute != nil {
			c.Code = codeAttribute.String()
		}
		// get code system oid from the codeSystem XML attribute (ex. 2.16.840.1.113883.6.96)
		//     code system from the corresponding code system name for the codeSystem XML attribute (ex. "SNOMED-CT")
		if codeSystemAttribute := codeElement.Attribute("codeSystem"); codeSystemAttribute != nil {
			c.CodeSystemOid = codeSystemAttribute.String()
			c.CodeSystem = CodeSystemFor(codeSystemAttribute.String())
		}
		// get code system name from the codeSystemName XML attribute (ex. SNOMED CT)
		if codeSystemNameAttribute := codeElement.Attribute("codeSystemName"); codeSystemNameAttribute != nil {
			c.CodeSystemName = codeSystemNameAttribute.String()
		}
	}
}
