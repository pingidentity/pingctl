package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAgreementResource{}
)

type PingoneAgreementResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneAgreementResource
func Agreement(clientInfo *connector.SDKClientInfo) *PingoneAgreementResource {
	return &PingoneAgreementResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAgreementResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllAgreements"

	embedded, err := GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())
	for _, agreement := range embedded.GetAgreements() {
		agreementId, agreementIdOk := agreement.GetIdOk()
		agreementName, agreementNameOk := agreement.GetNameOk()

		if agreementIdOk && agreementNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *agreementName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *agreementId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAgreementResource) ResourceType() string {
	return "pingone_agreement"
}
