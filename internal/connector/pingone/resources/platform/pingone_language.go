package platform

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	pingoneresources "github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneLanguageResource{}
)

type PingoneLanguageResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneLanguageResource
func Language(clientInfo *connector.SDKClientInfo) *PingoneLanguageResource {
	return &PingoneLanguageResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneLanguageResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.LanguagesApi.ReadLanguages(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadLanguages"

	embedded, err := pingoneresources.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, language := range embedded.GetLanguages() {
		if language.Language != nil {
			languageId, languageIdOk := language.Language.GetIdOk()
			languageName, languageNameOk := language.Language.GetNameOk()

			if languageIdOk && languageNameOk {
				importBlocks = append(importBlocks, connector.ImportBlock{
					ResourceType: r.ResourceType(),
					ResourceName: *languageName,
					ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *languageId),
				})
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneLanguageResource) ResourceType() string {
	return "pingone_language"
}
