package config_internal

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test RunInternalConfigUnset function with empty args
func Test_RunInternalConfigUnset_EmptyArgs(t *testing.T) {
	expectedErrorPattern := `^unable to unset configuration: no key given in unset command$`
	err := RunInternalConfigUnset([]string{})
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigUnset function with invalid key
func Test_RunInternalConfigUnset_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^unable to unset configuration: key 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := RunInternalConfigUnset([]string{"pingctl.invalid"})
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigUnset function with valid key
func Test_RunInternalConfigUnset_ValidKey(t *testing.T) {
	testutils_helpers.InitVipers(t)

	err := RunInternalConfigUnset([]string{profiles.ColorOption.ViperKey})
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigUnset function with too many args
func Test_RunInternalConfigUnset_TooManyArgs(t *testing.T) {
	testutils_helpers.InitVipers(t)

	err := RunInternalConfigUnset([]string{profiles.ColorOption.ViperKey, profiles.WorkerClientIDOption.ViperKey})
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test parseUnsetArgs function with empty args
func Test_parseUnsetArgs_EmptyArgs(t *testing.T) {
	expectedErrorPattern := `^unable to unset configuration: no key given in unset command$`
	_, err := parseUnsetArgs([]string{})
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test parseUnsetArgs function with valid args
func Test_parseUnsetArgs_ValidArgs(t *testing.T) {
	_, err := parseUnsetArgs([]string{profiles.ColorOption.ViperKey})
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test parseUnsetArgs function with too many args
func Test_parseUnsetArgs_TooManyArgs(t *testing.T) {
	_, err := parseUnsetArgs([]string{profiles.ColorOption.ViperKey, profiles.WorkerClientIDOption.ViperKey})
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test unsetValue function with invalid value type
func Test_unsetValue_InvalidValueType(t *testing.T) {
	expectedErrorPattern := `^unable to unset configuration: variable type for key 'pingctl\.color' is not recognized$`
	err := UnsetValue(profiles.ColorOption.ViperKey, "invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test unsetValue function with valid value type
func Test_unsetValue_ValidValueType(t *testing.T) {
	testutils_helpers.InitVipers(t)

	err := UnsetValue(profiles.ColorOption.ViperKey, profiles.ENUM_BOOL)
	testutils_helpers.CheckExpectedError(t, err, nil)
}
