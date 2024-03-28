package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneResourceResource{}
)

type PingoneResourceResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneResourceResource
func Resource(clientInfo *connector.SDKClientInfo) *PingoneResourceResource {
	return &PingoneResourceResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneResourceResource) ExportAll() (*[]connector.ImportBlock, error) {
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
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *resourceName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *resourceId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneResourceResource) ResourceType() string {
	return "pingone_resource"
}