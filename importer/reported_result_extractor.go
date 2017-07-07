package importer

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
	"github.com/pebbe/util"
)

type Cat3Results struct {
	SupplementalData SupplementalData  `json:"supplemental_data"`
	IPP              string            `json:"IPP"`
	DENOM            string            `json:"DENOM"`
	NUMER            string            `json:"NUMER"`
	PR               map[string]string `json:"PR"`
	DENEX            string            `json:"DENEX"`
	PopIds           PopIds            `json:"population_ids"`
}

type SupplementalData struct {
	IPP   SupDataElem `json:"IPP"`
	DENOM SupDataElem `json:"DENOM"`
	NUMER SupDataElem `json:"NUMER"`
	DENEX SupDataElem `json:"DENEX"`
}

type SupDataElem struct {
	RACE      map[string]string `json:"RACE"`
	ETHNICITY map[string]string `json:"ETHNICITY"`
	SEX       map[string]string `json:"SEX"`
	PAYER     map[string]string `json:"PAYER"`
}

type PopIds struct {
	IPP   string `json:"IPP"`
	DENOM string `json:"DENOM"`
	NUMER string `json:"NUMER"`
	DENEX string `json:"DENEX"`
}

func ExtractResultsByIds(measureID string, ids map[string]string, document string) string {

	doc, err := xml.Parse([]byte(document), nil, nil, 0, xml.DefaultEncodingBytes)
	util.CheckErr(err)

	xp := doc.DocXPathCtx()
	xp.RegisterNamespace("cda", "urn:hl7-org:v3")
	defer doc.Free()

	stratification := ids["stratification"]
	if stratification == "" {
		stratification = ids["STRAT"]
	}

	node, ok := findMeasureNode(measureID, doc)

	if !ok {
		return "{}"
	}

	var results Cat3Results
	results = getMeasureComponents(node, ids, stratification)

	resultsJSON, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
	}

	return string(resultsJSON)
}

func findMeasureNode(measureID string, doc *xml.XmlDocument) (node xml.Node, found bool) {
	measureXPath := xpath.Compile(fmt.Sprintf(`/cda:ClinicalDocument/cda:component/cda:structuredBody/cda:component/cda:section/cda:entry/cda:organizer[./cda:templateId[@root = '2.16.840.1.113883.10.20.27.3.1'] and ./cda:reference/cda:externalDocument/cda:id[@extension='%s' and @root='2.16.840.1.113883.4.738']]`, strings.ToUpper(measureID)))
	measureNodes, err := doc.Root().Search(measureXPath)
	if err != nil {
		log.Fatal(err)
	}
	if len(measureNodes) == 0 {
		return nil, false
	}
	return measureNodes[0], true
}

func getMeasureComponents(node xml.Node, ids map[string]string, stratification string) Cat3Results {
	var results Cat3Results

	for key, value := range ids {
		var val string
		var sup SupDataElem
		var pr map[string]string
		if key == "OBSERV" {
			msrpopl := ids["MSRPOPL"]
			val, sup = extractCVValue(node, value, msrpopl, stratification)
		} else {
			val, sup, pr = extractComponentValue(node, key, value, stratification)
		}
		if val != "" {
			switch key {
			case "IPP":
				results.IPP = val
				results.SupplementalData.IPP = sup
			case "DENOM":
				results.DENOM = val
				results.SupplementalData.DENOM = sup
			case "NUMER":
				results.NUMER = val
				results.SupplementalData.NUMER = sup
			case "DENEX":
				results.DENEX = val
				results.SupplementalData.DENEX = sup
			}
		}
		if pr != nil {
			results.PR = pr
		}
	}
	return results
}

func extractCVValue(node xml.Node, id string, msrpopl string, stratification string) (val string, sup SupDataElem) {
	observationXPath := xpath.Compile(fmt.Sprintf(`cda:component/cda:observation[./cda:value[@code = "MSRPOPL"] and ./cda:reference/cda:externalObservation/cda:id[@root='%s']]`, strings.ToUpper(msrpopl)))
	cv := FirstElement(observationXPath, node)
	if cv == nil {
		return "", SupDataElem{}
	}

	if stratification != "" {
		stratXPath := xpath.Compile(fmt.Sprintf(`cda:entryRelationship[@typeCode="COMP"]/cda:observation[./cda:templateId[@root = "2.16.840.1.113883.10.20.27.3.4"]  and ./cda:reference/cda:externalObservation/cda:id[@root=%s]]`, strings.ToUpper(stratification)))
		stratNode := FirstElement(stratXPath, node)
		val := getCVValue(stratNode, id)
		return val, SupDataElem{}
	}
	val = getCVValue(cv, id)
	sup = extractSupplementalData(cv)
	return val, sup
}

func extractComponentValue(node xml.Node, code string, id string, stratification string) (val string, sup SupDataElem, perfRate map[string]string) {
	observationXPath := xpath.Compile(fmt.Sprintf(`cda:component/cda:observation[./cda:value[@code = "%s"] and ./cda:reference/cda:externalObservation/cda:id["@root"='%s']]`, code, strings.ToUpper(id)))
	cv := FirstElement(observationXPath, node)

	if cv == nil {
		return "", SupDataElem{}, nil
	}

	if stratification != "" {
		stratXPath := xpath.Compile(fmt.Sprintf(`cda:entryRelationship[@typeCode="COMP"]/cda:observation[./cda:templateId[@root = "2.16.840.1.113883.10.20.27.3.4"]  and ./cda:reference/cda:externalObservation/cda:id[@root"='%s']]`, strings.ToUpper(stratification)))
		stratNode := FirstElement(stratXPath, node)
		val := getAggregateCount(stratNode)
		return val, SupDataElem{}, nil
	}
	val = getAggregateCount(cv)

	if code == "NUMER" && stratification == "" {
		perfRate = extractPerformanceRate(node, code, id)
	}

	sup = extractSupplementalData(cv)
	return val, sup, perfRate
}

func extractPerformanceRate(node xml.Node, code string, id string) (perfRateValue map[string]string) {
	perfRateXPath := xpath.Compile(fmt.Sprintf(`cda:component/cda:observation[./cda:templateId[@root = "2.16.840.1.113883.10.20.27.3.14"] and ./cda:reference/cda:externalObservation/cda:id["@root"='%s']]/cda:value`, strings.ToUpper(id)))
	perfNode := FirstElement(perfRateXPath, node)
	if perfNode != nil {
		nfXPath := xpath.Compile("./@nullFlavor")
		nfNode := FirstElement(nfXPath, perfNode)
		if nfNode != nil {
			perfRateValue["nullFlavor"] = "NA"
			return perfRateValue
		}
		value := perfNode.Attribute("value").String()
		perfRateValue["value"] = value
		return perfRateValue
	}
	return nil
}

func getCVValue(node xml.Node, cvID string) string {
	cvXPath := xpath.Compile(fmt.Sprintf(`cda:entryRelationship/cda:observation[./cda:templateId[@root="2.16.840.1.113883.10.20.27.3.2"] and ./cda:reference/cda:externalObservation/cda:id["@root"='%s']]/cda:value`, strings.ToUpper(cvID)))
	valueNode := FirstElement(cvXPath, node)
	var value string
	if valueNode != nil {
		value = valueNode.Attribute("value").String()
	}
	return value
}

func getAggregateCount(node xml.Node) string {
	valueXPath := xpath.Compile(`cda:entryRelationship/cda:observation[./cda:templateId[@root="2.16.840.1.113883.10.20.27.3.3"]]/cda:value`)
	valueNode := FirstElement(valueXPath, node)
	var value string
	if valueNode != nil {
		value = valueNode.Attribute("value").String()
	}
	return value
}

func extractSupplementalData(node xml.Node) (suppDataElem SupDataElem) {
	var suppDataMap = map[string]string{
		"RACE":      "2.16.840.1.113883.10.20.27.3.8",
		"ETHNICITY": "2.16.840.1.113883.10.20.27.3.7",
		"SEX":       "2.16.840.1.113883.10.20.27.3.6",
		"PAYER":     "2.16.840.1.113883.10.20.27.3.9",
	}
	var resultMap = make(map[string]map[string]string)
	for name, oid := range suppDataMap {
		var keyMap = make(map[string]string)
		countXPath := xpath.Compile(fmt.Sprintf("cda:entryRelationship/cda:observation[cda:templateId[@root='%s']]", oid))
		countNodes, err := node.Search(countXPath)
		if err != nil {
			log.Fatal(err)
		}
		if len(countNodes) != 0 {
			for _, node := range countNodes {
				valueXPath := xpath.Compile("cda:value")
				valueNode := FirstElement(valueXPath, node)
				count := getAggregateCount(node)
				nfXPath := xpath.Compile("./@nullFlavor")
				nfNode := FirstElement(nfXPath, node)
				if nfNode != nil {
					keyMap["UNK"] = count
				} else {
					keyMap[valueNode.Attribute("code").String()] = count
				}
			}
		}
		resultMap[name] = keyMap
	}
	suppDataElem = SupDataElem{RACE: resultMap["RACE"], ETHNICITY: resultMap["ETHNICITY"], SEX: resultMap["SEX"], PAYER: resultMap["PAYER"]}
	return suppDataElem
}
