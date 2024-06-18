package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

func TestBrandingThemeExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	resource := resources.BrandingTheme(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_branding_theme",
			ResourceName: "test_slate_2",
			ResourceID:   fmt.Sprintf("%s/a3e0fc98-a7bf-4750-9778-2397fc0a3586", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_branding_theme",
			ResourceName: "Ping Default",
			ResourceID:   fmt.Sprintf("%s/b02d49a3-c468-462a-9fd0-659e0f3dde96", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_branding_theme",
			ResourceName: "Slate",
			ResourceID:   fmt.Sprintf("%s/fbf0886a-fb1f-41c2-ad42-e7dc601dabb3", testutils_helpers.GetEnvironmentID()),
		},
	}

	testutils_helpers.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
