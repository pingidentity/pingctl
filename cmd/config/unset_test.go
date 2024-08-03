package config_test

import (
	"os"
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Config Unset Command Executes without issue
func TestConfigUnsetCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "unset", profiles.ColorOption.ViperKey)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Set Command Fails when provided too few arguments
func TestConfigUnsetCmd_TooFewArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingctl config unset': command accepts 1 arg\(s\), received 0$`
	err := testutils_cobra.ExecutePingctl(t, "config", "unset")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command Fails when provided too many arguments
func TestConfigUnsetCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingctl config unset': command accepts 1 arg\(s\), received 2$`
	err := testutils_cobra.ExecutePingctl(t, "config", "unset", profiles.ColorOption.ViperKey, profiles.OutputOption.ViperKey)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Unset Command Fails when an invalid key is provided
func TestConfigUnsetCmd_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^unable to unset configuration: key 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := testutils_cobra.ExecutePingctl(t, "config", "unset", "pingctl.invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Unset Command for key 'pingone.worker.clientId' updates viper configuration
func TestConfigUnsetCmd_CheckViperConfig(t *testing.T) {
	viperKey := profiles.PingOneWorkerClientIDOption.ViperKey
	viperOldValue := os.Getenv(profiles.PingOneWorkerClientIDOption.EnvVar)

	err := testutils_cobra.ExecutePingctl(t, "config", "unset", viperKey)
	testutils.CheckExpectedError(t, err, nil)

	viperNewValue := profiles.GetProfileViper().GetString(viperKey)
	if viperOldValue == viperNewValue {
		t.Errorf("Expected viper configuration value to be updated. Old: %s, New: %s", viperOldValue, viperNewValue)
	}

	if viperNewValue != "" {
		t.Errorf("Expected viper configuration value to be empty. Got: %s", viperNewValue)
	}
}
