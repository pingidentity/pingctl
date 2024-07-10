package sso_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test --generate-config-out for the Application resource
func TestApplicationTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	applicationResource := resources.Application(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, applicationResource, nil)
}

// Test --generate-config-out for the ApplicationAttributeMapping resource
func TestApplicationAttributeMappingTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	applicationAttributeMappingResource := resources.ApplicationAttributeMapping(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, applicationAttributeMappingResource, nil)
}

// Test --generate-config-out for the ApplicationFlowPolicyAssignment resource
func TestApplicationFlowPolicyAssignmentTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	applicationFlowPolicyAssignmentResource := resources.ApplicationFlowPolicyAssignment(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, applicationFlowPolicyAssignmentResource, nil)
}

// Test --generate-config-out for the ApplicationResourceGrant resource
func TestApplicationResourceGrantTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	applicationResourceGrantResource := resources.ApplicationResourceGrant(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, applicationResourceGrantResource, nil)
}

// Test --generate-config-out for the ApplicationRoleAssignment resource
func TestApplicationRoleAssignmentTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	applicationRoleAssignmentResource := resources.ApplicationRoleAssignment(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, applicationRoleAssignmentResource, nil)
}

// Test --generate-config-out for the ApplicationSecret resource
func TestApplicationSecretTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	applicationSecretResource := resources.ApplicationSecret(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, applicationSecretResource, nil)
}

// Test --generate-config-out for the ApplicationSignOnPolicyAssignment resource
func TestApplicationSignOnPolicyAssignmentTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	applicationSignOnPolicyAssignmentResource := resources.ApplicationSignOnPolicyAssignment(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, applicationSignOnPolicyAssignmentResource, nil)
}

// Test --generate-config-out for the Group resource
func TestGroupTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	groupResource := resources.Group(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, groupResource, nil)
}

// Test --generate-config-out for the GroupNesting resource
func TestGroupNestingTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	groupNestingResource := resources.GroupNesting(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, groupNestingResource, nil)
}

// Test --generate-config-out for the GroupRoleAssignment resource
func TestGroupRoleAssignmentTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	groupRoleAssignmentResource := resources.GroupRoleAssignment(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, groupRoleAssignmentResource, nil)
}

// Test --generate-config-out for the IdentityProvider resource
func TestIdentityProviderTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	identityProviderResource := resources.IdentityProvider(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, identityProviderResource, nil)
}

// Test --generate-config-out for the IdentityProviderAttribute resource
func TestIdentityProviderAttributeTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	identityProviderAttributeResource := resources.IdentityProviderAttribute(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, identityProviderAttributeResource, nil)
}

// Test --generate-config-out for the PasswordPolicy resource
func TestPasswordPolicyTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	passwordPolicyResource := resources.PasswordPolicy(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, passwordPolicyResource, nil)
}

// Test --generate-config-out for the Population resource
func TestPopulationTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	populationResource := resources.Population(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, populationResource, nil)
}

// Test --generate-config-out for the PopulationDefault resource
func TestPopulationDefaultTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	populationDefaultResource := resources.PopulationDefault(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, populationDefaultResource, nil)
}

// Test --generate-config-out for the Resource resource
func TestResourceTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	resourceResource := resources.Resource(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, resourceResource, nil)
}

// Test --generate-config-out for the ResourceAttribute resource
func TestResourceAttributeTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	resourceAttributeResource := resources.ResourceAttribute(sdkClientInfo)

	testutils_helpers.ValidateTerraformPlan(t, resourceAttributeResource, nil)
}

// Test --generate-config-out for the ResourceScope resource
func TestResourceScopeTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	resourceScopeResource := resources.ResourceScope(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, resourceScopeResource, nil)
}

// Test --generate-config-out for the ResourceScopeOpenId resource
func TestResourceScopeOpenIdTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	resourceScopeOpenIdResource := resources.ResourceScopeOpenId(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, resourceScopeOpenIdResource, nil)
}

// Test --generate-config-out for the ResourceScopePingOneApi resource
func TestResourceScopePingOneApiTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	resourceScopePingOneApiResource := resources.ResourceScopePingOneApi(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, resourceScopePingOneApiResource, nil)
}

// Test --generate-config-out for the SchemaAttribute resource
func TestSchemaAttributeTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	schemaAttributeResource := resources.SchemaAttribute(sdkClientInfo)
	ignoreErrors := []string{
		"Error: Data Loss Protection",
	}

	testutils_helpers.ValidateTerraformPlan(t, schemaAttributeResource, ignoreErrors)
}

// Test --generate-config-out for the SignOnPolicy resource
func TestSignOnPolicyTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	signOnPolicyResource := resources.SignOnPolicy(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, signOnPolicyResource, nil)
}

// Test --generate-config-out for the SignOnPolicyAction resource
func TestSignOnPolicyActionTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	signOnPolicyActionResource := resources.SignOnPolicyAction(sdkClientInfo)
	ignoreErrors := []string{
		"Error: Conflicting configuration arguments",
	}
	testutils_helpers.ValidateTerraformPlan(t, signOnPolicyActionResource, ignoreErrors)
}
