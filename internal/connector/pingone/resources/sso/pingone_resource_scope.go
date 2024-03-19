package sso

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneResourceScopeResource{}
)

type PingoneResourceScopeResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneResourceScopeResource
func ResourceScope(clientInfo *connector.SDKClientInfo) *PingoneResourceScopeResource {
	return &PingoneResourceScopeResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneResourceScopeResource) ExportAll() (*[]connector.ImportBlock, error) {
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

		if resourceIdOk && resourceNameOk {
			apiResourceScopesExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.ResourceScopesApi.ReadAllResourceScopes(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *resourceId).Execute
			apiResourceScopesFunctionName := "ReadAllResourceScopes"

			embeddedResourceScopes, err := common.GetManagementEmbedded(apiResourceScopesExecuteFunc, apiResourceScopesFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, scope := range embeddedResourceScopes.GetScopes() {
				scopeId, scopeIdOk := scope.GetIdOk()
				scopeName, scopeNameOk := scope.GetNameOk()
				if scopeIdOk && scopeNameOk {
					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType: r.ResourceType(),
						ResourceName: fmt.Sprintf("%s_%s", *resourceName, *scopeName),
						ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *resourceId, *scopeId),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneResourceScopeResource) ResourceType() string {
	return "pingone_resource_scope"
}
