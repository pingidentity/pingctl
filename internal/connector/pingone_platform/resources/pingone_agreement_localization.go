package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAgreementLocalizationResource{}
)

type PingoneAgreementLocalizationResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneAgreementLocalizationResource
func AgreementLocalization(clientInfo *connector.SDKClientInfo) *PingoneAgreementLocalizationResource {
	return &PingoneAgreementLocalizationResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAgreementLocalizationResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_agreement_localization resources...")

	agreementEntityArray, response, err := r.clientInfo.ApiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()
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

	l.Debug().Msgf("Generating Import Blocks for all pingone_agreement_localization resources...")
	for _, agreement := range agreementEmbedded.GetAgreements() {
		agreementId, agreementIdOk := agreement.GetIdOk()
		agreementName, agreementNameOk := agreement.GetNameOk()

		if agreementIdOk && agreementNameOk {
			agreementLanguageEntityArray, response, err := r.clientInfo.ApiClient.ManagementAPIClient.AgreementLanguagesResourcesApi.ReadAllAgreementLanguages(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *agreement.Id).Execute()
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

			for _, languageWrapper := range agreementLanguageEmbedded.GetLanguages() {
				if languageWrapper.AgreementLanguage != nil {
					agreementLanguage := languageWrapper.AgreementLanguage

					agreementLanguageLocale, agreementLanguageLocaleOk := agreementLanguage.GetLocaleOk()
					agreementLanguageId, agreementLanguageIdOk := agreementLanguage.GetIdOk()

					if agreementLanguageLocaleOk && agreementLanguageIdOk {
						importBlocks = append(importBlocks, connector.ImportBlock{
							ResourceType: r.ResourceType(),
							ResourceName: fmt.Sprintf("%s_%s", *agreementName, *agreementLanguageLocale),
							ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *agreementId, *agreementLanguageId),
						})
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
