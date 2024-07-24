package profile_internal

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test RunInternalConfigProfileSetActive function with valid args
func Test_RunInternalConfigProfileSetActive_ValidArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigProfileSetActive("production")
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigProfileSetActive function with invalid profile name
func Test_RunInternalConfigProfileSetActive_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to set active profile: invalid profile name: '.*'. name must contain only alphanumeric characters, underscores, and dashes$"
	err := RunInternalConfigProfileSetActive("invalid&*^*&")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileSetActive function with non-existent profile
func Test_RunInternalConfigProfileSetActive_NonExistentProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to set active profile: invalid profile name: '.*' profile does not exist$"
	err := RunInternalConfigProfileSetActive("invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileSetActive function with active profile
func Test_RunInternalConfigProfileSetActive_ActiveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigProfileSetActive("default")
	testutils.CheckExpectedError(t, err, nil)
}

func Example_runInternalConfigProfileSetActive() {
	testutils_viper.InitVipers(&testing.T{})

	_ = RunInternalConfigProfileSetActive("production")

	// Output:
	// Active configuration profile set to 'production' - Success
}
