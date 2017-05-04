package models

import (
	"strconv"
)

type Medication struct {
	Entry                  `bson:",inline"`
	AdministrationTiming   AdministrationTiming `json:"administrationTiming,omitempty"`
	AllowedAdministrations *int64               `json:"allowed_administrations,omitempty"`
	Route                  *CodedConcept        `json:"route,omitempty"`
	Dose                   Scalar               `json:"dose,omitempty"`
	AnatomicalApproach     *CodedConcept        `json:"anatomical_approach,omitempty"`
	DoseRestriction        doseRestriction      `json:"dose_restriction,omitempty"`
	ProductForm            *CodedConcept        `json:"productForm,omitempty"`
	DeliveryMethod         *CodedConcept        `json:"delivery_method,omitempty"`
	TypeOfMedication       *CodedConcept        `json:"type_of_medication,omitempty"`
	Indication             *CodedConcept        `json:"indication,omitempty"`
	Vehicle                *CodedConcept        `json:"vehicle,omitempty"`
	FulfillmentHistory     []FulfillmentHistory `json:"fulfillmentHistory,omitempty"`
	OrderInformation       []OrderInformation   `json:"orderInformation,omitempty"`
}

type AdministrationTiming struct {
	InstitutionSpecified bool   `json:"institutionSpecified,omitempty"`
	Period               Scalar `json:"period,omitempty"`
}

type doseRestriction struct {
	Numerator   Scalar `json:"numerator,omitempty"`
	Denominator Scalar `json:"denominator,omitempty"`
}

type FulfillmentHistory struct {
	PrescriptionNumber string `json:"prescription_number,omitempty"`
	DispenseDate       *int64 `json:"dispense_date,omitempty"`
	QuantityDispensed  Scalar `json:"quantityDispensed,omitempty"`
	FillNumber         int64  `json:"fill_number,omitempty"`
	FillStatus         string `json:"fill_status"`
}

type OrderInformation struct {
	OrderNumber     string `json:"order_number,omitempty"`
	Fills           int64  `json:"fills,omitempty"`
	QuantityOrdered Scalar `json:"quantity_ordered,omitempty"`
	OrderExpiration int64  `json:"order_expiration_date_time,omitempty"`
	OrderDate       *int64 `json:"order_date_time,omitempty"`
}

func (med *Medication) GetEntry() *Entry {
	return &med.Entry
}

func (med *Medication) HasSetAdministrationTiming() bool {
	return med.AdministrationTiming != AdministrationTiming{}
}

func (med *Medication) HasSetDose() bool {
	return med.Dose != Scalar{}
}

func (med *Medication) DoseQuantity() string {
	if med.Codes["RxNorm"] != nil || med.Codes["CVX"] != nil {
		if med.Dose.Unit != "" {
			return "value=\"1\" unit=\"" + ucumForDoseQuantity(med.Dose.Unit) + "\""
		}
		return "value=\"1\""
	}
	return "value=\"" + med.Dose.Scalar + "\" unit=\"" + med.Dose.Units + "\""
}

func ucumForDoseQuantity(dose string) string {
	switch dose {
	case "capsule(s)":
		return "{Capsule}"
	case "tablet(s)":
		return "{tbl}"
	default:
		return dose
	}
}

func (med *Medication) FulfillmentQuantity(fulfill FulfillmentHistory) string {
	if med.Codes["RxNorm"] != nil {
		var doses int64
		quantityDispensedVal, err1 := strconv.ParseFloat(fulfill.QuantityDispensed.Value, 64)
		doseVal, err2 := strconv.ParseFloat(med.Dose.Value, 64)
		if (err1 == nil && err2 == nil && fulfill.QuantityDispensed != Scalar{} && doseVal != 0) {
			doses = int64(quantityDispensedVal / doseVal)
		} else {
			doses = int64(0)
		}
		return "value=\"" + strconv.FormatInt(doses, 10) + "\" unit=\"1\""
	}
	return "value=\"" + fulfill.QuantityDispensed.Value + "\" unit=\"" + fulfill.QuantityDispensed.Unit + "\""
}
