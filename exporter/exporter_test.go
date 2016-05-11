package exporter

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"fmt"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

// This test is essentially a noop but it's useful to see what you're exporting.
// More functional tests are in the go-cda-repo where we run the exports through
//  HDS validation.
func (s *MySuite) TestExport(c *C) {
	patientData, err := ioutil.ReadFile("../fixtures/records/barry_berry.json")
	util.CheckErr(err)

	measureData, err := ioutil.ReadFile("../fixtures/measures/CMS9v4a.json")
	measureData2, err := ioutil.ReadFile("../fixtures/measures/CMS26v3.json")
	valueSetData, err := ioutil.ReadFile("../fixtures/value_sets/cms9_26.json")
	measureData = append([]byte("["), append(append(measureData, append([]byte(","), measureData2...)...), []byte("]")...)...)
	util.CheckErr(err)

	startDate := int64(1451606400)
	endDate := int64(1483228799)

	fmt.Print(GenerateCat1(patientData, measureData, valueSetData, startDate, endDate))
}

func (s *MySuite) TestGetEntriesForDataCriteria(c *C) {
	patientData, err := ioutil.ReadFile("../fixtures/records/1_n_n_ami.json")
	util.CheckErr(err)

	measureData, err := ioutil.ReadFile("../fixtures/measures/CMS9v4a.json")
	util.CheckErr(err)

	patient := &models.Record{}
	measure := &models.Measure{}
	valueSetData, err := ioutil.ReadFile("../fixtures/value_sets/cms9_26.json")
	var vs []models.ValueSet
	json.Unmarshal(valueSetData, &vs)
	initializeVsMap(vs)

	json.Unmarshal(patientData, patient)
	json.Unmarshal(measureData, measure)

	var entries []interface{}
	for _, crit := range measure.HQMFDocument.DataCriteria {
		if crit.HQMFOid != "" {
			entries = append(entries, entriesForDataCriteria(crit, *patient))
		}
	}
	// TODO: This test will have to change when we get a new export of CMS9v4a with all the HQMFOid fields filled.
	c.Assert(len(entries), Equals, 1)
}

func (s *MySuite) TestImportHQMFTemplateJSON(c *C) {
	var origID = "2.16.840.1.113883.10.20.28.3.19"
	var def = GetTemplateDefinition(origID, true)
	c.Assert(def.Definition, Equals, "diagnosis")
	c.Assert(def.Status, Equals, "resolved")
	var id = GetID(def, true)
	c.Assert(id, Equals, origID)
}

func (s *MySuite) TestGetAllDataCriteriaForOneMeasure(c *C) {
	mes := make([]models.Measure, 1)
	measureData, err := ioutil.ReadFile("../fixtures/measures/CMS9v4a.json")
	util.CheckErr(err)
	measureData = append([]byte("["), append(measureData, []byte("]")...)...)
	json.Unmarshal(measureData, &mes)
	c.Assert(len(allDataCriteria(mes)), Equals, 27)
}

func (s *MySuite) TestGetallDatacriteriaForMultipleMeasures(c *C) {
	mes := make([]models.Measure, 2)
	measureData, err := ioutil.ReadFile("../fixtures/measures/CMS9v4a.json")
	measureData2, err := ioutil.ReadFile("../fixtures/measures/CMS26v3.json")
	util.CheckErr(err)
	measureData = append([]byte("["), append(append(measureData, append([]byte(","), measureData2...)...), []byte("]")...)...)
	json.Unmarshal(measureData, &mes)

	c.Assert(len(allDataCriteria(mes)), Equals, 47)
}

func (s *MySuite) TestGetUniqueDataCriteriaForOneMeasure(c *C) {
	mes := make([]models.Measure, 1)
	measureData, err := ioutil.ReadFile("../fixtures/measures/CMS9v4a.json")
	util.CheckErr(err)
	measureData = append([]byte("["), append(measureData, []byte("]")...)...)
	json.Unmarshal(measureData, &mes)
	c.Assert(len(uniqueDataCriteria(allDataCriteria(mes))), Equals, 15)
}
