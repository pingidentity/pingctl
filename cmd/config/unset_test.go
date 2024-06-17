package config_test

import (
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
	err := testutils.ExecutePingctl("config", "unset")
	if err == nil {
		t.Errorf("Expected error for no arguments provided")
	}
}

// Test Config Unset Command Fails when an invalid key is provided
func TestConfigUnsetCmd_InvalidKey(t *testing.T) {
	err := testutils.ExecutePingctl("config", "unset", "pingctl.invalid")
	if err == nil {
		t.Errorf("Expected error for invalid key provided")
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
