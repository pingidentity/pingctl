package profile_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Config Profile Describe Command Executes without issue
func TestConfigProfileDescribeCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "describe", "production")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Profile Describe Command fails when provided too few arguments
func TestConfigProfileDescribeCmd_TooFewArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingctl config profile describe': command accepts 1 arg\(s\), received 0$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "describe")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Describe Command fails when provided too many arguments
func TestConfigProfileDescribeCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingctl config profile describe': command accepts 1 arg\(s\), received 2$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "describe", "production", "extra-arg")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Describe Command fails when provided an non-existent profile name
func TestConfigProfileDescribeCmd_NonExistentProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to describe profile: invalid profile name: '.*' profile does not exist$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "describe", "invalid-profile")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Describe Command fails when provided an invalid flag
func TestConfigProfileDescribeCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "describe", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Describe Command fails when provided an invalid profile name
func TestConfigProfileDescribeCmd_InvalidProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to describe profile: invalid profile name: '.*'. name must contain only alphanumeric characters, underscores, and dashes$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "describe", "invalid$++$#@#$")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Describe Command executes successfully when provided the help flag
func TestConfigProfileDescribeCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "describe", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingctl(t, "config", "profile", "describe", "-h")
	testutils.CheckExpectedError(t, err, nil)
}
