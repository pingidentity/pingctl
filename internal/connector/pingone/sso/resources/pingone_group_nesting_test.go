package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestGroupNestingExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.GroupNesting(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_group_nesting",
			ResourceName: "My parent group_My nested group",
			ResourceID:   fmt.Sprintf("%s/298cf355-6806-4058-b87e-1ae92c7fb13b/d12ae346-c596-438c-95e3-3d76f364d527", testutils.GetEnvironmentID()),
		},
	}

	expectedImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range expectedImportBlocks {
		expectedImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocksMap)
}