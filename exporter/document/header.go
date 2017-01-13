package document

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/projectcypress/cdatools/models"
)

type Header struct {
	Authenticator models.Authenticator
	Authors       []models.Author
	Custodian     models.Author
}

func NewHeader() Header {
	var atime1 = new(int64)
	var atime2 = new(int64)
	*atime1 = 1449686219
	*atime2 = 1449686219
	h := &Header{
		Authors: []models.Author{
			models.Author{
				Time: atime1,
				Entity: models.Entity{
					Ids: []models.CDAIdentifier{
						models.CDAIdentifier{
							Root:      "authorRoot",
							Extension: "authorExtension",
						},
					},
					Addresses: []models.Address{
						models.Address{
							Street: []string{
								"202 Burlington Road",
								"Apartment 1",
							},
							City:    "Bedford",
							State:   "MA",
							Zip:     "01730",
							Country: "USA",
							Use:     "PUB",
						},
					},
					Telecoms: []models.Telecom{
						models.Telecom{
							Use:   "WP",
							Value: "1(781)2712000",
						},
					},
				},
				Device: models.Device{
					Name:  "deviceName",
					Model: "deviceModel",
				},
				Organization: models.Organization{
					Entity: models.Entity{
						Ids: []models.CDAIdentifier{
							models.CDAIdentifier{
								Root:      "authorsOrganizationRoot",
								Extension: "authorsOrganizationExt",
							},
						},
					},
					Name:    "authorsOrganization",
					TagName: "representedOrganization",
				},
			},
		},
		Custodian: models.Author{
			Entity: models.Entity{
				Ids: []models.CDAIdentifier{
					models.CDAIdentifier{
						Root:      "custodianRoot",
						Extension: "custodianExtension",
					},
				},
			},
			Person: models.Person{
				First: "",
				Last:  "",
			},
			Organization: models.Organization{
				Entity: models.Entity{
					Ids: []models.CDAIdentifier{
						models.CDAIdentifier{
							Root:      "custodianOrganizationRoot",
							Extension: "custodianOrganzationExt",
						},
					},
					Addresses: []models.Address{
						models.Address{
							Street: []string{
								"202 Burlington Road",
								"Apartment 1",
							},
							City:    "Bedford",
							State:   "MA",
							Zip:     "01730",
							Country: "USA",
							Use:     "PUB",
						},
					},
					Telecoms: []models.Telecom{
						models.Telecom{
							Use:   "WP",
							Value: "1(781)2712000",
						},
					},
				},
				Name:    "CustodianOrganization",
				TagName: "representedCustodianOrganization",
			},
		},
		Authenticator: models.Authenticator{
			Author: models.Author{
				Entity: models.Entity{
					Ids: []models.CDAIdentifier{
						models.CDAIdentifier{
							Root:      "legalAuthenticatorRoot",
							Extension: "legalAuthenticatorExt",
						},
					},
					Addresses: []models.Address{
						models.Address{
							Street: []string{
								"202 Burlington Road",
								"Apartment 1",
							},
							City:    "Bedford",
							State:   "MA",
							Zip:     "01730",
							Country: "USA",
							Use:     "PUB",
						},
					},
					Telecoms: []models.Telecom{
						models.Telecom{
							Use:   "WP",
							Value: "1(781)2712000",
						},
					},
				},
				Time: atime2,
				Person: models.Person{
					First: "Legal",
					Last:  "Authenticator",
				},
				Organization: models.Organization{
					Entity: models.Entity{
						Ids: []models.CDAIdentifier{
							models.CDAIdentifier{
								Root:      "legalAuthenticatorOrgRoot",
								Extension: "legalAuthenticatorOrgExt",
							},
						},
					},
					Name:    "LegalAuthenticatorOrg",
					TagName: "representedOrganization",
				},
			},
		},
	}
	return *h
}

func (h Header) Print() string {
	tmpl, err := template.New("").Parse(h.cat1Template())
	if err != nil {
		fmt.Println("error making template:")
		fmt.Println(err)
		return ""
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, h)
	if err != nil {
		fmt.Println("error executing template:")
		fmt.Println(err)
		return ""
	}
	return b.String()
}

func (h Header) cat1Template() string {
	t := `    {{if .Header}}
    {{range .Header.Authors}}
      {{template "_author.xml" .}}
    {{end}}
    <!-- SHALL have 1..* author. MAY be device or person.
      The author of the CDA document in this example is a device at a data submission vendor/registry. -->

    <!-- The custodian of the CDA document is the same as the legal authenticator in this
    example and represents the reporting organization. -->
    <!-- SHALL -->
    <custodian>
      <assignedCustodian>
        {{template "_organization.xml" .Header.Custodian.Organization}}
      </assignedCustodian>
    </custodian>

    <!-- The legal authenticator of the CDA document is a single person who is at the
      same organization as the custodian in this example. This element must be present. -->
    <!-- SHALL -->
    <legalAuthenticator>
      <!-- SHALL -->
      <time value="{{escape .Header.Authenticator.Time}}"/>
      <!-- SHALL -->
      <signatureCode code="S"/>
      <assignedEntity>
        <!-- SHALL ID -->
        {{range .Header.Authenticator.Ids}}
          {{template "_id.xml" .}}
        {{end}}
        {{range .Header.Authenticator.Addresses}}
          {{template "_address.xml" .}}
        {{end}}
        {{range .Header.Authenticator.Telecoms}}
          {{template "_telecom.xml" .}}
        {{end}}
        <assignedPerson>
          <name>
             <given>{{escape .Header.Authenticator.Person.First}}</given>
             <family>{{escape .Header.Authenticator.Person.Last}}</family>
          </name>
       </assignedPerson>
        {{template "_organization.xml" .Header.Authenticator.Organization}}
      </assignedEntity>
    </legalAuthenticator>

  {{ else }}
    <author>
    <time value="{{ timeNow }}"/>
    <assignedAuthor>
      <!-- id extension="Cypress" root="2.16.840.1.113883.19.5"/ -->
      <!-- NPI -->
      <id extension="FakeNPI" root="2.16.840.1.113883.4.6"/>
      <addr>
        <streetAddressLine>202 Burlington Rd.</streetAddressLine>
        <city>Bedford</city>
        <state>MA</state>
        <postalCode>01730</postalCode>
        <country>US</country>
      </addr>
      <telecom use="WP" value="tel:(781)271-3000"/>
      <assignedAuthoringDevice>
        <manufacturerModelName>Cypress</manufacturerModelName >
        <softwareName>Cypress</softwareName >
      </assignedAuthoringDevice >
    </assignedAuthor>
  </author>
  <custodian>
    <assignedCustodian>
      <representedCustodianOrganization>
        <id root="2.16.840.1.113883.19.5"/>
        <name>Cypress Test Deck</name>
        <telecom use="WP" value="tel:(781)271-3000"/>
        <addr>
          <streetAddressLine>202 Burlington Rd.</streetAddressLine>
          <city>Bedford</city>
          <state>MA</state>
          <postalCode>01730</postalCode>
          <country>US</country>
        </addr>
      </representedCustodianOrganization>
    </assignedCustodian>
  </custodian>
  <legalAuthenticator>
    <time value="{{ timeNow }}"/>
    <signatureCode code="S"/>
    <assignedEntity>
      <id root="bc01a5d1-3a34-4286-82cc-43eb04c972a7"/>
      <addr>
        <streetAddressLine>202 Burlington Rd.</streetAddressLine>
        <city>Bedford</city>
        <state>MA</state>
        <postalCode>01730</postalCode>
        <country>US</country>
      </addr>
      <telecom use="WP" value="tel:(781)271-3000"/>
      <assignedPerson>
        <name>
           <given>Henry</given>
           <family>Seven</family>
        </name>
     </assignedPerson>
      <representedOrganization>
        <id root="2.16.840.1.113883.19.5"/>
        <name>Cypress</name>
      </representedOrganization>
    </assignedEntity>
  </legalAuthenticator>
  {{end}}`
	return t
}
