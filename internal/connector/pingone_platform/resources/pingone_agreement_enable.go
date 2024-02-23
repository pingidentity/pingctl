package resources

import (
	"context"
	"fmt"

	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAgreementEnableResource{}
)

type PingoneAgreementEnableResource struct {
	context       context.Context
	apiClient     *sdk.Client
	environmentID string
}

// Utility method for creating a PingoneAgreementResource
func AgreementEnableResource(ctx context.Context, apiClient *sdk.Client, environmentID string) *PingoneAgreementEnableResource {
	return &PingoneAgreementEnableResource{
		context:       ctx,
		apiClient:     apiClient,
		environmentID: environmentID,
	}
}

func (r *PingoneAgreementEnableResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_agreement_enable resources...")

	entityArray, response, err := r.apiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(r.context, r.environmentID).Execute()
	defer response.Body.Close()
	if err != nil {
		l.Error().Err(err).Msgf("ReadAllAgreements Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, err
	}

	if entityArray == nil {
		l.Error().Msgf("Returned ReadAllAgreements() entityArray is nil.")
		l.Error().Msgf("ReadAllAgreements Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_agreement_enable resources via ReadAllAgreements()")
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		l.Error().Msgf("Returned ReadAllAgreements() embedded data is nil.")
		l.Error().Msgf("ReadAllAgreements Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_agreement_enable resources via ReadAllAgreements()")
	}

	importBlocks := []connector.ImportBlock{}

	agreements, agreementsOk := embedded.GetAgreementsOk()

	if agreementsOk {
		l.Debug().Msgf("Generating Import Blocks for all pingone_agreement_enable resources...")
		for _, agreement := range agreements {
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
					ResourceName: fmt.Sprintf("%s_enable", *agreementName),
					ResourceID:   fmt.Sprintf("%s/%s", *agreementEnvironmentId, *agreementId),
				})
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAgreementEnableResource) ResourceType() string {
	return "pingone_agreement_enable"
}
