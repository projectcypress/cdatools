package exporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"text/template"
	"time"

	"github.com/pborman/uuid"
	"github.com/projectcypress/cdatools/exporter/cat3"
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
		"identifierForResultValue":                   identifierForResultValue,
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
		"hasAnatomicalLocation":                      hasAnatomicalLocation,
		"hasSeverity":                                hasSeverity,
		"codeDisplayWithPreferredCode":               vsMap.CodeDisplayWithPreferredCode,
		"codeDisplayWithPreferredCodeForField":       vsMap.CodeDisplayWithPreferredCodeForField,
		"codeDisplayWithPreferredCodeForResultValue": vsMap.CodeDisplayWithPreferredCodeForResultValue,
		"codeDisplayWithPreferredCodeAndLaterality":  vsMap.CodeDisplayWithPreferredCodeAndLaterality,
		"negationIndicator":                          negationIndicator,
		"isNil":                                      isNil,
		"derefBool":                                  derefBool,
		"emptyMdc":                                   models.EmptyMdc,
		"newRecordTarget":                            newRecordTarget,
		"sampleCodeSystemInCoded":                    sampleCodeSystemInCoded,
		"firstCodeInCodeSystem":                      firstCodeInCodeSystem,
	}
}

// Global Measure Data for a batch of patients
var m []models.Measure

// Global Value Set Data for a batch of patients
var vs []models.ValueSet

// Global ValueSetMap for a batch of patients
var vsMap models.ValueSetMap

func LoadMeasuresAndValueSets(measures []byte, valueSets []byte) {
	json.Unmarshal(measures, &m)
	json.Unmarshal(valueSets, &vs)
	vsMap = models.NewValueSetMap(vs)
}

// GenerateCat1 generates a cat1 xml string for export
func GenerateCat1(patient []byte, startDate int64, endDate int64, qrdaVersion string, cmsCompatibility bool) string {
	var failed bool
	var errmsg string
	result := generateCat1(patient, startDate, endDate, qrdaVersion, cmsCompatibility, &failed, &errmsg)
	if failed {
		return "Export Failed: " + errmsg
	}
	return result
}

func generateCat1(patient []byte, startDate int64, endDate int64, qrdaVersion string, cmsCompatibility bool, failed *bool, errmsg *string) string {
	defer catchPanics(failed, errmsg)

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
	cat1Template.Funcs(exporterFuncMap(cat1Template, vsMap))

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
	if len(m) > 0 && m[0].Type == "ep" {
		reportingProgram = "PQRS_MU_INDIVIDUAL"
	}

	c1d := cat1data{Record: *p, Header: h, Measures: m, ValueSets: vs, StartDate: startDate, EndDate: endDate, EntryInfos: p.EntryInfosForPatient(m, vsMap, qrdaVersion), CMSCompatibility: cmsCompatibility, ReportingProgram: reportingProgram}

	var b bytes.Buffer

	err = cat1Template.ExecuteTemplate(&b, "cat1.xml", c1d)

	if err != nil {
		fmt.Println(err)
	}

	return b.String()
}

func GenerateCat3(measures []byte, measure_results []byte, effectiveDate int64, startDate int64, endDate int64, version string) string {
	var failed bool
	var errmsg string
	result := generateCat3(measures, measure_results, effectiveDate, startDate, endDate, version, &failed, &errmsg)
	if failed {
		return "Export Failed: " + errmsg
	}
	return result
}

func generateCat3(measures []byte, measure_results []byte, effectiveDate int64, startDate int64, endDate int64, version string, failed *bool, errmsg *string) string {
	defer catchPanics(failed, errmsg)

	m := models.Measure{}
	mr := cat3.MeasureResults{}

	json.Unmarshal(measures, &m)
	json.Unmarshal(measure_results, &mr)

	aggCount := cat3.AggregateCount{
		Populations:      mr.Populations,
		PopulationGroups: mr.PopulationGroups,
	}
	aggCounts := map[string]cat3.AggregateCount{}
	aggCounts[m.HQMFID] = aggCount
	log.Println(aggCounts)
	ms := &cat3.MeasureSection{
		Measure: models.Measure{
			ID:        "measure test id",
			HQMFID:    m.HQMFID,
			Name:      m.Name,
			HQMFSetID: m.HQMFSetID,
		},
		Results: aggCounts,
	}

	// TODO: make header an argument to GenerateCat3()
	h := &models.Header{}
	h = nil

	if version == "" {
		version = "r2"
	}

	d := cat3.NewDoc(*h, *ms, m, startDate, endDate)

	return cat3.Print(d.Template(), d)
}

// This is to catch any panics in the code. If there is a panic, the status
// will be set to true and errmsg will contain the message passed into the panic.
// This function should only be used in the library API functions and a non-nil pointer
// should be passed in.
func catchPanics(failed *bool, errmsg *string) {
	if r := recover(); r != nil {
		*failed = true
		*errmsg = fmt.Sprintf("caught panic in cdatools: %v", r)
	}
}
