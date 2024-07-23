package profile_internal

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test RunInternalConfigProfileDescribe function with no args
func Test_RunInternalConfigProfileDescribe_NoArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to describe profile: profile name is required$"
	err := RunInternalConfigProfileDescribe([]string{})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileDescribe function with multiple args
func Test_RunInternalConfigProfileDescribe_MultipleArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigProfileDescribe([]string{"production", "extra-arg"})
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigProfileDescribe function with valid args
func Test_RunInternalConfigProfileDescribe_ValidArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigProfileDescribe([]string{"production"})
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigProfileDescribe function with invalid profile name
func Test_RunInternalConfigProfileDescribe_InvalidProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to describe profile: invalid profile name: '.*'. name must contain only alphanumeric characters, underscores, and dashes$"
	err := RunInternalConfigProfileDescribe([]string{"invalid&*^*&"})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileDescribe function with non-existent profile
func Test_RunInternalConfigProfileDescribe_NonExistentProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := "^failed to describe profile: invalid profile name: '.*' profile does not exist$"
	err := RunInternalConfigProfileDescribe([]string{"invalid"})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileDescribe function with active profile
func Test_RunInternalConfigProfileDescribe_ActiveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigProfileDescribe([]string{"default"})
	testutils.CheckExpectedError(t, err, nil)
}

// Check output of RunInternalConfigProfileDescribe function
func Example_runInternalConfigProfileDescribe() {
	testutils_viper.InitVipers(&testing.T{})

	_ = RunInternalConfigProfileDescribe([]string{"production"})

	// Output:
	// Profile Name: production
	// Description: test profile description
	//
	// Set Options:
	//  - pingctl.output: text
	//  - pingctl.color: true
	//
	// Unset Options:
	//  - pingone.export.environmentID
	//  - pingone.worker.environmentID
	//  - pingone.worker.clientID
	//  - pingone.worker.clientSecret
	//  - pingone.region
}

// Test parseDescArgs function with no args
func Test_parseDescArgs_NoArgs(t *testing.T) {
	expectedErrorPattern := "^profile name is required$"
	_, err := parseDescArgs([]string{})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test parseDescArgs function with multiple args
func Test_parseDescArgs_MultipleArgs(t *testing.T) {
	_, err := parseDescArgs([]string{"production", "extra-arg"})
	testutils.CheckExpectedError(t, err, nil)
}

// Test parseDescArgs function with one arg
func Test_parseDescArgs_OneArg(t *testing.T) {
	parsedArg, err := parseDescArgs([]string{"production"})
	testutils.CheckExpectedError(t, err, nil)

	if parsedArg != "production" {
		t.Fatalf("Expected parsed arg to be 'production', got '%s'", parsedArg)
	}
}
