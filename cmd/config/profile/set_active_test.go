package profile_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Config Profile Set-Active Command Executes without issue
func TestConfigProfileSetActiveCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "set-active", "production")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Profile Set-Active Command fails when provided an non-existent profile name
func TestConfigProfileSetActiveCmd_NonExistentProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to set active profile: invalid profile name: '.*' profile does not exist$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "set-active", "invalid-profile")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Set-Active Command fails when no argument is provided
func TestConfigProfileSetActiveCmd_NoProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to set active profile: profile name is required$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "set-active")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Set-Active Command executes successfully when too many arguments are provided
func TestConfigProfileSetActiveCmd_TooManyArgs(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "production", "production", "extra-arg1", "extra-arg2")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Profile Set-Active Command executes successfully when provided the already active profile name
func TestConfigProfileSetActiveCmd_ActiveProfileName(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "set-active", "default")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Profile Set-Active Command fails when provided an invalid flag
func TestConfigProfileSetActiveCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "set-active", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Set-Active Command fails when provided an invalid profile name
func TestConfigProfileSetActiveCmd_InvalidProfileName(t *testing.T) {
	expectedErrorPattern := `^failed to set active profile: invalid profile name: '.*'. name must contain only alphanumeric characters, underscores, and dashes$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "set-active", "invalid$++$#@#$")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile Set-Active Command executes successfully when provided the help flag
func TestConfigProfileSetActiveCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "set-active", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingctl(t, "config", "profile", "set-active", "-h")
	testutils.CheckExpectedError(t, err, nil)
}
