package config_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Config Set Command Executes without issue
func TestConfigSetCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "set", fmt.Sprintf("%s=false", profiles.ColorOption.ViperKey))
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Set Command Fails when no arguments are provided
func TestConfigSetCmd_NoArgs(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: no 'key=value' assignment given in set command$`
	err := testutils_cobra.ExecutePingctl(t, "config", "set")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command Fails when an invalid key is provided
func TestConfigSetCmd_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: key 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := testutils_cobra.ExecutePingctl(t, "config", "set", "pingctl.invalid=true")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command Fails when an invalid value type is provided
func TestConfigSetCmd_InvalidValueType(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingctl\.color' must be a boolean\. Use 'true' or 'false'$`
	err := testutils_cobra.ExecutePingctl(t, "config", "set", fmt.Sprintf("%s=invalid", profiles.ColorOption.ViperKey))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command Fails when no value is provided
func TestConfigSetCmd_NoValueProvided(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingctl\.color' is empty\. Use 'pingctl config unset pingctl\.color' to unset the key$`
	err := testutils_cobra.ExecutePingctl(t, "config", "set", fmt.Sprintf("%s=", profiles.ColorOption.ViperKey))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Set Command Executes normally when too many arguments are provided
func TestConfigSetCmd_TooManyArgs(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "set", fmt.Sprintf("%s=false", profiles.ColorOption.ViperKey), fmt.Sprintf("%s=json", profiles.OutputOption.ViperKey))
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Set Command for key 'pingone.worker.clientId' updates viper configuration
func TestConfigSetCmd_CheckViperConfig(t *testing.T) {
	viperKey := profiles.WorkerClientIDOption.ViperKey
	viperNewUUID := "12345678-1234-1234-1234-123456789012"

	err := testutils_cobra.ExecutePingctl(t, "config", "set", fmt.Sprintf("%s=%s", viperKey, viperNewUUID))
	testutils.CheckExpectedError(t, err, nil)

	viperNewValue := profiles.GetProfileViper().GetString(viperKey)
	if viperNewValue != viperNewUUID {
		t.Errorf("Expected viper configuration value to be updated")
	}
}
