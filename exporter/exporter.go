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
	Record models.Record
	Header models.Header
}

func timeToCdaFormat(t int64) string {
	parsedTime := time.Unix(t, 0)
	return parsedTime.Format("20060102")
}

//export generate_cat1
func Generate_cat1(patient []byte) string {

	data, err := AssetDir("templates/cat1")
	if err != nil {
		fmt.Println(err)
	}

	funcMap := template.FuncMap{
		"timeNow":         time.Now().UTC().Unix,
		"newRandom":       uuid.NewRandom,
		"timeToCdaFormat": timeToCdaFormat,
	}

	cat1_template := template.New("cat1").Funcs(funcMap)

	for _, d := range data {
		asset, _ := Asset("templates/cat1/" + d)
		template.Must(cat1_template.New(d).Parse(string(asset)))
	}

	p := &models.Record{}
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

	c1d := cat1data{Record: *p, Header: *h}

	var b bytes.Buffer

	err = cat1_template.ExecuteTemplate(&b, "cat1.xml", c1d)

	if err != nil {
		fmt.Println(err)
	}

	return b.String()
}
