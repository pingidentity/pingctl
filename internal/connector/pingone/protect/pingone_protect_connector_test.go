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
			name:     "RiskPolicy",
			resource: resources.RiskPolicy(sdkClientInfo),
			ignoredErrors: []string{
				`Error: attribute "default_result": attribute "type" is required`,
				`Error: attribute "policy_scores": attribute "policy_threshold_medium": attribute "max_score" is required`,
				`Error: attribute "policy_scores": attribute "policy_threshold_high": attribute "max_score" is required`,
				`Error: attribute "policy_scores": attribute "predictors": incorrect set element type: attribute "predictor_reference_value" is required`,
			},
		},
		{
			name:     "RiskPredictor",
			resource: resources.RiskPredictor(sdkClientInfo),
			ignoredErrors: []string{
				`Error: attribute "predictor_velocity": attributes "by", "every", "fallback", "sliding_window", and "use" are required`,
				`Error: attribute "predictor_user_location_anomaly": attribute "days" is required`,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testutils_terraform.ValidateTerraformPlan(t, tc.resource, tc.ignoredErrors)
		})
	}
}
