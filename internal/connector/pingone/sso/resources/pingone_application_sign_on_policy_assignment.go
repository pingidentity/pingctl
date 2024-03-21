package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneApplicationSignOnPolicyAssignmentResource{}
)

type PingoneApplicationSignOnPolicyAssignmentResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneApplicationSignOnPolicyAssignmentResource
func ApplicationSignOnPolicyAssignment(clientInfo *connector.SDKClientInfo) *PingoneApplicationSignOnPolicyAssignmentResource {
	return &PingoneApplicationSignOnPolicyAssignmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneApplicationSignOnPolicyAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
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
			appId     *string
			appIdOk   bool
			appName   *string
			appNameOk bool
		)

		switch {
		case app.ApplicationOIDC != nil:
			appId, appIdOk = app.ApplicationOIDC.GetIdOk()
			appName, appNameOk = app.ApplicationOIDC.GetNameOk()
		case app.ApplicationSAML != nil:
			appId, appIdOk = app.ApplicationSAML.GetIdOk()
			appName, appNameOk = app.ApplicationSAML.GetNameOk()
		case app.ApplicationExternalLink != nil:
			appId, appIdOk = app.ApplicationExternalLink.GetIdOk()
			appName, appNameOk = app.ApplicationExternalLink.GetNameOk()
		case app.ApplicationWSFED != nil:
			appId, appIdOk = app.ApplicationWSFED.GetIdOk()
			appName, appNameOk = app.ApplicationWSFED.GetNameOk()
		default:
			continue
		}

		if appIdOk && appNameOk {
			apiExecutePoliciesFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationSignOnPolicyAssignmentsApi.ReadAllSignOnPolicyAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *appId).Execute
			apiPolicyFunctionName := "ReadAllSignOnPolicyAssignments"

			signOnPolicyEmbedded, err := common.GetManagementEmbedded(apiExecutePoliciesFunc, apiPolicyFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, signOnPolicy := range signOnPolicyEmbedded.GetSignOnPolicyAssignments() {
				signOnPolicyId, signOnPolicyIdOk := signOnPolicy.GetIdOk()
				if signOnPolicyIdOk {
					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType: r.ResourceType(),
						ResourceName: fmt.Sprintf("%s_%s", *appName, *signOnPolicyId),
						ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *appId, *signOnPolicyId),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneApplicationSignOnPolicyAssignmentResource) ResourceType() string {
	return "pingone_application_sign_on_policy_assignment"
}
