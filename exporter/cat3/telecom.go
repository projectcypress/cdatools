package cat3

import (
	"github.com/projectcypress/cdatools/models"
)

type Telecoms []models.Telecom

func (te Telecoms) Template() string {
	t := `{{range . -}}
<telecom use="{{escape .Use}}" value="tel:+{{escape .Value}}" />
{{end -}}`
	return t
}
