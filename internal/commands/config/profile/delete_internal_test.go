package profile_internal

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test RunInternalConfigProfileDelete function with no args
func Test_RunInternalConfigProfileDelete_NoArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to delete profile: profile name is required$"
	err := RunInternalConfigProfileDelete([]string{})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)

}

// Test RunInternalConfigProfileDelete function with multiple args
func Test_RunInternalConfigProfileDelete_MultipleArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigProfileDelete([]string{"production", "extra-arg"})
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigProfileDelete function with valid args
func Test_RunInternalConfigProfileDelete_ValidArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigProfileDelete([]string{"production"})
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigProfileDelete function with invalid profile name
func Test_RunInternalConfigProfileDelete_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to delete profile: invalid profile name: '.*'. name must contain only alphanumeric characters, underscores, and dashes$"
	err := RunInternalConfigProfileDelete([]string{"invalid&*^*&"})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileDelete function with non-existent profile
func Test_RunInternalConfigProfileDelete_NonExistentProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to delete profile: invalid profile name: '.*' profile does not exist$"
	err := RunInternalConfigProfileDelete([]string{"invalid"})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileDelete function with active profile
func Test_RunInternalConfigProfileDelete_ActiveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to delete profile: '.*' is the active profile and cannot be deleted$"
	err := RunInternalConfigProfileDelete([]string{"default"})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Check output of RunInternalConfigProfileDelete function
func Example_runInternalConfigProfileDelete() {
	testutils_viper.InitVipers(&testing.T{})

	_ = RunInternalConfigProfileDelete([]string{"production"})

	// Output:
	// Profile 'production' deleted successfully - Success
}

// Test parseDeleteArgs function with no args
func Test_parseDeleteArgs_NoArgs(t *testing.T) {
	expectedErrorPattern := "^profile name is required$"
	_, err := parseDeleteArgs([]string{})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test parseDeleteArgs function with multiple args
func Test_parseDeleteArgs_MultipleArgs(t *testing.T) {
	_, err := parseDeleteArgs([]string{"production", "extra-arg"})
	testutils.CheckExpectedError(t, err, nil)
}

// Test parseDeleteArgs function with one arg
func Test_parseDeleteArgs_OneArg(t *testing.T) {
	parsedArg, err := parseDeleteArgs([]string{"production"})
	testutils.CheckExpectedError(t, err, nil)

	if parsedArg != "production" {
		t.Fatalf("Expected parsedArg to be 'production', got '%s'", parsedArg)
	}
}
