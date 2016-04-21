package models

import (
	"github.com/pebbe/util"
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

type ProcedureSuite struct {
}

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&ProcedureSuite{})

func (p *ProcedureSuite) TestMarshalOrdinality(c *C) {
	procedure := &Procedure{
		Ordinality: Coded{
			Codes: map[string][]string{"SNOMED-CT": []string{"63161005"}},
		},
	}

	expectedOrdinality := bson.M{
		"SNOMED-CT": []interface{}{"63161005"},
	}

	data, err := bson.Marshal(procedure) // call GetBSON() for procedure
	util.CheckErr(err)

	// unmarshal data back to a map so this map can be checked against the expected
	var procedureMap bson.M
	err = bson.Unmarshal(data, &procedureMap)
	util.CheckErr(err)

	c.Assert(procedureMap["ordinality"], DeepEquals, expectedOrdinality)
}

func (p *ProcedureSuite) TestUnmarshalOrdinality(c *C) {

	// marshal bson map to BSON bytestream
	data, err := bson.Marshal(bson.M{
		"ordinality": bson.M{
			"SNOMED-CT": []interface{}{"63161005"},
		},
		"incision_time": int64(5), // to test that all procedure attributes are coppied
	})
	util.CheckErr(err)

	// expect procedure to have Codes nested under Ordinality after SetBSON
	expectedProcedure := Procedure{
		Ordinality: Coded{
			Codes: map[string][]string{"SNOMED-CT": []string{"63161005"}},
		},
		IncisionTime: int64(5),
	}

	var procedure Procedure
	err = bson.Unmarshal(data, &procedure) // call SetBSON() for procedure
	util.CheckErr(err)

	c.Assert(procedure, DeepEquals, expectedProcedure)
}
