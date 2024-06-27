package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneUserRoleAssignmentResource{}
)

type PingoneUserRoleAssignmentResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneUserRoleAssignmentResource
func UserRoleAssignment(clientInfo *connector.SDKClientInfo) *PingoneUserRoleAssignmentResource {
	return &PingoneUserRoleAssignmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneUserRoleAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.UsersApi.ReadAllUsers(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllUsers"

	usersEmbedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, user := range usersEmbedded.GetUsers() {
		userId, userIdOk := user.GetIdOk()
		userName, userNameOk := user.GetUsernameOk()

		if userIdOk && userNameOk {
			apiExecuteFunc = r.clientInfo.ApiClient.ManagementAPIClient.UserRoleAssignmentsApi.ReadUserRoleAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *userId).Execute
			apiFunctionName = "ReadUserRoleAssignments"

			userRoleAssignmentsEmbedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}
			for userRoleAssignmentIndex, userRoleAssignment := range userRoleAssignmentsEmbedded.GetRoleAssignments() {
				// if the role assignment comes from a group, skip it
				_, userRoleAssignmentGroupRoleAssignmentOk := userRoleAssignment.GetGroupOk()
				if userRoleAssignmentGroupRoleAssignmentOk {
					continue
				}

				userRoleAssignmentId, userRoleAssignmentIdOk := userRoleAssignment.GetIdOk()
				userRoleAssignmentRole, userRoleAssignmentRoleOk := userRoleAssignment.GetRoleOk()
				userRoleAssignmentScope, userRoleAssignmentScopeOk := userRoleAssignment.GetScopeOk()

				var (
					userRoleAssignmentRoleId   *string
					userRoleAssignmentRoleIdOk bool

					userRoleAssignmentScopeType   *management.EnumRoleAssignmentScopeType
					userRoleAssignmentScopeTypeOk bool
				)

				if userRoleAssignmentRoleOk {
					userRoleAssignmentRoleId, userRoleAssignmentRoleIdOk = userRoleAssignmentRole.GetIdOk()
				}

				if userRoleAssignmentScopeOk {
					userRoleAssignmentScopeType, userRoleAssignmentScopeTypeOk = userRoleAssignmentScope.GetTypeOk()
				}

				if userRoleAssignmentIdOk && userRoleAssignmentRoleOk && userRoleAssignmentRoleIdOk && userRoleAssignmentScopeOk && userRoleAssignmentScopeTypeOk {
					role, response, err := r.clientInfo.ApiClient.ManagementAPIClient.RolesApi.ReadOneRole(r.clientInfo.Context, *userRoleAssignmentRoleId).Execute()
					err = common.HandleClientResponse(response, err, "ReadOneRole", r.ResourceType())
					if err != nil {
						return nil, err
					}

					if role != nil {
						roleName, roleNameOk := role.GetNameOk()
						if roleNameOk {
							commentData := map[string]string{
								"Resource Type":                   r.ResourceType(),
								"Username":                        *userName,
								"Role Name":                       string(*roleName),
								"User Role Assignment Scope Type": string(*userRoleAssignmentScopeType),
								"Export Environment ID":           r.clientInfo.ExportEnvironmentID,
								"User ID":                         *userId,
								"User Role Assignment ID":         *userRoleAssignmentId,
							}

							importBlocks = append(importBlocks, connector.ImportBlock{
								ResourceType:       r.ResourceType(),
								ResourceName:       fmt.Sprintf("%s_%s_%s_%d", *userName, *roleName, *userRoleAssignmentScopeType, (userRoleAssignmentIndex + 1)),
								ResourceID:         fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *userId, *userRoleAssignmentId),
								CommentInformation: common.GenerateCommentInformation(commentData),
							})
						}
					}
				}
			}

		}
	}

	return &importBlocks, nil
}

func (r *PingoneUserRoleAssignmentResource) ResourceType() string {
	return "pingone_user_role_assignment"
}
