package exporter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"testing"
	"text/template"

	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
	"github.com/stretchr/testify/assert"
)

// func TestCode(t *testing.T) {
// }

// test _2.16.840.1.113883.10.20.24.3.23.xml
func TestEncounterPerformed(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.23"
	fmt.Print(generateQrdaOidXML(qrdaOid))

	assert.Equal(t, 1, 1) // should make better test
}

// func generateXML(fileName string, qrdaOid string) string {
// }

func generateQrdaOidXML(qrdaOid string) string {
	templateFileName := "_" + qrdaOid + ".xml"

	var p models.Record
	var m []models.Measure
	var vs []models.ValueSet
	setPatientMeasuresAndValueSets(&p, &m, &vs)

	cat1Template := makeTemplate()

	var buf bytes.Buffer
	entryInfo, err := getEntryInfo(p, m, qrdaOid)
	if err != nil {
		util.CheckErr(err)
	}
	if err := cat1Template.ExecuteTemplate(&buf, templateFileName, entryInfo); err != nil {
		util.CheckErr(err)
	}
	return buf.String()
}

func makeTemplate() *template.Template {
	temp := template.New("cat1")
	temp.Funcs(exporterFuncMap(temp))
	fileNames, err := AssetDir("templates/cat1")
	if err != nil {
		util.CheckErr(err)
	}
	for _, fileName := range fileNames {
		asset, err := Asset("templates/cat1/" + fileName)
		util.CheckErr(err)
		template.Must(temp.New(fileName).Parse(string(asset)))
	}
	return temp
}

func setPatientMeasuresAndValueSets(patient *models.Record, measures *[]models.Measure, valueSets *[]models.ValueSet) {
	patientData, err := ioutil.ReadFile("../fixtures/records/1_n_n_ami.json")
	util.CheckErr(err)

	measureData, err := ioutil.ReadFile("../fixtures/measures/CMS9v4a.json")
	util.CheckErr(err)
	measureData2, err := ioutil.ReadFile("../fixtures/measures/CMS26v3.json")
	util.CheckErr(err)
	measureData = append([]byte("["), append(append(measureData, append([]byte(","), measureData2...)...), []byte("]")...)...)

	valueSetData, err := ioutil.ReadFile("../fixtures/value_sets/cms9_26.json")
	util.CheckErr(err)

	json.Unmarshal(patientData, patient)
	json.Unmarshal(measureData, measures)
	json.Unmarshal(valueSetData, valueSets)
	initializeVsMap(*valueSets)
}

func getEntryInfo(patient models.Record, measures []models.Measure, qrdaOid string) (entryInfo, error) {
	entryInfos := entryInfosForPatient(patient, measures)
	for _, ei := range entryInfos {
		if qrdaOid == HqmfToQrdaOid(models.ExtractEntry(&ei.EntrySection).Oid) {
			return ei, nil
		}
	}
	if len(entryInfos) == 0 {
		return entryInfo{}, errors.New("no entry infos found for patient and measures")
	}
	return entryInfo{}, errors.New(fmt.Sprintf("no entry info found with qrda oid \"%s\"", qrdaOid))
}
