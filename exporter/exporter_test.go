package exporter

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/pebbe/util"
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
	measureData = append([]byte("["), append(append(measureData, append([]byte(","), measureData2...)...), []byte("]")...)...)
	util.CheckErr(err)

	startDate := int64(1451606400)
	endDate := int64(1483228799)

	fmt.Println(GenerateCat1(patientData, measureData, startDate, endDate))
}

func (s *MySuite) TestImportHQMFTemplateJSON(c *C) {
	// c.Skip("skipped for the moment")
	var origID = "2.16.840.1.113883.10.20.28.3.19"
	var def = GetTemplateDefinition(origID)
	c.Assert(def.Definition, Equals, "diagnosis")
	c.Assert(def.Status, Equals, "resolved")
	var id = GetID(def)
	c.Assert(id, Equals, origID)
}
