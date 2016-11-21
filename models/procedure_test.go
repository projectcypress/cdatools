package models

import (
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
	"log"
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
	if err != nil {
		log.Fatalln(err)
	}

	// unmarshal data back to a map so this map can be checked against the expected
	var procedureMap bson.M
	err = bson.Unmarshal(data, &procedureMap)
	if err != nil {
		log.Fatalln(err)
	}

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
	if err != nil {
		log.Fatalln(err)
	}

	// expect procedure to have Codes nested under Ordinality after SetBSON
	expectedProcedure := Procedure{
		Ordinality: Coded{
			Codes: map[string][]string{"SNOMED-CT": []string{"63161005"}},
		},
		IncisionTime: int64(5),
	}

	var procedure Procedure
	err = bson.Unmarshal(data, &procedure) // call SetBSON() for procedure
	if err != nil {
		log.Fatalln(err)
	}

	c.Assert(procedure, DeepEquals, expectedProcedure)
}
