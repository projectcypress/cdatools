package exporter

import (
	"testing"

	"github.com/projectcypress/cdatools/models"
	"github.com/stretchr/testify/assert"
)

func TestHqmfToQrdaOid(t *testing.T) {
	// ["Device, Applied", "Encounter, Performed", "Diagnostic Study, Intolerance"]
	hqmfOids := []string{"2.16.840.1.113883.3.560.1.10", "2.16.840.1.113883.3.560.1.79", "2.16.840.1.113883.3.560.1.39"}
	qrdaOids := []string{"2.16.840.1.113883.10.20.24.3.7", "2.16.840.1.113883.10.20.24.3.23", "2.16.840.1.113883.10.20.24.3.16"}
	for i, hqmfOid := range hqmfOids {
		assert.Equal(t, qrdaOids[i], HqmfToQrdaOid(hqmfOid))
	}
}

func TestCodeDisplayForQrdaOid(t *testing.T) {
	// invalid qrda oid
	codeDisplays := codeDisplayForQrdaOid("not a valid qrda oid")
	assert.Equal(t, 0, len(codeDisplays))

	// qrda oid with multiple code displays
	codeDisplays = codeDisplayForQrdaOid("2.16.840.1.113883.10.20.24.3.23")
	assert.Equal(t, 3, len(codeDisplays))
	assert.Equal(t, "entryCode", codeDisplays[0].CodeType)
	assert.Equal(t, "code", codeDisplays[0].TagName)
	assert.Equal(t, false, codeDisplays[0].ExcludeNullFlavor)
	assert.Equal(t, []string{"SNOMED-CT", "ICD-9-CM", "ICD-10-CM", "CPT"}, codeDisplays[0].PreferredCodeSets)
}

func TestIsR2Compatible(t *testing.T) {
	// invalid hqmf oid
	entry := models.Encounter{Entry: &models.Entry{Oid: "not a valid hqmf r2 oid"}}
	assert.Equal(t, false, IsR2Compatible(entry))

	// valid hqmf oid
	entry = models.Encounter{Entry: &models.Entry{Oid: "2.16.840.1.113883.3.560.1.10"}}
	assert.Equal(t, true, IsR2Compatible(entry))
}
