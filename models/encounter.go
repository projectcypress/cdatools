package models

type Encounter struct {
	Entry                `bson:",inline"`
	AdmitTime            *int64             `json:"admitTime,omitempty"`
	DischargeTime        *int64             `json:"discharge_time,omitempty"`
	DischargeDisposition map[string]string  `json:"dischargeDisposition,omitempty"`
	TransferTo           Transfer           `json:"transferTo,omitempty"`
	TransferFrom         Transfer           `json:"transferFrom,omitempty"`
	Facility             Facility           `json:"facility,omitempty"`
	PrincipalDiagnosis   PrincipalDiagnosis `json:"principalDiagnosis,omitempty"`
	Diagnosis            Diagnosis          `json:"diagnosis,omitempty"`
}

type Facility struct {
	Name      string        `json:"name,omitempty"`
	Code      *CodedConcept `json:"code,omitempty"`
	StartTime *int64        `json:"start_time,omitempty"`
	EndTime   *int64        `json:"end_time,omitempty"`
	Addresses []Address     `json:"addresses,omitempty"`
	Telecoms  []Telecom     `json:"telecoms,omitempty"`
}

type Transfer struct {
	Coded
	CodedConcept
	Time *int64 `json:"time,omitempty"`
}

type PrincipalDiagnosis struct {
	Coded
	CodedConcept
}

type Diagnosis struct {
	Coded
	CodedConcept
}

func (enc *Encounter) GetEntry() *Entry {
	return &enc.Entry
}

// IsEmpty determines if any fields in a Facility have been set to their non-zero values
func (f Facility) IsEmpty() bool {
	return f.Name == "" && f.Code == nil && f.StartTime == nil && f.EndTime == nil &&
		len(f.Addresses) == 0 && len(f.Telecoms) == 0
}
