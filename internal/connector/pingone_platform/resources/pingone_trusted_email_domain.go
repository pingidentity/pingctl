package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneTrustedEmailDomainResource{}
)

type PingoneTrustedEmailDomainResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneTrustedEmailDomainResource
func TrustedEmailDomain(clientInfo *connector.SDKClientInfo) *PingoneTrustedEmailDomainResource {
	return &PingoneTrustedEmailDomainResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneTrustedEmailDomainResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.TrustedEmailDomainsApi.ReadAllTrustedEmailDomains(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllTrustedEmailDomains"

	embedded, err := GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, key := range embedded.() {
		keyId, keyIdOk := key.GetIdOk()
		keyName, keyNameOk := key.GetNameOk()

		if keyIdOk && keyNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *keyName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *keyId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneTrustedEmailDomainResource) ResourceType() string {
	return "domain_name"
}
