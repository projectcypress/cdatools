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

// Global template for Cat1 r3
var cat1r3Template *template.Template

// Global template for Cat1 r3_1
//var cat1r3_1Template *template.Template

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cat1r3Template = compileTemplates("r3")
	//cat1r3_1Template = compileTemplates("r3_1")
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

// Global Measure Data for a batch of patients
var m []models.Measure

// Global Value Set Data for a batch of patients
var vs []models.ValueSet

// Global ValueSetMap for a batch of patients
var vsMap models.ValueSetMap

// ExporterCat1Init initializes structures out of the measures and value sets
// given to be used for GenerateCat1.
func ExporterCat1Init(measures []byte, valueSets []byte) {
	json.Unmarshal(measures, &m)
	json.Unmarshal(valueSets, &vs)

	vsMap = models.NewValueSetMap(vs)
}

// GenerateCat1 creates the cat1 xml document for the patient given.
func GenerateCat1(patient []byte, startDate int64, endDate int64, qrdaVersion string) string {
	fmt.Println("length of measures: ", len(m))
	fmt.Println("length of value sets: ", len(vs))
	fmt.Println("qrda version given: ", qrdaVersion)

	p := &models.Record{}
	json.Unmarshal(patient, p)

	if qrdaVersion == "" {
		fmt.Println("qrdaVersion set to r3")
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

	var err error
	if qrdaVersion == "r3" {
		fmt.Println("printing r3 version of cat1")
		err = cat1r3Template.ExecuteTemplate(&b, "cat1.xml", c1d)
	} 
// else {
// 		fmt.Println("printing r3_1 version of cat1")
// 		err = cat1r3_1Template.ExecuteTemplate(&b, "cat1.xml", c1d)
// 	}
	if err != nil {
		fmt.Println(err)
	}

	return b.String()
}

func compileTemplates(version string) *template.Template {
	tmpl := template.New("cat1" + version)
	tmpl.Funcs(exporterFuncMap(tmpl, vsMap))

	data, err := AssetDir("templates/cat1/" + version)
	if err != nil {
		fmt.Println(err)
	}

	for _, d := range data {
		asset, _ := Asset("templates/cat1/" + version + "/" + d)
		template.Must(tmpl.New(d).Parse(string(asset)))
	}

	return tmpl
}
