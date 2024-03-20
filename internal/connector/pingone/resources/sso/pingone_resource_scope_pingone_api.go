package sso

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneResourceScopePingOneApiResource{}
)

type PingoneResourceScopePingOneApiResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneResourceScopePingOneApiResource
func ResourceScopePingOneApi(clientInfo *connector.SDKClientInfo) *PingoneResourceScopePingOneApiResource {
	return &PingoneResourceScopePingOneApiResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneResourceScopePingOneApiResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.ResourcesApi.ReadAllResources(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllResources"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, resource := range embedded.GetResources() {
		resourceId, resourceIdOk := resource.GetIdOk()
		resourceName, resourceNameOk := resource.GetNameOk()
		resourceType, resourceTypeOk := resource.GetTypeOk()

		if resourceIdOk && resourceNameOk && resourceTypeOk {
			apiResourceScopePingOneApisExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.ResourceScopesApi.ReadAllResourceScopes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *resourceId).Execute
			apiResourceScopePingOneApisFunctionName := "ReadAllResourceScopes"

			embeddedResourceScopePingOneApis, err := common.GetManagementEmbedded(apiResourceScopePingOneApisExecuteFunc, apiResourceScopePingOneApisFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, scopePingOneApi := range embeddedResourceScopePingOneApis.GetScopes() {
				scopePingOneApiId, scopePingOneApiIdOk := scopePingOneApi.GetIdOk()
				scopePingOneApiName, scopePingOneApiNameOk := scopePingOneApi.GetNameOk()
				isPingOneApiResource := strings.Contains(string(*resourceType), "PINGONE_API")
				if scopePingOneApiIdOk && scopePingOneApiNameOk && isPingOneApiResource {
					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType: r.ResourceType(),
						ResourceName: fmt.Sprintf("%s_%s", *resourceName, *scopePingOneApiName),
						ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *scopePingOneApiId),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneResourceScopePingOneApiResource) ResourceType() string {
	return "pingone_resource_scope_pingone_api"
}
