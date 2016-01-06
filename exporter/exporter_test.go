package exporter

import (
	"fmt"
	"testing"
)

// This test is essentially a noop but it's useful to see what you're exporting.
// More functional tests are in the go-cda-repo where we run the exports through
//  HDS validation.
func TestExport(t *testing.T) {
	fmt.Println(GenerateCat1([]byte("Foo"), []byte("")))
}
