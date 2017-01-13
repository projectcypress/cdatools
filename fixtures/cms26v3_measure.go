package fixtures

var Cms26v3 = []byte(`
{
  "id": "40280381-4B9A-3825-014B-BD8FA6B2062E",
  "nqf_id": "0338",
  "hqmf_id": "40280381-4B9A-3825-014B-BD8FA6B2062E",
  "hqmf_set_id": "E1CB05E0-97D5-40FC-B456-15C5DBF44309",
  "hqmf_version_number": 3,
  "cms_id": "CMS26v3",
  "name": "Home Management Plan of Care (HMPC) Document Given to Patient/Caregiver",
  "description": "An assessment that there is documentation in the medical record that a Home Management Plan of Care (HMPC) document was given to the pediatric asthma patient/caregiver.",
  "type": "eh",
  "category": "Asthma",
  "map_fn": "function() {\n          var patient = this;\n          var effective_date = <%= effective_date %>;\n          var enable_logging = <%= enable_logging %>;\n          var enable_rationale = <%= enable_rationale %>;\n          var short_circuit = <%= short_circuit %>;\n\n        <% if (!test_id.nil? && test_id.class==Moped::BSON::ObjectId) %>\n          var test_id = new ObjectId(\"<%= test_id %>\");\n        <% else %>\n          var test_id = null;\n        <% end %>\n\n          hqmfjs = {}\n          <%= init_js_frameworks %>\n\n          hqmfjs.effective_date = effective_date;\n          hqmfjs.test_id = test_id;\n      \n          \n        var patient_api = new hQuery.Patient(patient);\n\n        \n        // #########################\n        // ##### DATA ELEMENTS #####\n        // #########################\n\n        hqmfjs.nqf_id = '0338';\n        hqmfjs.hqmf_id = '40280381-4B9A-3825-014B-BD8FA6B2062E';\n        hqmfjs.sub_id = null;\n        if (typeof(test_id) == 'undefined') hqmfjs.test_id = null;\n\n        OidDictionary = {'2.16.840.1.113883.3.117.1.7.1.271':{'ICD-10-CM':['J45.41','J45.42','J45.990','J45.991','J45.20','J45.21','J45.31','J45.901','J45.32','J45.40','J45.51','J45.902','J45.52','J45.30','J45.50','J45.998','J45.909','J45.22'],'SNOMED-CT':['233683003','389145006','405944004','409663006','55570000','195949008','233678006','426979002','427679007','195977004','233679003','281239006','30352005','12428000','424643009','426656000','445427006','63088003','370219009','370218001','370220003','423889005','425969006','225057002','266361008','304527002','427354000','427603009','195967001','31387002','370221004','427295004'],'ICD-9-CM':['493.01','493.02','493.11','493.90','493.10','493.12','493.81','493.82','493.00','493.91','493.92']},'2.16.840.1.113883.3.117.1.7.1.131':{'LOINC':['69981-9']},'2.16.840.1.113883.3.117.1.7.1.82':{'SNOMED-CT':['306691003','306692005','306705005','306690002','65537008','86400002','10161009','306689006']},'2.16.840.1.113883.3.117.1.7.1.93':{'SNOMED-CT':['182901005','182903008','183944003','183945002','413312003','182890002','182897004','406149000','105480006','182895007','182896008','182900006','443390004','182898009','183947005','183948000','371138003','416432009','275936005','385648002']},'2.16.840.1.113762.1.4.1':{'AdministrativeGender':['M','F']},'2.16.840.1.114222.4.11.836':{'CDC Race':['2076-8','1002-5','2131-1','2106-3','2028-9','2054-5']},'2.16.840.1.114222.4.11.837':{'CDC Race':['2135-2','2186-5']},'2.16.840.1.114222.4.11.3591':{'Source of Payment Typology':['521','84','6','331','3119','953','3222','512','349','37','41','523','3116','312','3113','5','32126','212','3115','3211','54','112','611','311','333','21','122','39','822','332','32122','82','73','322','32125','3711','121','389','3','511','342','36','3712','59','3221','379','62','43','3223','123','119','3212','32121','52','81','55','34','69','8','821','98','3112','519','3114','79','3811','32123','25','38','613','35','2','94','85','99','91','3123','321','3229','3813','83','3713','24','951','213','522','129','61','1','3122','64','612','334','529','22','33','31','72','219','619','3812','92','211','29','4','51','313','63','341','9','514','3819','513','515','95','44','42','53','823','96','113','93','343','3121','362','89','959','11','32','32124','381','111','23','19','954','12','372','9999','71','361','7','382','369','3111','371']},'2.16.840.1.113883.3.666.5.307':{'SNOMED-CT':['8715000','183452005','32485007']},'2.16.840.1.113883.3.117.1.7.1.70':{'SNOMED-CT':['3950001']},'2.16.840.1.113883.3.117.1.7.1.14':{'SNOMED-CT':['63161005']}};\n        \n        // Measure variables\nvar MeasurePeriod = {\n  \"low\": new TS(\"201201010000\", true),\n  \"high\": new TS(\"201212312359\", true)\n}\nhqmfjs.MeasurePeriod = function(patient) {\n  return [new hQuery.CodedEntry(\n    {\n      \"start_time\": MeasurePeriod.low.asDate().getTime()/1000,\n      \"end_time\": MeasurePeriod.high.asDate().getTime()/1000,\n      \"codes\": {}\n    }\n  )];\n}\nif (typeof effective_date === 'number') {\n  MeasurePeriod.high.date = new Date(1000*effective_date);\n  // add one minute before pulling off the year.  This turns 12-31-2012 23:59 into 1-1-2013 00:00 => 1-1-2012 00:00\n  MeasurePeriod.low.date = new Date(1000*(effective_date+60));\n  MeasurePeriod.low.date.setFullYear(MeasurePeriod.low.date.getFullYear()-1);\n}\n\n// Data critera\nhqmfjs.GROUP_variable_CHILDREN_12 = function(patient, initialSpecificContext) {\n  var events = UNION(\n    hqmfjs.GROUP_satisfiesAll_CHILDREN_10(patient, initialSpecificContext)\n  );\n  // record the result of the source of the variable to the rationale\n  if(Logger.enable_rationale) Logger.record('GROUP_variable_CHILDREN_12',events);\n  events.specific_occurrence = 'GROUP_variable_CHILDREN_12';\n\n  events.specificContext=new hqmf.SpecificOccurrence(Row.buildForDataCriteria(events.specific_occurrence, events))\n  return events;\n}\n\nhqmfjs.PatientCharacteristicSexOncAdministrativeSex = function(patient, initialSpecificContext) {\n  var value = patient.gender() || null;\n  matching = matchingValue(value, new CD(\"M\", \"Administrative Sex\"));\n  matching.specificContext=hqmf.SpecificsManager.identity();\n  return matching;\n}\n\nhqmfjs.PatientCharacteristicRaceRace = function(patient, initialSpecificContext) {\n  var value = patient.race() || null;\n  if (value === null) {\n    matching = new Boolean(false);\n  } else {\n    matching = new Boolean(value.includedIn({\"CDC Race\":[\"2076-8\",\"1002-5\",\"2131-1\",\"2028-9\",\"2054-5\",\"2106-3\"]}));\n  }\n  matching.specificContext=hqmf.SpecificsManager.identity();\n  return matching;\n}\n\nhqmfjs.PatientCharacteristicEthnicityEthnicity = function(patient, initialSpecificContext) {\n  var value = patient.ethnicity() || null;\n  matching = matchingValue(value, null);\n  matching.specificContext=hqmf.SpecificsManager.identity();\n  return matching;\n}\n\nhqmfjs.PatientCharacteristicPayerPayer = function(patient, initialSpecificContext) {\n  var value = patient.payer() || null;\n  if (value === null) {\n    matching = new Boolean(false);\n  } else {\n    matching = new Boolean(value.includedIn({\"Source of Payment Typology\":[\"523\",\"41\",\"512\",\"953\",\"37\",\"212\",\"331\",\"6\",\"84\",\"521\",\"3115\",\"3119\",\"3222\",\"5\",\"312\",\"3116\",\"3113\",\"349\",\"32126\",\"121\",\"39\",\"333\",\"311\",\"3\",\"611\",\"389\",\"3711\",\"21\",\"32122\",\"32125\",\"122\",\"322\",\"73\",\"112\",\"54\",\"332\",\"82\",\"822\",\"3211\",\"3712\",\"62\",\"379\",\"119\",\"3221\",\"511\",\"43\",\"36\",\"123\",\"342\",\"59\",\"3223\",\"98\",\"8\",\"69\",\"3811\",\"35\",\"34\",\"55\",\"52\",\"79\",\"32121\",\"613\",\"519\",\"81\",\"38\",\"3114\",\"821\",\"25\",\"3212\",\"3112\",\"32123\",\"22\",\"213\",\"3229\",\"64\",\"3122\",\"321\",\"3123\",\"24\",\"91\",\"3713\",\"951\",\"529\",\"99\",\"85\",\"94\",\"1\",\"2\",\"83\",\"129\",\"3813\",\"522\",\"61\",\"334\",\"612\",\"29\",\"51\",\"211\",\"514\",\"9\",\"341\",\"3812\",\"31\",\"619\",\"33\",\"4\",\"63\",\"219\",\"92\",\"72\",\"313\",\"113\",\"89\",\"96\",\"362\",\"959\",\"823\",\"3121\",\"95\",\"53\",\"343\",\"515\",\"3819\",\"42\",\"44\",\"93\",\"513\",\"381\",\"371\",\"361\",\"3111\",\"71\",\"32124\",\"23\",\"111\",\"382\",\"9999\",\"369\",\"12\",\"372\",\"954\",\"19\",\"32\",\"7\",\"11\"]}));\n  }\n  matching.specificContext=hqmf.SpecificsManager.identity();\n  return matching;\n}\n\nhqmfjs.EncounterPerformedEncounterInpatient_precondition_2 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"encounters\", \"statuses\": [\"performed\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.666.5.307\"};\n  var events = patient.getEvents(eventCriteria);\n  events = filterEventsByField(events, \"lengthOfStay\", new IVL_PQ(null, new PQ(120, \"d\", true)));\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.EncounterPerformedEncounterInpatient_precondition_3 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"encounters\", \"statuses\": [\"performed\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.666.5.307\"};\n  var events = patient.getEvents(eventCriteria);\n  if (events.length > 0 || !Logger.short_circuit) events = EDU(events, hqmfjs.MeasurePeriod(patient));\n  if (events.length == 0) events.specificContext=hqmf.SpecificsManager.empty();\n  return events;\n}\n\nhqmfjs.GROUP_satisfiesAll_CHILDREN_4 = function(patient, initialSpecificContext) {\n  var events = INTERSECT(\n    hqmfjs.EncounterPerformedEncounterInpatient_precondition_2(patient, initialSpecificContext),\n    hqmfjs.EncounterPerformedEncounterInpatient_precondition_3(patient, initialSpecificContext)\n  );\n\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.GROUP_variable_CHILDREN_6 = function(patient, initialSpecificContext) {\n  var events = UNION(\n    hqmfjs.GROUP_satisfiesAll_CHILDREN_4(patient, initialSpecificContext)\n  );\n  // record the result of the source of the variable to the rationale\n  if(Logger.enable_rationale) Logger.record('GROUP_variable_CHILDREN_6',events);\n\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.EncounterPerformedEncounterInpatient_precondition_8 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"encounters\", \"statuses\": [\"performed\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.666.5.307\"};\n  var events = patient.getEvents(eventCriteria);\n  events = filterEventsByField(events, \"lengthOfStay\", new IVL_PQ(null, new PQ(120, \"d\", true)));\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.EncounterPerformedEncounterInpatient_precondition_9 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"encounters\", \"statuses\": [\"performed\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.666.5.307\"};\n  var events = patient.getEvents(eventCriteria);\n  if (events.length > 0 || !Logger.short_circuit) events = EDU(events, hqmfjs.MeasurePeriod(patient));\n  if (events.length == 0) events.specificContext=hqmf.SpecificsManager.empty();\n  return events;\n}\n\nhqmfjs.GROUP_satisfiesAll_CHILDREN_10 = function(patient, initialSpecificContext) {\n  var events = INTERSECT(\n    hqmfjs.EncounterPerformedEncounterInpatient_precondition_8(patient, initialSpecificContext),\n    hqmfjs.EncounterPerformedEncounterInpatient_precondition_9(patient, initialSpecificContext)\n  );\n\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.GROUP_variable_CHILDREN_12 = function(patient, initialSpecificContext) {\n  var events = UNION(\n    hqmfjs.GROUP_satisfiesAll_CHILDREN_10(patient, initialSpecificContext)\n  );\n  // record the result of the source of the variable to the rationale\n  if(Logger.enable_rationale) Logger.record('GROUP_variable_CHILDREN_12',events);\n  events.specific_occurrence = 'GROUP_variable_CHILDREN_12';\n\n  events.specificContext=new hqmf.SpecificOccurrence(Row.buildForDataCriteria(events.specific_occurrence, events))\n  return events;\n}\n\nhqmfjs.PatientCharacteristicBirthdateBirthdate_precondition_14 = function(patient, initialSpecificContext) {\n  var value = patient.birthtime() || null;\n  var events = value ? [value] : [];\n  events = SBS(events, hqmfjs.GROUP_variable_CHILDREN_12(patient), new IVL_PQ(new PQ(2, \"a\", true), null));\n  events.specificContext=events.specificContext||hqmf.SpecificsManager.identity();\n  return events;\n}\n\nhqmfjs.PatientCharacteristicBirthdateBirthdate_precondition_15 = function(patient, initialSpecificContext) {\n  var value = patient.birthtime() || null;\n  var events = value ? [value] : [];\n  events = SBS(events, hqmfjs.GROUP_variable_CHILDREN_12(patient), new IVL_PQ(null, new PQ(17, \"a\", true)));\n  events.specificContext=events.specificContext||hqmf.SpecificsManager.identity();\n  return events;\n}\n\nhqmfjs.EncounterPerformedEncounterInpatient_precondition_18 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"encounters\", \"statuses\": [\"performed\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.666.5.307\"};\n  var events = patient.getEvents(eventCriteria);\n  events = filterEventsByField(events, \"dischargeDisposition\", new CodeList(getCodes(\"2.16.840.1.113883.3.117.1.7.1.82\")));\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.GROUP_INTERSECT_CHILDREN_19 = function(patient, initialSpecificContext) {\n  var events = INTERSECT(\n    hqmfjs.GROUP_variable_CHILDREN_12(patient, initialSpecificContext),\n    hqmfjs.EncounterPerformedEncounterInpatient_precondition_18(patient, initialSpecificContext)\n  );\n\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.DiagnosisActiveAsthma_precondition_20 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"allProblems\", \"statuses\": [\"active\"], \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.117.1.7.1.271\"};\n  var events = patient.getEvents(eventCriteria);\n  events = filterEventsByField(events, \"ordinality\", new CodeList(getCodes(\"2.16.840.1.113883.3.117.1.7.1.14\")));\n  if (events.length > 0 || !Logger.short_circuit) events = SDU(events, hqmfjs.GROUP_variable_CHILDREN_12(patient));\n  if (events.length == 0) events.specificContext=hqmf.SpecificsManager.empty();\n  return events;\n}\n\nhqmfjs.CommunicationFromProviderToPatientAsthmaManagementPlan_precondition_23 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"communications\", \"includeEventsWithoutStatus\": true, \"valueSetId\": \"2.16.840.1.113883.3.117.1.7.1.131\"};\n  var events = patient.getEvents(eventCriteria);\n  events = filterEventsByCommunicationDirection(events, 'communication_from_provider_to_patient');\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.CommunicationFromProviderToPatientAsthmaManagementPlan_precondition_24 = function(patient, initialSpecificContext) {\n  var eventCriteria = {\"type\": \"communications\", \"includeEventsWithoutStatus\": true, \"negated\": true, \"negationValueSetId\": \"2.16.840.1.113883.3.117.1.7.1.93\", \"valueSetId\": \"2.16.840.1.113883.3.117.1.7.1.131\"};\n  var events = patient.getEvents(eventCriteria);\n  events = filterEventsByCommunicationDirection(events, 'communication_from_provider_to_patient');\n  hqmf.SpecificsManager.setIfNull(events);\n  return events;\n}\n\nhqmfjs.GROUP_UNION_CHILDREN_25 = function(patient, initialSpecificContext) {\n  var events = UNION(\n    hqmfjs.CommunicationFromProviderToPatientAsthmaManagementPlan_precondition_23(patient, initialSpecificContext),\n    hqmfjs.CommunicationFromProviderToPatientAsthmaManagementPlan_precondition_24(patient, initialSpecificContext)\n  );\n\n  if (events.length > 0 || !Logger.short_circuit) events = SDU(events, hqmfjs.GROUP_variable_CHILDREN_12(patient));\n  if (events.length == 0) events.specificContext=hqmf.SpecificsManager.empty();\n  return events;\n}\n\n\n\n        // #########################\n        // ##### MEASURE LOGIC #####\n        // #########################\n        \n        hqmfjs.initializeSpecifics = function(patient_api, hqmfjs) { hqmf.SpecificsManager.initialize(patient_api,hqmfjs,{\"id\":\"GROUP_variable_CHILDREN_12\",\"type\":\"OCCURRENCE_A_OF_ENCOUNTERINPATIENT\",\"function\":\"GROUP_variable_CHILDREN_12\"}) }\n\n        // INITIAL PATIENT POPULATION\n        hqmfjs.IPP = function(patient, initialSpecificContext) {\n  population_criteria_fn = allTrue('IPP', patient, initialSpecificContext,\n    allTrue('26', patient, initialSpecificContext, hqmfjs.DiagnosisActiveAsthma_precondition_20, hqmfjs.PatientCharacteristicBirthdateBirthdate_precondition_14, hqmfjs.PatientCharacteristicBirthdateBirthdate_precondition_15\n    )\n  );\n  if (typeof(population_criteria_fn) == 'function') {\n  \treturn population_criteria_fn();\n  } else {\n  \treturn population_criteria_fn;\n  }\n};\n\n\n        // STRATIFICATION\n        hqmfjs.STRAT=null;\n        // DENOMINATOR\n        hqmfjs.DENOM = function(patient, initialSpecificContext) {\n  population_criteria_fn = allTrue('DENOM', patient, initialSpecificContext,\n    allTrue('30', patient, initialSpecificContext, hqmfjs.GROUP_INTERSECT_CHILDREN_19\n    )\n  );\n  if (typeof(population_criteria_fn) == 'function') {\n  \treturn population_criteria_fn();\n  } else {\n  \treturn population_criteria_fn;\n  }\n};\n\n\n        // NUMERATOR\n        hqmfjs.NUMER = function(patient, initialSpecificContext) {\n  population_criteria_fn = allTrue('NUMER', patient, initialSpecificContext,\n    allTrue('32', patient, initialSpecificContext, hqmfjs.GROUP_UNION_CHILDREN_25\n    )\n  );\n  if (typeof(population_criteria_fn) == 'function') {\n  \treturn population_criteria_fn();\n  } else {\n  \treturn population_criteria_fn;\n  }\n};\n\n\n        hqmfjs.DENEX = function(patient) { return new Boolean(false); }\n        hqmfjs.DENEXCEP = function(patient) { return new Boolean(false); }\n        // CV\n        hqmfjs.MSRPOPL = function(patient) { return new Boolean(true); }\n        hqmfjs.MSRPOPLEX = function(patient) { return new Boolean(false); }\n        hqmfjs.OBSERV = function(patient) { return new Boolean(false); }\n        // VARIABLES\n        hqmfjs.VARIABLES = function(patient, initialSpecificContext) {\nhqmfjs.GROUP_variable_CHILDREN_6(patient, initialSpecificContext);\nreturn false;\n}\n        \n        \n        var occurrenceId = [\"GROUP_variable_CHILDREN_12\"];\n\n        hqmfjs.initializeSpecifics(patient_api, hqmfjs)\n        \n        var population = function() {\n          return executeIfAvailable(hqmfjs.IPP, patient_api);\n        }\n        var stratification = null;\n        if (hqmfjs.STRAT) {\n          stratification = function() {\n            return hqmf.SpecificsManager.setIfNull(executeIfAvailable(hqmfjs.STRAT, patient_api));\n          }\n        }\n        var denominator = function() {\n          return executeIfAvailable(hqmfjs.DENOM, patient_api);\n        }\n        var numerator = function() {\n          return executeIfAvailable(hqmfjs.NUMER, patient_api);\n        }\n        var exclusion = function() {\n          return executeIfAvailable(hqmfjs.DENEX, patient_api);\n        }\n        var denexcep = function() {\n          return executeIfAvailable(hqmfjs.DENEXCEP, patient_api);\n        }\n        var msrpopl = function() {\n          return executeIfAvailable(hqmfjs.MSRPOPL, patient_api);\n        }\n        var msrpoplex = function() {\n          return executeIfAvailable(hqmfjs.MSRPOPLEX, patient_api);\n        }\n        var observ = function(specific_context) {\n          \n          var observFunc = hqmfjs.OBSERV\n          if (typeof(observFunc)==='function')\n            return observFunc(patient_api, specific_context);\n          else\n            return [];\n        }\n        \n        var variables = function() {\n          if (Logger.enable_rationale) {\n            return executeIfAvailable(hqmfjs.VARIABLES, patient_api);\n          }\n        }\n        \n        var executeIfAvailable = function(optionalFunction, patient_api) {\n          if (typeof(optionalFunction)==='function') {\n            result = optionalFunction(patient_api);\n            \n            return result;\n          } else {\n            return false;\n          }\n        }\n\n        \n        if (typeof Logger != 'undefined') {\n          // clear out logger\n          Logger.logger = [];\n          Logger.rationale={};\n          if (typeof short_circuit == 'undefined') short_circuit = true;\n        \n          // turn on logging if it is enabled\n          if (enable_logging || enable_rationale) {\n            injectLogger(hqmfjs, enable_logging, enable_rationale, short_circuit);\n          } else {\n            Logger.enable_rationale = false;\n            Logger.short_circuit = short_circuit;\n          }\n        }\n\n        try {\n          map(patient, population, denominator, numerator, exclusion, denexcep, msrpopl, msrpoplex, observ, occurrenceId,false,stratification, variables);\n        } catch(err) {\n          print(err.stack);\n          throw err;\n        }\n\n        \n        };\n        ",
  "continuous_variable": false,
  "episode_of_care": true,
  "hqmf_document": {
    "id": "0338",
    "hqmf_id": "40280381-4B9A-3825-014B-BD8FA6B2062E",
    "hqmf_set_id": "E1CB05E0-97D5-40FC-B456-15C5DBF44309",
    "hqmf_version_number": 3,
    "title": "Home Management Plan of Care (HMPC) Document Given to Patient/Caregiver",
    "description": "An assessment that there is documentation in the medical record that a Home Management Plan of Care (HMPC) document was given to the pediatric asthma patient/caregiver.",
    "cms_id": "CMS26v3",
    "population_criteria": {
      "IPP": {
        "conjunction?": true,
        "type": "IPP",
        "title": "Initial Patient Population",
        "hqmf_id": "D88D0F2F-3176-44D8-99CE-63145C8BABB6",
        "preconditions": [
          {
            "id": 26,
            "preconditions": [
              {
                "id": 20,
                "reference": "DiagnosisActiveAsthma_precondition_20"
              },
              {
                "id": 14,
                "reference": "PatientCharacteristicBirthdateBirthdate_precondition_14"
              },
              {
                "id": 15,
                "reference": "PatientCharacteristicBirthdateBirthdate_precondition_15"
              }
            ],
            "conjunction_code": "allTrue"
          }
        ]
      },
      "DENOM": {
        "conjunction?": true,
        "type": "DENOM",
        "title": "Denominator",
        "hqmf_id": "BF63669B-FF4B-47E9-94DA-50F923F374CF",
        "preconditions": [
          {
            "id": 30,
            "preconditions": [
              {
                "id": 16,
                "reference": "GROUP_INTERSECT_CHILDREN_19"
              }
            ],
            "conjunction_code": "allTrue"
          }
        ]
      },
      "NUMER": {
        "conjunction?": true,
        "type": "NUMER",
        "title": "Numerator",
        "hqmf_id": "71B40FB7-90D3-41E4-AEFE-CCC1A174E368",
        "preconditions": [
          {
            "id": 32,
            "preconditions": [
              {
                "id": 21,
                "reference": "GROUP_UNION_CHILDREN_25"
              }
            ],
            "conjunction_code": "allTrue"
          }
        ]
      }
    },
    "data_criteria": {
      "PatientCharacteristicSexOncAdministrativeSex": {
        "title": "ONC Administrative Sex",
        "description": "Patient Characteristic Sex: ONC Administrative Sex",
        "code_list_id": "2.16.840.1.113762.1.4.1",
        "property": "gender",
        "type": "characteristic",
        "definition": "patient_characteristic_gender",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicSexOncAdministrativeSex",
        "variable": false,
        "value": {
          "type": "CD",
          "system": "Administrative Sex",
          "code": "M"
        }
      },
      "PatientCharacteristicRaceRace": {
        "title": "Race",
        "description": "Patient Characteristic Race: Race",
        "code_list_id": "2.16.840.1.114222.4.11.836",
        "property": "race",
        "type": "characteristic",
        "definition": "patient_characteristic_race",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicRaceRace",
        "variable": false,
        "inline_code_list": {
          "CDC Race": [
            "2076-8",
            "1002-5",
            "2131-1",
            "2028-9",
            "2054-5",
            "2106-3"
          ]
        }
      },
      "PatientCharacteristicEthnicityEthnicity": {
        "title": "Ethnicity",
        "description": "Patient Characteristic Ethnicity: Ethnicity",
        "code_list_id": "2.16.840.1.114222.4.11.837",
        "property": "ethnicity",
        "type": "characteristic",
        "definition": "patient_characteristic_ethnicity",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicEthnicityEthnicity",
        "variable": false,
        "inline_code_list": {
          "CDC Race": [
            "2135-2",
            "2186-5"
          ]
        }
      },
      "PatientCharacteristicPayerPayer": {
        "title": "Payer",
        "description": "Patient Characteristic Payer: Payer",
        "code_list_id": "2.16.840.1.114222.4.11.3591",
        "property": "payer",
        "type": "characteristic",
        "definition": "patient_characteristic_payer",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicPayerPayer",
        "variable": false,
        "inline_code_list": {
          "Source of Payment Typology": [
            "523",
            "41",
            "512",
            "953",
            "37",
            "212",
            "331",
            "6",
            "84",
            "521",
            "3115",
            "3119",
            "3222",
            "5",
            "312",
            "3116",
            "3113",
            "349",
            "32126",
            "121",
            "39",
            "333",
            "311",
            "3",
            "611",
            "389",
            "3711",
            "21",
            "32122",
            "32125",
            "122",
            "322",
            "73",
            "112",
            "54",
            "332",
            "82",
            "822",
            "3211",
            "3712",
            "62",
            "379",
            "119",
            "3221",
            "511",
            "43",
            "36",
            "123",
            "342",
            "59",
            "3223",
            "98",
            "8",
            "69",
            "3811",
            "35",
            "34",
            "55",
            "52",
            "79",
            "32121",
            "613",
            "519",
            "81",
            "38",
            "3114",
            "821",
            "25",
            "3212",
            "3112",
            "32123",
            "22",
            "213",
            "3229",
            "64",
            "3122",
            "321",
            "3123",
            "24",
            "91",
            "3713",
            "951",
            "529",
            "99",
            "85",
            "94",
            "1",
            "2",
            "83",
            "129",
            "3813",
            "522",
            "61",
            "334",
            "612",
            "29",
            "51",
            "211",
            "514",
            "9",
            "341",
            "3812",
            "31",
            "619",
            "33",
            "4",
            "63",
            "219",
            "92",
            "72",
            "313",
            "113",
            "89",
            "96",
            "362",
            "959",
            "823",
            "3121",
            "95",
            "53",
            "343",
            "515",
            "3819",
            "42",
            "44",
            "93",
            "513",
            "381",
            "371",
            "361",
            "3111",
            "71",
            "32124",
            "23",
            "111",
            "382",
            "9999",
            "369",
            "12",
            "372",
            "954",
            "19",
            "32",
            "7",
            "11"
          ]
        }
      },
      "EncounterPerformedEncounterInpatient_precondition_2": {
        "title": "Encounter Inpatient",
        "description": "Encounter, Performed: Encounter Inpatient",
        "code_list_id": "2.16.840.1.113883.3.666.5.307",
        "type": "encounters",
        "definition": "encounter",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "EncounterPerformedEncounterInpatient",
        "variable": false,
        "field_values": {
          "LENGTH_OF_STAY": {
            "type": "IVL_PQ",
            "high": {
              "type": "PQ",
              "unit": "d",
              "value": "120",
              "inclusive?": true,
              "derived?": false
            }
          }
        }
      },
      "EncounterPerformedEncounterInpatient_precondition_3": {
        "title": "Encounter Inpatient",
        "description": "Encounter, Performed: Encounter Inpatient",
        "code_list_id": "2.16.840.1.113883.3.666.5.307",
        "type": "encounters",
        "definition": "encounter",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "EncounterPerformedEncounterInpatient",
        "variable": false,
        "temporal_references": [
          {
            "type": "EDU",
            "reference": "MeasurePeriod"
          }
        ]
      },
      "GROUP_satisfiesAll_CHILDREN_4": {
        "title": "GROUP_satisfiesAll_CHILDREN_4",
        "description": "Encounter Inpatient : Encounter, Performed",
        "children_criteria": [
          "EncounterPerformedEncounterInpatient_precondition_2",
          "EncounterPerformedEncounterInpatient_precondition_3"
        ],
        "derivation_operator": "INTERSECT",
        "type": "derived",
        "definition": "satisfies_all",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "GROUP_satisfiesAll_CHILDREN_4",
        "variable": false
      },
      "GROUP_variable_CHILDREN_6": {
        "title": "GROUP_variable_CHILDREN_6",
        "description": "EncounterInpatient",
        "children_criteria": [
          "GROUP_satisfiesAll_CHILDREN_4"
        ],
        "derivation_operator": "UNION",
        "type": "derived",
        "definition": "derived",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "GROUP_variable_CHILDREN_6",
        "variable": true
      },
      "EncounterPerformedEncounterInpatient_precondition_8": {
        "title": "Encounter Inpatient",
        "description": "Encounter, Performed: Encounter Inpatient",
        "code_list_id": "2.16.840.1.113883.3.666.5.307",
        "type": "encounters",
        "definition": "encounter",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "EncounterPerformedEncounterInpatient",
        "variable": false,
        "field_values": {
          "LENGTH_OF_STAY": {
            "type": "IVL_PQ",
            "high": {
              "type": "PQ",
              "unit": "d",
              "value": "120",
              "inclusive?": true,
              "derived?": false
            }
          }
        }
      },
      "EncounterPerformedEncounterInpatient_precondition_9": {
        "title": "Encounter Inpatient",
        "description": "Encounter, Performed: Encounter Inpatient",
        "code_list_id": "2.16.840.1.113883.3.666.5.307",
        "type": "encounters",
        "definition": "encounter",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "EncounterPerformedEncounterInpatient",
        "variable": false,
        "temporal_references": [
          {
            "type": "EDU",
            "reference": "MeasurePeriod"
          }
        ]
      },
      "GROUP_satisfiesAll_CHILDREN_10": {
        "title": "GROUP_satisfiesAll_CHILDREN_10",
        "description": "Encounter Inpatient : Encounter, Performed",
        "children_criteria": [
          "EncounterPerformedEncounterInpatient_precondition_8",
          "EncounterPerformedEncounterInpatient_precondition_9"
        ],
        "derivation_operator": "INTERSECT",
        "type": "derived",
        "definition": "satisfies_all",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "GROUP_satisfiesAll_CHILDREN_10",
        "variable": false
      },
      "GROUP_variable_CHILDREN_12": {
        "title": "GROUP_variable_CHILDREN_12",
        "description": "Occurrence A of $EncounterInpatient",
        "children_criteria": [
          "GROUP_satisfiesAll_CHILDREN_10"
        ],
        "derivation_operator": "UNION",
        "type": "derived",
        "definition": "derived",
        "hard_status": false,
        "negation": false,
        "specific_occurrence": "A",
        "specific_occurrence_const": "OCCURRENCE_A_OF_ENCOUNTERINPATIENT",
        "source_data_criteria": "GROUP_variable_CHILDREN_12",
        "variable": true
      },
      "PatientCharacteristicBirthdateBirthdate_precondition_14": {
        "title": "Birthdate",
        "description": "Patient Characteristic Birthdate: Birthdate",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.70",
        "property": "birthtime",
        "type": "characteristic",
        "definition": "patient_characteristic_birthdate",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicBirthdateBirthdate",
        "variable": false,
        "inline_code_list": {
          "SNOMED-CT": [
            "3950001"
          ]
        },
        "temporal_references": [
          {
            "type": "SBS",
            "reference": "GROUP_variable_CHILDREN_12",
            "range": {
              "type": "IVL_PQ",
              "low": {
                "type": "PQ",
                "unit": "a",
                "value": "2",
                "inclusive?": true,
                "derived?": false
              }
            }
          }
        ]
      },
      "PatientCharacteristicBirthdateBirthdate_precondition_15": {
        "title": "Birthdate",
        "description": "Patient Characteristic Birthdate: Birthdate",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.70",
        "property": "birthtime",
        "type": "characteristic",
        "definition": "patient_characteristic_birthdate",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicBirthdateBirthdate",
        "variable": false,
        "inline_code_list": {
          "SNOMED-CT": [
            "3950001"
          ]
        },
        "temporal_references": [
          {
            "type": "SBS",
            "reference": "GROUP_variable_CHILDREN_12",
            "range": {
              "type": "IVL_PQ",
              "high": {
                "type": "PQ",
                "unit": "a",
                "value": "17",
                "inclusive?": true,
                "derived?": false
              }
            }
          }
        ]
      },
      "EncounterPerformedEncounterInpatient_precondition_18": {
        "title": "Encounter Inpatient",
        "description": "Encounter, Performed: Encounter Inpatient",
        "code_list_id": "2.16.840.1.113883.3.666.5.307",
        "type": "encounters",
        "definition": "encounter",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "EncounterPerformedEncounterInpatient",
        "variable": false,
        "field_values": {
          "DISCHARGE_STATUS": {
            "type": "CD",
            "code_list_id": "2.16.840.1.113883.3.117.1.7.1.82",
            "title": "Discharge To Home Or Police Custody"
          }
        }
      },
      "GROUP_INTERSECT_CHILDREN_19": {
        "title": "GROUP_INTERSECT_CHILDREN_19",
        "description": "",
        "children_criteria": [
          "GROUP_variable_CHILDREN_12",
          "EncounterPerformedEncounterInpatient_precondition_18"
        ],
        "derivation_operator": "INTERSECT",
        "type": "derived",
        "definition": "derived",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "GROUP_INTERSECT_CHILDREN_19",
        "variable": false
      },
      "DiagnosisActiveAsthma_precondition_20": {
        "title": "Asthma",
        "description": "Diagnosis, Active: Asthma",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.271",
        "type": "conditions",
        "definition": "diagnosis",
        "status": "active",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "DiagnosisActiveAsthma",
        "variable": false,
        "field_values": {
          "ORDINAL": {
            "type": "CD",
            "code_list_id": "2.16.840.1.113883.3.117.1.7.1.14",
            "title": "Principal"
          }
        },
        "temporal_references": [
          {
            "type": "SDU",
            "reference": "GROUP_variable_CHILDREN_12"
          }
        ]
      },
      "CommunicationFromProviderToPatientAsthmaManagementPlan_precondition_23": {
        "title": "Asthma Management Plan",
        "description": "Communication: From Provider to Patient: Asthma Management Plan",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.131",
        "type": "communications",
        "definition": "communication_from_provider_to_patient",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "CommunicationFromProviderToPatientAsthmaManagementPlan",
        "variable": false
      },
      "CommunicationFromProviderToPatientAsthmaManagementPlan_precondition_24": {
        "title": "Asthma Management Plan",
        "description": "Communication: From Provider to Patient: Asthma Management Plan",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.131",
        "type": "communications",
        "definition": "communication_from_provider_to_patient",
        "hard_status": false,
        "negation": true,
        "negation_code_list_id": "2.16.840.1.113883.3.117.1.7.1.93",
        "source_data_criteria": "CommunicationFromProviderToPatientAsthmaManagementPlan",
        "variable": false
      },
      "GROUP_UNION_CHILDREN_25": {
        "title": "GROUP_UNION_CHILDREN_25",
        "description": "",
        "children_criteria": [
          "CommunicationFromProviderToPatientAsthmaManagementPlan_precondition_23",
          "CommunicationFromProviderToPatientAsthmaManagementPlan_precondition_24"
        ],
        "derivation_operator": "UNION",
        "type": "derived",
        "definition": "derived",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "GROUP_UNION_CHILDREN_25",
        "variable": false,
        "temporal_references": [
          {
            "type": "SDU",
            "reference": "GROUP_variable_CHILDREN_12"
          }
        ]
      }
    },
    "source_data_criteria": {
      "DiagnosisActiveAsthma": {
        "title": "Asthma",
        "description": "Diagnosis, Active: Asthma",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.271",
        "type": "conditions",
        "definition": "diagnosis",
        "status": "active",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "DiagnosisActiveAsthma",
        "variable": false
      },
      "CommunicationFromProviderToPatientAsthmaManagementPlan": {
        "title": "Asthma Management Plan",
        "description": "Communication: From Provider to Patient: Asthma Management Plan",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.131",
        "type": "communications",
        "definition": "communication_from_provider_to_patient",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "CommunicationFromProviderToPatientAsthmaManagementPlan",
        "variable": false
      },
      "PatientCharacteristicBirthdateBirthdate": {
        "title": "Birthdate",
        "description": "Patient Characteristic Birthdate: Birthdate",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.70",
        "property": "birthtime",
        "type": "characteristic",
        "definition": "patient_characteristic_birthdate",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicBirthdateBirthdate",
        "variable": false,
        "inline_code_list": {
          "SNOMED-CT": [
            "3950001"
          ]
        }
      },
      "EncounterPerformedEncounterInpatient": {
        "title": "Encounter Inpatient",
        "description": "Encounter, Performed: Encounter Inpatient",
        "code_list_id": "2.16.840.1.113883.3.666.5.307",
        "type": "encounters",
        "definition": "encounter",
        "status": "performed",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "EncounterPerformedEncounterInpatient",
        "variable": false
      },
      "PatientCharacteristicEthnicityEthnicity": {
        "title": "Ethnicity",
        "description": "Patient Characteristic Ethnicity: Ethnicity",
        "code_list_id": "2.16.840.1.114222.4.11.837",
        "property": "ethnicity",
        "type": "characteristic",
        "definition": "patient_characteristic_ethnicity",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicEthnicityEthnicity",
        "variable": false,
        "inline_code_list": {
          "CDC Race": [
            "2135-2",
            "2186-5"
          ]
        }
      },
      "PatientCharacteristicExpiredExpired": {
        "title": "Expired",
        "description": "Patient Characteristic Expired: Expired",
        "code_list_id": "2.16.840.1.113883.3.117.1.7.1.309",
        "property": "expired",
        "type": "characteristic",
        "definition": "patient_characteristic_expired",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicExpiredExpired",
        "variable": false
      },
      "PatientCharacteristicSexOncAdministrativeSex": {
        "title": "ONC Administrative Sex",
        "description": "Patient Characteristic Sex: ONC Administrative Sex",
        "code_list_id": "2.16.840.1.113762.1.4.1",
        "property": "gender",
        "type": "characteristic",
        "definition": "patient_characteristic_gender",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicSexOncAdministrativeSex",
        "variable": false,
        "value": {
          "type": "CD",
          "system": "Administrative Sex",
          "code": "M"
        }
      },
      "PatientCharacteristicPayerPayer": {
        "title": "Payer",
        "description": "Patient Characteristic Payer: Payer",
        "code_list_id": "2.16.840.1.114222.4.11.3591",
        "property": "payer",
        "type": "characteristic",
        "definition": "patient_characteristic_payer",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicPayerPayer",
        "variable": false,
        "inline_code_list": {
          "Source of Payment Typology": [
            "523",
            "41",
            "512",
            "953",
            "37",
            "212",
            "331",
            "6",
            "84",
            "521",
            "3115",
            "3119",
            "3222",
            "5",
            "312",
            "3116",
            "3113",
            "349",
            "32126",
            "121",
            "39",
            "333",
            "311",
            "3",
            "611",
            "389",
            "3711",
            "21",
            "32122",
            "32125",
            "122",
            "322",
            "73",
            "112",
            "54",
            "332",
            "82",
            "822",
            "3211",
            "3712",
            "62",
            "379",
            "119",
            "3221",
            "511",
            "43",
            "36",
            "123",
            "342",
            "59",
            "3223",
            "98",
            "8",
            "69",
            "3811",
            "35",
            "34",
            "55",
            "52",
            "79",
            "32121",
            "613",
            "519",
            "81",
            "38",
            "3114",
            "821",
            "25",
            "3212",
            "3112",
            "32123",
            "22",
            "213",
            "3229",
            "64",
            "3122",
            "321",
            "3123",
            "24",
            "91",
            "3713",
            "951",
            "529",
            "99",
            "85",
            "94",
            "1",
            "2",
            "83",
            "129",
            "3813",
            "522",
            "61",
            "334",
            "612",
            "29",
            "51",
            "211",
            "514",
            "9",
            "341",
            "3812",
            "31",
            "619",
            "33",
            "4",
            "63",
            "219",
            "92",
            "72",
            "313",
            "113",
            "89",
            "96",
            "362",
            "959",
            "823",
            "3121",
            "95",
            "53",
            "343",
            "515",
            "3819",
            "42",
            "44",
            "93",
            "513",
            "381",
            "371",
            "361",
            "3111",
            "71",
            "32124",
            "23",
            "111",
            "382",
            "9999",
            "369",
            "12",
            "372",
            "954",
            "19",
            "32",
            "7",
            "11"
          ]
        }
      },
      "PatientCharacteristicRaceRace": {
        "title": "Race",
        "description": "Patient Characteristic Race: Race",
        "code_list_id": "2.16.840.1.114222.4.11.836",
        "property": "race",
        "type": "characteristic",
        "definition": "patient_characteristic_race",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "PatientCharacteristicRaceRace",
        "variable": false,
        "inline_code_list": {
          "CDC Race": [
            "2076-8",
            "1002-5",
            "2131-1",
            "2028-9",
            "2054-5",
            "2106-3"
          ]
        }
      },
      "GROUP_variable_CHILDREN_6": {
        "title": "GROUP_variable_CHILDREN_6",
        "description": "EncounterInpatient",
        "children_criteria": [
          "GROUP_satisfiesAll_CHILDREN_4"
        ],
        "derivation_operator": "UNION",
        "type": "derived",
        "definition": "derived",
        "hard_status": false,
        "negation": false,
        "source_data_criteria": "GROUP_variable_CHILDREN_6",
        "variable": true
      },
      "GROUP_variable_CHILDREN_12": {
        "title": "GROUP_variable_CHILDREN_12",
        "description": "Occurrence A of $EncounterInpatient",
        "children_criteria": [
          "GROUP_satisfiesAll_CHILDREN_10"
        ],
        "derivation_operator": "UNION",
        "type": "derived",
        "definition": "derived",
        "hard_status": false,
        "negation": false,
        "specific_occurrence": "A",
        "specific_occurrence_const": "OCCURRENCE_A_OF_ENCOUNTERINPATIENT",
        "source_data_criteria": "GROUP_variable_CHILDREN_12",
        "variable": true
      }
    },
    "attributes": [
      {
        "id": "NQF_ID_NUMBER",
        "code": "OTH",
        "value": "Not Applicable",
        "name": "NQF ID Number"
      },
      {
        "id": "COPYRIGHT",
        "code": "COPY",
        "value": "LOINC(R) is a registered trademark of the Regenstrief Institute.\n\nThis material contains SNOMED Clinical Terms (R) (SNOMED CT(C)) copyright 2004-2014 International Health Terminology Standards Development Organization. All rights reserved.",
        "name": "Copyright"
      },
      {
        "id": "DISCLAIMER",
        "code": "DISC",
        "value": "These performance measures are not clinical guidelines and do not establish a standard of medical care, and have not been tested for all potential applications. The measures and specifications are provided without warranty.",
        "name": "Disclaimer"
      },
      {
        "id": "MEASURE_SCORING",
        "code": "MSRSCORE",
        "value": "Proportion",
        "name": "Measure Scoring"
      },
      {
        "id": "MEASURE_TYPE",
        "code": "MSRTYPE",
        "value": "Process",
        "name": "Measure Type"
      },
      {
        "id": "STRATIFICATION",
        "code": "STRAT",
        "value": "None",
        "name": "Stratification"
      },
      {
        "id": "RISK_ADJUSTMENT",
        "code": "MSRADJ",
        "value": "None",
        "name": "Risk Adjustment"
      },
      {
        "id": "RATE_AGGREGATION",
        "code": "MSRAGG",
        "value": "None",
        "name": "Rate Aggregation"
      },
      {
        "id": "RATIONALE",
        "code": "RAT",
        "value": "Asthma is the most common chronic disease in children and a major cause of morbidity and health care costs nationally (Adams, et al, 2001). In 2005, 5.2% of children with asthma had at least one asthma attack in the previous year (3.8 million children). Nearly two of every three children who currently have asthma had at least one attack in the past 12 months. Chronic asthma in children can account for an annual loss of more than 14 million school days per year, according to the Asthma and Allergy Foundation (Asthma Facts and Figures).\n\nIt is clear from multiple sources of evidence including the National Heart Lung and Blood Institute (NHLBI) Guidelines that actual self-management of asthma by the patient or caregiver leads to more positive outcomes. Appropriate self-management is completely reliant upon patient education. Patient education is more effective when it aims at training self-management skills that will alter behavior (Norris, et al, 2001).\n\nNHLBI notes that review of asthma management by expert clinicians is necessary but not sufficient to improve outcomes. Active learning, participating and verbalization of understanding are all strategies that a healthcare organization must incorporate with parents or caregivers of asthmatic children in order for them to understand and make the appropriate changes that can impact the disease in the child in question. Education programs have been effective in improving lung function, feelings of self-esteem, and consequently decreased missed days of school in children and adolescents (Phipatanakul, 2004). Acute hospitalization follow up is imperative to a successful discharge from the hospital, providing the caretaker with the resource information needed to contact the follow up facility, medical office or clinic setting (Schatz, et al, 2009).\n\nEnvironmental control consists of removal of asthma triggers from the environment. Multiple studies support the positive correlation of household maintenance factors such as control of cockroach dust, and the number of acute asthma attacks in asthmatic children (McConnell, et al, 2005 and Eggleston, et al, 2005). Evidence from Carter et al, (2001) supported by the National Institute of Health (NIH) grant found specifically that reduction in triggers such as household conditions i.e. dust mites, cockroach, cats and presence of molds and fungus, resulted in a decrease in acute care visits and an overall positive outcome of children.\n\nRescue action education related to early recognition of symptoms and proper action to control incidence of asthma attacks is noted to have positive outcomes for asthmatic children (Ducharme and Bhogal, 2008).",
        "name": "Rationale"
      },
      {
        "id": "CLINICAL_RECOMMENDATION_STATEMENT",
        "code": "CRS",
        "value": "Under-treatment and/or inappropriate treatment of asthma are recognized as major contributors to asthma morbidity and mortality. National guidelines for the diagnosis and management of asthma in children, recommend establishing a plan for maintaining control of asthma and for establishing plans for managing exacerbation.",
        "name": "Clinical Recommendation Statement"
      },
      {
        "id": "IMPROVEMENT_NOTATION",
        "code": "IDUR",
        "value": "Improvement noted as an increase in rate.",
        "name": "Improvement Notation"
      },
      {
        "id": "REFERENCE",
        "code": "REF",
        "value": "Adams RJ, Fuhlbrigge A, Finkelstein JA, Lozano P, Livingston JM, Weiss KB, and Weiss ST (2001). Use of Inhaled Anti-inflammatory Medication in Children with Asthma in Managed Care Settings. Archives of Pediatrics and Adolescent Medicine, 155, 501-507.American College of Chest Physicians (ACCP), 10th Annual ACCP Community Asthma and COPD Coalitions Symposium: A Physician's Perspective. Retrieved on February 17, 2015 from: http://69.36.35.38/accp/perspective/10th_Asthma Asthma Facts and Figures. Retrieved on February 17, 2015 from http://www.aafa.org/display.cfm?id=9&sub=42#_ftn20.Ducharme FM, Bhogal SK. The role of written action plans in childhood asthma.  Curr Opin Allergy Clin Immunol. 2008 Apr;8(2):177-88.\nEggleston, P.A., Butz, A., Rand, C., Curtin-Brosnan, J, Kanchanaraksa, S, Swartz, L, Breysse, P, Buckley, T., Diette, G., Merriman, B., Krishnan, J. (2005). Home Environmental Intervention in Inner-City Asthma: A Randomized Controlled Clinical Trial. Ann Allergy Asthma Immunology; 95 p518-524McConnell, R., Milam, J., Richardson, J., Galvan, J., Jones, C., Thorne, P.S., and Berhane, K. (2005). Educational Intervention to Control Cockroach allergen Exposure in The homes of Hispanic Children in Los Angeles; Results of the La Casa Study. Clinical Exp Allergy; 35, p426-433National Asthma Education and Prevention Program, http://www.nhlbi.nih.gov/about/org/naepp.National Heart Lung and Blood Institute (NHLBI). Guidelines for the Diagnosis and Management of Asthma (EPR-3). Retrieved on February 17, 2015 from:  http://www.nhlbi.nih.gov/health-pro/guidelines/current/asthma-guidelines.Phipatanakul, W. (2004). Effects of Educational Interventions for Self-Management of Asthma in Children and Adolsecents:Systematic Review and Meta-Analysis. Pediatrics; 114; 530. Schatz, M, Rachelefsky, G, Krishnan, J.A. (2009). Follow-up after Acute Asthma Episodes. What Improves Future Outcomes. Proc Am Thorac Soc; Vol 6 p 386-393.Zemek, R.L., Bhogal, S.K., Ducharme, F. M. (2008). Systematic Review of Randomized Controlled Trials Examining Written Action Plans in Children: What is the Plan? (Reprinted) Arch Pediatr Adolesc Med/ vol 162 No 2. Retrieved on February 17, 2015 from http://archpedi.jamanetwork.com/article.aspx?articleid=379087.",
        "name": "Reference"
      },
      {
        "id": "DEFINITION",
        "code": "DEF",
        "value": "None",
        "name": "Definition"
      },
      {
        "id": "GUIDANCE",
        "code": "GUIDE",
        "value": "The home management plan of care document should be a separate and patient-specific written instruction. The document must be present in the form of an explicit and separate document specific to the patient rather than components or segments of the plan spread across discharge instruction sheets, discharge orders, education sheets, or other instruction sheets.\n\nThe home management plan of care is represented in the eMeasure logic by a LOINC code for an asthma action plan document. This form, or equivalent, contains most of the components required for the home management plan of care, including information on:\n*\tMethods and timing of rescue actions: the home management plan of care addresses what to do if asthma symptoms worsen after discharge, including all of the following: 1) When to take action, i.e., assessment of severity (eg, peak flow meter reading, signs and symptoms to watch for); 2) What specific steps to take, i.e., initial treatment instructions (eg, inhaled relievers up to three treatments of 2-4 puffs by MDI at 20-minute intervals or single nebulizer treatment); 3) Contact information to be used, when an asthma attack occurs or is about to occur.\n*\tAppropriate use of long-term asthma medications (controllers), including the medication name, dose, frequency, and method of administration.\n*\tAppropriate use of rescue, quick-relief, or short acting medications of choice to quickly relieve asthma exacerbations (relievers), including the medication name, dose, frequency, and method of administration.\n*\tEnvironmental control and control of other triggers: information on avoidance or mitigation of environmental and other triggers.\n\nIn addition to the information outlined in the asthma action plan form (or equivalent document), the home management plan of care is required to include information regarding arrangements for referral or follow-up care with a healthcare provider, namely:\n*\tIf an appointment for referral or follow-up care with a healthcare provider has been made, the home management plan of care is required to include the provider/clinic/office name, as well as the date and time of the appointment.\n*\tIf an appointment for referral of follow-up care with a healthcare provider has NOT been made, the home management plan of care is required to include information for the patient/caregiver to be able to make arrangements for follow-up care, i.e., provider/clinic/office name, telephone number and time frame for appointment for follow-up care (eg, 7-10 days).\n\nThe home management plan of care can only be considered to comply with the criteria outlined in the measure logic if it meets the requirements outlined above and is appropriately filled-out with information specific to the patient.\n\nPatient refusal includes refusal by a caregiver. The caregiver is defined as the patient's family or any other person (eg, home health, VNA provider, prison official or other law enforcement personnel) who will be responsible for care of the patient after discharge.\n\nThe \"Discharge To Home Or Police Custody\" value set also intends to capture the following discharge disposition values:\n* Assisted Living Facilities\n* Court/Law Enforcement - includes detention facilities, jails, and prison\n* Home - includes board and care, foster or residential care, group or personal care homes, and homeless shelters\n* Home with Home Health Services\n* Outpatient Services including outpatient procedures at another hospital, Outpatient Chemical Dependency Programs and Partial Hospitalization.\n\nThe unit of measurement for this measure is an inpatient episode of care. Each distinct hospitalization should be reported, regardless of whether the same patient is admitted for inpatient care more than once during the measurement period. In addition, the eMeasure logic intends to represent events within or surrounding a single occurrence of an inpatient hospitalization.",
        "name": "Guidance"
      },
      {
        "id": "TRANSMISSION_FORMAT",
        "code": "OTH",
        "value": "TBD",
        "name": "Transmission Format"
      },
      {
        "id": "DENOMINATOR",
        "code": "DENOM",
        "value": "Pediatric asthma inpatients with an age of 2 through 17 years, length of stay less than or equal to 120 days, and discharged to home or police custody.",
        "name": "Denominator"
      },
      {
        "id": "DENOMINATOR_EXCLUSIONS",
        "code": "OTH",
        "value": "None",
        "name": "Denominator Exclusions"
      },
      {
        "id": "NUMERATOR",
        "code": "NUMER",
        "value": "Pediatric asthma inpatients with documentation that they or their caregivers were given a written Home Management Plan of Care (HMPC) document that addresses all of the following: \n1. Arrangements for follow-up care \n2. Environmental control and control of other triggers\n3. Method and timing of rescue actions\n4. Use of controllers \n5. Use of relievers",
        "name": "Numerator"
      },
      {
        "id": "NUMERATOR_EXCLUSIONS",
        "code": "OTH",
        "value": "Not applicable",
        "name": "Numerator Exclusions"
      },
      {
        "id": "DENOMINATOR_EXCEPTIONS",
        "code": "DENEXCEP",
        "value": "None",
        "name": "Denominator Exceptions"
      },
      {
        "id": "MEASURE_POPULATION",
        "code": "MSRPOPL",
        "value": "Not applicable",
        "name": "Measure Population"
      },
      {
        "id": "MEASURE_OBSERVATIONS",
        "code": "OTH",
        "value": "Not applicable",
        "name": "Measure Observations"
      },
      {
        "id": "SUPPLEMENTAL_DATA_ELEMENTS",
        "code": "OTH",
        "value": "For every patient evaluated by this measure also identify payer, race, ethnicity and sex.",
        "name": "Supplemental Data Elements"
      }
    ],
    "populations": [
      {
        "IPP": "IPP",
        "DENOM": "DENOM",
        "NUMER": "NUMER"
      }
    ],
    "measure_period": {
      "type": "IVL_TS",
      "low": {
        "type": "TS",
        "value": "201201010000",
        "inclusive?": true,
        "derived?": false
      },
      "high": {
        "type": "TS",
        "value": "201212312359",
        "inclusive?": true,
        "derived?": false
      },
      "width": {
        "type": "PQ",
        "unit": "a",
        "value": "1",
        "inclusive?": true,
        "derived?": false
      }
    }
  },
  "oids": [
    "2.16.840.1.113883.3.117.1.7.1.271",
    "2.16.840.1.113883.3.117.1.7.1.131",
    "2.16.840.1.113883.3.117.1.7.1.82",
    "2.16.840.1.113883.3.117.1.7.1.93",
    "2.16.840.1.113762.1.4.1",
    "2.16.840.1.114222.4.11.836",
    "2.16.840.1.114222.4.11.837",
    "2.16.840.1.114222.4.11.3591",
    "2.16.840.1.113883.3.666.5.307",
    "2.16.840.1.113883.3.117.1.7.1.70",
    "2.16.840.1.113883.3.117.1.7.1.14"
  ],
  "population_ids": {
    "IPP": "D88D0F2F-3176-44D8-99CE-63145C8BABB6",
    "DENOM": "BF63669B-FF4B-47E9-94DA-50F923F374CF",
    "NUMER": "71B40FB7-90D3-41E4-AEFE-CCC1A174E368"
  }
}
`)
