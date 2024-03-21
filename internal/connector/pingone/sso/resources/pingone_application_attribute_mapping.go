package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneApplicationAttributeMappingResource{}
)

type PingoneApplicationAttributeMappingResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneApplicationAttributeMappingResource
func ApplicationAttributeMapping(clientInfo *connector.SDKClientInfo) *PingoneApplicationAttributeMappingResource {
	return &PingoneApplicationAttributeMappingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneApplicationAttributeMappingResource) ExportAll() (*[]connector.ImportBlock, error) {
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
		case app.ApplicationWSFED != nil:
			appId, appIdOk = app.ApplicationWSFED.GetIdOk()
			appName, appNameOk = app.ApplicationWSFED.GetNameOk()
		default:
			continue
		}

		if appIdOk && appNameOk {
			apiExecuteAttributeMappingFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationAttributeMappingApi.ReadAllApplicationAttributeMappings(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *appId).Execute
			apiAttributeMappingFunctionName := "ReadAllApplicationAttributeMappings"

			attributeMappingsEmbedded, err := common.GetManagementEmbedded(apiExecuteAttributeMappingFunc, apiAttributeMappingFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, attributeMapping := range attributeMappingsEmbedded.GetAttributes() {
				if attributeMapping.ApplicationAttributeMapping == nil {
					continue
				}

				attributeMappingId, attributeMappingIdOk := attributeMapping.ApplicationAttributeMapping.GetIdOk()
				attributeMappingName, attributeMappingNameOk := attributeMapping.ApplicationAttributeMapping.GetNameOk()

				if attributeMappingIdOk && attributeMappingNameOk {
					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType: r.ResourceType(),
						ResourceName: fmt.Sprintf("%s_%s", *appName, *attributeMappingName),
						ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *appId, *attributeMappingId),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneApplicationAttributeMappingResource) ResourceType() string {
	return "pingone_application_attribute_mapping"
}
