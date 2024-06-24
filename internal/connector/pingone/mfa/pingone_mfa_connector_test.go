package mfa_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector/pingone/mfa/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test --generate-config-out for the MFAApplicationPushCredential resource
func TestMFAConnectorTerraformPlanMFAApplicationPushCredentialResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the MFAApplicationPushCredential resource
	mfaApplicationPushCredentialResource := resources.MFAApplicationPushCredential(sdkClientInfo)

	// Run terraform plan --generate-config-out on the MFAApplicationPushCredential resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, mfaApplicationPushCredentialResource, sdkClientInfo)
}

// Test --generate-config-out for the MFAFido2Policy resource
func TestMFAConnectorTerraformPlanMFAFido2PolicyResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the MFAFido2Policy resource
	mfaFido2PolicyResource := resources.MFAFido2Policy(sdkClientInfo)

	// Run terraform plan --generate-config-out on the MFAFido2Policy resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, mfaFido2PolicyResource, sdkClientInfo)
}

// Test --generate-config-out for the MFAPolicy resource
func TestMFAConnectorTerraformPlanMFAPolicyResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the MFAPolicy resource
	mfaPolicyResource := resources.MFAPolicy(sdkClientInfo)

	// Run terraform plan --generate-config-out on the MFAPolicy resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, mfaPolicyResource, sdkClientInfo)
}

// Test --generate-config-out for the MFASettings resource
func TestMFAConnectorTerraformPlanMFASettingsResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the MFASettings resource
	mfaSettingsResource := resources.MFASettings(sdkClientInfo)

	// Run terraform plan --generate-config-out on the MFASettings resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, mfaSettingsResource, sdkClientInfo)
}
