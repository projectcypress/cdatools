package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProviders(t *testing.T) {
	pps := []ProviderPerformance{
		ProviderPerformance{
			Provider: Provider{
				CDAIdentifiers: []CDAIdentifier{
					CDAIdentifier{
						Root:      "test",
						Extension: "value1",
					},
					CDAIdentifier{
						Root:      "notvalid",
						Extension: "notthisvalue",
					},
				},
			},
		},
		ProviderPerformance{
			Provider: Provider{
				CDAIdentifiers: []CDAIdentifier{
					CDAIdentifier{
						Root:      "test",
						Extension: "value2",
					},
					CDAIdentifier{
						Root:      "notvalid",
						Extension: "notthisvalueeither",
					},
				},
			},
		},
	}

	provs := GetProviders(pps)

	assert.Equal(t, 2, len(provs))

	ids := IdentifiersForRoot(provs, "test")

	assert.Equal(t, 2, len(ids))
	assert.Contains(t, ids, "value1")
	assert.Contains(t, ids, "value2")
}
