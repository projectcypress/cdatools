package models

type Coded interface {
	Codes() map[string][]string
	SetCodes(codes map[string][]string)
}

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
	Name string
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
	MedicalRecordNumber   string      `json:"medical_record_number"`
	MedicalRecordAssigner string      `json:"medical_record_assigner"`
	Encounters            []Encounter `json:"encounters"`
	Diagnoses             []Diagnosis `json:"conditions"`
	Languages             []Language  `json:"languages"`
}

type ResultValue struct {
	Scalar string              `json:"scalar"`
	Units  string              `json:"units"`
	codes  map[string][]string `json:"codes"`
}

func (rv *ResultValue) SetCodes(codes map[string][]string) {
	rv.codes = codes
}

func NewResultValue() *ResultValue {
	rv := new(ResultValue)
	rv.codes = make(map[string][]string)
	return rv
}

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
}

type CDAIdentifier struct {
	Root      string `json:"root"`
	Extension string `json:"extension"`
}

func NewEntry() *Entry {
	entry := new(Entry)
	entry.Codes = make(map[string][]string)
	return entry
}

func (entry *Entry) AddResultValue(rv *ResultValue) {
	entry.Values = append(entry.Values, *rv)
}

func (entry *Entry) SetCodes(codes map[string][]string) {
	entry.Codes = codes
}

var oidMap = map[string]string{
	"2.16.840.1.113883.6.12":  "CPT",
	"2.16.840.1.113883.6.1":   "LOINC",
	"2.16.840.1.113883.6.96":  "SNOMED-CT",
	"2.16.840.1.113883.6.88":  "RxNorm",
	"2.16.840.1.113883.6.103": "ICD-9-CM",
	"2.16.840.1.113883.6.104": "ICD-9-PCS",
	"2.16.840.1.113883.6.4":   "ICD-10-PCS",
	"2.16.840.1.113883.6.90":  "ICD-10-CM",
}

func CodeSystemFor(oid string) string {
	return oidMap[oid]
}

func AddCode(coded Coded, code, codeSystem string) {
	codeSystemName := CodeSystemFor(codeSystem)
	coded.Codes()[codeSystemName] = append(coded.Codes()[codeSystemName], code)
}

type Encounter struct {
	Entry     `bson:",inline"`
	AdmitTime int64 `json:"admitTime"`
}

type Diagnosis struct {
	Entry `bson:",inline"`
}
