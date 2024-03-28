package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestCertificateExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.Certificate(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_certificate",
			ResourceName: "common name",
			ResourceID:   fmt.Sprintf("%s/b9eb2b6e-381e-4b1c-86d3-096d951787f4", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_certificate",
			ResourceName: "terraform",
			ResourceID:   fmt.Sprintf("%s/fa8f15d6-1c62-4db1-920e-d22f6dd68ba8", testutils.GetEnvironmentID()),
		},
	}

	expectedImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range expectedImportBlocks {
		expectedImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocksMap)
}