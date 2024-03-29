package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestUserExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.User(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_user",
			ResourceName: "test_user",
			ResourceID:   fmt.Sprintf("%s/8012b4a3-1874-42d3-803d-76d4a378335c", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_user",
			ResourceName: "test_user2",
			ResourceID:   fmt.Sprintf("%s/7570f416-5503-4eb0-9fc3-9a485bddd411", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_user",
			ResourceName: "testing",
			ResourceID:   fmt.Sprintf("%s/68cb3634-0ed4-4044-85d1-576eb3a55360", testutils.GetEnvironmentID()),
		},
	}

	expectedImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range expectedImportBlocks {
		expectedImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocksMap)
}
