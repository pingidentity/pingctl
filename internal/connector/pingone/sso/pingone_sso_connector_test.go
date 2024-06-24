package sso_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector/pingone/sso/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test --generate-config-out for the application resource
func TestSSOConnectorTerraformPlanApplicationResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the application resource
	applicationResource := resources.Application(sdkClientInfo)

	// Run terraform plan --generate-config-out on the application resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, applicationResource, sdkClientInfo)
}

// Test --generate-config-out for the application attribute mapping resource
func TestSSOConnectorTerraformPlanApplicationAttributeMappingResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the application attribute mapping resource
	applicationAttributeMappingResource := resources.ApplicationAttributeMapping(sdkClientInfo)

	// Run terraform plan --generate-config-out on the application attribute mapping resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, applicationAttributeMappingResource, sdkClientInfo)
}

// Test --generate-config-out for the application flow policy assignment resource
func TestSSOConnectorTerraformPlanApplicationFlowPolicyAssignmentResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the application flow policy assignment resource
	applicationFlowPolicyAssignmentResource := resources.ApplicationFlowPolicyAssignment(sdkClientInfo)

	// Run terraform plan --generate-config-out on the application flow policy assignment resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, applicationFlowPolicyAssignmentResource, sdkClientInfo)
}

// Test --generate-config-out for the application resource grant resource
func TestSSOConnectorTerraformPlanApplicationResourceGrantResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the application resource grant resource
	applicationResourceGrantResource := resources.ApplicationResourceGrant(sdkClientInfo)

	// Run terraform plan --generate-config-out on the application resource grant resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, applicationResourceGrantResource, sdkClientInfo)
}

// Test --generate-config-out for the application role assignment resource
func TestSSOConnectorTerraformPlanApplicationRoleAssignmentResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the application role assignment resource
	applicationRoleAssignmentResource := resources.ApplicationRoleAssignment(sdkClientInfo)

	// Run terraform plan --generate-config-out on the application role assignment resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, applicationRoleAssignmentResource, sdkClientInfo)
}

// Test --generate-config-out for the application secret resource
func TestSSOConnectorTerraformPlanApplicationSecretResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the application secret resource
	applicationSecretResource := resources.ApplicationSecret(sdkClientInfo)

	// Run terraform plan --generate-config-out on the application secret resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, applicationSecretResource, sdkClientInfo)
}

// Test --generate-config-out for the application sign-on policy assignment resource
func TestSSOConnectorTerraformPlanApplicationSignOnPolicyAssignmentResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the application sign-on policy assignment resource
	applicationSignOnPolicyAssignmentResource := resources.ApplicationSignOnPolicyAssignment(sdkClientInfo)

	// Run terraform plan --generate-config-out on the application sign-on policy assignment resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, applicationSignOnPolicyAssignmentResource, sdkClientInfo)
}

// Test --generate-config-out for the group resource
func TestSSOConnectorTerraformPlanGroupResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the group resource
	groupResource := resources.Group(sdkClientInfo)

	// Run terraform plan --generate-config-out on the group resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, groupResource, sdkClientInfo)
}

// Test --generate-config-out for the group nesting resource
func TestSSOConnectorTerraformPlanGroupNestingResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the group nesting resource
	groupNestingResource := resources.GroupNesting(sdkClientInfo)

	// Run terraform plan --generate-config-out on the group nesting resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, groupNestingResource, sdkClientInfo)
}

// Test --generate-config-out for the group role assignment resource
func TestSSOConnectorTerraformPlanGroupRoleAssignmentResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the group role assignment resource
	groupRoleAssignmentResource := resources.GroupRoleAssignment(sdkClientInfo)

	// Run terraform plan --generate-config-out on the group role assignment resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, groupRoleAssignmentResource, sdkClientInfo)
}

// Test --generate-config-out for the identity provider resource
func TestSSOConnectorTerraformPlanIdentityProviderResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the identity provider resource
	identityProviderResource := resources.IdentityProvider(sdkClientInfo)

	// Run terraform plan --generate-config-out on the identity provider resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, identityProviderResource, sdkClientInfo)
}

// Test --generate-config-out for the identity provider attribute resource
func TestSSOConnectorTerraformPlanIdentityProviderAttributeResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the identity provider attribute resource
	identityProviderAttributeResource := resources.IdentityProviderAttribute(sdkClientInfo)

	// Run terraform plan --generate-config-out on the identity provider attribute resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, identityProviderAttributeResource, sdkClientInfo)
}

// Test --generate-config-out for the password policy resource
func TestSSOConnectorTerraformPlanPasswordPolicyResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the password policy resource
	passwordPolicyResource := resources.PasswordPolicy(sdkClientInfo)

	// Run terraform plan --generate-config-out on the password policy resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, passwordPolicyResource, sdkClientInfo)
}

// Test --generate-config-out for the population resource
func TestSSOConnectorTerraformPlanPopulationResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the population resource
	populationResource := resources.Population(sdkClientInfo)

	// Run terraform plan --generate-config-out on the population resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, populationResource, sdkClientInfo)
}

// Test --generate-config-out for the population default resource
func TestSSOConnectorTerraformPlanPopulationDefaultResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the population default resource
	populationDefaultResource := resources.PopulationDefault(sdkClientInfo)

	// Run terraform plan --generate-config-out on the population default resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, populationDefaultResource, sdkClientInfo)
}

// Test --generate-config-out for the resource resource
func TestSSOConnectorTerraformPlanResourceResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the resource resource
	resourceResource := resources.Resource(sdkClientInfo)

	// Run terraform plan --generate-config-out on the resource resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, resourceResource, sdkClientInfo)
}

// Test --generate-config-out for the resource attribute resource
func TestSSOConnectorTerraformPlanResourceAttributeResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the resource attribute resource
	resourceAttributeResource := resources.ResourceAttribute(sdkClientInfo)

	// Run terraform plan --generate-config-out on the resource attribute resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, resourceAttributeResource, sdkClientInfo)
}

// Test --generate-config-out for the resource scope resource
func TestSSOConnectorTerraformPlanResourceScopeResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the resource scope resource
	resourceScopeResource := resources.ResourceScope(sdkClientInfo)

	// Run terraform plan --generate-config-out on the resource scope resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, resourceScopeResource, sdkClientInfo)
}

// Test --generate-config-out for the resource scope OpenID resource
func TestSSOConnectorTerraformPlanResourceScopeOpenIDResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the resource scope OpenID resource
	resourceScopeOpenIDResource := resources.ResourceScopeOpenId(sdkClientInfo)

	// Run terraform plan --generate-config-out on the resource scope OpenID resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, resourceScopeOpenIDResource, sdkClientInfo)
}

// Test --generate-config-out for the resource scope PingOne API resource

func TestSSOConnectorTerraformPlanResourceScopePingOneAPIResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the resource scope PingOne API resource
	resourceScopePingOneAPIResource := resources.ResourceScopePingOneApi(sdkClientInfo)

	// Run terraform plan --generate-config-out on the resource scope PingOne API resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, resourceScopePingOneAPIResource, sdkClientInfo)
}

// Test --generate-config-out for the schema attribute resource
func TestSSOConnectorTerraformPlanSchemaAttributeResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the schema attribute resource
	schemaAttributeResource := resources.SchemaAttribute(sdkClientInfo)

	// Run terraform plan --generate-config-out on the schema attribute resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, schemaAttributeResource, sdkClientInfo)
}

// Test --generate-config-out for the sign-on policy resource
func TestSSOConnectorTerraformPlanSignOnPolicyResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the sign-on policy resource
	signOnPolicyResource := resources.SignOnPolicy(sdkClientInfo)

	// Run terraform plan --generate-config-out on the sign-on policy resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, signOnPolicyResource, sdkClientInfo)
}

// Test --generate-config-out for the sign-on policy action resource
func TestSSOConnectorTerraformPlanSignOnPolicyActionResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the sign-on policy action resource
	signOnPolicyActionResource := resources.SignOnPolicyAction(sdkClientInfo)

	// Run terraform plan --generate-config-out on the sign-on policy action resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, signOnPolicyActionResource, sdkClientInfo)
}

// Test --generate-config-out for the user resource
func TestSSOConnectorTerraformPlanUserResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the user resource
	userResource := resources.User(sdkClientInfo)

	// Run terraform plan --generate-config-out on the user resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, userResource, sdkClientInfo)
}

// Test --generate-config-out for the user group assignment resource
func TestSSOConnectorTerraformPlanUserGroupAssignmentResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the user group assignment resource
	userGroupAssignmentResource := resources.UserGroupAssignment(sdkClientInfo)

	// Run terraform plan --generate-config-out on the user group assignment resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, userGroupAssignmentResource, sdkClientInfo)
}
