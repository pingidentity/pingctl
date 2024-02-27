package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneCustomDomainResource{}
)

type PingoneCustomDomainResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneCustomDomainResource
func CustomDomain(clientInfo *connector.SDKClientInfo) *PingoneCustomDomainResource {
	return &PingoneCustomDomainResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneCustomDomainResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_custom_domain resources...")

	entityArray, response, err := r.clientInfo.ApiClient.ManagementAPIClient.CustomDomainsApi.ReadAllDomains(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
	defer response.Body.Close()
	if err != nil {
		l.Error().Err(err).Msgf("ReadAllDomains Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, err
	}

	if entityArray == nil {
		l.Error().Msgf("Returned ReadAllDomains() entity array is nil.")
		l.Error().Msgf("ReadAllDomains Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_custom_domain resources via ReadAllDomains()")
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		l.Error().Msgf("Returned ReadAllDomains() embedded data is nil.")
		l.Error().Msgf("ReadAllDomains Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_custom_domain resources via ReadAllDomains()")
	}

	importBlocks := []connector.ImportBlock{}

	for _, customDomain := range embedded.GetCustomDomains() {
		customDomainName, customDomainNameOk := customDomain.GetDomainNameOk()
		customDomainId, customDomainIdOk := customDomain.GetIdOk()

		if customDomainIdOk && customDomainNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *customDomainName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *customDomainId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneCustomDomainResource) ResourceType() string {
	return "pingone_custom_domain"
}
