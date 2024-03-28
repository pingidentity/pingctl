package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestSystemApplicationExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.SystemApplication(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_system_application",
			ResourceName: "PingOne Application Portal",
			ResourceID:   fmt.Sprintf("%s/92a3765c-e135-4afa-8b12-4469672ac8a9", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_system_application",
			ResourceName: "PingOne Self-Service - MyAccount",
			ResourceID:   fmt.Sprintf("%s/4ce54d01-5138-4c56-8175-4f02f69278f5", testutils.GetEnvironmentID()),
		},
	}

	expectedImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range expectedImportBlocks {
		expectedImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocksMap)
}
