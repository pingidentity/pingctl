package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestGatewayExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.Gateway(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_gateway",
			ResourceName: "random",
			ResourceID:   fmt.Sprintf("%s/0b1d882c-5c71-4600-a9fb-befdad921df2", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway",
			ResourceName: "PingFederate LDAP Gateway",
			ResourceID:   fmt.Sprintf("%s/3b7b5d9d-1820-4b21-bb29-a5336af65352", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway",
			ResourceName: "PF TF Provider",
			ResourceID:   fmt.Sprintf("%s/554257ac-76ca-447a-a210-722343328312", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway",
			ResourceName: "Local Test",
			ResourceID:   fmt.Sprintf("%s/5cd3f6b7-35f0-4873-ac64-f32118bf3102", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_gateway",
			ResourceName: "TestGateway",
			ResourceID:   fmt.Sprintf("%s/bc37814f-b3a9-4149-b880-0ed457bbb5c5", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
