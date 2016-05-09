package models

// InsuranceProvider for a Record. Provides Payer information
type InsuranceProvider struct {
	Entry                       `bson:",inline"`
	Type                        string            `json:"type,omitempty"`
	MemberID                    string            `json:"member_id,omitempty"`
	Payer                       Organization      `json:"payer,omitempty"`
	Guarantors                  []Guarantor       `json:"guarantors,omitempty"`
	Subscriber                  Person            `json:"subscriber,omitempty"`
	FinancialResponsibilityType map[string]string `json:"financial_responsiblity_type,omitempty"`
	Relationship                map[string]string `json:"relationship,omitempty"`
}

// Guarantor organization/person for an insurance provider for a Record
type Guarantor struct {
	Organization Organization `json:"organization,omitempty"`
	Person       Person       `json:"person,omitempty"`
	Time         int64        `json:"time,omitempty"`
	StartTime    int64        `json:"start_time,omitempty"`
	EndTime      int64        `json:"end_time,omitempty"`
}

// ShiftDates adds dateDiff to start/end times/other times for Insurance Providers
func (p *InsuranceProvider) ShiftDates(dateDiff int64) {
	p.StartTime = shiftDate(p.StartTime, dateDiff)
	p.EndTime = shiftDate(p.EndTime, dateDiff)
	p.Time = shiftDate(p.Time, dateDiff)
	for _, g := range p.Guarantors {
		g.ShiftDates(dateDiff)
	}
}

// ShiftDates adds dateDiff to start/end times for Guarantors
func (g *Guarantor) ShiftDates(dateDiff int64) {
	g.StartTime = shiftDate(g.StartTime, dateDiff)
	g.EndTime = shiftDate(g.EndTime, dateDiff)
}

func shiftDate(date int64, dateDiff int64) int64 {
	if date != 0 {
		return date + dateDiff
	}
	return 0
}
