package models

type Transfer struct {
  Coded
  Reason
  Time  int64 `json:"time,omitempty"`
}
