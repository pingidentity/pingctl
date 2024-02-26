package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneAgreementLocalizationRevisionResource{}
)

type PingoneAgreementLocalizationRevisionResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneAgreementLocalizationRevisionResource
func AgreementLocalizationRevisionResource(clientInfo *connector.SDKClientInfo) *PingoneAgreementLocalizationRevisionResource {
	return &PingoneAgreementLocalizationRevisionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAgreementLocalizationRevisionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all pingone_agreement_localization_revision resources...")

	agreementEntityArray, response, err := r.clientInfo.ApiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(r.clientInfo.Context, r.clientInfo.EnvironmentID).Execute()
	defer response.Body.Close()
	if err != nil {
		l.Error().Err(err).Msgf("ReadAllAgreements Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, err
	}

	if agreementEntityArray == nil {
		l.Error().Msgf("Returned ReadAllAgreements() entityArray is nil.")
		l.Error().Msgf("ReadAllAgreements Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_agreement_localization_revision resources via ReadAllAgreements()")
	}

	agreementEmbedded, agreementEmbeddedOk := agreementEntityArray.GetEmbeddedOk()
	if !agreementEmbeddedOk {
		l.Error().Msgf("Returned ReadAllAgreements() embedded data is nil.")
		l.Error().Msgf("ReadAllAgreements Response Code: %s\nResponse Body: %s", response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch pingone_agreement_localization_revision resources via ReadAllAgreements()")
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all pingone_agreement_localization_revision resources...")
	for _, agreement := range agreementEmbedded.GetAgreements() {
		agreementId, agreementIdOk := agreement.GetIdOk()
		agreementName, agreementNameOk := agreement.GetNameOk()
		agreementEnvironment, agreementEnvironmentOk := agreement.GetEnvironmentOk()
		var agreementEnvironmentId *string
		var agreementEnvironmentIdOk = false
		if agreementEnvironmentOk {
			agreementEnvironmentId, agreementEnvironmentIdOk = agreementEnvironment.GetIdOk()
		}

		if agreementIdOk && agreementNameOk && agreementEnvironmentOk && agreementEnvironmentIdOk {
			agreementLanguageEntityArray, response, err := r.clientInfo.ApiClient.ManagementAPIClient.AgreementLanguagesResourcesApi.ReadAllAgreementLanguages(r.clientInfo.Context, r.clientInfo.EnvironmentID, *agreement.Id).Execute()
			defer response.Body.Close()
			if err != nil {
				l.Error().Err(err).Msgf("ReadAllAgreementLanguages Response Code: %s\nResponse Body: %s", response.Status, response.Body)
				return nil, err
			}

			if agreementLanguageEntityArray == nil {
				l.Error().Msgf("Returned ReadAllAgreementLanguages() entityArray is nil.")
				l.Error().Msgf("ReadAllAgreementLanguages Response Code: %s\nResponse Body: %s", response.Status, response.Body)
				return nil, fmt.Errorf("failed to fetch pingone_agreement_localization_revision resources via ReadAllAgreementLanguages()")
			}

			agreementLanguageEmbedded, agreementLanguageEmbeddedOk := agreementLanguageEntityArray.GetEmbeddedOk()
			if !agreementLanguageEmbeddedOk {
				l.Error().Msgf("Returned ReadAllAgreementLanguages() embedded data is nil.")
				l.Error().Msgf("ReadAllAgreementLanguages Response Code: %s\nResponse Body: %s", response.Status, response.Body)
				return nil, fmt.Errorf("failed to fetch pingone_agreement_localization_revision resources via ReadAllAgreementLanguages()")
			}

			for _, languageWrapper := range agreementLanguageEmbedded.GetLanguages() {
				if languageWrapper.AgreementLanguage != nil {
					agreementLanguage := languageWrapper.AgreementLanguage

					agreementLanguageLocale, agreementLanguageLocaleOk := agreementLanguage.GetLocaleOk()
					agreementLanguageId, agreementLanguageIdOk := agreementLanguage.GetIdOk()

					if agreementLanguageLocaleOk && agreementLanguageIdOk {
						agreementLanguageRevisionEntityArray, response, err := r.clientInfo.ApiClient.ManagementAPIClient.AgreementRevisionsResourcesApi.ReadAllAgreementLanguageRevisions(r.clientInfo.Context, r.clientInfo.EnvironmentID, *agreementId, *agreementLanguageId).Execute()
						defer response.Body.Close()
						if err != nil {
							l.Error().Err(err).Msgf("ReadAllAgreementLanguageRevisions Response Code: %s\nResponse Body: %s", response.Status, response.Body)
							return nil, err
						}

						if agreementLanguageRevisionEntityArray == nil {
							l.Error().Msgf("Returned ReadAllAgreementLanguageRevisions() entityArray is nil.")
							l.Error().Msgf("ReadAllAgreementLanguageRevisions Response Code: %s\nResponse Body: %s", response.Status, response.Body)
							return nil, fmt.Errorf("failed to fetch pingone_agreement_localization_revision resources via ReadAllAgreementLanguageRevisions()")
						}

						agreementLanguageRevisionEmbedded, agreementLanguageRevisionEmbeddedOk := agreementLanguageRevisionEntityArray.GetEmbeddedOk()
						if !agreementLanguageRevisionEmbeddedOk {
							l.Error().Msgf("Returned ReadAllAgreementLanguageRevisions() embedded data is nil.")
							l.Error().Msgf("ReadAllAgreementLanguageRevisions Response Code: %s\nResponse Body: %s", response.Status, response.Body)
							return nil, fmt.Errorf("failed to fetch pingone_agreement_localization_revision resources via ReadAllAgreementLanguageRevisions()")
						}

						for revisionIndex, revision := range agreementLanguageRevisionEmbedded.GetRevisions() {
							revisionId, revisionIdOk := revision.GetIdOk()

							if revisionIdOk {
								importBlocks = append(importBlocks, connector.ImportBlock{
									ResourceType: r.ResourceType(),
									ResourceName: fmt.Sprintf("%s_%s_%d", *agreementName, *agreementLanguageLocale, (revisionIndex + 1)),
									ResourceID:   fmt.Sprintf("%s/%s/%s/%s", *agreementEnvironmentId, *agreementId, *agreementLanguageId, *revisionId),
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

func (r *PingoneAgreementLocalizationRevisionResource) ResourceType() string {
	return "pingone_agreement_localization_revision"
}
