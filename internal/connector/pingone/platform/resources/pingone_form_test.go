package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestFormExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.Form(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_form",
			ResourceName: "Test Form",
			ResourceID:   fmt.Sprintf("%s/837c2f09-f8dc-41e9-b2f2-5f7bf1efded1", testutils.GetEnvironmentID()),
		},
	}

	expectedImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range expectedImportBlocks {
		expectedImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocksMap)
}