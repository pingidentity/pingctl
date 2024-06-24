package protect_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/connector/pingone/protect/resources"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test --generate-config-out for the RiskPolicy resource
func TestProtectConnectorTerraformPlanRiskPolicyResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the RiskPolicy resource
	riskPolicyResource := resources.RiskPolicy(sdkClientInfo)

	// Run terraform plan --generate-config-out on the RiskPolicy resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, riskPolicyResource, sdkClientInfo)
}

// Test --generate-config-out for the RiskPredictor resource
func TestProtectConnectorTerraformPlanRiskPredictorResource(t *testing.T) {
	// Get an instance of the PingOne SDK Client
	sdkClientInfo := testutils_helpers.GetPingOneSDKClientInfo(t)

	// Create an instance of the RiskPredictor resource
	riskPredictorResource := resources.RiskPredictor(sdkClientInfo)

	// Run terraform plan --generate-config-out on the RiskPredictor resource
	testutils_helpers.TestSingleResourceTerraformPlanGenerateConfigOut(t, riskPredictorResource, sdkClientInfo)
}
