package config_internal

import (
	"os"
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
	"github.com/pingidentity/pingctl/internal/viperconfig"
	"github.com/spf13/viper"
)

// Test RunInternalConfigSet function
func Test_RunInternalConfigSet_NoArgs(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: no 'key=value' assignment given in set command$`
	err := RunInternalConfigSet([]string{})
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
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

	err := RunInternalConfigSet([]string{"pingone.worker.clientId=12345678-1234-1234-1234-123456789012"})
	testutils_helpers.CheckExpectedError(t, err, nil)

	// Clean up
	if err := os.RemoveAll(configDir); err != nil {
		t.Errorf("Error removing config directory: %s", err.Error())
	}
}

// Test RunInternalConfigSet function with invalid key
func Test_RunInternalConfigSet_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: key 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := RunInternalConfigSet([]string{"pingctl.invalid=true"})
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
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

	err := RunInternalConfigSet([]string{"pingone.worker.clientId=12345678-1234-1234-1234-123456789012", "pingone.worker.environmentId=12345678-1234-1234-1234-123456789012"})
	testutils_helpers.CheckExpectedError(t, err, nil)

	// Clean up
	if err := os.RemoveAll(configDir); err != nil {
		t.Errorf("Error removing config directory: %s", err.Error())
	}
}

// Test RunInternalConfigSet function with empty value
func Test_RunInternalConfigSet_EmptyValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingone\.worker\.clientId' is empty\. Use 'pingctl config unset pingone\.worker\.clientId' to unset the key$`
	err := RunInternalConfigSet([]string{"pingone.worker.clientId="})
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSet function with invalid value
func Test_RunInternalConfigSet_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingone\.worker\.clientId' must be a valid UUID$`
	err := RunInternalConfigSet([]string{"pingone.worker.clientId=invalid"})
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSet function with invalid value type
func Test_RunInternalConfigSet_InvalidValueType(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingctl\.color' must be a boolean. Use 'true' or 'false'$`
	err := RunInternalConfigSet([]string{"pingctl.color=notboolean"})
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test parseSetArgs() function with no args
func Test_parseSetArgs_NoArgs(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: no 'key=value' assignment given in set command$`
	_, _, err := parseSetArgs([]string{})
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test parseSetArgs() function with invalid assignment format
func Test_parseSetArgs_InvalidAssignmentFormat(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: invalid assignment format 'pingone\.worker\.clientId'. Expect 'key=value' format$`
	_, _, err := parseSetArgs([]string{"pingone.worker.clientId"})
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test parseSetArgs() function with valid assignment format
func Test_parseSetArgs_ValidAssignmentFormat(t *testing.T) {
	_, _, err := parseSetArgs([]string{"pingone.worker.clientId=12345678-1234-1234-1234-123456789012"})
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test parseSetArgs() function with too many args
func Test_parseSetArgs_TooManyArgs(t *testing.T) {
	_, _, err := parseSetArgs([]string{"pingone.worker.clientId=12345678-1234-1234-1234-123456789012", "pingone.worker.environmentId=12345678-1234-1234-1234-123456789012"})
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test setValue() function with valid value
func Test_setValue_ValidValue(t *testing.T) {
	err := setValue("pingctl.color", "false", viperconfig.ENUM_BOOL)
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test setValue() function with invalid value
func Test_setValue_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingone\.worker\.clientId' must be a valid UUID$`
	err := setValue("pingone.worker.clientId", "invalid", viperconfig.ENUM_ID)
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test setValue() function with invalid value type
func Test_setValue_InvalidValueType(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: variable type for key 'pingctl\.color' is not recognized$`
	err := setValue("pingctl.color", "false", "invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test setBool() function with valid value
func Test_setBool_ValidValue(t *testing.T) {
	err := setBool("pingctl.color", "false")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test setBool() function with invalid value
func Test_setBool_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingctl\.color' must be a boolean. Use 'true' or 'false'$`
	err := setBool("pingctl.color", "invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test setUUID() function with valid value
func Test_setUUID_ValidValue(t *testing.T) {
	err := setUUID("pingone.worker.clientId", "12345678-1234-1234-1234-123456789012")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test setUUID() function with invalid value
func Test_setUUID_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingone\.worker\.clientId' must be a valid UUID$`
	err := setUUID("pingone.worker.clientId", "invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test setOutputFormat() function with valid value
func Test_setOutputFormat_ValidValue(t *testing.T) {
	err := setOutputFormat("pingctl.output", "json")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test setOutputFormat() function with invalid value
func Test_setOutputFormat_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: unrecognized Output Format: 'invalid'\. Must be one of: [a-z\s,]+$`
	err := setOutputFormat("pingctl.output", "invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test setPingOneRegion() function with valid value
func Test_setPingOneRegion_ValidValue(t *testing.T) {
	err := setPingOneRegion("pingone.region", "AsiaPacific")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test setPingOneRegion() function with invalid value
func Test_setPingOneRegion_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: unrecognized PingOne Region: 'invalid'\. Must be one of: [A-Za-z\s,]+$`
	err := setPingOneRegion("pingone.region", "invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}
