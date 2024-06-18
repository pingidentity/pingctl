package config_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils"
	"github.com/spf13/viper"
)

// Test Config Set Command Executes without issue
func TestConfigSetCmd_Execute(t *testing.T) {
	err := testutils.ExecutePingctl("config", "set", "pingctl.color=false")
	if err != nil {
		t.Errorf("Error executing config set command: %s", err.Error())
	}
}

// Test Config Set Command Fails when no arguments are provided
func TestConfigSetCmd_NoArgs(t *testing.T) {
	regex := regexp.MustCompile(`^failed to set configuration: no 'key=value' assignment given in set command$`)
	err := testutils.ExecutePingctl("config", "set")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Config Set command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test Config Set Command Fails when an invalid key is provided
func TestConfigSetCmd_InvalidKey(t *testing.T) {
	regex := regexp.MustCompile(`^failed to set configuration: key 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`)
	err := testutils.ExecutePingctl("config", "set", "pingctl.invalid=true")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Config Set command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test Config Set Command Fails when an invalid value type is provided
func TestConfigSetCmd_InvalidValueType(t *testing.T) {
	regex := regexp.MustCompile(`^failed to set configuration: value for key 'pingctl\.color' must be a boolean\. Use 'true' or 'false'$`)
	err := testutils.ExecutePingctl("config", "set", "pingctl.color=invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Config Set command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test Config Set Command Fails when no value is provided
func TestConfigSetCmd_NoValueProvided(t *testing.T) {
	regex := regexp.MustCompile(`^failed to set configuration: value for key 'pingctl\.color' is empty\. Use 'pingctl config unset pingctl\.color' to unset the key$`)
	err := testutils.ExecutePingctl("config", "set", "pingctl.color=")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Config Set command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test Config Set Command Executes normally when too many arguments are provided
func TestConfigSetCmd_TooManyArgs(t *testing.T) {
	err := testutils.ExecutePingctl("config", "set", "pingctl.color=false", "pingctl.color=true")
	if err != nil {
		t.Errorf("Error executing config set command: %s", err.Error())
	}
}

// Test Config Set Command for key 'pingone.worker.clientId' updates viper configuration
func TestConfigSetCmd_CheckViperConfig(t *testing.T) {
	viperKey := "pingone.worker.clientId"
	viperNewUUID := "12345678-1234-1234-1234-123456789012"
	viperOldValue := viper.GetString(viperKey)

	err := testutils.ExecutePingctl("config", "set", fmt.Sprintf("%s=%s", viperKey, viperNewUUID))
	if err != nil {
		t.Errorf("Error executing config set command: %s", err.Error())
	}

	viperNewValue := viper.GetString(viperKey)
	if viperNewValue != viperNewUUID {
		t.Errorf("Expected viper configuration value to be updated")
	}

	if viperOldValue == viperNewValue {
		t.Errorf("Expected viper configuration value to be updated")
	}
}
