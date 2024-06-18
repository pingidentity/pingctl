package config_internal

import (
	"os"
	"regexp"
	"testing"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/viperconfig"
	"github.com/spf13/viper"
)

// Test RunInternalConfigSet function
func Test_RunInternalConfigSet_NoArgs(t *testing.T) {
	regex := regexp.MustCompile(`^failed to set configuration: no 'key=value' assignment given in set command$`)
	err := RunInternalConfigSet([]string{})

	if !regex.MatchString(err.Error()) {
		t.Errorf("RunInternalConfigSet() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
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
	regex := regexp.MustCompile(`^failed to set configuration: key 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`)
	err := RunInternalConfigSet([]string{"pingctl.invalid=invalid"})

	if !regex.MatchString(err.Error()) {
		t.Errorf("RunInternalConfigSet() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
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
	regex := regexp.MustCompile(`^failed to set configuration: value for key 'pingone\.worker\.clientId' is empty. Use 'pingctl config unset pingone\.worker\.clientId' to unset the key$`)
	err := RunInternalConfigSet([]string{"pingone.worker.clientId="})

	if !regex.MatchString(err.Error()) {
		t.Errorf("RunInternalConfigSet() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test RunInternalConfigSet function with invalid value
func Test_RunInternalConfigSet_InvalidValue(t *testing.T) {
	regex := regexp.MustCompile(`^failed to set configuration: value for key 'pingone\.worker\.clientId' must be a valid UUID$`)
	err := RunInternalConfigSet([]string{"pingone.worker.clientId=invalid"})

	if !regex.MatchString(err.Error()) {
		t.Errorf("RunInternalConfigSet() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test RunInternalConfigSet function with invalid value type
func Test_RunInternalConfigSet_InvalidValueType(t *testing.T) {
	regex := regexp.MustCompile(`^failed to set configuration: value for key 'pingctl\.color' must be a boolean. Use 'true' or 'false'$`)
	err := RunInternalConfigSet([]string{"pingctl.color=notboolean"})

	if !regex.MatchString(err.Error()) {
		t.Errorf("RunInternalConfigSet() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test parseSetArgs() function with no args
func Test_parseSetArgs_NoArgs(t *testing.T) {
	regex := regexp.MustCompile(`^failed to set configuration: no 'key=value' assignment given in set command$`)
	_, _, err := parseSetArgs([]string{})

	if !regex.MatchString(err.Error()) {
		t.Errorf("parseSetArgs() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test parseSetArgs() function with invalid assignment format
func Test_parseSetArgs_InvalidAssignmentFormat(t *testing.T) {
	regex := regexp.MustCompile(`^failed to set configuration: invalid assignment format 'pingone\.worker\.clientId'. Expect 'key=value' format$`)
	_, _, err := parseSetArgs([]string{"pingone.worker.clientId"})

	if !regex.MatchString(err.Error()) {
		t.Errorf("parseSetArgs() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
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
	regex := regexp.MustCompile(`^failed to set configuration: value for key 'pingone\.worker\.clientId' must be a valid UUID$`)
	err := setValue("pingone.worker.clientId", "invalid", viperconfig.ENUM_ID)

	if !regex.MatchString(err.Error()) {
		t.Errorf("setValue() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test setValue() function with invalid value type
func Test_setValue_InvalidValueType(t *testing.T) {
	regex := regexp.MustCompile(`^unable to set configuration: variable type for key 'pingctl\.color' is not recognized$`)
	err := setValue("pingctl.color", "false", "invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("setValue() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
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
	regex := regexp.MustCompile(`^failed to set configuration: value for key 'pingctl\.color' must be a boolean. Use 'true' or 'false'$`)
	err := setBool("pingctl.color", "invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("setBool() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
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
	regex := regexp.MustCompile(`^failed to set configuration: value for key 'pingone\.worker\.clientId' must be a valid UUID$`)
	err := setUUID("pingone.worker.clientId", "invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("setUUID() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
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
	regex := regexp.MustCompile(`^failed to set configuration: unrecognized Output Format: 'invalid'\. Must be one of: [a-z\s,]+$`)
	err := setOutputFormat("pingctl.output", "invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("setOutputFormat() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
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
	regex := regexp.MustCompile(`^failed to set configuration: unrecognized PingOne Region: 'invalid'\. Must be one of: [A-Za-z\s,]+$`)
	err := setPingOneRegion("pingone.region", "invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("setPingOneRegion() error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}
