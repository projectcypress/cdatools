package document

import (
	"text/template"
	"time"

	"github.com/pborman/uuid"
)

func ExporterFuncMapCat3(cat3Template *template.Template) template.FuncMap {
	return template.FuncMap{
		"timeNow":                     time.Now().UTC().Unix,
		"newRandom":                   uuid.NewRandom,
		"timeToFormat":                timeToFormat,
		"identifierForInt":            identifierForInt,
		"identifierForIntp":           identifierForIntp,
		"identifierForString":         identifierForString,
		"escape":                      escape,
		"executeTemplateForEntry":     generateExecuteTemplateForEntry(cat3Template),
		"condAssign":                  condAssign,
		"valueOrNullFlavor":           valueOrNullFlavor,
		"dischargeDispositionDisplay": dischargeDispositionDisplay,
		"sdtcValueSetAttribute":       sdtcValueSetAttribute,
		"getTransferOid":              getTransferOid,
		"identifierForInterface":      identifierForInterface,
		"valueOrDefault":              valueOrDefault,
		"oidForCodeSystem":            oidForCodeSystem,
		"codeDisplayAttributeIsCodes": codeDisplayAttributeIsCodes,
		"hasPreferredCode":            hasPreferredCode,
		"hasLaterality":               hasLaterality,
		"negationIndicator":           negationIndicator,
		"isNil":                       isNil,
		"derefBool":                   derefBool,
	}
}
