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

// Utility method for creating a Pingone Trusted Email Domain Resource
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

	for _, emailDomain := range embedded.GetEmailDomains() {
		emailDomainId, emailDomainIdOk := emailDomain.GetIdOk()
		emailDomainName, emailDomainNameOk := emailDomain.GetDomainNameOk()

		if emailDomainIdOk && emailDomainNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *emailDomainName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *emailDomainId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneTrustedEmailDomainResource) ResourceType() string {
	return "pingone_trusted_email_domain"
}
