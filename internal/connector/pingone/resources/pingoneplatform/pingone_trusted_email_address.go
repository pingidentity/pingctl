package pingoneplatformresources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	pingoneresources "github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneTrustedEmailAddressResource{}
)

type PingoneTrustedEmailAddressResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneTrustedEmailAddressResource
func TrustedEmailAddress(clientInfo *connector.SDKClientInfo) *PingoneTrustedEmailAddressResource {
	return &PingoneTrustedEmailAddressResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneTrustedEmailAddressResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.TrustedEmailDomainsApi.ReadAllTrustedEmailDomains(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllTrustedEmailDomains"

	emailDomainEmbedded, err := pingoneresources.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, trustedEmailDomain := range emailDomainEmbedded.GetEmailDomains() {
		trustedEmailDomainId, trustedEmailDomainIdOk := trustedEmailDomain.GetIdOk()
		trustedEmailDomainName, trustedEmailDomainNameOk := trustedEmailDomain.GetDomainNameOk()

		if trustedEmailDomainIdOk && trustedEmailDomainNameOk {
			apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.TrustedEmailAddressesApi.ReadAllTrustedEmailAddresses(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *trustedEmailDomainId).Execute
			apiFunctionName := "ReadAllTrustedEmailAddresses"

			trustedEmailAddressEmbedded, err := pingoneresources.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, trustedEmailAddress := range trustedEmailAddressEmbedded.GetTrustedEmails() {
				if trustedEmailAddress.EmailAddress != "" {
					trustedEmailAddressId, trustedEmailAddressIdOk := trustedEmailAddress.GetIdOk()

					trustedEmailAddressDomainId, trustedEmailAddressDomainIdOk := trustedEmailAddress.GetDomainIdOk()

					if trustedEmailAddressIdOk && trustedEmailAddressDomainIdOk {
						importBlocks = append(importBlocks, connector.ImportBlock{
							ResourceType: r.ResourceType(),
							ResourceName: fmt.Sprintf("%s_%s", *trustedEmailDomainName, trustedEmailAddress.EmailAddress),
							ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *trustedEmailAddressDomainId, *trustedEmailAddressId),
						})
					}
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneTrustedEmailAddressResource) ResourceType() string {
	return "pingone_trusted_email_address"
}
