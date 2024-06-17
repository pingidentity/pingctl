package config_internal

import (
	"os"
	"testing"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/viperconfig"
	"github.com/spf13/viper"
)

// Test RunInternalConfigSet function
func Test_RunInternalConfigSet_NoArgs(t *testing.T) {
	args := []string{}
	if err := RunInternalConfigSet(args); err == nil {
		t.Errorf("Expected error running internal config set")
	}
}

// Test RunInternalConfigSet function with args
func Test_RunInternalConfigSet_WithArgs(t *testing.T) {
	// This is the happy path, so we need a valid config file to write to
	// Create a valid config file
	configDir := os.TempDir() + "/pingctlTestRunInternalConfigSetWithArgs"
	configFile := configDir + "/config.yaml"
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Errorf("Error creating config directory: %s", err.Error())
	}
	if _, err := os.Create(configFile); err != nil {
		t.Errorf("Error creating config file: %s", err.Error())
	}

	// Set the config file
	viper.SetConfigFile(configFile)

	args := []string{"pingone.worker.clientId=12345678-1234-1234-1234-123456789012"}
	if err := RunInternalConfigSet(args); err != nil {
		t.Errorf("Error running internal config set: %s", err.Error())
	}

	// Clean up
	if err := os.RemoveAll(configDir); err != nil {
		t.Errorf("Error removing config directory: %s", err.Error())
	}
}

// Test RunInternalConfigSet function with invalid key
func Test_RunInternalConfigSet_InvalidKey(t *testing.T) {
	args := []string{"pingctl.invalid=invalid"}
	if err := RunInternalConfigSet(args); err == nil {
		t.Errorf("Expected error running internal config set")
	}
}

// Test RunInternalConfigSet function with too many args
func Test_RunInternalConfigSet_TooManyArgs(t *testing.T) {
	// This is the happy path, so we need a valid config file to write to
	// Create a valid config file
	configDir := os.TempDir() + "/pingctlTestRunInternalConfigSetTooManyArgs"
	configFile := configDir + "/config.yaml"
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Errorf("Error creating config directory: %s", err.Error())
	}
	if _, err := os.Create(configFile); err != nil {
		t.Errorf("Error creating config file: %s", err.Error())
	}

	// Set the config file
	viper.SetConfigFile(configFile)

	args := []string{"pingone.worker.clientId=12345678-1234-1234-1234-123456789012", "pingone.worker.environmentId=12345678-1234-1234-1234-123456789012"}
	if err := RunInternalConfigSet(args); err != nil {
		t.Errorf("Error running internal config set: %s", err.Error())
	}

	// Clean up
	if err := os.RemoveAll(configDir); err != nil {
		t.Errorf("Error removing config directory: %s", err.Error())
	}
}

// Test RunInternalConfigSet function with empty value
func Test_RunInternalConfigSet_EmptyValue(t *testing.T) {
	args := []string{"pingone.worker.clientId="}
	if err := RunInternalConfigSet(args); err == nil {
		t.Errorf("Expected error running internal config set")
	}
}

// Test RunInternalConfigSet function with invalid value
func Test_RunInternalConfigSet_InvalidValue(t *testing.T) {
	args := []string{"pingone.worker.clientId=invalid"}
	if err := RunInternalConfigSet(args); err == nil {
		t.Errorf("Expected error running internal config set")
	}
}

// Test RunInternalConfigSet function with invalid value type
func Test_RunInternalConfigSet_InvalidValueType(t *testing.T) {
	args := []string{"pingctl.color=notboolean"}
	if err := RunInternalConfigSet(args); err == nil {
		t.Errorf("Expected error running internal config set")
	}
}

// Test parseSetArgs() function with no args
func Test_parseSetArgs_NoArgs(t *testing.T) {
	args := []string{}
	if _, _, err := parseSetArgs(args); err == nil {
		t.Errorf("Expected error parsing set args")
	}
}

// Test parseSetArgs() function with invalid assignment format
func Test_parseSetArgs_InvalidAssignmentFormat(t *testing.T) {
	args := []string{"pingone.worker.clientId"}
	if _, _, err := parseSetArgs(args); err == nil {
		t.Errorf("Expected error parsing set args")
	}
}

// Test parseSetArgs() function with valid assignment format
func Test_parseSetArgs_ValidAssignmentFormat(t *testing.T) {
	args := []string{"pingone.worker.clientId=12345678-1234-1234-1234-123456789012"}
	if _, _, err := parseSetArgs(args); err != nil {
		t.Errorf("Error parsing set args: %s", err.Error())
	}
}

// Test parseSetArgs() function with too many args
func Test_parseSetArgs_TooManyArgs(t *testing.T) {
	args := []string{"pingone.worker.clientId=12345678-1234-1234-1234-123456789012", "pingone.worker.environmentId=12345678-1234-1234-1234-123456789012"}
	if _, _, err := parseSetArgs(args); err != nil {
		t.Errorf("Error parsing set args: %s", err.Error())
	}
}

// Test setValue() function with valid value
func Test_setValue_ValidValue(t *testing.T) {
	if err := setValue("pingone.worker.clientId", "12345678-1234-1234-1234-123456789012", viperconfig.ENUM_ID); err != nil {
		t.Errorf("Error setting value: %s", err.Error())
	}
}

// Test setValue() function with invalid value
func Test_setValue_InvalidValue(t *testing.T) {
	if err := setValue("pingone.worker.clientId", "invalid", viperconfig.ENUM_ID); err == nil {
		t.Errorf("Expected error setting value")
	}
}

// Test setValue() function with invalid value type
func Test_setValue_InvalidValueType(t *testing.T) {
	if err := setValue("pingctl.color", "false", "invalid"); err == nil {
		t.Errorf("Expected error setting value")
	}
}

// Test setBool() function with valid value
func Test_setBool_ValidValue(t *testing.T) {
	if err := setBool("pingctl.color", "false"); err != nil {
		t.Errorf("Error setting bool: %s", err.Error())
	}
}

// Test setBool() function with invalid value
func Test_setBool_InvalidValue(t *testing.T) {
	if err := setBool("pingctl.color", "invalid"); err == nil {
		t.Errorf("Expected error setting bool")
	}
}

// Test setUUID() function with valid value
func Test_setUUID_ValidValue(t *testing.T) {
	if err := setUUID("pingone.worker.clientId", "12345678-1234-1234-1234-123456789012"); err != nil {
		t.Errorf("Error setting UUID: %s", err.Error())
	}
}

// Test setUUID() function with invalid value
func Test_setUUID_InvalidValue(t *testing.T) {
	if err := setUUID("pingone.worker.clientId", "invalid"); err == nil {
		t.Errorf("Expected error setting UUID")
	}
}

// Test setOutputFormat() function with valid value
func Test_setOutputFormat_ValidValue(t *testing.T) {
	if err := setOutputFormat("pingctl.output", "json"); err != nil {
		t.Errorf("Error setting output format: %s", err.Error())
	}
}

// Test setOutputFormat() function with invalid value
func Test_setOutputFormat_InvalidValue(t *testing.T) {
	if err := setOutputFormat("pingctl.output", "invalid"); err == nil {
		t.Errorf("Expected error setting output format")
	}
}

// Test setPingOneRegion() function with valid value
func Test_setPingOneRegion_ValidValue(t *testing.T) {
	if err := setPingOneRegion("pingone.region", customtypes.ENUM_PINGONE_REGION_AP); err != nil {
		t.Errorf("Error setting PingOne region: %s", err.Error())
	}
}

// Test setPingOneRegion() function with invalid value
func Test_setPingOneRegion_InvalidValue(t *testing.T) {
	if err := setPingOneRegion("pingone.region", "invalid"); err == nil {
		t.Errorf("Expected error setting PingOne region")
	}
}
