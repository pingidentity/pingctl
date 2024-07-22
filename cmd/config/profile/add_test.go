package profile_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Config Profile Add Command Executes without issue
func TestConfigProfileAddCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "add", "--name", "new-test-profile-name", "--set-active", "--description", "test-description")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Profile Add Command fails when provided an invalid profile name
func TestConfigProfileAddCmd_InvalidProfileName(t *testing.T) {
	expectedErrorPattern := `^invalid profile name: '.*'. name must contain only alphanumeric characters, underscores, and dashes$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "add", "--name", "invalid$++$#@#$", "--set-active", "--description", "test-description")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Add Command fails when provided an existing profile name
func TestConfigProfileAddCmd_ExistingProfileName(t *testing.T) {
	expectedErrorPattern := `^invalid profile name: '.*' profile already exists$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "add", "--name", "default", "--set-active", "--description", "test-description")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Add Command fails when name flag is not provided a value
func TestConfigProfileAddCmd_NoProfileName(t *testing.T) {
	expectedErrorPattern := `^flag needs an argument: --name$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "add", "--set-active", "--description", "test-description", "--name")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Add Command fails when description flag is not provided a value
func TestConfigProfileAddCmd_NoProfileDescription(t *testing.T) {
	expectedErrorPattern := `^flag needs an argument: --description$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "add", "--name", "new-test-profile-name", "--set-active", "--description")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Add Command fails when set-active flag is provided an invalid value
func TestConfigProfileAddCmd_InvalidSetActiveValue(t *testing.T) {
	expectedErrorPattern := `^invalid argument "invalid" for "--set-active" flag: strconv.ParseBool: parsing "invalid": invalid syntax$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "add", "--name", "new-test-profile-name", "--set-active=invalid", "--description", "test-description")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Add Command fails when provided an invalid flag
func TestConfigProfileAddCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "add", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Add Command executes successfully when provided too many arguments
func TestConfigProfileAddCmd_TooManyArgs(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "add", "--name", "new-test-profile-name", "--set-active", "--description", "test-description", "extra-arg1", "extra-arg2")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Profile Add Command executes successfully when provided help flag
func TestConfigProfileAddCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "add", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingctl(t, "config", "profile", "add", "-h")
	testutils.CheckExpectedError(t, err, nil)
}
