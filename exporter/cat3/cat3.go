package cat3

import (
	"time"

	"github.com/projectcypress/cdatools/models"
)

type Authenticator struct {
	models.Authenticator
	Author
}

func NewAuthenticator(a models.Authenticator) Authenticator {
	return Authenticator{
		Authenticator: a,
		Author:        NewAuthor(a.Author),
	}
}

type Header struct {
	models.Header
	Authenticator
	Authors
}

func NewHeader(h models.Header) Header {
	return Header{
		Header:        h,
		Authenticator: NewAuthenticator(h.Authenticator),
		Authors:       NewAuthors(h.Authors),
	}
}

type Doc struct {
	Header         Header
	Measures       models.Measure
	MeasureSection MeasureSection
	StartDate      int64
	EndDate        int64
	Timestamp      int64
}

func NewDoc(h models.Header, ms MeasureSection, m models.Measure, start int64, end int64) Doc {
	timeNow := time.Now().UTC().Unix()
	return Doc{
		Header:         NewHeader(h),
		Measures:       m,
		MeasureSection: ms,
		StartDate:      start,
		EndDate:        end,
		Timestamp:      timeNow,
	}
}

func (d Doc) Template() string {
	t := `<?xml version="1.0" encoding="utf-8"?>
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
	<id {{if .Header.Identifier.Root}}root="{{escape .Header.Identifier.Root}}"{{end}} extension="{{escape .Header.Identifier.Extension}}" />

	<!-- SHALL QRDA III document type code -->
	<code code="55184-6" codeSystem="2.16.840.1.113883.6.1" codeSystemName="LOINC"
		displayName="Quality Reporting Document Architecture Calculated Summary Report"/>
	<!-- SHALL Title, SHOULD have this content -->
	<title>QRDA Calculated Summary Report</title>
	<!-- SHALL  -->
	<effectiveTime value="{{timeToFormat .Timestamp "20060102"}}"/>
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

	{{Print .Header.Authors.Template .Header.Authors}}
	<!-- SHALL have 1..* author. MAY be device or person.
		The author of the CDA document in this example is a device at a data submission vendor/registry. -->

	<!-- The custodian of the CDA document is the same as the legal authenticator in this
	example and represents the reporting organization. -->
	<!-- SHALL -->
	<custodian>
		<assignedCustodian>
		{{Print .Header.Organization.Template .Header.Organization}} <!--TagName "representedCustodianOrganization"-->
		</assignedCustodian>
	</custodian>
	<!-- The legal authenticator of the CDA document is a single person who is at the
		same organization as the custodian in this example. This element must be present. -->
	<!-- SHALL -->
	<legalAuthenticator>
		<!-- SHALL -->
		<time value="{{.Header.Authenticator.Author.Time}}"/>
		<!-- SHALL -->
		<signatureCode code="S"/>
		<assignedEntity>
			<!-- SHALL ID -->
			{{Print .Header.Authenticator.Author.Ids.Template .Header.Authenticator.Author.Ids}}
			{{Print .Header.Authenticator.Author.Addresses.Template .Header.Authenticator.Author.Addresses}}
			<assignedPerson>
				<name>
				<given>{{escape .Header.Authenticator.Author.Person.First}}</given>
				<family>{{escape .Header.Authenticator.Author.Person.Last}}</family>
				</name>
			</assignedPerson>

			{{Print .Header.Authenticator.Author.Organization.Template .Header.Authenticator.Author.Organization}} <!--TagName "representedOrganization"-->
		</assignedEntity>
	</legalAuthenticator>

	{{Print .MeasureSection.Template .MeasureSection}}
</ClinicalDocument>`

	return t
}
