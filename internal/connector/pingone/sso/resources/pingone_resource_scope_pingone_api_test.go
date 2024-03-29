package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestResourceScopePingOneApiExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.ResourceScopePingOneApi(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:read:user",
			ResourceID:   fmt.Sprintf("%s/089adcde-be64-4e7e-9a5a-dda60ce38a9f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:read:user:2",
			ResourceID:   fmt.Sprintf("%s/83d8ee1d-938f-4287-9792-aa808dc0cad9", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:update:user",
			ResourceID:   fmt.Sprintf("%s/d5bd66de-8044-41c5-aed2-278b6cf47dad", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:update:userMfaEnabled",
			ResourceID:   fmt.Sprintf("%s/2a8c4a72-b7fb-4093-9e9f-4b2d6040749a", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:validate:userPassword",
			ResourceID:   fmt.Sprintf("%s/a550ce20-1b72-4bfe-a756-20f21647d1cf", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:reset:userPassword",
			ResourceID:   fmt.Sprintf("%s/a112f3ef-cfa2-4cde-89d6-965ca88096c5", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:read:userPassword",
			ResourceID:   fmt.Sprintf("%s/ac9a1304-419b-4786-9113-1b5c5df9db11", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:read:device",
			ResourceID:   fmt.Sprintf("%s/194a59ec-fae0-462f-ad73-a3fbbdf2430e", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:update:device",
			ResourceID:   fmt.Sprintf("%s/cb98b74e-ea38-4ae1-adcd-e141d5cf7a2a", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:create:device",
			ResourceID:   fmt.Sprintf("%s/23be7b0b-044c-480f-9d77-2ef3844e49b0", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:delete:device",
			ResourceID:   fmt.Sprintf("%s/1291c71c-ae3f-4f3b-b307-51c5636a6814", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:read:userLinkedAccounts",
			ResourceID:   fmt.Sprintf("%s/5f99b914-f610-4bc6-b4cb-29e439d94b7e", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:delete:userLinkedAccounts",
			ResourceID:   fmt.Sprintf("%s/cc9a5dad-9ca9-45c4-9b48-2a7ff78802d5", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:create:pairingKey",
			ResourceID:   fmt.Sprintf("%s/896118f4-f27c-4603-b348-70a602786d8a", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:delete:pairingKey",
			ResourceID:   fmt.Sprintf("%s/c417b04a-34ec-417d-a52d-d871abc5f411", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:read:pairingKey",
			ResourceID:   fmt.Sprintf("%s/78a51481-f1e9-426c-a6de-71a05fc7148a", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:read:sessions",
			ResourceID:   fmt.Sprintf("%s/4be047cf-70b8-4d0b-84c9-33d0fc9c15e8", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:delete:sessions",
			ResourceID:   fmt.Sprintf("%s/d06d54c1-e54f-4047-97a0-c7d42fa75dbe", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:read:userConsent",
			ResourceID:   fmt.Sprintf("%s/c5ffe0bf-8936-415b-a40d-a7c3818a3864", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:verify:user",
			ResourceID:   fmt.Sprintf("%s/7499ab71-a894-488c-9593-de96f74dd0df", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:read:oauthConsent",
			ResourceID:   fmt.Sprintf("%s/90abdaa3-fa8c-4034-aa36-42d1812c8df7", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope_pingone_api",
			ResourceName: "PingOne API_p1:update:oauthConsent",
			ResourceID:   fmt.Sprintf("%s/28281333-293e-4018-82f4-3295db0e0bf8", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
