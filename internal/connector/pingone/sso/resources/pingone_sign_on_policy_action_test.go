package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestSignOnPolicyActionExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.SignOnPolicyAction(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_sign_on_policy_action",
			ResourceName: "testing_LOGIN",
			ResourceID:   fmt.Sprintf("%s/0667e65d-fcdf-4049-b1b4-9d59392ee8bc/8d6fbf89-6913-403d-ab16-1470af9be22f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_sign_on_policy_action",
			ResourceName: "testing_AGREEMENT",
			ResourceID:   fmt.Sprintf("%s/0667e65d-fcdf-4049-b1b4-9d59392ee8bc/23a73045-e9a7-4557-83c7-8aa3b7c7fb2e", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_sign_on_policy_action",
			ResourceName: "testing_IDENTITY_PROVIDER",
			ResourceID:   fmt.Sprintf("%s/0667e65d-fcdf-4049-b1b4-9d59392ee8bc/e975d90d-8355-45a2-94ba-3757734cc64b", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_sign_on_policy_action",
			ResourceName: "test_LOGIN",
			ResourceID:   fmt.Sprintf("%s/50cff7e5-7c95-4d1d-9fce-c9cdc7d6f6a3/8114540e-8deb-408b-9307-fa74f00d2683", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_sign_on_policy_action",
			ResourceName: "Single_Factor_LOGIN",
			ResourceID:   fmt.Sprintf("%s/b1fdc38d-ea0c-47b1-9d83-c48105bd6806/6cc634a8-a89f-4632-8e84-45b976a18473", testutils.GetEnvironmentID()),
		},
	}

	expectedImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range expectedImportBlocks {
		expectedImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocksMap)
}