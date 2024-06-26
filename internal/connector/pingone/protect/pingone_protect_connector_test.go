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
	testutils_helpers.ValidateTerraformPlan(t, riskPolicyResource, nil)
}

// Test --generate-config-out for the RiskPredictor resource
func TestRiskPredictorTerraformPlan(t *testing.T) {
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)
	riskPredictorResource := resources.RiskPredictor(sdkClientInfo)
	testutils_helpers.ValidateTerraformPlan(t, riskPredictorResource, nil)
}
