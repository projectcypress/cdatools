package models

type Header struct {
	Authenticator Authenticator
	Authors       []Author
	Custodian     Author
}

type Author struct {
	Time int64
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
	Ids       []ID
	Addresses []Address
	Telecoms  []Telecom
}

type Person struct {
	Entity
	First     string    `json:"first"`
	Last      string    `json:"last"`
	Gender    string    `json:"gender"`
	Birthdate int64     `json:"birthdate"`
	Race      Race      `json:"race"`
	Ethnicity Ethnicity `json:"ethnicity"`
}

type Race struct {
	Code    string `json:"code"`
	CodeSet string `json:"code_set"`
	Name    string `json:"name"`
}

type Ethnicity struct {
	Code    string `json:"code"`
	CodeSet string `json:"code_set"`
	Name    string `json:"name"`
}

type Organization struct {
	Entity
	Name    string
	TagName string
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

type Language struct {
	Codes map[string][]string
}

type ID struct {
	Root      string
	Extension string
}

type Record struct {
	Person
	MedicalRecordNumber   string                `json:"medical_record_number"`
	MedicalRecordAssigner string                `json:"medical_record_assigner"`
	Encounters            []Encounter           `json:"encounters"`
	Diagnoses             []Diagnosis           `json:"conditions"`
	LabResults            []LabResult           `json:"results"`
	Languages             []Language            `json:"languages"`
	ProviderPerformances  []ProviderPerformance `json:"provider_performances"`
}

type ResultValue struct {
	Scalar string              `json:"scalar"`
	Units  string              `json:"units"`
	Codes  map[string][]string `json:"codes"`
}

type Reason struct {
	Code           string `json:"code"`
	CodeSystem     string `json:"code_system"`
	CodeSystemName string `json:"codeSystemName"`
}

type CDAIdentifier struct {
	Root      string `json:"root"`
	Extension string `json:"extension"`
}

type ProviderPerformance struct {
	Entry     `bson:",inline"`
	StartDate int64 `json:"startDate"`
	EndDate   int64 `json:"endDate"`
}
