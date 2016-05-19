package models

type Entry struct {
	Coded
	StartTime      int64               `json:"start_time,omitempty"`
	BSONID         string              `json:"bson_id,omitempty"`
	EndTime        int64               `json:"end_time,omitempty"`
	Time           int64               `json:"time,omitempty"`
	ID             CDAIdentifier       `json:"cda_identifier,omitempty"`
	Oid            string              `json:"oid,omitempty"`
	Description    string              `json:"description,omitempty"`
	NegationInd    bool                `json:"negationInd,omitempty"`
	NegationReason CodedConcept        `json:"negationReason,omitempty"`
	Values         []ResultValue       `json:"values,omitempty"`
	StatusCode     map[string][]string `json:"status_code,omitempty"`
	Reason         CodedConcept        `json:"reason,omitempty"`
	TransferTo     Transfer            `json:"transferTo,omitempty"`
	TransferFrom   Transfer            `json:"transferFrom,omitempty"`
	References     []Reference         `json:"references,omitempty"`
}

func ExtractEntry(entry interface{}) Entry {
	switch t := entry.(type) {
	case Encounter:
		return t.Entry
	case LabResult:
		return t.Entry
	case InsuranceProvider:
		return t.Entry
	case Procedure:
		return t.Entry
	case Allergy:
		return t.Entry
	case Medication:
		return t.Entry
	case Communication:
		return t.Entry
	case Condition:
		return t.Entry
	case Entry:
		return t
	default:
		panic("We don't know how to extract an Entry from this type")
	}
}

func (e *Entry) AddScalarValue(value string, units string) {
	val := ResultValue{}
	val.Scalar = value
	val.Units = units
	e.Values = append(e.Values, val)
}
