package sso

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneSignOnPolicyResource{}
)

type PingoneSignOnPolicyResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneSignOnPolicyResource
func SignOnPolicy(clientInfo *connector.SDKClientInfo) *PingoneSignOnPolicyResource {
	return &PingoneSignOnPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneSignOnPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
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
		signOnPolicyId, signOnPolicyIdOk := signOnPolicy.GetIdOk()
		signOnPolicyName, signOnPolicyNameOk := signOnPolicy.GetNameOk()

		if signOnPolicyIdOk && signOnPolicyNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *signOnPolicyName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *signOnPolicyId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneSignOnPolicyResource) ResourceType() string {
	return "pingone_sign_on_policy"
}
