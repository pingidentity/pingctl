package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

func TestNotificationPolicyExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	resource := resources.NotificationPolicy(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_notification_policy",
			ResourceName: "Test",
			ResourceID:   fmt.Sprintf("%s/32cc413d-0ec8-4be9-823c-a9e06f5a5830", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_policy",
			ResourceName: "Default Notification Policy",
			ResourceID:   fmt.Sprintf("%s/54606af4-72a6-4b38-bfb8-75034097af9a", testutils_helpers.GetEnvironmentID()),
		},
	}

	testutils_helpers.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
