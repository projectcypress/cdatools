package importer

import (
	"fmt"
	"io/ioutil"

	"github.com/pebbe/util"

	. "gopkg.in/check.v1"
)

type ResultsSuite struct{}

var _ = Suite(&ResultsSuite{})

func (i *ResultsSuite) TestCat3ResultsExtractor(c *C) {
	data, err := ioutil.ReadFile("../fixtures/cat3.xml")
	util.CheckErr(err)
	doc := string(data)

	measureID := "40280381-4C18-79DF-014C-291EF3F90654"
	ids := map[string]string{
		"IPP":     "EAD808CB-A6FA-4824-A204-74F299839396",
		"MSRPOPL": "7462E67A-5ECB-41D6-AE14-2E89BB55BBDE",
		"OBSERV":  "2D084067-703B-4072-9F43-D50F938F4F9C",
	}

	result := ExtractResultsByIds(measureID, ids, doc)
	fmt.Println(result)
	c.Assert(0, Equals, 1)
}
