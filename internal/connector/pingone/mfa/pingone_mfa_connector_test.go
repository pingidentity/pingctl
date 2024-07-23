package mfa_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/mfa/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_terraform"
)

func TestMFATerraformPlan(t *testing.T) {
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)

	testutils_terraform.InitTerraform(t)

	testCases := []struct {
		name          string
		resource      connector.ExportableResource
		ignoredErrors []string
	}{
		{
			name:     "MFAApplicationPushCredential",
			resource: resources.MFAApplicationPushCredential(sdkClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Combination",
			},
		},
		{
			name:          "MFAFido2Policy",
			resource:      resources.MFAFido2Policy(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "MFADevicePolicy",
			resource:      resources.MFADevicePolicy(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "MFASettings",
			resource:      resources.MFASettings(sdkClientInfo),
			ignoredErrors: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_terraform.ValidateTerraformPlan(t, tc.resource, tc.ignoredErrors)
		})
	}
}
