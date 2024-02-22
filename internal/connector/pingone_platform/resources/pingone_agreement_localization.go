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

	agreementEntityArray, response, err := r.apiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(r.context, r.environmentID).Execute()
	defer response.Body.Close()
	if err != nil {
		l.Error().Err(err).Msgf("ReadAllAgreements Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	if agreementEntityArray != nil && agreementEntityArray.Embedded != nil && agreementEntityArray.Embedded.Agreements != nil {
		l.Debug().Msgf("Generating Import Blocks for all pingone_agreement_localization resources...")
		for _, agreement := range agreementEntityArray.Embedded.Agreements {
			if agreement.Id != nil && agreement.Name != "" && agreement.Environment != nil && agreement.Environment.Id != nil {
				agreementLanguageEntityArray, response, err := r.apiClient.ManagementAPIClient.AgreementLanguagesResourcesApi.ReadAllAgreementLanguages(r.context, r.environmentID, *agreement.Id).Execute()
				defer response.Body.Close()
				if err != nil {
					l.Error().Err(err).Msgf("ReadAllAgreementLanguages Response Code: %s\nResponse Body: %s", response.Status, response.Body)
					return nil, err
				}

				if agreementLanguageEntityArray != nil && agreementLanguageEntityArray.Embedded != nil &&
					agreementLanguageEntityArray.Embedded.Languages != nil {

					for _, languageWrapper := range agreementLanguageEntityArray.Embedded.Languages {
						if languageWrapper.AgreementLanguage != nil {
							agreementLanguage := languageWrapper.AgreementLanguage

							if agreementLanguage.Locale != "" && agreementLanguage.Id != nil {
								importBlocks = append(importBlocks, connector.ImportBlock{
									ResourceType: r.ResourceType(),
									ResourceName: fmt.Sprintf("%s_%s", agreement.Name, agreementLanguage.Locale),
									ResourceID:   fmt.Sprintf("%s/%s/%s", *agreement.Environment.Id, *agreement.Id, *agreementLanguage.Id),
								})
							}
						}
					}
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneAgreementLocalizationResource) ResourceType() string {
	return "pingone_agreement_localization"
}
