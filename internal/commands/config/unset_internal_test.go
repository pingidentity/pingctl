package config_internal

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test RunInternalConfigUnset function with empty args
func Test_RunInternalConfigUnset_EmptyArgs(t *testing.T) {
	expectedErrorPattern := `^unable to unset configuration: no key given in unset command$`
	err := RunInternalConfigUnset([]string{})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigUnset function with invalid key
func Test_RunInternalConfigUnset_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^unable to unset configuration: key 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := RunInternalConfigUnset([]string{"pingctl.invalid"})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigUnset function with valid key
func Test_RunInternalConfigUnset_ValidKey(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigUnset([]string{profiles.ColorOption.ViperKey})
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigUnset function with too many args
func Test_RunInternalConfigUnset_TooManyArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigUnset([]string{profiles.ColorOption.ViperKey, profiles.WorkerClientIDOption.ViperKey})
	testutils.CheckExpectedError(t, err, nil)
}

// Test parseUnsetArgs function with empty args
func Test_parseUnsetArgs_EmptyArgs(t *testing.T) {
	expectedErrorPattern := `^unable to unset configuration: no key given in unset command$`
	_, err := parseUnsetArgs([]string{})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test parseUnsetArgs function with valid args
func Test_parseUnsetArgs_ValidArgs(t *testing.T) {
	_, err := parseUnsetArgs([]string{profiles.ColorOption.ViperKey})
	testutils.CheckExpectedError(t, err, nil)
}

// Test parseUnsetArgs function with too many args
func Test_parseUnsetArgs_TooManyArgs(t *testing.T) {
	_, err := parseUnsetArgs([]string{profiles.ColorOption.ViperKey, profiles.WorkerClientIDOption.ViperKey})
	testutils.CheckExpectedError(t, err, nil)
}
