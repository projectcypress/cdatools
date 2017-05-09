package exporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"text/template"
	"time"

	"github.com/pborman/uuid"
	"github.com/projectcypress/cdatools/models"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type cat1data struct {
	EntryInfos []models.EntryInfo
	Record     models.Record
	Header     models.Header
	Measures   []models.Measure
	ValueSets  []models.ValueSet
	StartDate  int64
	EndDate    int64
}

func exporterFuncMap(cat1Template *template.Template, vsMap models.ValueSetMap) template.FuncMap {
	return template.FuncMap{
		"timeNow":                                    time.Now().UTC().Unix,
		"newRandom":                                  uuid.NewRandom,
		"timeToFormat":                               timeToFormat,
		"identifierForInt":                           identifierForInt,
		"identifierForIntp":                          identifierForIntp,
		"identifierForString":                        identifierForString,
		"escape":                                     escape,
		"executeTemplateForEntry":                    generateExecuteTemplateForEntry(cat1Template),
		"condAssign":                                 condAssign,
		"valueOrNullFlavor":                          valueOrNullFlavor,
		"dischargeDispositionDisplay":                dischargeDispositionDisplay,
		"sdtcValueSetAttribute":                      sdtcValueSetAttribute,
		"getTransferOid":                             getTransferOid,
		"identifierForInterface":                     identifierForInterface,
		"valueOrDefault":                             valueOrDefault,
		"oidForCodeSystem":                           oidForCodeSystem,
		"oidForCode":                                 vsMap.OidForCode,
		"codeDisplayAttributeIsCodes":                codeDisplayAttributeIsCodes,
		"hasPreferredCode":                           hasPreferredCode,
		"hasLaterality":                              hasLaterality,
		"codeDisplayWithPreferredCode":               vsMap.CodeDisplayWithPreferredCode,
		"codeDisplayWithPreferredCodeForResultValue": vsMap.CodeDisplayWithPreferredCodeForResultValue,
		"codeDisplayWithPreferredCodeAndLaterality":  vsMap.CodeDisplayWithPreferredCodeAndLaterality,
		"negationIndicator":                          negationIndicator,
		"isNil":                                      isNil,
		"derefBool":                                  derefBool,
		"emptyMdc":                                   models.EmptyMdc,
	}
}

// Global template for Cat1
var cat1Template *template.Template

// Global Measure Data for a batch of patients
var m []models.Measure

// Global Value Set Data for a batch of patients
var vs []models.ValueSet

// Global ValueSetMap for a batch of patients
var vsMap models.ValueSetMap

//export ExporterCat1Init
func ExporterCat1Init(measures []byte, valueSets []byte, version string) {
	json.Unmarshal(measures, &m)
	json.Unmarshal(valueSets, &vs)

	vsMap = models.NewValueSetMap(vs)

	cat1Template = template.New("cat1" + version)
	cat1Template.Funcs(exporterFuncMap(cat1Template, vsMap))

	data, err := AssetDir("templates/cat1/" + version)
	if err != nil {
		fmt.Println(err)
	}

	for _, d := range data {
		asset, _ := Asset("templates/cat1/" + version + "/" + d)
		template.Must(cat1Template.New(d).Parse(string(asset)))
	}
}

//export GenerateCat1
func GenerateCat1(patient []byte, startDate int64, endDate int64, qrdaVersion string) string {

	p := &models.Record{}
	json.Unmarshal(patient, p)

	if qrdaVersion == "" {
		qrdaVersion = "r3"
	}

	var atime1 = new(int64)
	var atime2 = new(int64)
	*atime1 = 1449686219
	*atime2 = 1449686219
	h := &models.Header{
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

	c1d := cat1data{Record: *p, Header: *h, Measures: m, ValueSets: vs, StartDate: startDate, EndDate: endDate, EntryInfos: p.EntryInfosForPatient(m, vsMap, qrdaVersion)}

	var b bytes.Buffer

	err := cat1Template.ExecuteTemplate(&b, "cat1.xml", c1d)

	if err != nil {
		fmt.Println(err)
	}

	return b.String()
}
