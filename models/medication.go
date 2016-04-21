package models

type Medication struct {
	Entry                `bson:",inline"`
	AdministrationTiming administrationTiming `json:"administration_timing,omitempty"`
	Route                Coded                `json:"route,omitempty"`
	Dose                 Scalar               `json:"dose,omitempty"`
	AnatomicalApproach   Coded                `json:"anatomical_approach,omitempty"`
	DoseRestriction      doseRestriction      `json:"dose_restriction,omitempty"`
	ProductForm          Coded                `json:"product_form,omitempty"`
	DeliveryMethod       Coded                `json:"delivery_method,omitempty"`
	TypeOfMedication     Coded                `json:"type_of_medication,omitempty"`
	Indication           Coded                `json:"indication,omitempty"`
	Vehicle              Coded                `json:"vehicle,omitempty"`
	FulfillmentHistory   []fulfillmentHistory `json:"fulfillmentHistory,omitempty"`
	OrderInformation     []orderInformation   `json:"orderInformation,omitempty"`
}

type administrationTiming struct {
	InstitutionSpecified bool   `json:"institutionSpecified,omitempty"`
	Period               Scalar `json:"period,omitempty"`
}

type doseRestriction struct {
	Numerator   Scalar `json:"numerator,omitempty"`
	Denominator Scalar `json:"denominator,omitempty"`
}

type fulfillmentHistory struct {
	PrescriptionNumber string `json:"prescription_number,omitempty"`
	DispenseDate       int64  `json:"dispense_date,omitempty"`
	QuantityDispensed  Scalar `json,"quantity_dispensed,omitempty"`
	FillNumber         int64  `json:"fill_number,omitempty"`
	FillStatus         string `json:"fill_status"`
}

type OrderInformation struct {
	OrderNumber     string `json:"order_number,omitempty"`
	Fills           int64  `json:"fills,omitempty"`
	QuantityOrdered Scalar `json:"quantity_ordered,omitempty"`
	OrderExpiration int64  `json:"order_expiration_date_time,omitempty"`
	OrderDate       int64  `json:"order_date_time,omitempty"`
}
