package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestResourceScopeExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.ResourceScope(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "test_testing",
			ResourceID:   fmt.Sprintf("%s/4b9ef858-62ce-4bd0-9186-997b8527529d/99bda6e7-f34b-4218-8fb0-221f5414e0db", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "test_oidc",
			ResourceID:   fmt.Sprintf("%s/4b9ef858-62ce-4bd0-9186-997b8527529d/9f2c9b87-a190-446e-bf6b-d97b7f8b1a70", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "testing_test",
			ResourceID:   fmt.Sprintf("%s/52afd89f-f3c0-4c78-b896-432c0a07329b/d9935d01-5baa-4843-970a-9df33b60439f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:read:user",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/089adcde-be64-4e7e-9a5a-dda60ce38a9f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:read:user:2",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/83d8ee1d-938f-4287-9792-aa808dc0cad9", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:update:user",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/d5bd66de-8044-41c5-aed2-278b6cf47dad", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:update:userMfaEnabled",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/2a8c4a72-b7fb-4093-9e9f-4b2d6040749a", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:validate:userPassword",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/a550ce20-1b72-4bfe-a756-20f21647d1cf", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:reset:userPassword",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/a112f3ef-cfa2-4cde-89d6-965ca88096c5", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:read:userPassword",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/ac9a1304-419b-4786-9113-1b5c5df9db11", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:read:device",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/194a59ec-fae0-462f-ad73-a3fbbdf2430e", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:update:device",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/cb98b74e-ea38-4ae1-adcd-e141d5cf7a2a", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:create:device",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/23be7b0b-044c-480f-9d77-2ef3844e49b0", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:delete:device",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/1291c71c-ae3f-4f3b-b307-51c5636a6814", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:read:userLinkedAccounts",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/5f99b914-f610-4bc6-b4cb-29e439d94b7e", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:delete:userLinkedAccounts",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/cc9a5dad-9ca9-45c4-9b48-2a7ff78802d5", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:create:pairingKey",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/896118f4-f27c-4603-b348-70a602786d8a", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:delete:pairingKey",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/c417b04a-34ec-417d-a52d-d871abc5f411", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:read:pairingKey",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/78a51481-f1e9-426c-a6de-71a05fc7148a", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:read:sessions",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/4be047cf-70b8-4d0b-84c9-33d0fc9c15e8", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:delete:sessions",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/d06d54c1-e54f-4047-97a0-c7d42fa75dbe", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:read:userConsent",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/c5ffe0bf-8936-415b-a40d-a7c3818a3864", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:verify:user",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/7499ab71-a894-488c-9593-de96f74dd0df", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:read:oauthConsent",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/90abdaa3-fa8c-4034-aa36-42d1812c8df7", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "PingOne API_p1:update:oauthConsent",
			ResourceID:   fmt.Sprintf("%s/95ed3610-7668-4a17-8334-b3db5ff9a875/28281333-293e-4018-82f4-3295db0e0bf8", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "openid_profile",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/5a2881ba-affc-4556-a9ff-ad662ea84e89", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "openid_newscope2",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/5f07b021-5f0e-47d0-a62b-1e983bdff753", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "openid_openid",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/6f095311-2cb9-4414-b30f-af8ee5e11e34", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "openid_newscope",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/792fa804-8aae-43c8-bea7-ea2dbbb1ca88", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "openid_email",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/a95eb903-b691-4aa9-91df-8b02d69816df", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "openid_test",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/d4213f0d-e1fc-42db-bcc6-dfad730f7be7", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "openid_phone",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/dad64f0c-187e-4991-a5b3-c4e53a4167e5", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "openid_testing",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/eb7e9feb-6076-4a2e-9e9e-5c9c0a503606", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_resource_scope",
			ResourceName: "openid_address",
			ResourceID:   fmt.Sprintf("%s/8c428665-3e68-4f3c-997d-16a97f8cbe80/fcd04665-fb97-4943-9c88-427331ebe930", testutils.GetEnvironmentID()),
		},
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocks)
}
