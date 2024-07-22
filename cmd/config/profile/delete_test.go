package profile_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Config Profile Delete Command Executes without issue
func TestConfigProfileDeleteCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "delete", "production")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Profile Delete Command fails when provided an non-existent profile name
func TestConfigProfileDeleteCmd_NonExistentProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to delete profile: invalid profile name: '.*' profile does not exist$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "delete", "invalid-profile")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Delete Command fails when no argument is provided
func TestConfigProfileDeleteCmd_NoProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to delete profile: profile name is required$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "delete")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Delete Command executes successfully when too many arguments are provided
func TestConfigProfileDeleteCmd_TooManyArgs(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "delete", "production", "extra-arg1", "extra-arg2")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Profile Delete Command fails when provided the active profile name
func TestConfigProfileDeleteCmd_ActiveProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to delete profile: '.*' is the active profile and cannot be deleted$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "delete", "default")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Delete Command fails when provided an invalid flag
func TestConfigProfileDeleteCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "delete", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Delete Command fails when provided an invalid profile name
func TestConfigProfileDeleteCmd_InvalidProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to delete profile: invalid profile name: '.*'. name must contain only alphanumeric characters, underscores, and dashes$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "delete", "invalid$++$#@#$")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Delete Command executes successfully when provided the help flag
func TestConfigProfileDeleteCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "delete", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingctl(t, "config", "profile", "delete", "-h")
	testutils.CheckExpectedError(t, err, nil)
}
