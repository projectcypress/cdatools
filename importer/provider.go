package importer

import (
	"fmt"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
	"github.com/projectcypress/cdatools/models"
)

func ProviderPerformanceExtractor(patientElement xml.Node, providerXPath *xpath.Expression) []models.ProviderPerformance {
	var providerPerformances []models.ProviderPerformance

	performerElements, err := patientElement.Search(providerXPath)
	if err != nil {
		fmt.Printf("Error extracting providers %v \n", err)
	}

	for _, performerElement := range performerElements {
		provider := ProviderExtractor(performerElement)
		pp := models.ProviderPerformance{StartDate: provider.Start, EndDate: provider.End, Provider: provider}
		providerPerformances = append(providerPerformances, pp)
	}

	return providerPerformances
}

func ProviderExtractor(performerElement xml.Node) models.Provider {
	var provider models.Provider

	entityXPath := xpath.Compile("cda:assignedEntity")
	entityElement := FirstElement(entityXPath, performerElement)

	cdaIDXPath := xpath.Compile("cda:id")
	cdaIDElements, err := entityElement.Search(cdaIDXPath)
	if err != nil {
		fmt.Printf("Error extracting cda identifiers %v \n", err)
	}

	for _, cdaIDElement := range cdaIDElements {
		provider.CDAIdentifiers = append(provider.CDAIdentifiers, models.CDAIdentifier{Root: cdaIDElement.Attr("root"), Extension: cdaIDElement.Attr("extension")})
	}

	nameXPath := xpath.Compile("cda:assignedPerson/cda:name")
	nameElement := FirstElement(nameXPath, entityElement)

	if nameElement != nil {
		titleXPath := xpath.Compile("cda:prefix")
		provider.Title = FirstElementContent(titleXPath, nameElement)

		givenNameXPath := xpath.Compile("cda:given[1]")
		provider.GivenName = FirstElementContent(givenNameXPath, nameElement)

		familyNameXPath := xpath.Compile("cda:family")
		provider.FamilyName = FirstElementContent(familyNameXPath, nameElement)

		specialtyXPath := xpath.Compile("cda:code/@code")
		provider.Specialty = FirstElementContent(specialtyXPath, nameElement)
	}

	organizationXPath := xpath.Compile("cda:representedOrganization")
	provider.Organization = OrganizationExtractor(FirstElement(organizationXPath, entityElement))

	startXPath := xpath.Compile("cda:time/cda:low/@value")
	provider.Start = GetTimestamp(startXPath, performerElement)

	endXPath := xpath.Compile("cda:time/cda:high/@value")
	provider.End = GetTimestamp(endXPath, performerElement)

	npiXPath := xpath.Compile("cda:id[@root='2.16.840.1.113883.4.6' or @root='2.16.840.1.113883.3.72.5.2']/@extension")
	provider.Npi = FirstElementContent(npiXPath, entityElement)

	addressXPath := xpath.Compile("cda:addr")
	addressElements, err := performerElement.Search(addressXPath)
	if err != nil {
		fmt.Printf("Error extracting addresses %v \n", err)
	}
	provider.Addresses = make([]models.Address, len(addressElements))
	for i, addressElement := range addressElements {
		provider.Addresses[i] = ImportAddress(addressElement)
	}

	telecomXPath := xpath.Compile("cda:telecom")
	telecomElements, err := performerElement.Search(telecomXPath)
	if err != nil {
		fmt.Printf("Error extracting telecoms %v \n", err)
	}
	provider.Telecoms = make([]models.Telecom, len(telecomElements))
	for i, telecomElement := range telecomElements {
		provider.Telecoms[i] = ImportTelecom(telecomElement)
	}

	return provider
}

func OrganizationExtractor(organzationElement xml.Node) models.Organization {
	var org models.Organization

	nameXPath := xpath.Compile("cda:name")
	org.Name = FirstElementContent(nameXPath, organzationElement)

	addressXPath := xpath.Compile("cda:addr")
	addressElements, err := organzationElement.Search(addressXPath)
	if err != nil {
		fmt.Printf("Error extracting addresses %v \n", err)
	}
	org.Addresses = make([]models.Address, len(addressElements))
	for i, addressElement := range addressElements {
		org.Addresses[i] = ImportAddress(addressElement)
	}

	telecomXPath := xpath.Compile("cda:telecom")
	telecomElements, err := organzationElement.Search(telecomXPath)
	if err != nil {
		fmt.Printf("Error extracting telecoms %v \n", err)
	}
	org.Telecoms = make([]models.Telecom, len(telecomElements))
	for i, telecomElement := range telecomElements {
		org.Telecoms[i] = ImportTelecom(telecomElement)
	}

	return org
}
