package cat3

import (
	"github.com/projectcypress/cdatools/models"
)

type Addresses []models.Address

func (a Addresses) Template() string {
	t := `{{range . -}}
<addr use="{{escape .Use}}">
	{{range .Street -}}
	<streetAddressLine>{{escape .}}</streetAddressLine>
	{{end -}}
	<city>{{escape .City}}</city>
	<state>{{escape .State}}</state>
	<postalCode>{{escape .Zip}}</postalCode>
	<country>{{escape .Country}}</country>
</addr>
{{end -}}`

	return t
}
