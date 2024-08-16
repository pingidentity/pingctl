package profile_internal

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test RunInternalConfigProfileDescribe function with valid arg
func Test_RunInternalConfigProfileDescribe_ValidArg(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigProfileDescribe("production")
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigProfileDescribe function with invalid profile name
func Test_RunInternalConfigProfileDescribe_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to describe profile: invalid profile name: '.*'. name must contain only alphanumeric characters, underscores, and dashes$"
	err := RunInternalConfigProfileDescribe("invalid&*^*&")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileDescribe function with non-existent profile
func Test_RunInternalConfigProfileDescribe_NonExistentProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to describe profile: invalid profile name: '.*' profile does not exist$"
	err := RunInternalConfigProfileDescribe("invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileDescribe function with active profile
func Test_RunInternalConfigProfileDescribe_ActiveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigProfileDescribe("default")
	testutils.CheckExpectedError(t, err, nil)
}

// Check output of RunInternalConfigProfileDescribe function
func Example_runInternalConfigProfileDescribe() {
	testutils_viper.InitVipers(&testing.T{})

	_ = RunInternalConfigProfileDescribe("production")

	// Output:
	// Profile Name: production
	// Description: test profile description
	//
	// Set Options:
	//  - pingctl.outputFormat: text
	//  - pingctl.color: true
	//  - pingfederate.clientCredentialsAuth.scopes: []
	//  - pingfederate.xBypassExternalValidationHeader: false
	//  - pingfederate.caCertificatePemFiles: []
	//  - pingfederate.insecureTrustAllTLS: false
	//
	// Unset Options:
	//  - pingone.export.environmentID
	//  - pingone.worker.environmentID
	//  - pingone.worker.clientID
	//  - pingone.worker.clientSecret
	//  - pingone.region
	//  - pingfederate.basicAuth.username
	//  - pingfederate.basicAuth.password
	//  - pingfederate.httpsHost
	//  - pingfederate.adminApiPath
	//  - pingfederate.clientCredentialsAuth.clientID
	//  - pingfederate.clientCredentialsAuth.clientSecret
	//  - pingfederate.clientCredentialsAuth.tokenURL
	//  - pingfederate.accessTokenAuth.accessToken
}
