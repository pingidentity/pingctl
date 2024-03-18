package platform

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/resources/common"
	"github.com/pingidentity/pingctl/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingoneKeyRotationPolicyResource{}
)

type PingoneKeyRotationPolicyResource struct {
	clientInfo *connector.SDKClientInfo
}

// Utility method for creating a PingoneKeyRotationPolicyResource
func KeyRotationPolicy(clientInfo *connector.SDKClientInfo) *PingoneKeyRotationPolicyResource {
	return &PingoneKeyRotationPolicyResource{
		clientInfo: clientInfo,
	}
}

func (r *PingoneKeyRotationPolicyResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()

	l.Debug().Msgf("Fetching all %s resources...", r.ResourceType())

	apiExecuteFunc := r.clientInfo.ApiClient.ManagementAPIClient.KeyRotationPoliciesApi.GetKeyRotationPolicies(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute
	apiFunctionName := "GetKeyRotationPolicies"

	embedded, err := common.GetManagementEmbedded(apiExecuteFunc, apiFunctionName, r.ResourceType())
	if err != nil {
		return nil, err
	}

	importBlocks := []connector.ImportBlock{}

	l.Debug().Msgf("Generating Import Blocks for all %s resources...", r.ResourceType())

	for _, keyRotationPolicy := range embedded.GetKeyRotationPolicies() {
		keyRotationPolicyId, keyRotationPolicyIdOk := keyRotationPolicy.GetIdOk()
		keyRotationPolicyName, keyRotationPolicyNameOk := keyRotationPolicy.GetNameOk()

		if keyRotationPolicyIdOk && keyRotationPolicyNameOk {
			importBlocks = append(importBlocks, connector.ImportBlock{
				ResourceType: r.ResourceType(),
				ResourceName: *keyRotationPolicyName,
				ResourceID:   fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, *keyRotationPolicyId),
			})
		}
	}

	return &importBlocks, nil
}

func (r *PingoneKeyRotationPolicyResource) ResourceType() string {
	return "pingone_key_rotation_policy"
}