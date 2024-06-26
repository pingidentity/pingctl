package mfa_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector/pingone/mfa/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test --generate-config-out for the MFAApplicationPushCredential resource
func TestMFAApplicationPushCredentialTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	mfaApplicationPushCredentialResource := resources.MFAApplicationPushCredential(sdkClientInfo)
	ignoreErrors := []string{
		"Error: Invalid Attribute Combination",
	}
	testutils_helpers.ValidateTerraformPlan(t, mfaApplicationPushCredentialResource, ignoreErrors)
}

// Test --generate-config-out for the MFAFido2Policy resource
func TestMFAFido2PolicyTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	mfaFido2PolicyResource := resources.MFAFido2Policy(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, mfaFido2PolicyResource, nil)
}

// Test --generate-config-out for the MFADevicePolicy resource
func TestMFADevicePolicyTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	mfaDevicePolicyResource := resources.MFADevicePolicy(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, mfaDevicePolicyResource, nil)
}

// Test --generate-config-out for the MFASettings resource
func TestMFASettingsTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	mfaSettingsResource := resources.MFASettings(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, mfaSettingsResource, nil)
}
