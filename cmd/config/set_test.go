package config_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils/testutils_command"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
	"github.com/spf13/viper"
)

// Test Config Set Command Executes without issue
func TestConfigSetCmd_Execute(t *testing.T) {
	expectedErrorPattern := "" //No error expected
	err := testutils_command.ExecutePingctl("config", "set", "pingctl.color=false")
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test Config Set Command Fails when no arguments are provided
func TestConfigSetCmd_NoArgs(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: no 'key=value' assignment given in set command$`
	err := testutils_command.ExecutePingctl("config", "set")
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test Config Set Command Fails when an invalid key is provided
func TestConfigSetCmd_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: key 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := testutils_command.ExecutePingctl("config", "set", "pingctl.invalid=true")
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test Config Set Command Fails when an invalid value type is provided
func TestConfigSetCmd_InvalidValueType(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingctl\.color' must be a boolean\. Use 'true' or 'false'$`
	err := testutils_command.ExecutePingctl("config", "set", "pingctl.color=invalid")
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test Config Set Command Fails when no value is provided
func TestConfigSetCmd_NoValueProvided(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingctl\.color' is empty\. Use 'pingctl config unset pingctl\.color' to unset the key$`
	err := testutils_command.ExecutePingctl("config", "set", "pingctl.color=")
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test Config Set Command Executes normally when too many arguments are provided
func TestConfigSetCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := "" //No error expected
	err := testutils_command.ExecutePingctl("config", "set", "pingctl.color=false", "pingctl.output=json")
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test Config Set Command for key 'pingone.worker.clientId' updates viper configuration
func TestConfigSetCmd_CheckViperConfig(t *testing.T) {
	viperKey := "pingone.worker.clientId"
	viperNewUUID := "12345678-1234-1234-1234-123456789012"
	viperOldValue := viper.GetString(viperKey)

	expectedErrorPattern := "" //No error expected
	err := testutils_command.ExecutePingctl("config", "set", fmt.Sprintf("%s=%s", viperKey, viperNewUUID))
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)

	viperNewValue := viper.GetString(viperKey)
	if viperNewValue != viperNewUUID {
		t.Errorf("Expected viper configuration value to be updated")
	}

	if viperOldValue == viperNewValue {
		t.Errorf("Expected viper configuration value to be updated")
	}
}
