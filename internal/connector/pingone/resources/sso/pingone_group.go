package sso

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneGroupResource{}
)

type PingoneGroupResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneGroupResource
func Group(clientInfo *connector.SDKClientInfo) *PingoneGroupResource {
	return &PingoneGroupResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneGroupResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.GroupsApi.ReadAllGroups(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllGroups"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, group := range embedded.GetGroups() {
		groupId, groupIdOk := group.GetIdOk()
		groupName, groupNameOk := group.GetNameOk()

		if groupIdOk && groupNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *groupName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *groupId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneGroupResource) ResourceType() string {
	return "pingone_group"
}
