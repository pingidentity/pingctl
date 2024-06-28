package protect_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector/pingone/protect/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test --generate-config-out for the RiskPolicy resource
func TestRiskPolicyTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	riskPolicyResource := resources.RiskPolicy(sdkClientInfo)

	// TODO - Remove this ignore error.
	// This test is failing due to a bug where computed values are failing
	// config generation as they are treated as required attributes.
	ignoreErrors := []string{
		`Error: attribute "default_result": attribute "type" is required`,
		`Error: attribute "policy_scores": attribute "policy_threshold_medium": attribute "max_score" is required`,
		`Error: attribute "policy_scores": attribute "policy_threshold_high": attribute "max_score" is required`,
		`Error: attribute "policy_scores": attribute "predictors": incorrect set element type: attribute "predictor_reference_value" is required`,
	}

	testutils_helpers.ValidateTerraformPlan(t, riskPolicyResource, ignoreErrors)
}

// Test --generate-config-out for the RiskPredictor resource
func TestRiskPredictorTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	riskPredictorResource := resources.RiskPredictor(sdkClientInfo)

	// TODO - Remove this ignore error.
	// This test is failing due to a bug where computed values are failing
	// config generation as they are treated as required attributes.
	ignoreErrors := []string{
		`Error: attribute "predictor_velocity": attributes "by", "every", "fallback", "sliding_window", and "use" are required`,
		`Error: attribute "predictor_user_location_anomaly": attribute "days" is required`,
	}

	testutils_helpers.ValidateTerraformPlan(t, riskPredictorResource, ignoreErrors)
}
