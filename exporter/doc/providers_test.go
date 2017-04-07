package doc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/projectcypress/cdatools/exporter/doc"
	"github.com/projectcypress/cdatools/models"
)

func TestProvidersPrint(t *testing.T) {
	startDate := int64(1) // with timeToFormat, converted to 19700101
	endDate := int64(1)   // with timeToFormat, converted to 19700101
	timestamp := int64(1)

	var tests = []struct {
		n        doc.ProviderPerformances
		expected string
	}{
		{
			doc.ProviderPerformances{
				Timestamp: timestamp,
				ProviderPerformances: []models.ProviderPerformance{
					models.ProviderPerformance{StartDate: &startDate, EndDate: &endDate,
						Provider: models.Provider{
							CDAIdentifiers: []models.CDAIdentifier{
								models.CDAIdentifier{Root: "root", Extension: "extension"},
								models.CDAIdentifier{Root: "2.16.840.1.113883.4.2", Extension: "extension2"}}}}}}, // TODO: Find out what's going on with the provider stuff
			fmt.Sprintf(providerPerformancesCat3TestTemplate, 19700101, 19700101, 19700101, 19700101, "root", "extension", "extension2"),
		},
	}

	for _, tt := range tests {
		actual := tt.n.Print()
		if strings.TrimSpace(actual) != strings.TrimSpace(tt.expected) {
			t.Errorf("ProviderPerformances.Print(): expected \n%s\n, actual \n%s", tt.expected, actual)
		}
	}
}

var providerPerformancesCat3TestTemplate = `<documentationOf typeCode="DOC">
	<serviceEvent classCode="PCPR"> <!-- care provision -->
		<effectiveTime>
			<low value="%d"/>
			<high value="%d"/>
		</effectiveTime>
		<!-- You can include multiple performers, each with an NPI, TIN, CCN. -->
		<performer typeCode="PRF">
			<time>
				<low value="%d"/>
				<high value="%d"/>
			</time>
			<assignedEntity>
				<id root="%s" extension="%s" />
				<representedOrganization>
					<id root="2.16.840.1.113883.4.2" extension="%s" />
					
				</representedOrganization>
			</assignedEntity>
		</performer>
		
	</serviceEvent>
</documentationOf>`
