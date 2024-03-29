package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestIdentityProviderAttributeExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.IdentityProviderAttribute(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_identity_provider_attribute",
			ResourceName: "Test IdP_username",
			ResourceID:   fmt.Sprintf("%s/a99df558-7090-4303-8f35-860ac660e371/51a036c6-41ed-44f7-bd1d-eacaa2a1feab", testutils.GetEnvironmentID()),
		},
	}

	expectedImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range expectedImportBlocks {
		expectedImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocksMap)
}
