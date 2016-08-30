package models

type Medication struct {
	Entry                `bson:",inline"`
	AdministrationTiming administrationTiming `json:"administration_timing,omitempty"`
	Route                *CodedConcept        `json:"route,omitempty"`
	Dose                 Scalar               `json:"dose,omitempty"`
	AnatomicalApproach   *CodedConcept        `json:"anatomical_approach,omitempty"`
	DoseRestriction      doseRestriction      `json:"dose_restriction,omitempty"`
	ProductForm          *CodedConcept        `json:"product_form,omitempty"`
	DeliveryMethod       *CodedConcept        `json:"delivery_method,omitempty"`
	TypeOfMedication     *CodedConcept        `json:"type_of_medication,omitempty"`
	Indication           *CodedConcept        `json:"indication,omitempty"`
	Vehicle              *CodedConcept        `json:"vehicle,omitempty"`
	FulfillmentHistory   []FulfillmentHistory `json:"fulfillmentHistory,omitempty"`
	OrderInformation     []OrderInformation   `json:"orderInformation,omitempty"`
	CumulativeDuration   cumulativeDur        `json:"cumulative_medication_duration,omitempty"`
}

type cumulativeDur struct {
	Scalar int64  `json:"scalar,omitempty"`
	Unit   string `json:"units,omitempty"`
}

type administrationTiming struct {
	InstitutionSpecified bool   `json:"institutionSpecified,omitempty"`
	Period               Scalar `json:"period,omitempty"`
}

type doseRestriction struct {
	Numerator   Scalar `json:"numerator,omitempty"`
	Denominator Scalar `json:"denominator,omitempty"`
}

type FulfillmentHistory struct {
	PrescriptionNumber string `json:"prescriptionNumber,omitempty"`
	DispenseDate       int64  `json:"dispenseDate,omitempty"`
	QuantityDispensed  Scalar `json:"quantityDispensed,omitempty"`
	FillNumber         int64  `json:"fillNumber,omitempty"`
	FillStatus         string `json:"fillStatus"`
}

type OrderInformation struct {
	OrderNumber     string `json:"order_number,omitempty"`
	Fills           int64  `json:"fills,omitempty"`
	QuantityOrdered Scalar `json:"quantity_ordered,omitempty"`
	OrderExpiration int64  `json:"order_expiration_date_time,omitempty"`
	OrderDate       int64  `json:"order_date_time,omitempty"`
}

func (med *Medication) GetEntry() *Entry {
	return &med.Entry
}
