package importer

import (
	"fmt"
	"log"
	"strings"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/gokogiri/xpath"
	"github.com/pebbe/util"
)

type Cat3Results struct {
	SupplementalData SupplementalData  `json:"supplemental_data"`
	IPP              int               `json:"IPP"`
	DENOM            int               `json:"DENOM"`
	NUMER            int               `json:"NUMER"`
	PR               map[string]string `json:"PR"`
	DENEX            int               `json:"DENEX"`
	PopIds           PopIds            `json:"population_ids"`
}

type SupplementalData struct {
	IPP   SupDataElem `json:"IPP"`
	DENOM SupDataElem `json:"DENOM"`
	NUMER SupDataElem `json:"NUMER"`
	DENEX SupDataElem `json:"DENEX"`
}

type SupDataElem struct {
	RACE      map[string]int    `json:"RACE"`
	ETHNICITY map[string]int    `json:"ETHNICITY"`
	SEX       map[string]int    `json:"SEX"`
	PAYER     map[string]string `json:"PAYER"`
}

type PopIds struct {
	IPP   string `json:"IPP"`
	DENOM string `json:"DENOM"`
	NUMER string `json:"NUMER"`
	DENEX string `json:"DENEX"`
}

func extractResultsByIdds(measureID string, ids map[string]string, document string) string {

	doc, err := xml.Parse([]byte(document), nil, nil, 0, xml.DefaultEncodingBytes)
	util.CheckErr(err)
	defer doc.Free()

	stratification := ids["stratification"]
	if stratification == "" {
		stratification = ids["STRAT"]
	}

	node, ok := findMeasureNode(measureID, doc)

	if !ok {
		return "{}"
	}

	var results map[string]string
	results = getMeasureComponents(node, ids, stratification)

	if len(results) == 0 {
		return "{}"
	}

	return ""
}

func findMeasureNode(measureID string, doc *xml.XmlDocument) (node xml.Node, found bool) {
	measureXPath := xpath.Compile(fmt.Sprintf(`/cda:ClinicalDocument/cda:component/cda:structuredBody/cda:component/cda:section/cda:entry/cda:organizer[ ./cda:templateId[@root = "2.16.840.1.113883.10.20.27.3.1"] and ./cda:reference/cda:externalDocument/cda:id[@extension='%s' and @root='2.16.840.1.113883.4.738']`, strings.ToUpper(measureID)))
	measureNodes, err := doc.Root().Search(measureXPath)
	if err != nil {
		log.Fatal(err)
	}
	if len(measureNodes) == 0 {
		return nil, false
	}
	return measureNodes[0], true
}

func getMeasureComponents(node xml.Node, ids map[string]string, stratification string) map[string]string {
	var results map[string]string
	results["supplemental_data"] = ""

	for key, value := range ids {
		var val string
		var sup map[string]map[string]string
		var pr map[string]string
		if key == "OBSERV" {
			msrpopl := ids["MSRPOPL"]
			val, sup := extractCVValue(node, value, msrpopl, stratification)
		} else {
			val, sup, pr := extractComponentValue(node, key, value, stratification)
		}
		if val != "" {
			results[key] = val
			results["supplemental_data"] = fmt.Sprintf(sup)
		}
		if pr != nil {
			results["PR"] = fmt.Sprintf(pr)
		}
	}
	return results
}

func extractCVValue(node xml.Node, id string, msrpopl string, stratification string) (val string, sup map[string]map[string]string) {
	observationXPath := xpath.Compile(fmt.Sprintf(`cda:component/cda:observation[./cda:value[@code = "MSRPOPL"] and ./cda:reference/cda:externalObservation/cda:id[#{translate("@root")}='%s']]`, strings.ToUpper(msrpopl)))
	cv := FirstElement(observationXPath, node)
	if cv == nil {
		return "", nil
	}

	if stratification != "" {
		stratXPath := xpath.Compile(fmt.Sprintf(`cda:entryRelationship[@typeCode="COMP"]/cda:observation[./cda:templateId[@root = "2.16.840.1.113883.10.20.27.3.4"]  and ./cda:reference/cda:externalObservation/cda:id[@root=%s]]`, strings.ToUpper(stratification)))
		stratNode := FirstElement(stratXPath, node)
		val := getCVValue(stratNode, id)
		return val, nil
	}
	val = getCVValue(cv, id)
	sup = extractSupplementalData(cv)
	return val, sup
}

func extractComponentValue(node xml.Node, code string, id string, stratification string) (val string, sup map[string]map[string]string, perfRate map[string]string) {
	observationXPath := xpath.Compile(fmt.Sprintf(`cda:component/cda:observation[./cda:value[@code = "%s"] and ./cda:reference/cda:externalObservation/cda:id["@root"='%s']]`, code, strings.ToUpper(id)))
	cv := FirstElement(observationXPath, node)

	if cv == nil {
		return "", nil, nil
	}

	if stratification != "" {
		stratXPath := xpath.Compile(fmt.Sprintf(`cda:entryRelationship[@typeCode="COMP"]/cda:observation[./cda:templateId[@root = "2.16.840.1.113883.10.20.27.3.4"]  and ./cda:reference/cda:externalObservation/cda:id[@root"='%s']]`, strings.ToUpper(stratification)))
		stratNode := FirstElement(stratXPath, node)
		val := getAggregateCount(stratNode)
		return val, nil, nil
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
		value := valueNode.Attribute("value").String()
	}
	return value
}

func getAggregateCount(node xml.Node) string {
	valueXPath := xpath.Compile(`cda:entryRelationship/cda:observation[./cda:templateId[@root="2.16.840.1.113883.10.20.27.3.3"]]/cda:value`)
	valueNode := FirstElement(valueXPath, node)
	var value string
	if valueNode != nil {
		value := valueNode.Attribute("value").String()
	}
	return value
}

func extractSupplementalData(node xml.Node) (resultMap map[string]map[string]string) {
	var suppDataMap = map[string]string{
		"RACE":      "2.16.840.1.113883.10.20.27.3.8",
		"ETHNICITY": "2.16.840.1.113883.10.20.27.3.7",
		"SEX":       "2.16.840.1.113883.10.20.27.3.6",
		"PAYER":     "2.16.840.1.113883.10.20.27.3.9",
	}
	for name, oid := range suppDataMap {
		var keyMap map[string]string
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
					keyMap[node.Attribute("value").String()] = count
				}
			}
		}
		resultMap[name] = keyMap
	}
	return resultMap
}
