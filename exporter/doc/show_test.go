package doc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/projectcypress/cdatools/exporter/doc"
	"github.com/projectcypress/cdatools/models"
)

func TestShowPrint(t *testing.T) {
	startDate := int64(1) // with timeToFormat, converted to 19700101
	endDate := int64(1)   // with timeToFormat, converted to 19700101
	timestamp := int64(1) // with timeToFormat, converted to 19700101
	var tests = []struct {
		n        doc.Cat3Data
		expected string
	}{
		{
			doc.Cat3Data{
				Header: doc.NewHeader(models.Header{
					Identifier: models.CDAIdentifier{}, Authenticator: models.Authenticator{Author: models.Author{Time: &startDate,
						Person: models.Person{First: "first", Last: "last"}},
					}}),
				Record: doc.Record{
					Record:               models.Record{},
					ProviderPerformances: doc.ProviderPerformances{Timestamp: timestamp},
				},
				Measures:  []models.Measure{},
				StartDate: startDate,
				EndDate:   endDate,
				Timestamp: timestamp,
			},
			fmt.Sprintf(showCat3TestTemplate, 19700101, 1, "first", "last", 19700101, 19700101),
		},
	}

	for _, tt := range tests {
		actual := tt.n.Print()
		if strings.TrimSpace(actual) != strings.TrimSpace(tt.expected) {
			t.Errorf("Cat3Data.Print(): expected \n%s\n, actual \n%s", tt.expected, actual)
		}
	}
}

var showCat3TestTemplate = `<?xml version="1.0" encoding="utf-8"?>
<ClinicalDocument xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
 xmlns="urn:hl7-org:v3"
 xmlns:cda="urn:hl7-org:v3">

	<!--
		********************************************************
		CDA Header
		********************************************************
	-->
	<realmCode code="US"/>
	<typeId root="2.16.840.1.113883.1.3" extension="POCD_HD000040"/>
	<!-- QRDA Category III template ID (this template ID differs from QRDA III comment only template ID). -->
	<templateId root="2.16.840.1.113883.10.20.27.1.1" extension="2016-09-01"/>
	<id  extension="" />

	<!-- SHALL QRDA III document type code -->
	<code code="55184-6" codeSystem="2.16.840.1.113883.6.1" codeSystemName="LOINC"
		displayName="Quality Reporting Document Architecture Calculated Summary Report"/>
	<!-- SHALL Title, SHOULD have this content -->
	<title>QRDA Calculated Summary Report</title>
	<!-- SHALL  -->
	<effectiveTime value="%d"/>
	<confidentialityCode codeSystem="2.16.840.1.113883.5.25" code="N"/>
	<languageCode code="en"/>
	<!-- SHOULD The version of the file being submitted. -->
	<versionNumber value="1"/>
	<!-- SHALL contain recordTarget and ID - but ID is nulled to NA. This is an aggregate summary report. Therefore CDA's required patient identifier is nulled. -->
	<recordTarget>
		<patientRole>
			<id nullFlavor="NA"/>
		</patientRole>
	</recordTarget>

	<!-- SHALL have 1..* author. MAY be device or person. 
    The author of the CDA document in this example is a device at a data submission vendor/registry. -->

	<!-- SHALL have 1..* author. MAY be device or person.
		The author of the CDA document in this example is a device at a data submission vendor/registry. -->

	<!-- The custodian of the CDA document is the same as the legal authenticator in this
	example and represents the reporting organization. -->
	<!-- SHALL -->
	<custodian>
		<assignedCustodian>
		<>
	<!-- Represents unique registry organization TIN -->
	
	<!-- Contains name - specific registry not required-->
	<name></name>
</> <!--TagName "representedCustodianOrganization"-->
		</assignedCustodian>
	</custodian>
	<!-- The legal authenticator of the CDA document is a single person who is at the
		same organization as the custodian in this example. This element must be present. -->
	<!-- SHALL -->
	<legalAuthenticator>
		<!-- SHALL -->
		<time value="%d"/>
		<!-- SHALL -->
		<signatureCode code="S"/>
		<assignedEntity>
			<!-- SHALL ID -->
			
			
			<assignedPerson>
				<name>
				<given>%s</given>
				<family>%s</family>
				</name>
			</assignedPerson>

			<>
	<!-- Represents unique registry organization TIN -->
	
	<!-- Contains name - specific registry not required-->
	<name></name>
</> <!--TagName "representedOrganization"-->
		</assignedEntity>
	</legalAuthenticator>

	<documentationOf typeCode="DOC">
	<serviceEvent classCode="PCPR"> <!-- care provision -->
		<!-- No provider data found in the patient record
			putting in a fake provider -->
		<effectiveTime>
			<low value="20020716"/>
			<high value="%d"/>
		</effectiveTime>
		<!-- You can include multiple performers, each with an NPI, TIN, CCN. -->
		<performer typeCode="PRF">
			<time>
				<low value="20020716"/>
				<high value="%d"/>
			</time>
			<assignedEntity>
				<!-- This is the provider NPI -->
				<id root="2.16.840.1.113883.4.6" extension="111111111" />
				<representedOrganization>
					<!-- This is the organization TIN -->
					<id root="2.16.840.1.113883.4.2" extension="1234567" />
					<!-- This is the organization CCN -->
					<id root="2.16.840.1.113883.4.336" extension="54321" />
				</representedOrganization>
			</assignedEntity>
		</performer>
	</serviceEvent>
</documentationOf>
</ClinicalDocument>`
