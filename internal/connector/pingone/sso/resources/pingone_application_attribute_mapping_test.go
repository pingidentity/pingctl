package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestApplicationAttributeMappingExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.ApplicationAttributeMapping(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_application_attribute_mapping",
			ResourceName: "Example OAuth App_sub",
			ResourceID:   fmt.Sprintf("%s/2a7c1b5d-415b-4fb5-a6c0-1e290f776785/f6d41400-e571-432e-9151-4ff06e0b51ce", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_attribute_mapping",
			ResourceName: "Getting Started Application_sub",
			ResourceID:   fmt.Sprintf("%s/3da7aae6-92e5-4295-a37c-8515d1f2cd86/f6d41400-e571-432e-9151-4ff06e0b51ce", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_attribute_mapping",
			ResourceName: "OAuth Worker App_sub",
			ResourceID:   fmt.Sprintf("%s/9d6c443b-6329-4d3c-949e-880eda3b9599/f6d41400-e571-432e-9151-4ff06e0b51ce", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_attribute_mapping",
			ResourceName: "PingOne DaVinci Connection_sub",
			ResourceID:   fmt.Sprintf("%s/7b621870-7124-4426-b432-6c675642afcb/f6d41400-e571-432e-9151-4ff06e0b51ce", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_attribute_mapping",
			ResourceName: "test app_sub",
			ResourceID:   fmt.Sprintf("%s/a4cbf57e-fa2c-452f-bbc8-f40b551da0e2/f6d41400-e571-432e-9151-4ff06e0b51ce", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_application_attribute_mapping",
			ResourceName: "Worker App_sub",
			ResourceID:   fmt.Sprintf("%s/c45c2f8c-dee0-4a12-b169-bae693a13d57/f6d41400-e571-432e-9151-4ff06e0b51ce", testutils.GetEnvironmentID()),
		},
	}

	expectedImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range expectedImportBlocks {
		expectedImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocksMap)
}