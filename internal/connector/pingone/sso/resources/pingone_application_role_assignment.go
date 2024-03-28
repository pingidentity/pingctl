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
	_ connector.ExportableResource = &PingoneApplicationRoleAssignmentResource{}
)

type PingoneApplicationRoleAssignmentResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneApplicationRoleAssignmentResource
func ApplicationRoleAssignment(clientInfo *connector.SDKClientInfo) *PingoneApplicationRoleAssignmentResource {
	return &PingoneApplicationRoleAssignmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneApplicationRoleAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
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
			appId                  *string
			appIdOk                bool
			appName                *string
			appNameOk              bool
			appAccessControlRole   *management.ApplicationAccessControlRole
			appAccessControlRoleOk bool
		)

		switch {
		case app.ApplicationOIDC != nil:
			appId, appIdOk = app.ApplicationOIDC.GetIdOk()
			appName, appNameOk = app.ApplicationOIDC.GetNameOk()
			if app.ApplicationOIDC.AccessControl != nil {
				appAccessControlRole, appAccessControlRoleOk = app.ApplicationOIDC.AccessControl.GetRoleOk()
			}
		case app.ApplicationSAML != nil:
			appId, appIdOk = app.ApplicationSAML.GetIdOk()
			appName, appNameOk = app.ApplicationSAML.GetNameOk()
			if app.ApplicationSAML.AccessControl != nil {
				appAccessControlRole, appAccessControlRoleOk = app.ApplicationSAML.AccessControl.GetRoleOk()
			}
		case app.ApplicationExternalLink != nil:
			appId, appIdOk = app.ApplicationExternalLink.GetIdOk()
			appName, appNameOk = app.ApplicationExternalLink.GetNameOk()
			if app.ApplicationExternalLink.AccessControl != nil {
				appAccessControlRole, appAccessControlRoleOk = app.ApplicationExternalLink.AccessControl.GetRoleOk()
			}
		default:
			continue
		}

		if appIdOk && appNameOk && appAccessControlRoleOk && appAccessControlRole.GetType() == management.ENUMAPPLICATIONACCESSCONTROLTYPE_ADMIN_USERS_ONLY {
			apiExecutePoliciesFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationRoleAssignmentsApi.ReadApplicationRoleAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *appId).Execute
			apiApplicationRoleAssignmentsFunctionName := "ReadApplicationRoleAssignments"

			appRoleAssignmentsEmbedded, err := common.GetManagementEmbedded(apiExecutePoliciesFunc, apiApplicationRoleAssignmentsFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, roleAssignment := range appRoleAssignmentsEmbedded.GetRoleAssignments() {
				roleAssignmentId, roleAssignmentIdOk := roleAssignment.GetIdOk()
				roleAssignmentRole, roleAssignmentRoleOk := roleAssignment.GetRoleOk()
				if roleAssignmentIdOk && roleAssignmentRoleOk {
					roleAssignmentRoleId, roleAssignmentRoleIdOk := roleAssignmentRole.GetIdOk()
					if roleAssignmentRoleIdOk {
						apiRole, resp, err := r.clientInfo.ApiClient.ManagementAPIClient.RolesApi.ReadOneRole(r.clientInfo.Context, *roleAssignmentRoleId).Execute()
						err = common.HandleClientResponse(resp, err, "ReadOneRole", r.ResourceType())
						if err != nil {
							return nil, err
						}
						if apiRole != nil {
							apiRoleName, apiRoleNameOk := apiRole.GetNameOk()
							if apiRoleNameOk {
								importBlocks = append(importBlocks, connector.ImportBlock{
									ResourceType: r.ResourceType(),
									ResourceName: fmt.Sprintf("%s_%s", *appName, *apiRoleName),
									ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *appId, *roleAssignmentId),
								})
							}
						}
					}
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneApplicationRoleAssignmentResource) ResourceType() string {
	return "pingone_application_role_assignment"
}