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
	EntryInfos       []models.EntryInfo
	Record           models.Record
	Header           *models.Header
	Measures         []models.Measure
	ValueSets        []models.ValueSet
	StartDate        int64
	EndDate          int64
	CMSCompatibility bool
	ReportingProgram string
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
		"newRecordTarget":                            newRecordTarget,
	}
}

// Global Cat1 Measure Data for a batch of patients
var cat1Measures []models.Measure

// Global Cat1 Value Set Data for a batch of patients
var cat1ValueSets []models.ValueSet

// Global Cat1 ValueSetMap for a batch of patients
var cat1ValueSetMap models.ValueSetMap

// ExporterCat1Init initializes structures out of the measures and value sets
// given to be used for GenerateCat1
func ExporterCat1Init(measures []byte, valueSets []byte) {
	json.Unmarshal(measures, &cat1Measures)
	json.Unmarshal(valueSets, &cat1ValueSets)
	cat1ValueSetMap = models.NewValueSetMap(cat1ValueSets)
}

// GenerateCat1 generates a cat1 xml string for export
func GenerateCat1(patient []byte, startDate int64, endDate int64, qrdaVersion string, cmsCompatibility bool) string {
	p := &models.Record{}
	json.Unmarshal(patient, p)

	if qrdaVersion == "" {
		qrdaVersion = "r3"
	}

	data, err := AssetDir("templates/cat1/" + qrdaVersion)
	if err != nil {
		fmt.Println(err)
	}

	cat1Template := template.New("cat1")
	cat1Template.Funcs(exporterFuncMap(cat1Template, cat1ValueSetMap))

	for _, d := range data {
		asset, _ := Asset("templates/cat1/" + qrdaVersion + "/" + d)
		template.Must(cat1Template.New(d).Parse(string(asset)))
	}
	var atime1 = new(int64)
	var atime2 = new(int64)
	*atime1 = 1449686219
	*atime2 = 1449686219

	// TODO: make header an argument to GenerateCat1()
	h := &models.Header{}
	h = nil

	reportingProgram := "HQR_EHR"
	if len(cat1Measures) > 0 && cat1Measures[0].Type == "ep" {
		reportingProgram = "PQRS_MU_INDIVIDUAL"
	}

	c1d := cat1data{Record: *p, Header: h, Measures: cat1Measures, ValueSets: cat1ValueSets, StartDate: startDate, EndDate: endDate, EntryInfos: p.EntryInfosForPatient(cat1Measures, cat1ValueSetMap, qrdaVersion), CMSCompatibility: cmsCompatibility, ReportingProgram: reportingProgram}

	var b bytes.Buffer

	err = cat1Template.ExecuteTemplate(&b, "cat1.xml", c1d)

	if err != nil {
		fmt.Println(err)
	}

	return b.String()
}
