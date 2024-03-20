package sso

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneApplicationSecretResource{}
)

type PingoneApplicationSecretResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneApplicationSecretResource
func ApplicationSecret(clientInfo *connector.SDKClientInfo) *PingoneApplicationSecretResource {
	return &PingoneApplicationSecretResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneApplicationSecretResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteApplicationsFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiApplicationFunctionName := "ReadAllApplications"

	embedded, err := common.GetManagementEmbedded(apiExecuteApplicationsFunc, apiApplicationFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, app := range embedded.GetApplications() {
		var (
			appId     *string
			appIdOk   bool
			appName   *string
			appNameOk bool
		)

		switch {
		case app.ApplicationOIDC != nil:
			appId, appIdOk = app.ApplicationOIDC.GetIdOk()
			appName, appNameOk = app.ApplicationOIDC.GetNameOk()
		case app.ApplicationSAML != nil:
			appId, appIdOk = app.ApplicationSAML.GetIdOk()
			appName, appNameOk = app.ApplicationSAML.GetNameOk()
		case app.ApplicationExternalLink != nil:
			appId, appIdOk = app.ApplicationExternalLink.GetIdOk()
			appName, appNameOk = app.ApplicationExternalLink.GetNameOk()
		default:
			continue
		}

		if appIdOk && appNameOk {
			apiExecuteSecretFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationSecretApi.ReadApplicationSecret(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *appId).Execute
			apiSecretFunctionName := "ReadApplicationSecret"

			applicationSecret, response, err := apiExecuteSecretFunc()
			defer response.Body.Close()

			if err != nil {
				l.Error().Err(err).Msgf("%s Response Code: %s\nResponse Body: %s", apiSecretFunctionName, response.Status, response.Body)
				return nil, err
			}

			validateApiResponseErr := common.ValidateApiResponse(l, response, apiSecretFunctionName, r.ResourceType())

			if validateApiResponseErr != nil {
				return nil, validateApiResponseErr
			}

			if *applicationSecret.Secret != "" {
				importBlocks = append(importBlocks, connector.ImportBlock{
					ResourceType: r.ResourceType(),
					ResourceName: fmt.Sprintf("%s_secret", *appName),
					ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *appId),
				})
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneApplicationSecretResource) ResourceType() string {
	return "pingone_application_secret"
}
