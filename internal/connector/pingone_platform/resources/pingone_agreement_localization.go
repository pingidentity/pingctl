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

	if agreementEntityArray == nil {
		l.Error().Msgf("Returned ReadAllAgreements() entityArray is nil.")
		l.Error().Msgf("ReadAllAgreements Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_agreement_localization resources via ReadAllAgreements()")
	}

	agreementEmbedded, agreementEmbeddedOk := agreementEntityArray.GetEmbeddedOk()
	if !agreementEmbeddedOk {
		l.Error().Msgf("Returned ReadAllAgreements() embedded data is nil.")
		l.Error().Msgf("ReadAllAgreements Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_agreement_localization resources via ReadAllAgreements()")
	}

	importBlocks := []connector.ImportBlock{}

	agreements, agreementsOk := agreementEmbedded.GetAgreementsOk()

	if agreementsOk {
		l.Debug().Msgf("Generating Import Blocks for all pingone_agreement_localization resources...")
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
				agreementLanguageEntityArray, response, err := r.apiClient.ManagementAPIClient.AgreementLanguagesResourcesApi.ReadAllAgreementLanguages(r.context, r.environmentID, *agreement.Id).Execute()
				defer response.Body.Close()
				if err != nil {
					l.Error().Err(err).Msgf("ReadAllAgreementLanguages Response Code: %s\nResponse Body: %s", response.Status, response.Body)
					return nil, err
				}

				if agreementLanguageEntityArray == nil {
					l.Error().Msgf("Returned ReadAllAgreementLanguages() entityArray is nil.")
					l.Error().Msgf("ReadAllAgreementLanguages Response Code: %s\nResponse Body: %s", response.Status, response.Body)
					return nil, fmt.Errorf("failed to fetch pingone_agreement_localization resources via ReadAllAgreementLanguages()")
				}

				agreementLanguageEmbedded, agreementLanguageEmbeddedOk := agreementLanguageEntityArray.GetEmbeddedOk()
				if !agreementLanguageEmbeddedOk {
					l.Error().Msgf("Returned ReadAllAgreementLanguages() embedded data is nil.")
					l.Error().Msgf("ReadAllAgreementLanguages Response Code: %s\nResponse Body: %s", response.Status, response.Body)
					return nil, fmt.Errorf("failed to fetch pingone_agreement_localization resources via ReadAllAgreementLanguages()")
				}

				languages, languagesOk := agreementLanguageEmbedded.GetLanguagesOk()

				if languagesOk {
					for _, languageWrapper := range languages {
						if languageWrapper.AgreementLanguage != nil {
							agreementLanguage := languageWrapper.AgreementLanguage

							agreementLanguageLocale, agreementLanguageLocaleOk := agreementLanguage.GetLocaleOk()
							agreementLanguageId, agreementLanguageIdOk := agreementLanguage.GetIdOk()

							if agreementLanguageLocaleOk && agreementLanguageIdOk {
								importBlocks = append(importBlocks, connector.ImportBlock{
									ResourceType: r.ResourceType(),
									ResourceName: fmt.Sprintf("%s_%s", *agreementName, *agreementLanguageLocale),
									ResourceID:   fmt.Sprintf("%s/%s/%s", *agreementEnvironmentId, *agreementId, *agreementLanguageId),
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
