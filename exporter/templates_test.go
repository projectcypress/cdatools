package exporter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"text/template"

	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/pebbe/util"
	"github.com/projectcypress/cdatools/models"
	"github.com/stretchr/testify/assert"
)

func TestCodeTemplate(t *testing.T) {
	// tag name is "code", has preferred code, attribute is not "codes"
	preferredCode := models.Concept{Code: "my_code", CodeSystem: "SNOMED-CT"}
	codeDisplay := models.CodeDisplay{CodeType: "my_code_type", TagName: "code", Attribute: "my_attr", PreferredCode: preferredCode, ExtraContent: "my_extra_content=\"extra_content_value\""}
	rootNode := xmlCodeRootNode(codeDisplay)
	assertXPath(t, rootNode, "//code", map[string]string{"code": "my_code", "codeSystem": "2.16.840.1.113883.6.96", "my_extra_content": "extra_content_value"}, nil)
	assertNoXPath(t, rootNode, "//code/originalText")

	// tag name is not "code"
	codeDisplay = models.CodeDisplay{CodeType: "my_code_type", TagName: "other_tag_name"}
	assertXPath(t, xmlCodeRootNode(codeDisplay), "//other_tag_name", nil, nil)

	// tag name is "code", no preferred code, exclude null flavor true
	codeDisplay = models.CodeDisplay{CodeType: "my_code_type", TagName: "code", ExcludeNullFlavor: true}
	assertXPath(t, xmlCodeRootNode(codeDisplay), "//code", nil, []string{"excludeNullFlavor"})

	// tag name is "code", no preferred code, exclude null flavor false
	codeDisplay = models.CodeDisplay{CodeType: "my_code_type", TagName: "code", Attribute: "my_attr", ExcludeNullFlavor: false, ExtraContent: "extra_stuff"}
	assertXPath(t, xmlCodeRootNode(codeDisplay), "//code", map[string]string{"nillFlavor": "UNK"}, nil)

	// attribute is "codes"
	codeDisplay = models.CodeDisplay{CodeType: "my_code_type", Attribute: "codes"}
	rootNode = xmlCodeRootNode(codeDisplay)
	assertXPath(t, rootNode, "//code/originalText", nil, nil)
	assertContent(t, rootNode, "//code/originalText", "my lil description")
}

func xmlCodeRootNode(codeDisplay models.CodeDisplay) *xml.ElementNode {
	codeDisplay.Description = "my lil description"
	xmlString := generateXML("_code.xml", codeDisplay)
	// printXmlString(xmlString)
	doc, err := xml.Parse([]byte(xmlString), nil, nil, 0, xml.DefaultEncodingBytes)
	util.CheckErr(err)
	return doc.Root()
}

func TestReasonTemplate(t *testing.T) {
	// do not negate reason, r2 compatable
	reason := models.CodedConcept{Code: "RESULT_CODE_1", CodeSystem: "2.16.840.1.113883.6.1"} // specified in cms9_26.json
	rootNode := xmlReasonRootNode(reason, false, true)
	assertXPath(t, rootNode, "//entryRelationship", map[string]string{"typeCode": "RSON"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation", map[string]string{"classCode": "OBS", "moodCode": "EVN"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation/templateId", map[string]string{"root": "2.16.840.1.113883.10.20.24.3.88"}, []string{"extension"})
	assertXPath(t, rootNode, "//entryRelationship/observation/id", map[string]string{"root": "1.3.6.1.4.1.115"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation/code", map[string]string{"code": "410666004", "codeSystem": "2.16.840.1.113883.6.96", "displayName": "reason", "codeSystemName": "SNOMED CT"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation/statusCode", map[string]string{"code": "completed"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation/effectiveTime", map[string]string{"value": "197001010000+0000"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation/value", map[string]string{"xsi:type": "CD", "code": "RESULT_CODE_1", "codeSystem": "2.16.840.1.113883.6.1", "sdtc:valueSet": "1.2.3.4.5.6.7.8.9.11"}, nil)

	// do not negate reason, not r2 compatable
	rootNode = xmlReasonRootNode(reason, false, false)
	assertXPath(t, rootNode, "//entryRelationship/observation/templateId", map[string]string{"root": "2.16.840.1.113883.10.20.24.3.88", "extension": "2014-12-01"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation/code", map[string]string{"code": "77301-0", "codeSystem": "2.16.840.1.113883.6.1", "displayName": "reason", "codeSystemName": "LOINC"}, nil)
	assertXPath(t, rootNode, "//entryRelationship/observation/effectiveTime/low", map[string]string{"value": "197001010000+0000"}, nil)

	// negate reason
	rootNode = xmlReasonRootNode(reason, true, true)
	assertXPath(t, rootNode, "//entryRelationship", nil, nil)

	// reason that is not specifed by a measure
	reason = models.CodedConcept{Code: "not_a_specified_code", CodeSystem: "¯\\_(ツ)_/¯"}
	xmlString := generateXML("_reason.xml", *getReasonData(reason, false, true))
	assert.Equal(t, "", strings.TrimSpace(xmlString))
}

func xmlReasonRootNode(reason models.CodedConcept, negateReason bool, r2CompatableQrdaOid bool) *xml.ElementNode {
	data := getReasonData(reason, negateReason, r2CompatableQrdaOid)
	setMapDataCriteria(data)
	xmlString := generateXML("_reason.xml", *data)
	// printXmlString(xmlString)
	doc, err := xml.Parse([]byte(xmlString), nil, nil, 0, xml.DefaultEncodingBytes)
	util.CheckErr(err)
	return doc.Root()
}

func getReasonData(reason models.CodedConcept, negateReason bool, r2CompatableQrdaOid bool) *entryInfo {
	encounter := models.Encounter{}
	if negateReason {
		encounter.Entry = &models.Entry{NegationReason: reason}
	} else {
		encounter.Entry = &models.Entry{Reason: reason}
	}
	encounter.StartTime = int64(0)
	if r2CompatableQrdaOid {
		encounter.Entry.Oid = "2.16.840.1.113883.3.560.1.79" // a valid hqmf oid (Encounter Performed)
	} else {
		encounter.Entry.Oid = "invalid_qrda_oid"
	}
	return &entryInfo{EntrySection: encounter}
}

func setMapDataCriteria(ei *entryInfo) {
	var fieldOids = map[string][]string{"REASON": []string{"1.2.3.4.5.6.7.8.9.11"},
		"ORDINAL":  []string{"1.2.3.4.5.6.7.8.9.10"},
		"SEVERITY": []string{"1.2.3.4.5.6.7.8.9.13"},
		"ROUTE":    []string{"1.2.3.4.5.6.7.8.9.12"}}
	var resultOids = []string{"1.2.3.4.5.6.7.8.9.14"}
	// var vsOid = "1.2.3.4.5.6.7.8.9"
	ei.MapDataCriteria = mdc{FieldOids: fieldOids, ResultOids: resultOids}
}

// test _2.16.840.1.113883.10.20.24.3.23.xml
func TestEncounterPerformedTemplate(t *testing.T) {
	qrdaOid := "2.16.840.1.113883.10.20.24.3.23"

	// only tests r2 compatable
	rootNode := xmlRootNodeForQrdaOid(qrdaOid)
	assertXPath(t, rootNode, "//entry/encounter", map[string]string{"classCode": "ENC", "moodCode": "ENV"}, nil)
	assertXPath(t, rootNode, "//entry/encounter/templateId[@root='2.16.840.1.113883.10.20.22.4.49']", nil, nil)
	assertXPath(t, rootNode, "//entry/encounter/templateId[@root='2.16.840.1.113883.10.20.24.3.23']", nil, nil)
	assertXPath(t, rootNode, "//entry/encounter/id", map[string]string{"root": "1.3.6.1.4.1.115"}, nil)
	assertContent(t, rootNode, "//entry/encounter/text", "Encounter, Performed: Encounter Inpatient")
	assertXPath(t, rootNode, "//entry/encounter/statusCode", map[string]string{"code": "completed"}, nil)
	assertXPath(t, rootNode, "//entry/encounter/effectiveTime", nil, nil)
	// assertXPath(t, rootNode, "//entry/encounter/effectiveTime/low", map[string]string{"value": "201408110415"}, nil)

	// test admit time vs start time for <low> tag. test discharge time vs end time for <high> tag
	ei := getDataForQrdaOid(qrdaOid)
	entrySection := ei.EntrySection.(models.Encounter)
	entrySection.AdmitTime = 1262462640 // is time 2010 01 02 1504 in EST
	entrySection.StartTime = 0
	entrySection.DischargeTime = 1293998640 // is time 2011 01 02 1504 in EST
	entrySection.EndTime = 0
	ei.EntrySection = entrySection
	rootNode = xmlRootNodeForQrdaOidWithData(qrdaOid, ei)
	assertXPath(t, rootNode, "//entry/encounter/effectiveTime/low", map[string]string{"value": "201001022004+0000"}, nil)
	assertXPath(t, rootNode, "//entry/encounter/effectiveTime/high", map[string]string{"value": "201101022004+0000"}, nil)
	entrySection.AdmitTime = 0
	entrySection.StartTime = 1293998640 // is time 2011 01 02 1504 in EST
	entrySection.DischargeTime = 0
	entrySection.EndTime = 1262462640 // is time 2010 01 02 1504 in EST
	ei.EntrySection = entrySection
	rootNode = xmlRootNodeForQrdaOidWithData(qrdaOid, ei)
	assertXPath(t, rootNode, "//entry/encounter/effectiveTime/low", map[string]string{"value": "201101022004+0000"}, nil)
	assertXPath(t, rootNode, "//entry/encounter/effectiveTime/high", map[string]string{"value": "201001022004+0000"}, nil)

	// continue testing here
}

// - - - - - - - - //
//   H E L P E R   //
// - - - - - - - - //

// asserts the xml path exists in xml string
// asserts that each expected attribute is on the tag
// asserts that each unexpected attribute is not on the tag
func assertXPath(t *testing.T, elem *xml.ElementNode, pathString string, expectedAttributes map[string]string, unexpectedAttributes []string) {
	path := xpath.Compile(pathString)
	nodes, err := elem.Search(path)
	util.CheckErr(err)
	assert.NotEqual(t, len(nodes), 0)
	for _, node := range nodes {
		for attr, val := range expectedAttributes {
			if attrVal := node.Attribute(attr); attrVal != nil {
				assert.Equal(t, val, attrVal.String())
			} else {
				assert.Fail(t, fmt.Sprintf("expected xml attribute %s was not found at xml path \"%s\"", attr, pathString))
			}
		}
		for _, attr := range unexpectedAttributes {
			assert.Nil(t, node.Attribute(attr))
		}
	}
}

// assert the xml path does not exist in the xml string
func assertNoXPath(t *testing.T, elem *xml.ElementNode, pathString string) {
	path := xpath.Compile(pathString)
	_, err := elem.Search(path)
	assert.Nil(t, err)
}

// assert all xml tags at the xml path do not contain the content
func assertContent(t *testing.T, elem *xml.ElementNode, pathString string, content string) {
	path := xpath.Compile(pathString)
	nodes, err := elem.Search(path)
	util.CheckErr(err)
	for _, node := range nodes {
		assert.Equal(t, content, node.Content())
	}
}

// just for debugging. remove later
func printXmlString(xmlString string) {
	fmt.Printf("\n====================\n")
	fmt.Printf(xmlString)
	fmt.Printf("\n====================\n\n")
}

// - - - - - - - - - - - - - - - - - - - //
//   G E N E R A T E   T E M P L A T E   //
// - - - - - - - - - - - - - - - - - - - //

func xmlRootNodeForQrdaOid(qrdaOid string) *xml.ElementNode {
	fileName := "_" + qrdaOid + ".xml"
	// printXmlString(generateXML(fileName, getDataForQrdaOid(qrdaOid)))
	return xmlRootNode(generateXML(fileName, getDataForQrdaOid(qrdaOid)))
}

// same as xmlRootNodeForQrdaOid() function but allows custom input data (should be an EntryInfo struct)
func xmlRootNodeForQrdaOidWithData(qrdaOid string, data interface{}) *xml.ElementNode {
	fileName := "_" + qrdaOid + ".xml"
	return xmlRootNode(generateXML(fileName, data))
}

func getDataForQrdaOid(qrdaOid string) entryInfo {
	var p models.Record
	var m []models.Measure
	var vs []models.ValueSet
	setPatientMeasuresAndValueSets(&p, &m, &vs)
	ei, err := getEntryInfo(p, m, qrdaOid) // ei stands for entry info
	if err != nil {
		util.CheckErr(err)
	}
	return ei
}

func xmlRootNode(xmlString string) *xml.ElementNode {
	doc, err := xml.Parse([]byte(xmlString), nil, nil, 0, xml.DefaultEncodingBytes)
	util.CheckErr(err)
	return doc.Root()
}

func generateXML(fileName string, templateData interface{}) string {
	var p models.Record
	var m []models.Measure
	var vs []models.ValueSet
	setPatientMeasuresAndValueSets(&p, &m, &vs)
	return generateTemplateForFile(makeTemplate(), fileName, templateData)
}

func setPatientMeasuresAndValueSets(patient *models.Record, measures *[]models.Measure, valueSets *[]models.ValueSet) {
	patientData, err := ioutil.ReadFile("../fixtures/records/1_n_n_ami.json")
	util.CheckErr(err)

	measureData, err := ioutil.ReadFile("../fixtures/measures/CMS9v4a.json")
	util.CheckErr(err)
	measureData2, err := ioutil.ReadFile("../fixtures/measures/CMS26v3.json")
	util.CheckErr(err)
	measureData = append([]byte("["), append(append(measureData, append([]byte(","), measureData2...)...), []byte("]")...)...)

	valueSetData, err := ioutil.ReadFile("../fixtures/value_sets/cms9_26.json")
	util.CheckErr(err)

	json.Unmarshal(patientData, patient)
	json.Unmarshal(measureData, measures)
	json.Unmarshal(valueSetData, valueSets)
	initializeVsMap(*valueSets)
}

func makeTemplate() *template.Template {
	temp := template.New("cat1")
	temp.Funcs(exporterFuncMap(temp))
	fileNames, err := AssetDir("templates/cat1")
	if err != nil {
		util.CheckErr(err)
	}
	for _, fileName := range fileNames {
		asset, err := Asset("templates/cat1/" + fileName)
		util.CheckErr(err)
		template.Must(temp.New(fileName).Parse(string(asset)))
	}
	return temp
}

func generateTemplateForFile(temp *template.Template, fileName string, templateData interface{}) string {
	var buf bytes.Buffer
	if err := temp.ExecuteTemplate(&buf, fileName, templateData); err != nil {
		util.CheckErr(err)
	}
	return buf.String()
}

func getEntryInfo(patient models.Record, measures []models.Measure, qrdaOid string) (entryInfo, error) {
	entryInfos := entryInfosForPatient(patient, measures)
	for _, ei := range entryInfos {
		if qrdaOid == HqmfToQrdaOid(models.ExtractEntry(&ei.EntrySection).Oid) {
			return ei, nil
		}
	}
	if len(entryInfos) == 0 {
		return entryInfo{}, errors.New("no entry infos found for patient and measures")
	}
	return entryInfo{}, errors.New(fmt.Sprintf("no entry info found with qrda oid \"%s\"", qrdaOid))
}
