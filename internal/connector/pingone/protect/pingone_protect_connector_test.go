package protect_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/protect/resources"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_terraform"
)

func TestProtectTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)

	testutils_terraform.InitTerraform(t)

	testCases := []struct {
		name          string
		resource      connector.ExportableResource
		ignoredErrors []string
	}{
		{
			name:          "RiskPolicy",
			resource:      resources.RiskPolicy(sdkClientInfo),
			ignoredErrors: nil,
		},
		{
			name:          "RiskPredictor",
			resource:      resources.RiskPredictor(sdkClientInfo),
			ignoredErrors: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_terraform.ValidateTerraformPlan(t, tc.resource, tc.ignoredErrors)
		})
	}
}
