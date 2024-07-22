package config_internal

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test RunInternalConfigSet function
func Test_RunInternalConfigSet_NoArgs(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: no 'key=value' assignment given in set command$`
	err := RunInternalConfigSet([]string{})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSet function with args
func Test_RunInternalConfigSet_WithArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	uuid := "12345678-1234-1234-1234-123456789012"
	err := RunInternalConfigSet([]string{fmt.Sprintf("%s=%s", profiles.WorkerClientIDOption.ViperKey, uuid)})
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigSet function with invalid key
func Test_RunInternalConfigSet_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: key 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := RunInternalConfigSet([]string{"pingctl.invalid=true"})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSet function with too many args
func Test_RunInternalConfigSet_TooManyArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	uuid := "12345678-1234-1234-1234-123456789012"
	err := RunInternalConfigSet([]string{fmt.Sprintf("%s=%s", profiles.WorkerClientIDOption.ViperKey, uuid), fmt.Sprintf("%s=%s", profiles.WorkerEnvironmentIDOption.ViperKey, uuid)})
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigSet function with empty value
func Test_RunInternalConfigSet_EmptyValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingone\.worker\.clientID' is empty\. Use 'pingctl config unset pingone\.worker\.clientID' to unset the key$`
	err := RunInternalConfigSet([]string{fmt.Sprintf("%s=", profiles.WorkerClientIDOption.ViperKey)})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSet function with invalid value
func Test_RunInternalConfigSet_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingone\.worker\.clientID' must be a valid UUID$`
	err := RunInternalConfigSet([]string{fmt.Sprintf("%s=invalid", profiles.WorkerClientIDOption.ViperKey)})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSet function with invalid value type
func Test_RunInternalConfigSet_InvalidValueType(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingctl\.color' must be a boolean. Use 'true' or 'false'$`
	err := RunInternalConfigSet([]string{fmt.Sprintf("%s=notboolean", profiles.ColorOption.ViperKey)})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test parseSetArgs() function with no args
func Test_parseSetArgs_NoArgs(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: no 'key=value' assignment given in set command$`
	_, _, err := parseSetArgs([]string{})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test parseSetArgs() function with invalid assignment format
func Test_parseSetArgs_InvalidAssignmentFormat(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: invalid assignment format 'pingone\.worker\.clientID'. Expect 'key=value' format$`
	_, _, err := parseSetArgs([]string{profiles.WorkerClientIDOption.ViperKey})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test parseSetArgs() function with valid assignment format
func Test_parseSetArgs_ValidAssignmentFormat(t *testing.T) {
	uuid := "12345678-1234-1234-1234-123456789012"
	_, _, err := parseSetArgs([]string{fmt.Sprintf("%s=%s", profiles.WorkerClientIDOption.ViperKey, uuid)})
	testutils.CheckExpectedError(t, err, nil)
}

// Test parseSetArgs() function with too many args
func Test_parseSetArgs_TooManyArgs(t *testing.T) {
	uuid := "12345678-1234-1234-1234-123456789012"
	_, _, err := parseSetArgs([]string{fmt.Sprintf("%s=%s", profiles.WorkerClientIDOption.ViperKey, uuid), fmt.Sprintf("%s=%s", profiles.WorkerEnvironmentIDOption.ViperKey, uuid)})
	testutils.CheckExpectedError(t, err, nil)
}

// Test setValue() function with valid value
func Test_setValue_ValidValue(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := setValue(profiles.ColorOption.ViperKey, "false", profiles.ENUM_BOOL)
	testutils.CheckExpectedError(t, err, nil)
}

// Test setValue() function with invalid value
func Test_setValue_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingone\.worker\.clientID' must be a valid UUID$`
	err := setValue(profiles.WorkerClientIDOption.ViperKey, "invalid", profiles.ENUM_ID)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test setValue() function with invalid value type
func Test_setValue_InvalidValueType(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: variable type for key 'pingctl\.color' is not recognized$`
	err := setValue(profiles.ColorOption.ViperKey, "false", "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test setBool() function with valid value
func Test_setBool_ValidValue(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := setBool(profiles.ColorOption.ViperKey, "false")
	testutils.CheckExpectedError(t, err, nil)
}

// Test setBool() function with invalid value
func Test_setBool_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingctl\.color' must be a boolean. Use 'true' or 'false'$`
	err := setBool(profiles.ColorOption.ViperKey, "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test setUUID() function with valid value
func Test_setUUID_ValidValue(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := setUUID(profiles.WorkerClientIDOption.ViperKey, "12345678-1234-1234-1234-123456789012")
	testutils.CheckExpectedError(t, err, nil)
}

// Test setUUID() function with invalid value
func Test_setUUID_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingone\.worker\.clientID' must be a valid UUID$`
	err := setUUID(profiles.WorkerClientIDOption.ViperKey, "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test setOutputFormat() function with valid value
func Test_setOutputFormat_ValidValue(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := setOutputFormat(profiles.OutputOption.ViperKey, "json")
	testutils.CheckExpectedError(t, err, nil)
}

// Test setOutputFormat() function with invalid value
func Test_setOutputFormat_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: unrecognized Output Format: 'invalid'\. Must be one of: [a-z\s,]+$`
	err := setOutputFormat(profiles.OutputOption.ViperKey, "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test setPingOneRegion() function with valid value
func Test_setPingOneRegion_ValidValue(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := setPingOneRegion(profiles.RegionOption.ViperKey, "AsiaPacific")
	testutils.CheckExpectedError(t, err, nil)
}

// Test setPingOneRegion() function with invalid value
func Test_setPingOneRegion_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: unrecognized PingOne Region: 'invalid'\. Must be one of: [A-Za-z\s,]+$`
	err := setPingOneRegion(profiles.RegionOption.ViperKey, "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
