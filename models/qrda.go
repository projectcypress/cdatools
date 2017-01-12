package models

type Header struct {
	Authenticator Authenticator
	Authors       []Author
	Custodian     Author
}

type Author struct {
	Time *int64
	Person
	Entity
	Device
	Organization
}

type Authenticator struct {
	Author
}

type Device struct {
	Model string
	Name  string
}

type Entity struct {
	Ids       []CDAIdentifier `json:"ids,omitempty"`
	Addresses []Address       `json:"addresses,omitempty"`
	Telecoms  []Telecom       `json:"telecoms,omitempty"`
}

type Organization struct {
	Entity
	Name    string `json:"name,omitempty"`
	TagName string `json:"tag_name,omitempty"`
}

type Address struct {
	Street  []string `json:"street"`
	City    string   `json:"city"`
	State   string   `json:"state"`
	Zip     string   `json:"zip"`
	Country string   `json:"country"`
	Use     string   `json:"use"`
}

type Telecom struct {
	Use   string `json:"use"`
	Value string `json:"value"`
}

type CDAIdentifier struct {
	Root      string `json:"root,omitempty"`
	Extension string `json:"extension,omitempty"`
}

type Scalar struct {
	Unit  string `json:"unit,omitempty"`
	Value string `json:"value,omitempty"`
}
