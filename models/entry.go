package models

type Entry struct {
	StartTime   int64               `json:"start_time,omitempty"`
	EndTime     int64               `json:"end_time,omitempty"`
	Time        int64               `json:"time,omitempty"`
	ID          CDAIdentifier       `json:"cda_identifier,omitempty"`
	Oid         string              `json:"oid,omitempty"`
	Description string              `json:"description,omitempty"`
	Codes       map[string][]string `json:"codes,omitempty"`
	NegationInd bool                `json:"negationInd,omitempty"`
	Values      []ResultValue       `bson:"values,omitempty"`
	StatusCode  map[string][]string `json:"status_code,omitempty"`
	Reason      Reason              `json:"reason,omitempty"`
}

func (e *Entry) AddCode(code string, codeSystem string) {
	if _, ok := e.Codes[codeSystem]; ok {
		e.Codes[codeSystem] = append(e.Codes[codeSystem], code)
	} else {
		e.Codes[codeSystem] = []string{code}
	}
}
