package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

func TestLanguageExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	resource := resources.Language(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_language",
			ResourceName: "Afar",
			ResourceID:   fmt.Sprintf("%s/c5f97303-c5c1-459a-b8ee-3fcb183eb52a", testutils_helpers.GetEnvironmentID()),
		},
	}

	testutils_helpers.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
