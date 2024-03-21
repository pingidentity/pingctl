package resources

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneSignOnPolicyActionResource{}
)

type PingoneSignOnPolicyActionResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneSignOnPolicyActionResource
func SignOnPolicyAction(clientInfo *connector.SDKClientInfo) *PingoneSignOnPolicyActionResource {
	return &PingoneSignOnPolicyActionResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneSignOnPolicyActionResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.SignOnPoliciesApi.ReadAllSignOnPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "ReadAllSignOnPolicies"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, signOnPolicy := range embedded.GetSignOnPolicies() {
		var (
			actionId   *string
			actionIdOk bool
		)

		signOnPolicyId, signOnPolicyIdOk := signOnPolicy.GetIdOk()
		signOnPolicyName, signOnPolicyNameOk := signOnPolicy.GetNameOk()

		if signOnPolicyIdOk && signOnPolicyNameOk {

			apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.SignOnPolicyActionsApi.ReadAllSignOnPolicyActions(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID, *signOnPolicyId).Execute
			apiFunctionName := "ReadAllSignOnPolicyActions"

			actionEmbedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
			if err != nil {
				return nil, err
			}

			for _, action := range actionEmbedded.GetActions() {
				switch {
				case action.SignOnPolicyActionAgreement != nil:
					actionId, actionIdOk = action.SignOnPolicyActionAgreement.GetIdOk()
				case action.SignOnPolicyActionCommon != nil:
					actionId, actionIdOk = action.SignOnPolicyActionCommon.GetIdOk()
				case action.SignOnPolicyActionIDFirst != nil:
					actionId, actionIdOk = action.SignOnPolicyActionIDFirst.GetIdOk()
				case action.SignOnPolicyActionIDP != nil:
					actionId, actionIdOk = action.SignOnPolicyActionIDP.GetIdOk()
				case action.SignOnPolicyActionLogin != nil:
					actionId, actionIdOk = action.SignOnPolicyActionLogin.GetIdOk()
				case action.SignOnPolicyActionMFA != nil:
					actionId, actionIdOk = action.SignOnPolicyActionMFA.GetIdOk()
				case action.SignOnPolicyActionPingIDWinLoginPasswordless != nil:
					actionId, actionIdOk = action.SignOnPolicyActionPingIDWinLoginPasswordless.GetIdOk()
				case action.SignOnPolicyActionProgressiveProfiling != nil:
					actionId, actionIdOk = action.SignOnPolicyActionProgressiveProfiling.GetIdOk()
				default:
					continue
				}

				if actionIdOk {
					importBlocks = append(importBlocks, connector.ImportBlock{
						ResourceType: r.ResourceType(),
						ResourceName: fmt.Sprintf("%s_%s", *signOnPolicyName, *actionId),
						ResourceID:   fmt.Sprintf("%s/%s/%s", r.clientInfo.ExportEnvironmentID, *signOnPolicyId, *actionId),
					})
				}
			}
		}
	}

	return &importBlocks, nil
}

func (r *PingoneSignOnPolicyActionResource) ResourceType() string {
	return "pingone_sign_on_policy_action"
}
