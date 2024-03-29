package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestApplicationSecretExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.ApplicationSecret(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_application_secret",
			ResourceName: "Example OAuth App_secret",
			ResourceID:   fmt.Sprintf("%s/2a7c1b5d-415b-4fb5-a6c0-1e290f776785", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_secret",
			ResourceName: "Getting Started Application_secret",
			ResourceID:   fmt.Sprintf("%s/3da7aae6-92e5-4295-a37c-8515d1f2cd86", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_secret",
			ResourceName: "OAuth Worker App_secret",
			ResourceID:   fmt.Sprintf("%s/9d6c443b-6329-4d3c-949e-880eda3b9599", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_secret",
			ResourceName: "PingOne DaVinci Connection_secret",
			ResourceID:   fmt.Sprintf("%s/7b621870-7124-4426-b432-6c675642afcb", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_secret",
			ResourceName: "test app_secret",
			ResourceID:   fmt.Sprintf("%s/a4cbf57e-fa2c-452f-bbc8-f40b551da0e2", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_secret",
			ResourceName: "Worker App_secret",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57", testutils.GetEnvironmentID()),
		},
	}

	expectedImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range expectedImportBlocks {
		expectedImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocksMap)
}
