package exporter

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"

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

func timeToCdaFormat(t int64) string {
	parsedTime := time.Unix(t, 0)
	return parsedTime.Format("20060102")
}

func timeToFormat(t int64, f string) string {
	parsedTime := time.Unix(t, 0)
	return parsedTime.Format(f)
}

func identifierForInt(objs ...int64) string {
	b := make([]byte, len(objs))
	for _, val := range objs {
		b = append(b, []byte(strconv.FormatInt(val, 10))...)
	}
	md := md5.Sum(b)
	return strings.ToUpper(hex.EncodeToString(md[:]))
}

func identifierForString(objs ...string) string {
	b := strings.Join(objs, ",")
	md := md5.Sum([]byte(b))
	return strings.ToUpper(hex.EncodeToString(md[:]))
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
		"timeToCdaFormat":     timeToCdaFormat,
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

	return b.String()
}
