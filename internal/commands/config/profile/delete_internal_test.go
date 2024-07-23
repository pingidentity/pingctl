package profile_internal

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test RunInternalConfigProfileDelete function with valid arg
func Test_RunInternalConfigProfileDelete_ValidArg(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigProfileDelete("production")
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigProfileDelete function with invalid profile name
func Test_RunInternalConfigProfileDelete_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to delete profile: invalid profile name: '.*'. name must contain only alphanumeric characters, underscores, and dashes$"
	err := RunInternalConfigProfileDelete("invalid&*^*&")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileDelete function with non-existent profile
func Test_RunInternalConfigProfileDelete_NonExistentProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to delete profile: invalid profile name: '.*' profile does not exist$"
	err := RunInternalConfigProfileDelete("invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileDelete function with active profile
func Test_RunInternalConfigProfileDelete_ActiveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to delete profile: '.*' is the active profile and cannot be deleted$"
	err := RunInternalConfigProfileDelete("default")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Check output of RunInternalConfigProfileDelete function
func Example_runInternalConfigProfileDelete() {
	testutils_viper.InitVipers(&testing.T{})

	_ = RunInternalConfigProfileDelete("production")

	// Output:
	// Profile 'production' deleted successfully - Success
}
