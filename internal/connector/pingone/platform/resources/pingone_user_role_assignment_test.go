package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

func TestRoleAssignmentUserExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	resource := resources.UserRoleAssignment(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_user_role_assignment",
			ResourceName: "test_user_Identity Data Admin_POPULATION_1",
			ResourceID:   fmt.Sprintf("%s/8012b4a3-1874-42d3-803d-76d4a378335c/37cb92f1-858f-4cce-afa9-4d1bbd2a6b4e", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_user_role_assignment",
			ResourceName: "test_user_Identity Data Admin_POPULATION_2",
			ResourceID:   fmt.Sprintf("%s/8012b4a3-1874-42d3-803d-76d4a378335c/ed3ecdd9-973c-4cab-b6d1-e5474e25309e", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_user_role_assignment",
			ResourceName: "test_user_Identity Data Admin_POPULATION_3",
			ResourceID:   fmt.Sprintf("%s/8012b4a3-1874-42d3-803d-76d4a378335c/2c2a1bb4-9d07-4c29-86ab-143b2cbb3bb6", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_user_role_assignment",
			ResourceName: "test_user2_Identity Data Read Only_ENVIRONMENT_1",
			ResourceID:   fmt.Sprintf("%s/7570f416-5503-4eb0-9fc3-9a485bddd411/71219163-1a98-45ed-827c-d4ad6e3ff1d4", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_user_role_assignment",
			ResourceName: "testing_Identity Data Read Only_ENVIRONMENT_1",
			ResourceID:   fmt.Sprintf("%s/68cb3634-0ed4-4044-85d1-576eb3a55360/224da3de-c879-478e-b219-6255feb78a59", testutils_helpers.GetEnvironmentID()),
		},
	}

	testutils_helpers.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
