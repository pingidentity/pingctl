package config_test

import (
	"fmt"
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
	err := testutils.ExecutePingctl("config", "set")
	if err == nil {
		t.Errorf("Expected error for no arguments provided")
	}
}

// Test Config Set Command Fails when an invalid key is provided
func TestConfigSetCmd_InvalidKey(t *testing.T) {
	err := testutils.ExecutePingctl("config", "set", "pingctl.invalid=true")
	if err == nil {
		t.Errorf("Expected error for providing invalid key")
	}
}

// Test Config Set Command Fails when an invalid value type is provided
func TestConfigSetCmd_InvalidValueType(t *testing.T) {
	err := testutils.ExecutePingctl("config", "set", "pingctl.color=invalid")
	if err == nil {
		t.Errorf("Expected error for providing invalid value type")
	}
}

// Test Config Set Command Fails when no value is provided
func TestConfigSetCmd_NoValueProvided(t *testing.T) {
	err := testutils.ExecutePingctl("config", "set", "pingctl.color=")
	if err == nil {
		t.Errorf("Expected error for not providing a value")
	}
}

// Test Config Set Command Executes normally when too many arguments are provided
func TestConfigSetCmd_TooManyArgs(t *testing.T) {
	err := testutils.ExecutePingctl("config", "set", "pingctl.color=false", "pingctl.color=true")
	if err != nil {
		t.Errorf("Expected error for too many arguments provided")
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
