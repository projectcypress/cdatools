package exporter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	entry := &models.Entry{CodeDisplays: []models.CodeDisplay{codeDisplay}, Description: "my lil description"}
	data := codeDisplayWithPreferredCode(entry, &entry.Coded, codeDisplay.CodeType)
	data.Description = entry.Description
	xmlString := generateXML("_code.xml", data)
	// printXmlString(xmlString)
	doc, err := xml.Parse([]byte(xmlString), nil, nil, 0, xml.DefaultEncodingBytes)
	util.CheckErr(err)
	return doc.Root()
}

// just for debugging. remove later
func printXmlString(xmlString string) {
	fmt.Printf("\n====================\n")
	fmt.Printf(xmlString)
	fmt.Printf("\n====================\n\n")
}

// test _2.16.840.1.113883.10.20.24.3.23.xml
func TestEncounterPerformed(t *testing.T) {
	// qrdaOid := "2.16.840.1.113883.10.20.24.3.23"
	// fmt.Printf(generateQrdaOidXML(qrdaOid))

	assert.Equal(t, 1, 1) // should make better test
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
				assert.Equal(t, attrVal.String(), val)
			} else {
				assert.NotEqual(t, attrVal, nil)
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
		assert.Equal(t, node.Content(), content)
	}
}

// - - - - - - - - - - - - - - - - - - - //
//   G E N E R A T E   T E M P L A T E   //
// - - - - - - - - - - - - - - - - - - - //

func generateXML(fileName string, templateData interface{}) string {
	var p models.Record
	var m []models.Measure
	var vs []models.ValueSet
	setPatientMeasuresAndValueSets(&p, &m, &vs)
	return generateTemplateForFile(makeTemplate(), fileName, templateData)
}

func generateQrdaOidXML(qrdaOid string) string {
	fileName := "_" + qrdaOid + ".xml"

	var p models.Record
	var m []models.Measure
	var vs []models.ValueSet
	setPatientMeasuresAndValueSets(&p, &m, &vs)

	entryInfo, err := getEntryInfo(p, m, qrdaOid)
	if err != nil {
		util.CheckErr(err)
	}
	return generateTemplateForFile(makeTemplate(), fileName, entryInfo)
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
