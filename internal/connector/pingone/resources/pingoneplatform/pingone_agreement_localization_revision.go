package pingoneplatformresources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	pingoneresources "github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
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
func AgreementLocalizationRevision(clientInfo *connector.SDKClientInfo) *PingoneAgreementLocalizationRevisionResource {
	return &PingoneAgreementLocalizationRevisionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneAgreementLocalizationRevisionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.AgreementsResourcesApi.ReadAllAgreements(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllAgreements"

	agreementEmbedded, err := pingoneresources.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())
	for _, agreement := range agreementEmbedded.GetAgreements() {
		agreementId, agreementIdOk := agreement.GetIdOk()
		agreementName, agreementNameOk := agreement.GetNameOk()

		if agreementIdOk && agreementNameOk {
			apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.AgreementLanguagesResourcesApi.ReadAllAgreementLanguages(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *agreement.Id).Execute
			apiFunctionName := "ReadAllAgreementLanguages"

			agreementLanguageEmbedded, err := pingoneresources.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, languageWrapper := range agreementLanguageEmbedded.GetLanguages() {
				if languageWrapper.AgreementLanguage != nil {
					agreementLanguage := languageWrapper.AgreementLanguage

					agreementLanguageLocale, agreementLanguageLocaleOk := agreementLanguage.GetLocaleOk()
					agreementLanguageId, agreementLanguageIdOk := agreementLanguage.GetIdOk()

					if agreementLanguageLocaleOk && agreementLanguageIdOk {
						apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.AgreementRevisionsResourcesApi.ReadAllAgreementLanguageRevisions(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *agreementId, *agreementLanguageId).Execute
						apiFunctionName := "ReadAllAgreementLanguageRevisions"

						agreementLanguageRevisionEmbedded, err := pingoneresources.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
						if err != nil {
							return nil, err
						}

						for revisionIndex, revision := range agreementLanguageRevisionEmbedded.GetRevisions() {
							revisionId, revisionIdOk := revision.GetIdOk()

							if revisionIdOk {
								importBlocks = append(importBlocks, connector.ImportBlock{
									ResourceType: r.ResourceType(),
									ResourceName: fmt.Sprintf("%s_%s_%d", *agreementName, *agreementLanguageLocale, (revisionIndex + 1)),
									ResourceID:   fmt.Sprintf("%s/%s/%s/%s", r.clientInfo.ExportEnvironmentID, *agreementId, *agreementLanguageId, *revisionId),
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
