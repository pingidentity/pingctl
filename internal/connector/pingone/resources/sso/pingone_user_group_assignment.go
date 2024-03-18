package sso

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneUserGroupAssignmentResource{}
)

type PingoneUserGroupAssignmentResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneUserGroupAssignmentResource
func UserGroupAssignment(clientInfo *connector.SDKClientInfo) *PingoneUserGroupAssignmentResource {
	return &PingoneUserGroupAssignmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneUserGroupAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteUsersFunc := r.clientInfo.ApiClient.ManagementAPIClient.UsersApi.ReadAllUsers(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiUserFunctionName := "ReadAllUsers"

	embedded, err := common.GetManagementEmbedded(apiExecuteUsersFunc, apiUserFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, user := range embedded.GetUsers() {
		userId, userIdOk := user.GetIdOk()
		username, usernameOk := user.GetUsernameOk()
		if userIdOk && usernameOk {
			apiUserGroupMembershipFunc := r.clientInfo.ApiClient.ManagementAPIClient.GroupMembershipApi.ReadAllGroupMembershipsForUser(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *userId).Execute
			apiUserGroupMembershipFunctionName := "ReadAllGroupMembershipsForUser"

			userGroupMembershipEmbedded, err := common.GetManagementEmbedded(apiUserGroupMembershipFunc, apiUserGroupMembershipFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, groupMembership := range userGroupMembershipEmbedded.GetGroupMemberships() {
				groupMembershipId, groupMembershipIdOk := groupMembership.GetIdOk()
				groupMembershipName, groupMembershipNameOk := groupMembership.GetNameOk()
				if groupMembershipIdOk && groupMembershipNameOk {
					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType: r.ResourceType(),
						ResourceName: fmt.Sprintf("%s_%s", *username, *groupMembershipName),
						ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *userId, *groupMembershipId),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneUserGroupAssignmentResource) ResourceType() string {
	return "pingone_user_group_assignment"
}
