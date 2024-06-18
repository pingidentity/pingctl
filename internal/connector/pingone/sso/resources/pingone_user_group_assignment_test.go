package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

func TestUserGroupAssignmentExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	resource := resources.UserGroupAssignment(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_user_group_assignment",
			ResourceName: "test_user2_testing",
			ResourceID:   fmt.Sprintf("%s/7570f416-5503-4eb0-9fc3-9a485bddd411/b6924f30-73ca-4d3c-964b-90c77adce6a7", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_user_group_assignment",
			ResourceName: "testing_testing",
			ResourceID:   fmt.Sprintf("%s/68cb3634-0ed4-4044-85d1-576eb3a55360/b6924f30-73ca-4d3c-964b-90c77adce6a7", testutils_helpers.GetEnvironmentID()),
		},
	}

	testutils_helpers.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
