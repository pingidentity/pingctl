package sso_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_terraform"
)

func TestSSOTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)

	testutils_terraform.InitTerraform(t)

	testCases := []struct {
		name          string
		resource      connector.ExportableResource
		ignoredErrors []string
	}{
		{
			name:     "Application",
			resource: resources.Application(sdkClientInfo),
			ignoredErrors: []string{
				`Error: attribute "oidc_options": attribute "client_id" is required`,
			},
		},
		{
			name:          "ApplicationAttributeMapping",
			resource:      resources.ApplicationAttributeMapping(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationFlowPolicyAssignment",
			resource:      resources.ApplicationFlowPolicyAssignment(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationResourceGrant",
			resource:      resources.ApplicationResourceGrant(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationRoleAssignment",
			resource:      resources.ApplicationRoleAssignment(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationSecret",
			resource:      resources.ApplicationSecret(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ApplicationSignOnPolicyAssignment",
			resource:      resources.ApplicationSignOnPolicyAssignment(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Group",
			resource:      resources.Group(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "GroupNesting",
			resource:      resources.GroupNesting(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "GroupRoleAssignment",
			resource:      resources.GroupRoleAssignment(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "IdentityProvider",
			resource:      resources.IdentityProvider(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "IdentityProviderAttribute",
			resource:      resources.IdentityProviderAttribute(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PasswordPolicy",
			resource:      resources.PasswordPolicy(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Population",
			resource:      resources.Population(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PopulationDefault",
			resource:      resources.PopulationDefault(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "Resource",
			resource:      resources.Resource(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ResourceAttribute",
			resource:      resources.ResourceAttribute(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ResourceScope",
			resource:      resources.ResourceScope(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ResourceScopeOpenId",
			resource:      resources.ResourceScopeOpenId(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "ResourceScopePingOneApi",
			resource:      resources.ResourceScopePingOneApi(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "SchemaAttribute",
			resource: resources.SchemaAttribute(sdkClientInfo),
			ignoredErrors: []string{
				"Error: Data Loss Protection",
			},
		},
		{
			name:          "SignOnPolicy",
			resource:      resources.SignOnPolicy(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "SignOnPolicyAction",
			resource: resources.SignOnPolicyAction(sdkClientInfo),
			ignoredErrors: []string{
				"Error: Conflicting configuration arguments",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_terraform.ValidateTerraformPlan(t, tc.resource, tc.ignoredErrors)
		})
	}
}
