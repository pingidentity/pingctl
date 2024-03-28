package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestAgreementExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.Agreement(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_agreement",
			ResourceName: "Test",
			ResourceID:   fmt.Sprintf("%s/37ab76b8-8eff-43ae-b499-a7dce9fe0e75", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_agreement",
			ResourceName: "Test2",
			ResourceID:   fmt.Sprintf("%s/38c0c463-b13d-4d22-8da5-f9fd8093d594", testutils.GetEnvironmentID()),
		},
	}

	expectedImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range expectedImportBlocks {
		expectedImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocksMap)
}