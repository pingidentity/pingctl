package profile_internal

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test RunInternalConfigProfileSetActive function with no args
func Test_RunInternalConfigProfileSetActive_NoArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to set active profile: profile name is required$"
	err := RunInternalConfigProfileSetActive([]string{})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileSetActive function with multiple args
func Test_RunInternalConfigProfileSetActive_MultipleArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigProfileSetActive([]string{"production", "extra-arg"})
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigProfileSetActive function with valid args
func Test_RunInternalConfigProfileSetActive_ValidArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigProfileSetActive([]string{"production"})
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigProfileSetActive function with invalid profile name
func Test_RunInternalConfigProfileSetActive_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to set active profile: invalid profile name: '.*'. name must contain only alphanumeric characters, underscores, and dashes$"
	err := RunInternalConfigProfileSetActive([]string{"invalid&*^*&"})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileSetActive function with non-existent profile
func Test_RunInternalConfigProfileSetActive_NonExistentProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to set active profile: invalid profile name: '.*' profile does not exist$"
	err := RunInternalConfigProfileSetActive([]string{"invalid"})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileSetActive function with active profile
func Test_RunInternalConfigProfileSetActive_ActiveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigProfileSetActive([]string{"default"})
	testutils.CheckExpectedError(t, err, nil)
}

func Example_runInternalConfigProfileSetActive() {
	testutils_viper.InitVipers(&testing.T{})

	_ = RunInternalConfigProfileSetActive([]string{"production"})

	// Output:
	// Active configuration profile set to 'production' - Success
}

// Test parseSetActiveArgs function with no args
func Test_parseSetActiveArgs_NoArgs(t *testing.T) {
	expectedErrorPattern := "^profile name is required$"
	_, err := parseSetActiveArgs([]string{})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test parseSetActiveArgs function with multiple args
func Test_parseSetActiveArgs_MultipleArgs(t *testing.T) {
	_, err := parseSetActiveArgs([]string{"production", "extra-arg"})
	testutils.CheckExpectedError(t, err, nil)
}

// Test parseSetActiveArgs function with one arg
func Test_parseSetActiveArgs_OneArg(t *testing.T) {
	parsedArg, err := parseSetActiveArgs([]string{"production"})
	testutils.CheckExpectedError(t, err, nil)

	if parsedArg != "production" {
		t.Fatalf("Expected profile name to be 'production', got '%s'", parsedArg)
	}
}
