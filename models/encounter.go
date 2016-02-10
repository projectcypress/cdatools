package models

type Encounter struct {
	Entry                `bson:",inline"`
	AdmitTime            int64               `json:"admitTime"`
	DischargeTime        int64               `json:"discharge_time"`
	DischargeDisposition map[string][]string `json:"discharge_disposition"`
}
