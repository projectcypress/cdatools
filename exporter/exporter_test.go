package exporter

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/pebbe/util"
)

// This test is essentially a noop but it's useful to see what you're exporting.
// More functional tests are in the go-cda-repo where we run the exports through
//  HDS validation.
func TestExport(t *testing.T) {
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
