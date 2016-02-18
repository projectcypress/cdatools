package models

type Entry struct {
	StartTime   int64               `json:"start_time"`
	EndTime     int64               `json:"end_time"`
	Time        int64               `json:"time"`
	ID          CDAIdentifier       `json:"cda_identifier"`
	Oid         string              `json:"oid"`
	Description string              `json:"description"`
	Codes       map[string][]string `json:"codes"`
	NegationInd bool                `json:"negationInd"`
	Values      []ResultValue       `bson:"values"`
	StatusCode  map[string][]string `json:"status_code"`
	Reason      Reason              `json:"reason"`
}

func (e *Entry) AddCode(code string, codeSystem string) {
	if _, ok := e.Codes[codeSystem]; ok {
		e.Codes[codeSystem] = append(e.Codes[codeSystem], code)
	} else {
		e.Codes[codeSystem] = []string{code}
	}
}
