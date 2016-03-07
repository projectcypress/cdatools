package exporter

import (
	"fmt"
	"testing"
	"github.com/pebbe/util"
	"io/ioutil"
)

// This test is essentially a noop but it's useful to see what you're exporting.
// More functional tests are in the go-cda-repo where we run the exports through
//  HDS validation.
func TestExport(t *testing.T) {
	patientData, err := ioutil.ReadFile("../fixtures/records/barry_berry.json")
	util.CheckErr(err)

	measureData, err := ioutil.ReadFile("../fixtures/measures/CMS9v4a.json")
	util.CheckErr(err)
	fmt.Println(string(measureData))

	fmt.Println(GenerateCat1(patientData, measureData))
}
