package sso

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneGroupRoleAssignmentResource{}
)

type PingoneGroupRoleAssignmentResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneGroupRoleAssignmentResource
func GroupRoleAssignment(clientInfo *connector.SDKClientInfo) *PingoneGroupRoleAssignmentResource {
	return &PingoneGroupRoleAssignmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneGroupRoleAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
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
			apiGroupRoleAssignmentExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.GroupRoleAssignmentsApi.ReadGroupRoleAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *groupId).Execute
			apiGroupRoleAssignmentFunctionName := "ReadGroupRoleAssignments"

			embeddedGroupRoleAssignments, err := common.GetManagementEmbedded(apiGroupRoleAssignmentExecuteFunc, apiGroupRoleAssignmentFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, roleAssignment := range embeddedGroupRoleAssignments.GetRoleAssignments() {
				roleAssignmentId, roleAssignmentIdOk := roleAssignment.GetIdOk()
				if roleAssignmentIdOk {
					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType: r.ResourceType(),
						ResourceName: fmt.Sprintf("%s_%s", *groupName, *roleAssignmentId),
						ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *groupId, *roleAssignmentId),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneGroupRoleAssignmentResource) ResourceType() string {
	return "pingone_group_role_assignment"
}
