package models

type Entry struct {
	Coded
	StartTime   int64               `json:"start_time,omitempty"`
	EndTime     int64               `json:"end_time,omitempty"`
	Time        int64               `json:"time,omitempty"`
	ID          CDAIdentifier       `json:"cda_identifier,omitempty"`
	Oid         string              `json:"oid,omitempty"`
	Description string              `json:"description,omitempty"`
	NegationInd bool                `json:"negationInd,omitempty"`
	Values      []ResultValue       `bson:"values,omitempty"`
	StatusCode  map[string][]string `json:"status_code,omitempty"`
	Reason      Reason              `json:"reason,omitempty"`
}

func (e *Entry) AddScalarValue(value int64, units string) {
	val := ResultValue{}
	val.Scalar = value
	val.Units = units
	e.Values = append(e.Values, val)
}

func (e *Entry) AddStringValue(value string, units string) {
	val := ResultValue{}
	val.Value = value
	val.Units = units
	e.Values = append(e.Values, val)
}
