package models

type Encounter struct {
	Entry                `bson:",inline"`
	AdmitTime            int64             `json:"admitTime,omitempty"`
	DischargeTime        int64             `json:"discharge_time,omitempty"`
	DischargeDisposition map[string]string `json:"discharge_disposition,omitempty"`
	Facility             Facility          `json:"facility,omitempty"`
}

type Facility struct {
	Concept
	Name      string  `json:"name,omitempty"`
	StartTime int64   `json:"start_time,omitempty"`
	EndTime   int64   `json:"end_time,omitempty"`
	Address   Address `json:"address,omitempty"`
	Telecoms  []Telecom
}

func (enc *Encounter) GetEntry() *Entry {
	return &enc.Entry
}
