package models

type Encounter struct {
	Entry                `bson:",inline"`
	AdmitTime            int64             `json:"admitTime,omitempty"`
	DischargeTime        int64             `json:"discharge_time,omitempty"`
	DischargeDisposition map[string]string `json:"discharge_disposition,omitempty"`
	TransferTo           Transfer          `json:"transferTo,omitempty"`
	TransferFrom         Transfer          `json:"transferFrom,omitempty"`
	Facility             Facility          `json:"facility,omitempty"`
	PrincipalDiagnosis   Coded             `json:"principal_diagnosis,omitempty"`
	Diagnosis            Coded             `json:"diagnosis,omitempty"`
}

type Facility struct {
	Name      string        `json:"name,omitempty"`
	Code      *CodedConcept `json:"code,omitempty"`
	StartTime int64         `json:"start_time,omitempty"`
	EndTime   int64         `json:"end_time,omitempty"`
	Addresses []Address     `json:"addresses,omitempty"`
	Telecoms  []Telecom     `json:"telecoms,omitempty"`
}

type Transfer struct {
	Coded
	CodedConcept
	Time int64 `json:"time,omitempty"`
}

func (enc *Encounter) GetEntry() *Entry {
	return &enc.Entry
}
