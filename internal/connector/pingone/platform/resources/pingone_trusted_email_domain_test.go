package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

func TestTrustedEmailDomainExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	resource := resources.TrustedEmailDomain(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_trusted_email_domain",
			ResourceName: "test.customdomain.com",
			ResourceID:   fmt.Sprintf("%s/47efb375-e9e8-40dc-b1ce-8598bf7b4aea", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_trusted_email_domain",
			ResourceName: "test.pingidentity.com",
			ResourceID:   fmt.Sprintf("%s/ff26c5c9-2e87-46d4-9cb0-077d162c7bcb", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_trusted_email_domain",
			ResourceName: "demo.bxretail.org",
			ResourceID:   fmt.Sprintf("%s/49f94864-f9c7-4778-ae37-839c2c546d1c", testutils_helpers.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_trusted_email_domain",
			ResourceName: "pioneerpalaceband.com",
			ResourceID:   fmt.Sprintf("%s/63d645d1-046a-4d53-a267-513cfc1d4213", testutils_helpers.GetEnvironmentID()),
		},
	}

	testutils_helpers.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
