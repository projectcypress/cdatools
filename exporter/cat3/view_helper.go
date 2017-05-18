package cat3

import (
	"bytes"
	"encoding/xml"
	"log"
	"text/template"
	"time"
)

func Print(tmpl string, data interface{}) string {
	t := template.New("")
	t, err := t.Funcs(exporterFuncMapCat3(t)).Parse(tmpl)
	if err != nil {
		log.Println("error making template:", err)
		return ""
	}
	var b bytes.Buffer
	err = t.Execute(&b, data)
	if err != nil {
		log.Println("error making template:", err)
		return ""
	}
	return b.String()
}

func exporterFuncMapCat3(cat3Template *template.Template) template.FuncMap {
	return template.FuncMap{
		"Print":        Print,
		"timeToFormat": timeToFormat,
		"escape":       escape,
	}
}

func escape(s string) string {
	b := new(bytes.Buffer)
	err := xml.EscapeText(b, []byte(s))
	if err != nil {
		log.Println("xml.EscapeText failed:", err)
		return ""
	}
	return b.String()
}

// timeToFormat parses time from a seconds since Epoch value, and spits out a string in the supplied format
func timeToFormat(t int64, f string) string {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Println("time.LoadLocation failed:", err)
		return ""
	}
	parsedTime := time.Unix(t, 0)
	return escape(parsedTime.In(utc).Format(f))
}
