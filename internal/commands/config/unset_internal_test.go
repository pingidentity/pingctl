package config_internal

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test RunInternalConfigUnset function with invalid key
func Test_RunInternalConfigUnset_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^unable to unset configuration: key 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := RunInternalConfigUnset("pingctl.invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigUnset function with valid key
func Test_RunInternalConfigUnset_ValidKey(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigUnset(profiles.ColorOption.ViperKey)
	testutils.CheckExpectedError(t, err, nil)
}
