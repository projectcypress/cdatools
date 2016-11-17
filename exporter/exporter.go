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
	EntryInfos []models.EntryInfo
	Record     models.Record
	Header     models.Header
	Measures   []models.Measure
	ValueSets  []models.ValueSet
	StartDate  int64
	EndDate    int64
}

func exporterFuncMap(cat1Template *template.Template) template.FuncMap {
	return template.FuncMap{
		"timeNow":                      time.Now().UTC().Unix,
		"newRandom":                    uuid.NewRandom,
		"timeToFormat":                 timeToFormat,
		"identifierForInt":             identifierForInt,
		"identifierForString":          identifierForString,
		"escape":                       escape,
		// mattrbianchi: I grepped for this and it wasn't used in any of the templates... should it really be in here?
		//"entryInfosForPatient":         models.Record.EntryInfosForPatient,
		"executeTemplateForEntry":      generateExecuteTemplateForEntry(cat1Template),
		"condAssign":                   condAssign,
		"valueOrNullFlavor":            valueOrNullFlavor,
		"dischargeDispositionDisplay":  dischargeDispositionDisplay,
		"sdtcValueSetAttribute":        sdtcValueSetAttribute,
		"getTransferOid":               getTransferOid,
		"identifierForInterface":       identifierForInterface,
		"codeToDisplay":                codeToDisplay,
		"valueOrDefault":               valueOrDefault,
		"reasonValueSetOid":            models.ReasonValueSetOid,
		"oidForCodeSystem":             oidForCodeSystem,
		"oidForCode":                   models.OidForCode,
		"codeDisplayAttributeIsCodes":  codeDisplayAttributeIsCodes,
		"hasPreferredCode":             hasPreferredCode,
		"codeDisplayWithPreferredCode": codeDisplayWithPreferredCode,
		"negationIndicator":            negationIndicator,
	}
}

//export GenerateCat1
func GenerateCat1(patient []byte, measures []byte, valueSets []byte, startDate int64, endDate int64, qrdaVersion string) string {

	if qrdaVersion == "" {
		qrdaVersion = "r3"
	}

	data, err := AssetDir("templates/cat1/" + qrdaVersion)
	if err != nil {
		fmt.Println(err)
	}

	cat1Template := template.New("cat1")
	cat1Template.Funcs(exporterFuncMap(cat1Template))

	for _, d := range data {
		asset, _ := Asset("templates/cat1/" + qrdaVersion + "/" + d)
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
				Time: 1449686219,
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

	json.Unmarshal(patient, p)
	json.Unmarshal(measures, &m)
	json.Unmarshal(valueSets, &vs)

	vsMap := models.InitializeVsMap(vs)

	c1d := cat1data{Record: *p, Header: *h, Measures: m, ValueSets: vs, StartDate: startDate, EndDate: endDate, EntryInfos: p.EntryInfosForPatient(m, vsMap)}

	var b bytes.Buffer

	err = cat1Template.ExecuteTemplate(&b, "cat1.xml", c1d)

	if err != nil {
		fmt.Println(err)
	}

	return b.String()
}
