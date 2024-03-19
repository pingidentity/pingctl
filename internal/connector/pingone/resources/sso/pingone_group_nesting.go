package sso

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneGroupNestingResource{}
)

type PingoneGroupNestingResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneGroupNestingResource
func GroupNesting(clientInfo *connector.SDKClientInfo) *PingoneGroupNestingResource {
	return &PingoneGroupNestingResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneGroupNestingResource) ExportAll() (*[]connector.ImportBlock, error) {
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

	for _, parentGroup := range embedded.GetGroups() {
		parentGroupId, parentGroupIdOk := parentGroup.GetIdOk()
		parentGroupName, parentGroupNameOk := parentGroup.GetNameOk()

		if parentGroupIdOk && parentGroupNameOk {
			apiGroupNestingExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.GroupsApi.ReadGroupNesting(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *parentGroupId).Execute
			apiGroupNestingFunctionName := "ReadGroupNesting"

			embeddedGroupNesting, err := common.GetManagementEmbedded(apiGroupNestingExecuteFunc, apiGroupNestingFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, nestedGroup := range embeddedGroupNesting.GetGroupMemberships() {
				nestedGroupId, nestedGroupIdOk := nestedGroup.GetIdOk()
				nestedGroupName, nestedGroupNameOk := nestedGroup.GetNameOk()
				if nestedGroupIdOk && nestedGroupNameOk {
					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType: r.ResourceType(),
						ResourceName: fmt.Sprintf("%s_%s", *parentGroupName, *nestedGroupName),
						ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *parentGroupId, *nestedGroupId),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneGroupNestingResource) ResourceType() string {
	return "pingone_group_nesting"
}