package exporter

import (
	"fmt"
	"testing"

	"github.com/projectcypress/cdatools/fixtures"
)

// This test is essentially a noop but it's useful to see what you're exporting.
// More functional tests are in the go-cda-repo where we run the exports through
//  HDS validation.
func TestExport(t *testing.T) {
	t.Skip()
	measureData := append([]byte("["), append(append(fixtures.Cms9v4a, append([]byte(","), fixtures.Cms26v3...)...), []byte("]")...)...)

	startDate := int64(1451606400)
	endDate := int64(1483228799)
	LoadMeasuresAndValueSets(measureData, fixtures.Cms9_26)
	fmt.Print(GenerateCat1(fixtures.TestPatientDataAmi, startDate, endDate, "r3", true))
}
