package models

type Encounter struct {
	*Entry               `bson:",inline"`
	AdmitTime            int64             `json:"admitTime,omitempty"`
	DischargeTime        int64             `json:"discharge_time,omitempty"`
	DischargeDisposition map[string]string `json:"discharge_disposition,omitempty"`
}
