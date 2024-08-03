package config_internal

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test RunInternalConfigSet function with args
func Test_RunInternalConfigSet_WithArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	uuid, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatalf("failed to generate UUID: %v", err)
	}

	err = RunInternalConfigSet(fmt.Sprintf("%s=%s", profiles.PingOneWorkerClientIDOption.ViperKey, uuid))
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigSet function with invalid key
func Test_RunInternalConfigSet_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: key 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := RunInternalConfigSet("pingctl.invalid=true")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSet function with empty value
func Test_RunInternalConfigSet_EmptyValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingone\.worker\.clientID' is empty\. Use 'pingctl config unset pingone\.worker\.clientID' to unset the key$`
	err := RunInternalConfigSet(fmt.Sprintf("%s=", profiles.PingOneWorkerClientIDOption.ViperKey))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSet function with invalid value
func Test_RunInternalConfigSet_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingone\.worker\.clientID' must be a valid UUID$`
	err := RunInternalConfigSet(fmt.Sprintf("%s=invalid", profiles.PingOneWorkerClientIDOption.ViperKey))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigSet function with invalid value type
func Test_RunInternalConfigSet_InvalidValueType(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingctl\.color' must be a boolean. Use 'true' or 'false'$`
	err := RunInternalConfigSet(fmt.Sprintf("%s=notboolean", profiles.ColorOption.ViperKey))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
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
	err := setValue(profiles.PingOneWorkerClientIDOption.ViperKey, "invalid", profiles.ENUM_ID)
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

	uuid, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatalf("failed to generate UUID: %v", err)
	}

	err = setUUID(profiles.PingOneWorkerClientIDOption.ViperKey, uuid)
	testutils.CheckExpectedError(t, err, nil)
}

// Test setUUID() function with invalid value
func Test_setUUID_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: value for key 'pingone\.worker\.clientID' must be a valid UUID$`
	err := setUUID(profiles.PingOneWorkerClientIDOption.ViperKey, "invalid")
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

	err := setPingOneRegion(profiles.PingOneRegionOption.ViperKey, "AsiaPacific")
	testutils.CheckExpectedError(t, err, nil)
}

// Test setPingOneRegion() function with invalid value
func Test_setPingOneRegion_InvalidValue(t *testing.T) {
	expectedErrorPattern := `^failed to set configuration: unrecognized PingOne Region: 'invalid'\. Must be one of: [A-Za-z\s,]+$`
	err := setPingOneRegion(profiles.PingOneRegionOption.ViperKey, "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
