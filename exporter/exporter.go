package exporter

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/pborman/uuid"
	"github.com/projectcypress/cdatools/models"
)

type cat1data struct {
	Record    models.Record
	Header    models.Header
	Measures  []models.Measure
	StartDate int64
	EndDate   int64
}

func timeToFormat(t int64, f string) string {
	parsedTime := time.Unix(t, 0)
	return parsedTime.Format(f)
}

func identifierFor(b []byte) string {
	md := md5.Sum(b)
	return strings.ToUpper(hex.EncodeToString(md[:]))
}

func identifierForInt(objs ...int64) string {
	b := make([]byte, len(objs))
	for _, val := range objs {
		b = append(b, []byte(strconv.FormatInt(val, 10))...)
	}
	return identifierFor(b)
}

func identifierForString(objs ...string) string {
	b := strings.Join(objs, ",")
	return identifierFor([]byte(b))
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
}

func uniqueDataCriteria(allDataCriteria []models.DataCriteria) map[dc]mdc {
	mappedDataCriteria := map[dc]mdc{}
	for _, dataCriteria := range allDataCriteria {
		oid := GetID(dataCriteria)
		dc := dc{DataCriteriaOid: oid, ValueSetOid: dataCriteria.CodeListID}
		var mappedDc = mappedDataCriteria[dc]
		if mappedDc.FieldOids == nil {
			mappedDc = mdc{DataCriteria: dataCriteria, FieldOids: make(map[string][]string)}
		}
		for field, descr := range dataCriteria.FieldValues {
			if descr.Type == "CD" {
				mappedDc.FieldOids[field] = append(mappedDc.FieldOids[field], descr.CodeListID)
			}
		}
		if dataCriteria.Negation {
			mappedDc.FieldOids["REASON"] = append(mappedDc.FieldOids["REASON"])
		}

		if dataCriteria.Value.Type == "CD" {
			mappedDc.ResultOids = append(mappedDc.ResultOids, dataCriteria.Value.CodeListID)
		}

		mappedDataCriteria[dc] = mappedDc
	}
	return mappedDataCriteria
}

//export GenerateCat1
func GenerateCat1(patient []byte, measures []byte, startDate int64, endDate int64) string {

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
	}

	cat1Template := template.New("cat1").Funcs(funcMap)

	for _, d := range data {
		asset, _ := Asset("templates/cat1/" + d)
		template.Must(cat1Template.New(d).Parse(string(asset)))
	}

	p := &models.Record{}
	m := []models.Measure{}
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

	c1d := cat1data{Record: *p, Header: *h, Measures: m, StartDate: startDate, EndDate: endDate}

	var b bytes.Buffer

	err = cat1Template.ExecuteTemplate(&b, "cat1.xml", c1d)

	if err != nil {
		fmt.Println(err)
	}

	spew.Dump(uniqueDataCriteria(allDataCriteria(m)))
	// uniqueDataCriteria(allDataCriteria(m))

	// return b.String()
	return ""
}
