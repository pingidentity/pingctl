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
func AgreementResource(clientInfo *connector.SDKClientInfo) *PingoneAgreementResource {
	return &PingoneAgreementResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAgreementResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_agreement resources...")

	entityArray, response, err := r.clientInfo.ApiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(r.clientInfo.Context, r.clientInfo.EnvironmentID).Execute()
	defer response.Body.Close()
	if err != nil {
		l.Error().Err(err).Msgf("ReadAllAgreements Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, err
	}

	if entityArray == nil {
		l.Error().Msgf("Returned ReadAllAgreements() entityArray is nil.")
		l.Error().Msgf("ReadAllAgreements Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_agreement resources via ReadAllAgreements()")
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		l.Error().Msgf("Returned ReadAllAgreements() embedded data is nil.")
		l.Error().Msgf("ReadAllAgreements Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_agreement resources via ReadAllAgreements()")
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all pingone_agreement resources...")
	for _, agreement := range embedded.GetAgreements() {
		agreementId, agreementIdOk := agreement.GetIdOk()
		agreementName, agreementNameOk := agreement.GetNameOk()
		agreementEnvironment, agreementEnvironmentOk := agreement.GetEnvironmentOk()
		var agreementEnvironmentId *string
		var agreementEnvironmentIdOk = false
		if agreementEnvironmentOk {
			agreementEnvironmentId, agreementEnvironmentIdOk = agreementEnvironment.GetIdOk()
		}

		if agreementIdOk && agreementNameOk && agreementEnvironmentOk && agreementEnvironmentIdOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *agreementName,
				ResourceID:   fmt.Sprintf("%s/%s", *agreementEnvironmentId, *agreementId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAgreementResource) ResourceType() string {
	return "pingone_agreement"
}
