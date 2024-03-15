package pingonessoresources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	pingoneresourcescommon "github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneApplicationFlowPolicyAssignmentResource{}
)

type PingoneApplicationFlowPolicyAssignmentResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneApplicationFlowPolicyAssignmentResource
func ApplicationFlowPolicyAssignment(clientInfo *connector.SDKClientInfo) *PingoneApplicationFlowPolicyAssignmentResource {
	return &PingoneApplicationFlowPolicyAssignmentResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneApplicationFlowPolicyAssignmentResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteApplicationsFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationsApi.ReadAllApplications(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiApplicationFunctionName := "ReadAllApplications"

	embedded, err := pingoneresourcescommon.GetManagementEmbedded(apiExecuteApplicationsFunc, apiApplicationFunctionName, r.ResourceType())
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
		default:
			continue
		}

		if appIdOk && appNameOk {
			apiExecutePoliciesFunc := r.clientInfo.ApiClient.ManagementAPIClient.ApplicationFlowPolicyAssignmentsApi.ReadAllFlowPolicyAssignments(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *appId).Execute
			apiPolicyFunctionName := "ReadAllFlowPolicyAssignments"

			policyEmbedded, err := pingoneresourcescommon.GetManagementEmbedded(apiExecutePoliciesFunc, apiPolicyFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, policy := range policyEmbedded.GetFlowPolicyAssignments() {
				policyId, policyIdOk := policy.GetIdOk()
				if policyIdOk {
					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType: r.ResourceType(),
						ResourceName: *appName,
						ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *appId, *policyId),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneApplicationFlowPolicyAssignmentResource) ResourceType() string {
	return "pingone_application_flow_policy_assignment"
}
