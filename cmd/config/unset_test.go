package config_test

import (
	"regexp"
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils"
	"github.com/spf13/viper"
)

// Test Config Unset Command Executes without issue
func TestConfigUnsetCmd_Execute(t *testing.T) {
	err := testutils.ExecutePingctl("config", "unset", "pingctl.color")
	if err != nil {
		t.Errorf("Error executing config unset command: %s", err.Error())
	}
}

// Test Config Unset Command Fails when no arguments are provided
func TestConfigUnsetCmd_NoArgs(t *testing.T) {
	regex := regexp.MustCompile(`^unable to unset configuration: no key given in unset command$`)
	err := testutils.ExecutePingctl("config", "unset")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Config Unset command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test Config Unset Command Fails when an invalid key is provided
func TestConfigUnsetCmd_InvalidKey(t *testing.T) {
	regex := regexp.MustCompile(`^unable to unset configuration: key 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`)
	err := testutils.ExecutePingctl("config", "unset", "pingctl.invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Config Unset command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test Config Unset Command Executes normally when too many arguments are provided
func TestConfigUnsetCmd_TooManyArgs(t *testing.T) {
	err := testutils.ExecutePingctl("config", "unset", "pingctl.color", "pingctl.output")
	if err != nil {
		t.Errorf("Error executing config unset command: %s", err.Error())
	}
}

// Test Config Unset Command for key 'pingone.worker.clientId' updates viper configuration
func TestConfigUnsetCmd_CheckViperConfig(t *testing.T) {
	viper.Set("pingone.worker.clientId", "12345678-1234-1234-1234-123456789012")

	viperKey := "pingone.worker.clientId"
	viperOldValue := viper.GetString(viperKey)

	err := testutils.ExecutePingctl("config", "unset", viperKey)
	if err != nil {
		t.Errorf("Error executing config unset command: %s", err.Error())
	}

	viperNewValue := viper.GetString(viperKey)
	if viperOldValue == viperNewValue {
		t.Errorf("Expected viper configuration value to be updated. Old: %s, New: %s", viperOldValue, viperNewValue)
	}

	if viperNewValue != "" {
		t.Errorf("Expected viper configuration value to be empty. Got: %s", viperNewValue)
	}
}
