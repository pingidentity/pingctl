package config_test

import (
	"os"
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_command"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test Config Unset Command Executes without issue
func TestConfigUnsetCmd_Execute(t *testing.T) {
	err := testutils_command.ExecutePingctl("config", "unset", profiles.ColorOption.ViperKey)
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Config Unset Command Fails when no arguments are provided
func TestConfigUnsetCmd_NoArgs(t *testing.T) {
	expectedErrorPattern := `^unable to unset configuration: no key given in unset command$`
	err := testutils_command.ExecutePingctl("config", "unset")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Unset Command Fails when an invalid key is provided
func TestConfigUnsetCmd_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^unable to unset configuration: key 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := testutils_command.ExecutePingctl("config", "unset", "pingctl.invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Unset Command Executes normally when too many arguments are provided
func TestConfigUnsetCmd_TooManyArgs(t *testing.T) {
	err := testutils_command.ExecutePingctl("config", "unset", profiles.ColorOption.ViperKey, profiles.OutputOption.ViperKey)
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Config Unset Command for key 'pingone.worker.clientId' updates viper configuration
func TestConfigUnsetCmd_CheckViperConfig(t *testing.T) {
	viperKey := profiles.WorkerClientIDOption.ViperKey
	viperOldValue := os.Getenv(profiles.WorkerClientIDOption.EnvVar)

	err := testutils_command.ExecutePingctl("config", "unset", viperKey)
	testutils_helpers.CheckExpectedError(t, err, nil)

	viperNewValue := profiles.GetProfileViper().GetString(viperKey)
	if viperOldValue == viperNewValue {
		t.Errorf("Expected viper configuration value to be updated. Old: %s, New: %s", viperOldValue, viperNewValue)
	}

	if viperNewValue != "" {
		t.Errorf("Expected viper configuration value to be empty. Got: %s", viperNewValue)
	}
}
