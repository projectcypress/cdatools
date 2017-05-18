package cat3

import (
	"github.com/projectcypress/cdatools/models"
)

type CDAIdentifiers []models.CDAIdentifier

func (i CDAIdentifiers) Template() string {
	t := `{{range . -}}
	<id {{if .Root}}root="{{escape .Root}}"{{end}} extension="{{escape .Extension}}" />
{{end -}}`
	return t
}
