package models

import "gopkg.in/mgo.v2/bson"

type Procedure struct {
	Entry            `bson:",inline"`
	Ordinality       Coded        `json:"ordinality,omitempty" bson:"ordinality,omitempty"`
	Performer        Performer    `json:"performer,omitempty" bson:"performer,omitempty"`
	AnatomicalTarget CodedConcept `json:"anatomical_target,omitempty" bson:"anatomical_target,omitempty"`
	IncisionTime     int64        `json:"incision_time,omitempty" bson:"incision_time,omitempty"`
}

type Performer struct {
}

func (proc *Procedure) GetEntry() *Entry {
	return &proc.Entry
}

// used so bson.Marshal() and bson.Unmarshal() will not recursively call GetBSON() and SetBSON() respectively
type ProcedureNoFunc Procedure

func (procedure Procedure) GetBSON() (interface{}, error) {

	// marshal and unmarshal to copy all attributes to procedureMap. this would happen by default if we did not define a custom GetBSON() function
	data, err := bson.Marshal(ProcedureNoFunc(procedure))
	if err != nil {
		return nil, err
	}
	var procedureMap bson.M
	if err := bson.Unmarshal(data, &procedureMap); err != nil {
		return nil, err
	}

	// set ordinality.codes to ordinality (for health data standards)
	procedureMap["ordinality"] = procedureMap["ordinality"].(bson.M)["codes"]
	return procedureMap, nil
}

func (procedure *Procedure) SetBSON(raw bson.Raw) error {

	// unmarshal all attributes into procedure. this would happen by default if we did not define a custom SetBSON() function
	procedureNoFunc := ProcedureNoFunc(*procedure)
	if err := raw.Unmarshal(&procedureNoFunc); err != nil {
		return err
	}
	*procedure = Procedure(procedureNoFunc)

	// find the ordinality attribute and unmarshal the value into the coded variable
	var coded map[string][]string
	var rawData bson.RawD
	if err := raw.Unmarshal(&rawData); err != nil {
		return err
	}
	for i := range rawData {
		switch rawData[i].Name {
		case "ordinality":
			if err := rawData[i].Value.Unmarshal(&coded); err != nil {
				return err
			}
		}
	}

	// set Ordinality with embeded Codes
	procedure.Ordinality.Codes = coded
	return nil
}
