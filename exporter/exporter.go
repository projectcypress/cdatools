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
		"createRecordTarget":                         createRecordTarget,
	}
}

// GenerateCat1 generates a cat1 xml string for export
func GenerateCat1(patient []byte, measures []byte, valueSets []byte, startDate int64, endDate int64, qrdaVersion string, cmsCompatibility bool) string {

	p := &models.Record{}
	m := []models.Measure{}
	vs := []models.ValueSet{}

	json.Unmarshal(patient, p)
	json.Unmarshal(measures, &m)
	json.Unmarshal(valueSets, &vs)

	vsMap := models.NewValueSetMap(vs)

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
