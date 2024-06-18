package config_internal

import (
	"os"
	"regexp"
	"testing"

	"github.com/pingidentity/pingctl/internal/viperconfig"
	"github.com/spf13/viper"
)

// Test RunInternalConfigUnset function with empty args
func Test_RunInternalConfigUnset_EmptyArgs(t *testing.T) {
	regex := regexp.MustCompile(`^unable to unset configuration: no key given in unset command$`)
	err := RunInternalConfigUnset([]string{})

	if !regex.MatchString(err.Error()) {
		t.Errorf("RunInternalConfigUnset() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test RunInternalConfigUnset function with invalid key
func Test_RunInternalConfigUnset_InvalidKey(t *testing.T) {
	regex := regexp.MustCompile(`^unable to unset configuration: key 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`)
	err := RunInternalConfigUnset([]string{"pingctl.invalid"})

	if !regex.MatchString(err.Error()) {
		t.Errorf("RunInternalConfigUnset() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test RunInternalConfigUnset function with valid key
func Test_RunInternalConfigUnset_ValidKey(t *testing.T) {
	// This is the happy path, so we need a valid config file to write to
	// Create a valid config file
	configDir := os.TempDir() + "/pingctlTestRunInternalConfigUnsetValidKey"
	configFile := configDir + "/config.yaml"
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Errorf("Error creating config directory: %s", err.Error())
	}
	if _, err := os.Create(configFile); err != nil {
		t.Errorf("Error creating config file: %s", err.Error())
	}

	// Set the config file
	viper.SetConfigFile(configFile)

	args := []string{"pingctl.color"}
	if err := RunInternalConfigUnset(args); err != nil {
		t.Errorf("Error running internal config unset: %s", err.Error())
	}

	// Clean up
	if err := os.RemoveAll(configDir); err != nil {
		t.Errorf("Error removing config directory: %s", err.Error())
	}
}

// Test RunInternalConfigUnset function with too many args
func Test_RunInternalConfigUnset_TooManyArgs(t *testing.T) {
	// This is the happy path, so we need a valid config file to write to
	// Create a valid config file
	configDir := os.TempDir() + "/pingctlTestRunInternalConfigUnsetTooManyArgs"
	configFile := configDir + "/config.yaml"
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Errorf("Error creating config directory: %s", err.Error())
	}
	if _, err := os.Create(configFile); err != nil {
		t.Errorf("Error creating config file: %s", err.Error())
	}

	// Set the config file
	viper.SetConfigFile(configFile)

	args := []string{"pingctl.color", "pingctl.arg", "pingctl.arg2"}
	if err := RunInternalConfigUnset(args); err != nil {
		t.Errorf("Error running internal config unset: %s", err.Error())
	}

	// Clean up
	if err := os.RemoveAll(configDir); err != nil {
		t.Errorf("Error removing config directory: %s", err.Error())
	}
}

// Test parseUnsetArgs function with empty args
func Test_parseUnsetArgs_EmptyArgs(t *testing.T) {
	regex := regexp.MustCompile(`^unable to unset configuration: no key given in unset command$`)
	_, err := parseUnsetArgs([]string{})

	if !regex.MatchString(err.Error()) {
		t.Errorf("parseUnsetArgs() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test parseUnsetArgs function with valid args
func Test_parseUnsetArgs_ValidArgs(t *testing.T) {
	args := []string{"pingctl.color"}
	if _, err := parseUnsetArgs(args); err != nil {
		t.Errorf("Error parsing unset args: %s", err.Error())
	}
}

// Test parseUnsetArgs function with too many args
func Test_parseUnsetArgs_TooManyArgs(t *testing.T) {
	args := []string{"pingctl.color", "pingctl.arg", "pingctl.arg2"}
	if _, err := parseUnsetArgs(args); err != nil {
		t.Errorf("Error parsing unset args: %s", err.Error())
	}
}

// Test unsetValue function with invalid value type
func Test_unsetValue_InvalidValueType(t *testing.T) {
	regex := regexp.MustCompile(`^unable to unset configuration: variable type for key 'pingctl\.color' is not recognized$`)
	err := unsetValue("pingctl.color", "invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("unsetValue() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test unsetValue function with valid value type
func Test_unsetValue_ValidValueType(t *testing.T) {
	if err := unsetValue("pingctl.color", viperconfig.ENUM_BOOL); err != nil {
		t.Errorf("Error unsetting value: %s", err.Error())
	}
}
