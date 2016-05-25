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
	First     string    `json:"first,omitempty"`
	Last      string    `json:"last,omitempty"`
	Birthdate int64     `json:"birthdate,omitempty"`
	Gender    string    `json:"gender,omitempty"`
	Race      Race      `json:"race,omitempty"`
	Ethnicity Ethnicity `json:"ethnicity,omitempty"`
}

type Race struct {
	Code    string `json:"code,omitempty"`
	CodeSet string `json:"code_set,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Ethnicity struct {
	Code    string `json:"code,omitempty"`
	CodeSet string `json:"code_set,omitempty"`
	Name    string `json:"name,omitempty"`
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
	Coded
}

type ID struct {
	Root      string
	Extension string
}

type Record struct {
	Person
	MedicalRecordNumber   string                `json:"medical_record_number,omitempty"`
	MedicalRecordAssigner string                `json:"medical_record_assigner,omitempty"`
	BirthDate             int64                 `json:"birthdate,omitempty"`
	DeathDate             int64                 `json:"deathdate,omitempty"`
	Expired               bool                  `json:"expired,omitempty"`
	Encounters            []Encounter           `json:"encounters,omitempty"`
	LabResults            []LabResult           `json:"results,omitempty"`
	Languages             []Language            `json:"languages,omitempty"`
	ProviderPerformances  []ProviderPerformance `json:"provider_performances,omitempty"`
	InsuranceProviders    []InsuranceProvider   `json:"insurance_providers,omitempty"`
	Procedures            []Procedure           `json:"procedures,omitempty"`
	Medications           []Medication          `json:"medications, omitempty"`
	Allergies             []Allergy             `json:"allergies,omitempty"`
	Conditions            []Condition           `json:"conditions,omitempty"`
	Communications        []Communication       `json:"communications,omitempty"`
	MedicalEquipment      []MedicalEquipment    `json:"medical_equipment,omitempty`
	CareGoals             []CareGoal            `json:"care_goals,omitempty"`
}

// Entries returns all the entries from the Encounters, Diagnoses, and LabResults for a Record
func (r *Record) Entries() []interface{} {
	var entries []interface{}

	// This whole "for loop for each of these things" is unavoidable, because elements must be copied individually to a []interface{}
	for _, en := range r.Encounters {
		entries = append(entries, en)
	}

	for _, lr := range r.LabResults {
		entries = append(entries, lr)
	}

	for _, ip := range r.InsuranceProviders {
		entries = append(entries, ip)
	}

	for _, pp := range r.ProviderPerformances {
		entries = append(entries, pp)
	}

	for _, pr := range r.Procedures {
		entries = append(entries, pr)
	}

	for _, md := range r.Medications {
		entries = append(entries, md)
	}

	for _, al := range r.Allergies {
		entries = append(entries, al)
	}

	for _, cn := range r.Conditions {
		entries = append(entries, cn)
	}

	return entries
}

// EntriesForOid returns all the entries which include the OID
func (r *Record) EntriesForOid(oid string) []interface{} {
	var matchedEntries []interface{}
	for _, entry := range r.Entries() {

		if ExtractEntry(&entry).Oid == oid {
			matchedEntries = append(matchedEntries, entry)
		}
	}
	return matchedEntries
}

type ResultValue struct {
	Coded
	Scalar    string `json:"scalar,omitempty"`
	Units     string `json:"units,omitempty"`
	StartTime int64  `json:"start_time,omitempty"`
	EndTime   int64  `json:"end_time,omitempty"`
}

type CDAIdentifier struct {
	Root      string `json:"root,omitempty"`
	Extension string `json:"extension,omitempty"`
}

type ProviderPerformance struct {
	Entry     `bson:",inline"`
	StartDate int64 `json:"startDate,omitempty"`
	EndDate   int64 `json:"endDate,omitempty"`
}

type Scalar struct {
	Unit  string `json:"unit,omitempty"`
	Value int64  `json:"value,omitempty"`
}
