package exporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
	"time"
	"github.com/pborman/uuid"
	"github.com/projectcypress/cdatools/models"
)

type cat1data struct {
	Record    models.Record
	Header    models.Header
	Measures  []models.Measure
	ValueSets []models.ValueSet
	StartDate int64
	EndDate   int64
}

func allDataCriteria(measures []models.Measure) []models.DataCriteria {
	var dc []models.DataCriteria
	for _, measure := range measures {
		for _, crit := range measure.HQMFDocument.DataCriteria {
			dc = append(dc, crit)
		}
	}
	return dc
}

type dc struct {
	DataCriteriaOid string
	ValueSetOid     string
}

type mdc struct {
	FieldOids    map[string][]string
	ResultOids   []string
	DataCriteria models.DataCriteria
	dc
}

func uniqueDataCriteria(allDataCriteria []models.DataCriteria) []mdc {
	mappedDataCriteria := map[dc]mdc{}
	for _, dataCriteria := range allDataCriteria {
		// Based on the data criteria, get the HQMF oid associated with it
		oid := GetID(dataCriteria, false)
		if oid == "" {
			oid = GetID(dataCriteria, true)
		}
		vsOid := dataCriteria.CodeListID

		// Special cases for the valueSet OID, taken from Health Data Standards
		if oid == "2.16.840.1.113883.3.560.1.71" {
			vsOid = dataCriteria.FieldValues["TRANSFER_FROM"].CodeListID
		} else if oid == "2.16.840.1.113883.3.560.1.72" {
			vsOid = dataCriteria.FieldValues["TRANSFER_TO"].CodeListID
		}

		// Generate the key for the mappedDataCriteria
		dc := dc{DataCriteriaOid: oid, ValueSetOid: vsOid}

		var mappedDc = mappedDataCriteria[dc]
		if mappedDc.FieldOids == nil {
			mappedDc = mdc{DataCriteria: dataCriteria, FieldOids: make(map[string][]string)}
		}

		// Add all the codedValues onto the list of field OIDs
		for field, descr := range dataCriteria.FieldValues {
			if descr.Type == "CD" {
				mappedDc.FieldOids[field] = append(mappedDc.FieldOids[field], descr.CodeListID)
			}
		}

		// If the data criteria has a negation, add the reason onto the returned FieldOids
		if dataCriteria.Negation {
			mappedDc.FieldOids["REASON"] = append(mappedDc.FieldOids["REASON"])
		}

		// If the data criteria has a value, and it's a "coded" type, added the CodeListId into the result OID set
		if dataCriteria.Value.Type == "CD" {
			mappedDc.ResultOids = append(mappedDc.ResultOids, dataCriteria.CodeListID)
		}

		mappedDataCriteria[dc] = mappedDc
	}

	// Add the key to the value to get what HDS would have returned
	var retDataCriteria []mdc
	for key, value := range mappedDataCriteria {
		value.DataCriteriaOid = key.DataCriteriaOid
		value.ValueSetOid = key.ValueSetOid
		retDataCriteria = append(retDataCriteria, value)
	}
	return retDataCriteria
}

//export GenerateCat1
func GenerateCat1(patient []byte, measures []byte, valueSets []byte, startDate int64, endDate int64) string {

	data, err := AssetDir("templates/cat1")
	if err != nil {
		fmt.Println(err)
	}

	funcMap := template.FuncMap{
		"timeNow":             time.Now().UTC().Unix,
		"newRandom":           uuid.NewRandom,
		"timeToFormat":        timeToFormat,
		"identifierForInt":    identifierForInt,
		"identifierForString": identifierForString,
		"escape":              escape,
		"entriesForPatient":   entriesForPatient,
	}

	cat1Template := template.New("cat1").Funcs(funcMap)

	for _, d := range data {
		asset, _ := Asset("templates/cat1/" + d)
		template.Must(cat1Template.New(d).Parse(string(asset)))
	}

	p := &models.Record{}
	m := []models.Measure{}
	vs := []models.ValueSet{}
	h := &models.Header{
		Authors: []models.Author{
			models.Author{
				Time: 1449686219,
				Entity: models.Entity{
					Ids: []models.ID{
						models.ID{
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
						Ids: []models.ID{
							models.ID{
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
				Ids: []models.ID{
					models.ID{
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
					Ids: []models.ID{
						models.ID{
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
					Ids: []models.ID{
						models.ID{
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
				Time: 1449686219,
				Person: models.Person{
					First: "Legal",
					Last:  "Authenticator",
				},
				Organization: models.Organization{
					Entity: models.Entity{
						Ids: []models.ID{
							models.ID{
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

	json.Unmarshal(patient, p)
	json.Unmarshal(measures, &m)
	json.Unmarshal(valueSets, &vs)

	initializeVsMap(vs)

	c1d := cat1data{Record: *p, Header: *h, Measures: m, ValueSets: vs, StartDate: startDate, EndDate: endDate}

	var b bytes.Buffer

	err = cat1Template.ExecuteTemplate(&b, "cat1.xml", c1d)

	if err != nil {
		fmt.Println(err)
	}

	return b.String()
}
