package pingfederate_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingfederate/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_terraform"
)

func TestPingFederateTerraformPlan(t *testing.T) {
	PingFederateClientInfo := testutils.GetPingFederateClientInfo(t)

	testutils_terraform.InitPingFederateTerraform(t)

	testCases := []struct {
		name          string
		resource      connector.ExportableResource
		ignoredErrors []string
	}{
		{
			name:     "PingFederateAuthenticationApiApplication",
			resource: resources.AuthenticationApiApplication(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Invalid Attribute Value", // TODO - Remove with PDI-1925 fix
			},
		},
		{
			name:          "PingFederateAuthenticationApiSettings",
			resource:      resources.AuthenticationApiSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateAuthenticationPolicies",
			resource: resources.AuthenticationPolicies(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Plugin did not respond",
				"Error: Request cancelled",
			},
		},
		{
			name:     "PingFederateAuthenticationPoliciesFragment",
			resource: resources.AuthenticationPoliciesFragment(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Reference to undeclared resource",
			},
		},
		{
			name:          "PingFederateAuthenticationPoliciesSettings",
			resource:      resources.AuthenticationPoliciesSettings(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "PingFederateAuthenticationPolicyContract",
			resource:      resources.AuthenticationPolicyContract(PingFederateClientInfo),
			ignoredErrors: nil,
		},
		{
			name:     "PingFederateAuthenticationSelector",
			resource: resources.AuthenticationSelector(PingFederateClientInfo),
			ignoredErrors: []string{
				"Error: Plugin did not respond",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_terraform.ValidateTerraformPlan(t, tc.resource, tc.ignoredErrors)
		})
	}
}
