package exporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"

	"github.com/projectcypress/cdatools/models"
)

type cat1data struct {
	Record           models.Record
	Header           models.Header
	OrganizationName string
}

//export generate_cat1
func Generate_cat1(patient []byte) string {

	data, err := AssetDir("templates/cat1")
	if err != nil {
		fmt.Println(err)
	}
	cat1_template := template.New("cat1")
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
					Name: "authorsOrganization",
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
					Name: "LegalAuthenticatorOrg",
				},
			},
		},
	}

	c1d := cat1data{Record: *p, Header: *h}

	json.Unmarshal(patient, p)

	var b bytes.Buffer

	cat1_template.ExecuteTemplate(&b, "cat1.xml", c1d)

	return b.String()
}
