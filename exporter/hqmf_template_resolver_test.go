package exporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHqmfToQrdaOid(t *testing.T) {
	// ["Device, Applied", "Encounter, Performed", "Diagnostic Study, Intolerance"]
	hqmfOids := []string{"2.16.840.1.113883.3.560.1.10", "2.16.840.1.113883.3.560.1.79", "2.16.840.1.113883.3.560.1.39", "2.16.840.1.113883.3.560.1.1001", "2.16.840.1.113883.3.560.1.1001"}
	vsOids := []string{"", "", "", "2.16.840.1.113883.3.117.1.7.1.403", "2.16.840.1.113883.3.526.3.1279"}
	qrdaOids := []string{"2.16.840.1.113883.10.20.24.3.7", "2.16.840.1.113883.10.20.24.3.23", "2.16.840.1.113883.10.20.24.3.16", "2.16.840.1.113883.10.20.24.3.101", "2.16.840.1.113883.10.20.24.3.103"}
	for i, hqmfOid := range hqmfOids {
		assert.Equal(t, qrdaOids[i], HqmfToQrdaOid(hqmfOid, vsOids[i]))
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
