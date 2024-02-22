package resources

import (
	"context"

	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAgreementEnableResource{}
)

type PingoneAgreementLocalizationResource struct {
	context       context.Context
	apiClient     *sdk.Client
	environmentID string
}

// Utility method for creating a PingoneAgreementResource
func AgreementLocalizationResource(ctx context.Context, apiClient *sdk.Client, environmentID string) *PingoneAgreementLocalizationResource {
	return &PingoneAgreementLocalizationResource{
		context:       ctx,
		apiClient:     apiClient,
		environmentID: environmentID,
	}
}

func (r *PingoneAgreementLocalizationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_agreement_localization resources...")

	entityArray, response, err := r.apiClient.ManagementAPIClient.LanguagesApi.ReadLanguages(r.context, r.environmentID).Execute()
	defer response.Body.Close()
	if err != nil {
		l.Error().Err(err).Msgf("ReadLanguages Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, err
	}

	l.Debug().Msgf("Generating Import Blocks for all pingone_agreement_localization resources...")

	var importBlocks []connector.ImportBlock
	for _, language := range entityArray.Embedded.Languages {
		importBlocks = append(importBlocks, connector.ImportBlock{
			ResourceType: r.ResourceType(),
			ResourceName: language.AgreementLanguage.Locale,
			ResourceID:   *language.AgreementLanguage.Id,
		})
	}

	return &importBlocks, nil
}

func (r *PingoneAgreementLocalizationResource) ResourceType() string {
	return "pingone_agreement_localization"
}
