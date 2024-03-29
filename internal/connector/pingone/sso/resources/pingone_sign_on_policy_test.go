package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestSignOnPolicyExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.SignOnPolicy(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_sign_on_policy",
			ResourceName: "testing",
			ResourceID:   fmt.Sprintf("%s/0667e65d-fcdf-4049-b1b4-9d59392ee8bc", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_sign_on_policy",
			ResourceName: "test",
			ResourceID:   fmt.Sprintf("%s/50cff7e5-7c95-4d1d-9fce-c9cdc7d6f6a3", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_sign_on_policy",
			ResourceName: "Single_Factor",
			ResourceID:   fmt.Sprintf("%s/b1fdc38d-ea0c-47b1-9d83-c48105bd6806", testutils.GetEnvironmentID()),
		},
	}

	expectedImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range expectedImportBlocks {
		expectedImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocksMap)
}
