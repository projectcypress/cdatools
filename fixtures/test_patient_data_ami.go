package fixtures

var TestPatientDataAmi = []byte(`
{
  "birthdate": -725619600,
  "conditions": [
    {
      "anatomical_location": null,
      "anatomical_target": null,
      "causeOfDeath": null,
      "codes": {
        "ICD-10-CM": [
          "I22.8"
        ],
        "SNOMED-CT": [
          "15990001"
        ],
        "ICD-9-CM": [
          "410.61"
        ]
      },
      "description": "Diagnosis, Active: Acute Myocardial Infarction (AMI)",
      "end_time": 1407917700,
      "laterality": null,
      "mood_code": "EVN",
      "name": null,
      "negationInd": null,
      "negationReason": null,
      "oid": "2.16.840.1.113883.3.560.1.2",
      "ordinality": {
        "code_system": "SNOMED-CT",
        "code": "63161005",
        "title": "Principal"
      },
      "priority": null,
      "reason": null,
      "severity": null,
      "specifics": null,
      "start_time": 1407741600,
      "status_code": {
        "SNOMED-CT": [
          "55561003"
        ],
        "HL7 ActStatus": [
          "active"
        ]
      },
      "time": null,
      "time_of_death": null,
      "type": null,
      "_type": "Condition"
    }
  ],
  "created_at": "2015-08-25T15:30:16.802Z",
  "description": "",
  "encounters": [
    {
      "admitTime": 1407740400,
      "admitType": null,
      "codes": {
        "SNOMED-CT": [
          "4525004"
        ]
      },
      "description": "Encounter, Performed: Emergency Department Visit",
      "diagnosis": null,
      "dischargeDisposition": null,
      "dischargeTime": 1407744900,
      "end_time": 1407744900,
      "mood_code": "EVN",
      "negationInd": null,
      "negationReason": null,
      "oid": "2.16.840.1.113883.3.560.1.79",
      "performer_id": null,
      "principalDiagnosis": null,
      "reason": null,
      "specifics": null,
      "start_time": 1407740400,
      "status_code": {
        "HL7 ActStatus": [
          "performed"
        ]
      },
      "time": null,
      "_type": "Encounter"
    },
    {
      "admitTime": null,
      "admitType": null,
      "codes": {
        "SNOMED-CT": [
          "32485007"
        ]
      },
      "code_list_id": "2.16.840.1.113883.3.666.5.307",
      "description": "Encounter, Performed: Encounter Inpatient",
      "diagnosis": null,
      "dischargeDisposition": null,
      "dischargeTime": null,
      "end_time": 1407917700,
      "mood_code": "EVN",
      "negationInd": null,
      "negationReason": null,
      "oid": "2.16.840.1.113883.3.560.1.79",
      "performer_id": null,
      "principalDiagnosis": {
        "codes": {
          "SNOMED-CT": [
            "8715000"
          ]
        }
      },
      "reason": null,
      "specifics": null,
      "start_time": 1407744900,
      "status_code": {
        "HL7 ActStatus": [
          "performed"
        ]
      },
      "time": null,
      "transferFrom": {
        "code": "434771000124107",
        "code_system": "SNOMED-CT",
        "codes": {
        },
        "time": 1407739500,
        "title": "Transfer From Outpatient"
      },
      "_type": "Encounter"
    }
  ],
  "ethnicity": {
    "code": "2186-5",
    "name": "Not Hispanic or Latino",
    "codeSystem": "CDC Race"
  },
  "expected_values": [
    {
      "measure_id": "7D374C6A-3821-4333-A1BC-4531005D77B8",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "7D374C6A-3821-4333-A1BC-4531005D77B8",
      "population_index": 1,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "E1CB05E0-97D5-40FC-B456-15C5DBF44309",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "EBFA203E-ACC1-4228-906C-855C4BF11310",
      "population_index": 0,
      "IPP": 1,
      "DENOM": 1,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 1
    },
    {
      "measure_id": "0924FBAE-3FDB-4D0A-AAB7-9F354E699FDE",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "3FD13096-2C8F-40B5-9297-B714E8DE9133",
      "population_index": 0,
      "IPP": 1,
      "MSRPOPL": 0,
      "OBSERV_UNIT": " mins"
    },
    {
      "measure_id": "3FD13096-2C8F-40B5-9297-B714E8DE9133",
      "population_index": 1,
      "STRAT": 0,
      "IPP": 0,
      "MSRPOPL": 0,
      "OBSERV_UNIT": " mins"
    },
    {
      "measure_id": "3FD13096-2C8F-40B5-9297-B714E8DE9133",
      "population_index": 2,
      "STRAT": 0,
      "IPP": 0,
      "MSRPOPL": 0,
      "OBSERV_UNIT": " mins"
    },
    {
      "measure_id": "3FD13096-2C8F-40B5-9297-B714E8DE9133",
      "population_index": 3,
      "STRAT": 1,
      "IPP": 1,
      "MSRPOPL": 0,
      "OBSERV_UNIT": " mins"
    },
    {
      "measure_id": "84B9D0B5-0CAF-4E41-B345-3492A23C2E9F",
      "population_index": 0,
      "IPP": 1,
      "DENOM": 1,
      "DENEX": 1,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "9A033274-3D9B-11E1-8634-00237D5BF174",
      "population_index": 0,
      "IPP": 1,
      "MSRPOPL": 1,
      "OBSERV_UNIT": " mins",
      "OBSERV": [
        75
      ]
    },
    {
      "measure_id": "9A033274-3D9B-11E1-8634-00237D5BF174",
      "population_index": 1,
      "STRAT": 1,
      "IPP": 1,
      "MSRPOPL": 1,
      "OBSERV_UNIT": " mins",
      "OBSERV": [
        75
      ]
    },
    {
      "measure_id": "9A033274-3D9B-11E1-8634-00237D5BF174",
      "population_index": 2,
      "STRAT": 0,
      "IPP": 0,
      "MSRPOPL": 0,
      "OBSERV_UNIT": " mins"
    },
    {
      "measure_id": "909CF4B4-7A85-4ABF-A1C7-CB597ED1C0B6",
      "population_index": 0,
      "IPP": 1,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "03876D69-085B-415C-AE9D-9924171040C2",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "93F3479F-75D8-4731-9A3F-B7749D8BCD37",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "6F069BB2-B3C4-4BF4-ADC5-F6DD424A10B7",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "2838875A-07B5-4BF0-BE04-C3EB99F53975",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "BB481284-30DD-4383-928C-82385BBF1B17",
      "population_index": 0,
      "IPP": 1,
      "DENOM": 1,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 1
    },
    {
      "measure_id": "7DC26160-E615-4CC2-879C-75985189EC1A",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "42BF391F-38A3-4C0F-9ECE-DCD47E9609D9",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "1F503318-BB8D-4B91-AF63-223AE0A2328E",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "217FDF0D-3D64-4720-9116-D5E5AFA27F2C",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "38B0B5EC-0F63-466F-8FE3-2CD20DDD1622",
      "population_index": 0,
      "IPP": 1,
      "DENOM": 1,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "BCCE43DD-08E3-46C3-BFDD-0B1B472690F0",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "7FE69617-FA28-4305-A2B8-CEB6BCD9693D",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "979F21BD-3F93-4CDD-8273-B23DFE9C0513",
      "population_index": 0,
      "IPP": 1,
      "MSRPOPL": 0,
      "OBSERV_UNIT": " mins"
    },
    {
      "measure_id": "979F21BD-3F93-4CDD-8273-B23DFE9C0513",
      "population_index": 1,
      "STRAT": 1,
      "IPP": 1,
      "MSRPOPL": 0,
      "OBSERV_UNIT": " mins"
    },
    {
      "measure_id": "979F21BD-3F93-4CDD-8273-B23DFE9C0513",
      "population_index": 2,
      "STRAT": 0,
      "IPP": 0,
      "MSRPOPL": 0,
      "OBSERV_UNIT": " mins"
    },
    {
      "measure_id": "FD7CA18D-B56D-4BCA-AF35-71CE36B15246",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "32CFC834-843A-4F45-B359-8E158EAC4396",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "D09ADD1D-30F5-462D-B677-3D17D9CCD664",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "D09ADD1D-30F5-462D-B677-3D17D9CCD664",
      "population_index": 1,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "D09ADD1D-30F5-462D-B677-3D17D9CCD664",
      "population_index": 2,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "D09ADD1D-30F5-462D-B677-3D17D9CCD664",
      "population_index": 3,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "D09ADD1D-30F5-462D-B677-3D17D9CCD664",
      "population_index": 4,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "D09ADD1D-30F5-462D-B677-3D17D9CCD664",
      "population_index": 5,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "D09ADD1D-30F5-462D-B677-3D17D9CCD664",
      "population_index": 6,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "D09ADD1D-30F5-462D-B677-3D17D9CCD664",
      "population_index": 7,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "FEEA3922-F61F-4B05-98F9-B72A11815F12",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "FEEA3922-F61F-4B05-98F9-B72A11815F12",
      "population_index": 1,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "FEEA3922-F61F-4B05-98F9-B72A11815F12",
      "population_index": 2,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "FEEA3922-F61F-4B05-98F9-B72A11815F12",
      "population_index": 3,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "FEEA3922-F61F-4B05-98F9-B72A11815F12",
      "population_index": 4,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "FEEA3922-F61F-4B05-98F9-B72A11815F12",
      "population_index": 5,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "FEEA3922-F61F-4B05-98F9-B72A11815F12",
      "population_index": 6,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "FEEA3922-F61F-4B05-98F9-B72A11815F12",
      "population_index": 7,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "D78CE034-8288-4012-A31E-7F485A74F2A9",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "FF796FD9-F99D-41FD-B8C2-57D0A59A5D8D",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "8243EAE0-BBD7-4107-920B-FC3DB04B9584",
      "population_index": 0,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "8243EAE0-BBD7-4107-920B-FC3DB04B9584",
      "population_index": 1,
      "IPP": 0,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    },
    {
      "measure_id": "FA91BA68-1E66-4A23-8EB2-BAA8E6DF2F2F",
      "population_index": 0,
      "IPP": 1,
      "DENOM": 0,
      "DENEX": 0,
      "NUMER": 0,
      "DENEXCEP": 0
    }
  ],
  "expired": false,
  "first": "1 N",
  "gender": "F",
  "insurance_providers": [
    {
      "codes": {
        "SOP": [
          "349"
        ]
      },
      "description": null,
      "end_time": null,
      "financial_responsibility_type": {
        "code": "SELF",
        "codeSystem": "HL7 Relationship Code"
      },
      "member_id": "1234567890",
      "mood_code": "EVN",
      "name": "Other",
      "negationInd": null,
      "negationReason": null,
      "oid": null,
      "payer": {
        "name": "Other"
      },
      "reason": null,
      "relationship": null,
      "specifics": null,
      "start_time": 1199163600,
      "status_code": null,
      "time": null,
      "type": "OT",
      "_type": "InsuranceProvider"
    }
  ],
  "is_shared": false,
  "languages": [

  ],
  "last": "N AMI",
  "measure_ids": [
    "979F21BD-3F93-4CDD-8273-B23DFE9C0513",
    "42BF391F-38A3-4C0F-9ECE-DCD47E9609D9",
    "7FE69617-FA28-4305-A2B8-CEB6BCD9693D",
    "93F3479F-75D8-4731-9A3F-B7749D8BCD37",
    "E1CB05E0-97D5-40FC-B456-15C5DBF44309",
    "6F069BB2-B3C4-4BF4-ADC5-F6DD424A10B7",
    "03876D69-085B-415C-AE9D-9924171040C2",
    "1F503318-BB8D-4B91-AF63-223AE0A2328E",
    "7DC26160-E615-4CC2-879C-75985189EC1A",
    "FD7CA18D-B56D-4BCA-AF35-71CE36B15246",
    "BCCE43DD-08E3-46C3-BFDD-0B1B472690F0",
    "0924FBAE-3FDB-4D0A-AAB7-9F354E699FDE",
    "38B0B5EC-0F63-466F-8FE3-2CD20DDD1622",
    "FA91BA68-1E66-4A23-8EB2-BAA8E6DF2F2F",
    "9A033274-3D9B-11E1-8634-00237D5BF174",
    "FF796FD9-F99D-41FD-B8C2-57D0A59A5D8D",
    "EBFA203E-ACC1-4228-906C-855C4BF11310",
    "2838875A-07B5-4BF0-BE04-C3EB99F53975",
    "3FD13096-2C8F-40B5-9297-B714E8DE9133",
    "FEEA3922-F61F-4B05-98F9-B72A11815F12",
    "D09ADD1D-30F5-462D-B677-3D17D9CCD664",
    "217FDF0D-3D64-4720-9116-D5E5AFA27F2C",
    "32CFC834-843A-4F45-B359-8E158EAC4396",
    "BB481284-30DD-4383-928C-82385BBF1B17",
    "909CF4B4-7A85-4ABF-A1C7-CB597ED1C0B6",
    "8243EAE0-BBD7-4107-920B-FC3DB04B9584",
    "84B9D0B5-0CAF-4E41-B345-3492A23C2E9F",
    "D78CE034-8288-4012-A31E-7F485A74F2A9",
    "7D374C6A-3821-4333-A1BC-4531005D77B8",
    null
  ],
  "medical_record_assigner": "2.16.840.1.113883.3.1257",
  "medical_record_number": "93c4d6904738d3e38fb781761f09a439",
  "medications": [
    {
      "active_datetime": null,
      "administrationTiming": {
        "period": {
          "value": "1",
          "unit": "d"
        }
      },
      "anatomical_approach": null,
      "codes": {
        "RxNorm": [
          "1191"
        ]
      },
      "cumulativeMedicationDuration": null,
      "deliveryMethod": null,
      "description": "Medication, Order: Aspirin ingredient specific",
      "dose": {
        "value": "81",
        "unit": "mg"
      },
      "doseIndicator": null,
      "doseRestriction": null,
      "end_time": 1407745200,
      "freeTextSig": null,
      "fulfillmentHistory": [

      ],
      "fulfillmentInstructions": null,
      "indication": null,
      "method": null,
      "mood_code": "EVN",
      "negationInd": true,
      "negationReason": {
        "code_system": "SNOMED-CT",
        "code": "182901005"
      },
      "oid": "2.16.840.1.113883.3.560.1.78",
      "patientInstructions": null,
      "productForm": null,
      "reaction": null,
      "reason": null,
      "route": null,
      "signed_datetime": null,
      "specifics": null,
      "start_time": 1407745200,
      "statusOfMedication": null,
      "status_code": {
        "HL7 ActStatus": [
          "ordered"
        ]
      },
      "time": null,
      "typeOfMedication": null,
      "vehicle": null,
      "_type": "Medication"
    },
    {
      "active_datetime": null,
      "administrationTiming": {
        "period": {
          "value": "12",
          "unit": "h"
        }
      },
      "anatomical_approach": null,
      "codes": {
        "RxNorm": [
          "41127"
        ]
      },
      "cumulativeMedicationDuration": null,
      "deliveryMethod": null,
      "description": "Medication, Order: Statin ingredient specific",
      "dose": {
        "value": "40",
        "unit": "mg"
      },
      "doseIndicator": null,
      "doseRestriction": null,
      "end_time": 1407745200,
      "freeTextSig": null,
      "fulfillmentHistory": [

      ],
      "fulfillmentInstructions": null,
      "indication": null,
      "method": null,
      "mood_code": "EVN",
      "negationInd": true,
      "negationReason": {
        "code_system": "SNOMED-CT",
        "code": "182901005"
      },
      "oid": "2.16.840.1.113883.3.560.1.78",
      "patientInstructions": null,
      "productForm": null,
      "reaction": null,
      "reason": null,
      "route": null,
      "signed_datetime": null,
      "specifics": null,
      "start_time": 1407745200,
      "statusOfMedication": null,
      "status_code": {
        "HL7 ActStatus": [
          "ordered"
        ]
      },
      "time": null,
      "typeOfMedication": null,
      "vehicle": null,
      "_type": "Medication"
    },
    {
      "active_datetime": null,
      "administrationTiming": {
        "period": {
          "value": "1",
          "unit": "d"
        }
      },
      "anatomical_approach": null,
      "codes": {
        "RxNorm": [
          "1191"
        ]
      },
      "cumulativeMedicationDuration": null,
      "deliveryMethod": null,
      "description": "Medication, Discharge: Aspirin ingredient specific",
      "dose": {
        "value": "81",
        "unit": "mg"
      },
      "doseIndicator": null,
      "doseRestriction": null,
      "end_time": 1407916800,
      "freeTextSig": null,
      "fulfillmentHistory": [

      ],
      "fulfillmentInstructions": null,
      "indication": null,
      "method": null,
      "mood_code": "EVN",
      "negationInd": true,
      "negationReason": {
        "code_system": "SNOMED-CT",
        "code": "182901005"
      },
      "oid": "2.16.840.1.113883.3.560.1.200",
      "patientInstructions": null,
      "productForm": null,
      "reaction": null,
      "reason": null,
      "route": null,
      "signed_datetime": null,
      "specifics": null,
      "start_time": 1407916800,
      "statusOfMedication": null,
      "status_code": {
        "HL7 ActStatus": [
          "discharge"
        ]
      },
      "time": null,
      "typeOfMedication": null,
      "vehicle": null,
      "_type": "Medication"
    },
    {
      "active_datetime": null,
      "administrationTiming": {
        "period": {
          "value": "12",
          "unit": "h"
        }
      },
      "anatomical_approach": null,
      "codes": {
        "RxNorm": [
          "41127"
        ]
      },
      "cumulativeMedicationDuration": null,
      "deliveryMethod": null,
      "description": "Medication, Discharge: Statin ingredient specific",
      "dose": {
        "value": "40",
        "unit": "mg"
      },
      "doseIndicator": null,
      "doseRestriction": null,
      "end_time": 1407916800,
      "freeTextSig": null,
      "fulfillmentHistory": [

      ],
      "fulfillmentInstructions": null,
      "indication": null,
      "method": null,
      "mood_code": "EVN",
      "negationInd": true,
      "negationReason": {
        "code_system": "SNOMED-CT",
        "code": "182901005"
      },
      "oid": "2.16.840.1.113883.3.560.1.200",
      "patientInstructions": null,
      "productForm": null,
      "reaction": null,
      "reason": null,
      "route": null,
      "signed_datetime": null,
      "specifics": null,
      "start_time": 1407916800,
      "statusOfMedication": null,
      "status_code": {
        "HL7 ActStatus": [
          "discharge"
        ]
      },
      "time": null,
      "typeOfMedication": null,
      "vehicle": null,
      "_type": "Medication"
    },
    {
      "active_datetime": null,
      "administrationTiming": null,
      "anatomical_approach": null,
      "codes": {
        "RxNorm": [
          "1439897"
        ]
      },
      "cumulativeMedicationDuration": null,
      "deliveryMethod": null,
      "description": "Medication, Administered: Statin",
      "dose": {
        "value": "1",
        "unit": "tablet(s)"
      },
      "doseIndicator": null,
      "doseRestriction": null,
      "end_time": 1407916800,
      "freeTextSig": null,
      "fulfillmentHistory": [

      ],
      "fulfillmentInstructions": null,
      "indication": null,
      "method": null,
      "mood_code": "EVN",
      "negationInd": null,
      "negationReason": null,
      "oid": "2.16.840.1.113883.3.560.1.14",
      "patientInstructions": null,
      "productForm": null,
      "reaction": null,
      "reason": null,
      "route": null,
      "signed_datetime": null,
      "specifics": null,
      "start_time": 1407916800,
      "statusOfMedication": null,
      "status_code": {
        "HL7 ActStatus": [
          "administered"
        ]
      },
      "time": null,
      "typeOfMedication": null,
      "vehicle": null,
      "_type": "Medication"
    }
  ],
  "notes": "",
  "origin_data": [

  ],
  "procedures": [
    {
      "anatomical_approach": null,
      "anatomical_target": null,
      "codes": {
        "SNOMED-CT": [
          "164847006"
        ]
      },
      "description": "Diagnostic Study, Performed: Electrocardiogram (ECG)",
      "end_time": 1407741300,
      "incisionTime": null,
      "method": null,
      "mood_code": "EVN",
      "negationInd": null,
      "negationReason": null,
      "oid": "2.16.840.1.113883.3.560.1.3",
      "ordinality": null,
      "performer_id": null,
      "radiation_dose": null,
      "radiation_duration": null,
      "reaction": null,
      "reason": null,
      "source": null,
      "specifics": null,
      "start_time": 1407741300,
      "status_code": {
        "HL7 ActStatus": [
          "performed"
        ]
      },
      "time": null,
      "values": [
        {
          "codes": {
            "ICD-10-CM": [
              "I22.8"
            ],
            "SNOMED-CT": [
              "15990001"
            ],
            "ICD-9-CM": [
              "410.61"
            ]
          },
          "description": "Acute or Evolving MI",
          "_type": "CodedResultValue"
        }
      ],
      "_type": "Procedure"
    },
    {
      "anatomical_approach": null,
      "anatomical_target": null,
      "codes": {
        "ICD-10-PCS": [
          "0270346"
        ],
        "SNOMED-CT": [
          "68466008"
        ],
        "ICD-9-PCS": [
          "00.66"
        ]
      },
      "description": "Procedure, Performed: PCI",
      "end_time": 1407747600,
      "incisionTime": null,
      "method": null,
      "mood_code": "EVN",
      "negationInd": null,
      "negationReason": null,
      "oid": "2.16.840.1.113883.3.560.1.6",
      "ordinality": {
        "code_system": "SNOMED-CT",
        "code": "399455000",
        "title": "Primary Procedure"
      },
      "performer_id": null,
      "radiation_dose": null,
      "radiation_duration": null,
      "reaction": null,
      "reason": null,
      "source": null,
      "specifics": null,
      "start_time": 1407745800,
      "status_code": {
        "HL7 ActStatus": [
          "performed"
        ]
      },
      "time": null,
      "_type": "Procedure"
    }
  ],
  "race": {
    "code": "2028-9",
    "name": "Asian",
    "codeSystem": "CDC Race"
  },
  "source_data_criteria": [
    {
      "negation": false,
      "definition": "encounter",
      "status": "performed",
      "title": "Emergency Department Visit",
      "description": "Encounter, Performed: Emergency Department Visit",
      "code_list_id": "2.16.840.1.113883.3.117.1.7.1.292",
      "type": "encounters",
      "id": "EncounterPerformedEmergencyDepartmentVisit",
      "start_date": 1407740400000,
      "end_date": 1407744900000,
      "value": [

      ],
      "references": null,
      "field_values": {
        "ADMISSION_DATETIME": {
          "type": "TS",
          "value": 1407740400000
        },
        "DISCHARGE_DATETIME": {
          "type": "TS",
          "value": 1407744900000
        }
      },
      "hqmf_set_id": "9A033274-3D9B-11E1-8634-00237D5BF174",
      "cms_id": "CMS55v4",
      "criteria_id": "151e9f44bf7vJ",
      "codes": {
        "SNOMED-CT": [
          "4525004"
        ]
      },
      "negation_code_list_id": "",
      "coded_entry_id": "56bbad6202d4055317002128",
      "code_source": "DEFAULT"
    },
    {
      "negation": false,
      "definition": "diagnostic_study",
      "status": "performed",
      "title": "Electrocardiogram (ECG)",
      "description": "Diagnostic Study, Performed: Electrocardiogram (ECG)",
      "code_list_id": "2.16.840.1.113883.3.666.5.735",
      "type": "diagnostic_studies",
      "id": "DiagnosticStudyPerformedElectrocardiogramEcg",
      "start_date": 1407741300000,
      "end_date": 1407741300000,
      "value": [
        {
          "type": "CD",
          "code_list_id": "2.16.840.1.113883.3.666.5.3022",
          "title": "Acute or Evolving MI"
        }
      ],
      "references": null,
      "field_values": {
      },
      "hqmf_set_id": "84B9D0B5-0CAF-4E41-B345-3492A23C2E9F",
      "cms_id": "CMS53v4",
      "criteria_id": "14f6aaceb91jw",
      "codes": {
        "SNOMED-CT": [
          "164847006"
        ]
      },
      "negation_code_list_id": "",
      "coded_entry_id": "56bbad6202d405531700212f",
      "code_source": "DEFAULT"
    },
    {
      "negation": false,
      "definition": "diagnosis",
      "status": "active",
      "title": "Acute Myocardial Infarction (AMI)",
      "description": "Diagnosis, Active: Acute Myocardial Infarction (AMI)",
      "code_list_id": "2.16.840.1.113883.3.666.5.3011",
      "type": "conditions",
      "id": "DiagnosisActiveAcuteMyocardialInfarctionAmi",
      "start_date": 1407741600000,
      "end_date": 1407917700000,
      "value": [

      ],
      "references": null,
      "field_values": {
        "ORDINAL": {
          "type": "CD",
          "code_list_id": "2.16.840.1.113883.3.117.1.7.1.14",
          "title": "Principal"
        }
      },
      "hqmf_set_id": "EBFA203E-ACC1-4228-906C-855C4BF11310",
      "cms_id": "CMS30v5",
      "criteria_id": "14f1da58fe828",
      "codes": {
        "ICD-10-CM": [
          "I22.8"
        ],
        "SNOMED-CT": [
          "15990001"
        ],
        "ICD-9-CM": [
          "410.61"
        ]
      },
      "negation_code_list_id": "",
      "coded_entry_id": "56bbad6202d4055317002127",
      "code_source": "DEFAULT"
    },
    {
      "negation": false,
      "definition": "encounter",
      "status": "performed",
      "title": "Encounter Inpatient",
      "description": "Encounter, Performed: Encounter Inpatient",
      "code_list_id": "2.16.840.1.113883.3.666.5.307",
      "type": "encounters",
      "id": "EncounterPerformedEncounterInpatient",
      "start_date": 1407744900000,
      "end_date": 1407917700000,
      "value": [

      ],
      "references": null,
      "field_values": {
        "TRANSFER_FROM": {
          "type": "CD",
          "code_list_id": "2.16.840.1.113883.3.67.1.101.950",
          "title": "Transfer From Outpatient"
        },
        "TRANSFER_FROM_DATETIME": {
          "type": "TS",
          "value": 1407739500000
        }
      },
      "hqmf_set_id": "84B9D0B5-0CAF-4E41-B345-3492A23C2E9F",
      "cms_id": "CMS53v4",
      "criteria_id": "14f6aafdfcbMD",
      "codes": {
        "SNOMED-CT": [
          "32485007"
        ]
      },
      "negation_code_list_id": "",
      "coded_entry_id": "56bbad6202d4055317002129",
      "code_source": "USER_DEFINED"
    },
    {
      "negation": true,
      "definition": "medication",
      "status": "ordered",
      "title": "Aspirin ingredient specific",
      "description": "Medication, Order: Aspirin ingredient specific",
      "code_list_id": "2.16.840.1.113762.1.4.1021.3",
      "type": "medications",
      "id": "MedicationOrderAspirinIngredientSpecific",
      "start_date": 1407745200000,
      "end_date": 1407745200000,
      "value": [

      ],
      "references": null,
      "field_values": {
      },
      "hqmf_set_id": "BB481284-30DD-4383-928C-82385BBF1B17",
      "cms_id": "CMS100v4",
      "criteria_id": "14f65bbe3d9bH",
      "codes": {
        "RxNorm": [
          "1191"
        ]
      },
      "fulfillments": null,
      "dose_value": "81",
      "dose_unit": "mg",
      "frequency_value": "1",
      "frequency_unit": "d",
      "negation_code_list_id": "2.16.840.1.113883.3.117.1.7.1.93",
      "coded_entry_id": "56bbad6202d405531700212a",
      "code_source": "DEFAULT"
    },
    {
      "negation": true,
      "definition": "medication",
      "status": "ordered",
      "title": "Statin ingredient specific",
      "description": "Medication, Order: Statin ingredient specific",
      "code_list_id": "2.16.840.1.113762.1.4.1021.7",
      "type": "medications",
      "id": "MedicationOrderStatinIngredientSpecific",
      "start_date": 1407745200000,
      "end_date": 1407745200000,
      "value": [

      ],
      "references": null,
      "field_values": {
      },
      "hqmf_set_id": "EBFA203E-ACC1-4228-906C-855C4BF11310",
      "cms_id": "CMS30v5",
      "criteria_id": "14f660e68devj",
      "codes": {
        "RxNorm": [
          "41127"
        ]
      },
      "fulfillments": null,
      "dose_value": "40",
      "dose_unit": "mg",
      "frequency_value": "12",
      "frequency_unit": "h",
      "negation_code_list_id": "2.16.840.1.113883.3.117.1.7.1.93",
      "coded_entry_id": "56bbad6202d405531700212b",
      "code_source": "DEFAULT"
    },
    {
      "negation": false,
      "definition": "procedure",
      "status": "performed",
      "title": "PCI",
      "description": "Procedure, Performed: PCI",
      "code_list_id": "2.16.840.1.113762.1.4.1045.67",
      "type": "procedures",
      "id": "ProcedurePerformedPci",
      "start_date": 1407745800000,
      "end_date": 1407747600000,
      "value": [

      ],
      "references": null,
      "field_values": {
        "ORDINAL": {
          "type": "CD",
          "code_list_id": "2.16.840.1.113762.1.4.1111.12",
          "title": "Primary Procedure"
        }
      },
      "hqmf_set_id": "84B9D0B5-0CAF-4E41-B345-3492A23C2E9F",
      "cms_id": "CMS53v4",
      "criteria_id": "14f6aadb0c0HA",
      "codes": {
        "ICD-10-PCS": [
          "0270346"
        ],
        "SNOMED-CT": [
          "68466008"
        ],
        "ICD-9-PCS": [
          "00.66"
        ]
      },
      "negation_code_list_id": "",
      "coded_entry_id": "56bbad6202d4055317002131",
      "code_source": "DEFAULT"
    },
    {
      "negation": true,
      "definition": "medication",
      "status": "discharge",
      "title": "Aspirin ingredient specific",
      "description": "Medication, Discharge: Aspirin ingredient specific",
      "code_list_id": "2.16.840.1.113762.1.4.1021.3",
      "type": "medications",
      "id": "MedicationDischargeAspirinIngredientSpecific",
      "start_date": 1407916800000,
      "end_date": 1407916800000,
      "value": [

      ],
      "references": null,
      "field_values": {
      },
      "hqmf_set_id": "BB481284-30DD-4383-928C-82385BBF1B17",
      "cms_id": "CMS100v4",
      "criteria_id": "14f65bcaac7pj",
      "codes": {
        "RxNorm": [
          "1191"
        ]
      },
      "fulfillments": null,
      "dose_value": "81",
      "dose_unit": "mg",
      "frequency_value": "1",
      "frequency_unit": "d",
      "negation_code_list_id": "2.16.840.1.113883.3.117.1.7.1.93",
      "coded_entry_id": "56bbad6202d405531700212c",
      "code_source": "DEFAULT"
    },
    {
      "negation": true,
      "definition": "medication",
      "status": "discharge",
      "title": "Statin ingredient specific",
      "description": "Medication, Discharge: Statin ingredient specific",
      "code_list_id": "2.16.840.1.113762.1.4.1021.7",
      "type": "medications",
      "id": "MedicationDischargeStatinIngredientSpecific",
      "start_date": 1407916800000,
      "end_date": 1407916800000,
      "value": [

      ],
      "references": null,
      "field_values": {
      },
      "hqmf_set_id": "EBFA203E-ACC1-4228-906C-855C4BF11310",
      "cms_id": "CMS30v5",
      "criteria_id": "14f660ecb24LE",
      "codes": {
        "RxNorm": [
          "41127"
        ]
      },
      "fulfillments": null,
      "dose_value": "40",
      "dose_unit": "mg",
      "frequency_value": "12",
      "frequency_unit": "h",
      "negation_code_list_id": "2.16.840.1.113883.3.117.1.7.1.93",
      "coded_entry_id": "56bbad6202d405531700212d",
      "code_source": "DEFAULT"
    },
    {
      "negation": false,
      "definition": "medication",
      "status": "administered",
      "title": "Statin",
      "description": "Medication, Administered: Statin",
      "code_list_id": "2.16.840.1.113883.3.117.1.7.1.225",
      "type": "medications",
      "id": "MedicationAdministeredStatin",
      "start_date": 1407916800000,
      "end_date": 1407916800000,
      "value": [

      ],
      "references": null,
      "field_values": {
      },
      "hqmf_set_id": "EBFA203E-ACC1-4228-906C-855C4BF11310",
      "cms_id": "CMS30v5",
      "criteria_id": "15247051129wR",
      "codes": {
        "RxNorm": [
          "1439897"
        ]
      },
      "fulfillments": null,
      "dose_value": "1",
      "dose_unit": "tablet(s)",
      "frequency_value": "",
      "frequency_unit": "h",
      "negation_code_list_id": "",
      "coded_entry_id": "56bbad6202d405531700212e",
      "code_source": "DEFAULT"
    },
    {
      "id": "MeasurePeriod",
      "start_date": 0,
      "end_date": 0,
      "references": null
    }
  ],
  "updated_at": "2016-02-10T21:36:34.324Z"
}
`)
