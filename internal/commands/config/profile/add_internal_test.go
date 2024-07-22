package profile_internal

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/manifoldco/promptui"
	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test RunInternalConfigProfileAdd function with no user input needed
func Test_RunInternalConfigProfileAdd_NoInput(t *testing.T) {
	testutils_viper.InitVipers(t)

	newProfileName := "test-profile"

	err := RunInternalConfigProfileAdd(newProfileName, "test-description", false, true, nil)
	testutils.CheckExpectedError(t, err, nil)

	// Check if the profile is created
	if err := profiles.ValidateExistingProfileName(newProfileName); err != nil {
		t.Fatalf("Expected profile '%s' to exist, got error: %v", newProfileName, err)
	}
}

// Test RunInternalConfigProfileAdd function with a valid name from user input
func Test_RunInternalConfigProfileAdd_ValidName(t *testing.T) {
	testutils_viper.InitVipers(t)

	newProfileName := "test-profile"

	reader := testutils.WriteStringToPipe(fmt.Sprintf("%s\n", newProfileName), t)
	defer reader.Close()

	err := RunInternalConfigProfileAdd("", "test-description", false, true, reader)
	testutils.CheckExpectedError(t, err, nil)

	// Check if the profile is created
	if err := profiles.ValidateExistingProfileName(newProfileName); err != nil {
		t.Fatalf("Expected profile '%s' to exist, got error: %v", newProfileName, err)
	}
}

// Test RunInternalConfigProfileAdd function with an invalid name from user input
func Test_RunInternalConfigProfileAdd_InvalidName(t *testing.T) {
	testutils_viper.InitVipers(t)

	reader := testutils.WriteStringToPipe("invalid$++$#@#$\n", t)
	defer reader.Close()

	promptuiErrMsg := promptui.ErrEOF.Error()
	expectedErrorPattern := fmt.Sprintf("^%s$", regexp.QuoteMeta(promptuiErrMsg))
	err := RunInternalConfigProfileAdd("", "test-description", false, true, reader)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileAdd function with an empty name from user input
func Test_RunInternalConfigProfileAdd_EmptyName(t *testing.T) {
	testutils_viper.InitVipers(t)

	reader := testutils.WriteStringToPipe("\n", t)
	defer reader.Close()

	promptuiErrMsg := promptui.ErrEOF.Error()
	expectedErrorPattern := fmt.Sprintf("^%s$", regexp.QuoteMeta(promptuiErrMsg))
	err := RunInternalConfigProfileAdd("", "test-description", false, true, reader)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigProfileAdd function with a valid description from user input
func Test_RunInternalConfigProfileAdd_ValidDescription(t *testing.T) {
	testutils_viper.InitVipers(t)

	newProfileName := "test-profile"

	reader := testutils.WriteStringToPipe("test-description\n", t)
	defer reader.Close()

	err := RunInternalConfigProfileAdd(newProfileName, "", false, true, reader)
	testutils.CheckExpectedError(t, err, nil)

	// Check if the profile is created
	if err := profiles.ValidateExistingProfileName(newProfileName); err != nil {
		t.Fatalf("Expected profile '%s' to exist, got error: %v", newProfileName, err)
	}
}

// Test RunInternalConfigProfileAdd function with an empty description from user input
func Test_RunInternalConfigProfileAdd_EmptyDescription(t *testing.T) {
	testutils_viper.InitVipers(t)

	newProfileName := "test-profile"

	reader := testutils.WriteStringToPipe("\n", t)
	defer reader.Close()

	err := RunInternalConfigProfileAdd(newProfileName, "", false, true, reader)
	testutils.CheckExpectedError(t, err, nil)

	// Check if the profile is created
	if err := profiles.ValidateExistingProfileName(newProfileName); err != nil {
		t.Fatalf("Expected profile '%s' to exist, got error: %v", newProfileName, err)
	}
}

// Test RunInternalConfigProfileAdd function with a 'y' user input to set active confirm prompt
func Test_RunInternalConfigProfileAdd_ActiveProfileYes(t *testing.T) {
	testutils_viper.InitVipers(t)

	newProfileName := "test-profile"

	reader := testutils.WriteStringToPipe("y\n", t)
	defer reader.Close()

	err := RunInternalConfigProfileAdd(newProfileName, "test-description", false, false, reader)
	testutils.CheckExpectedError(t, err, nil)

	// Check if the profile is created
	if err := profiles.ValidateExistingProfileName(newProfileName); err != nil {
		t.Fatalf("Expected profile '%s' to exist, got error: %v", newProfileName, err)
	}

	// Check if the profile is active
	profileName := profiles.GetConfigActiveProfile()
	if profileName != newProfileName {
		t.Fatalf("Expected active profile to be '%s', got '%s'", newProfileName, profileName)
	}
}

// Test RunInternalConfigProfileAdd function with a 'n' user input to set active confirm prompt
func Test_RunInternalConfigProfileAdd_ActiveProfileNo(t *testing.T) {
	testutils_viper.InitVipers(t)

	newProfileName := "test-profile"

	reader := testutils.WriteStringToPipe("n\n", t)
	defer reader.Close()

	err := RunInternalConfigProfileAdd(newProfileName, "test-description", false, false, reader)
	testutils.CheckExpectedError(t, err, nil)

	// Check if the profile is created
	if err := profiles.ValidateExistingProfileName(newProfileName); err != nil {
		t.Fatalf("Expected profile '%s' to exist, got error: %v", newProfileName, err)
	}

	// Check if the profile is not active
	profileName := profiles.GetConfigActiveProfile()
	if profileName == newProfileName {
		t.Fatalf("Expected active profile to not be '%s'", newProfileName)
	}
}

// Test RunInternalConfigProfileAdd function with an invalid set active confirm prompt user input
// promptui assumes non-'y' responses to be a 'n' response
func Test_RunInternalConfigProfileAdd_ActiveProfileInvalidInput(t *testing.T) {
	testutils_viper.InitVipers(t)

	newProfileName := "test-profile"

	reader := testutils.WriteStringToPipe("invalid\n", t)
	defer reader.Close()

	err := RunInternalConfigProfileAdd(newProfileName, "test-description", false, false, reader)
	testutils.CheckExpectedError(t, err, nil)

	// Check if the profile is created
	if err := profiles.ValidateExistingProfileName(newProfileName); err != nil {
		t.Fatalf("Expected profile '%s' to exist, got error: %v", newProfileName, err)
	}

	// Check if the profile is not active
	profileName := profiles.GetConfigActiveProfile()
	if profileName == newProfileName {
		t.Fatalf("Expected active profile to not be '%s'", newProfileName)
	}
}

// Test RunInternalConfigProfileAdd function with an existing profile name
func Test_RunInternalConfigProfileAdd_ExistingProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	newProfileName := "default"

	expectedErrorPattern := "^invalid profile name: '.*' profile already exists$"
	err := RunInternalConfigProfileAdd(newProfileName, "test-description", false, true, nil)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
