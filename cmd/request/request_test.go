package request_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Request Command Executes without issue
func TestRequestCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "request",
		"--service", "pingone",
		"--http-method", "GET",
		"environments",
	)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Request Command fails when provided too many arguments
func TestRequestCmd_Execute_TooManyArguments(t *testing.T) {
	expectedErrorPattern := `accepts 1 arg\(s\), received 2`
	err := testutils_cobra.ExecutePingctl(t, "request", "arg1", "arg2")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Request Command fails when provided invalid flag
func TestRequestCmd_Execute_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `unknown flag: --invalid`
	err := testutils_cobra.ExecutePingctl(t, "request", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Request Command --help, -h flag
func TestRequestCmd_Execute_Help(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "request", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingctl(t, "request", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Request Command with Invalid Service
func TestRequestCmd_Execute_InvalidService(t *testing.T) {
	expectedErrorPattern := `^invalid argument ".*" for "-s, --service" flag: unrecognized Request Service: '.*'. Must be one of: .*$`
	err := testutils_cobra.ExecutePingctl(t, "request",
		"--service", "invalid-service",
		"--http-method", "GET",
		"environments",
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Request Command with Invalid HTTP Method
func TestRequestCmd_Execute_InvalidHTTPMethod(t *testing.T) {
	expectedErrorPattern := `^invalid argument ".*" for "-m, --http-method" flag: unrecognized HTTP Method: '.*'. Must be one of: .*$`
	err := testutils_cobra.ExecutePingctl(t, "request",
		"--service", "pingone",
		"--http-method", "INVALID",
		"environments",
	)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Request Command with Missing Required Service Flag
func TestRequestCmd_Execute_MissingRequiredServiceFlag(t *testing.T) {
	expectedErrorPattern := `failed to send custom request: service is required`
	err := testutils_cobra.ExecutePingctl(t, "request", "environments")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
