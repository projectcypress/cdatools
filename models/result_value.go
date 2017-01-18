package models

import "strconv"

type ResultValue struct {
	Coded
	Scalar    string `json:"scalar,omitempty"`
	Units     string `json:"units,omitempty"`
	StartTime *int64 `json:"start_time,omitempty"`
	EndTime   *int64 `json:"end_time,omitempty"`
}

func (r *ResultValue) GetScalarType() string {
	if r.Scalar == "true" || r.Scalar == "false" {
		return "bool"
	}
	_, err := strconv.ParseFloat(r.Scalar, 64)
	if err == nil {
		return "num"
	} else {
		return "other"
	}
}

type ResultValueWrap struct {
	ResultValueEntry *Entry
	Values           []ResultValue
}
