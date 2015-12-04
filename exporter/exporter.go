package exporter

import (
	"bytes"
	"encoding/json"
	"github.com/projectcypress/cdatools/models"
)

//export generate_cat1
func Generate_cat1(patient []byte) string {

	p := &models.Record{}

	json.Unmarshal(patient, p)

	var buf bytes.Buffer
	err := Cat1Tmpl(&buf, p)

	if err != nil {
		panic(err)
	} else {
		// b, _ := json.Marshal(p)
		// return string(b)
		return buf.String()
	}
}
