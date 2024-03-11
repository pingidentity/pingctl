package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneRoleAssignmentUserResource{}
)

type PingoneRoleAssignmentUserResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneRoleAssignmentUserResource
func RoleAssignmentUser(clientInfo *connector.SDKClientInfo) *PingoneRoleAssignmentUserResource {
	return &PingoneRoleAssignmentUserResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneRoleAssignmentUserResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.UsersApi.ReadAllUsers(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllUsers"
	//UserRoleAssignmentsApi.ReadUserRoleAssignments()

	usersEmbedded, err := GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, user := range usersEmbedded.GetUsers() {
		userId, userIdOk := user.GetIdOk()
		userName, userNameOk := user.GetNameOk()
		userNameGiven, userNameGivenOk := userName.GetGivenOk()

		if userIdOk && userNameOk && userNameGivenOk {
			apiExecuteFunc = r.clientInfo.ApiClient.ManagementAPIClient.UserRoleAssignmentsApi.ReadUserRoleAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *userId).Execute
			apiFunctionName = "ReadUserRoleAssignments"

			userRoleAssignmentsEmbedded, err := GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}
			for index, userRoleAssignment := range userRoleAssignmentsEmbedded.GetRoleAssignments() {
				userRoleAssignmentId, userRoleAssignmentIdOk := userRoleAssignment.GetIdOk()

				if userRoleAssignmentIdOk {
					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType: r.ResourceType(),
						ResourceName: fmt.Sprintf("%s_role_assignment_%d", *userNameGiven, index),
						ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *userId, *userRoleAssignmentId),
					})
				}
			}

		}
	}

	return &importBlocks, nil
}

func (r *PingoneRoleAssignmentUserResource) ResourceType() string {
	return "pingone_role_assignment_user"
}
