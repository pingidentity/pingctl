package config_internal

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
	"github.com/spf13/viper"
)

// Test RunInternalConfigGet function
func Test_RunInternalConfigGet_NoArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigGet("")
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigGet function with args that are set
func Test_RunInternalConfigGet_WithArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigGet(profiles.ColorOption.ViperKey)
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigGet function with args that are not set
func Test_RunInternalConfigGet_WithArgs_NotSet(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigGet(profiles.PingOneWorkerClientIDOption.ViperKey)
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigGet function with invalid key
func Test_RunInternalConfigGet_InvalidKey(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^unable to get configuration: value 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := RunInternalConfigGet("pingctl.invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test PrintConfig() function
func Example_printConfig() {
	// set viper configuration key-value for testing
	profileViper := viper.New()
	profileViper.Set(profiles.ColorOption.ViperKey, true)
	profileViper.Set(profiles.OutputOption.ViperKey, "text")
	profileViper.Set(profiles.PingOneRegionOption.ViperKey, "test-region")
	profileViper.Set(profiles.PingOneWorkerClientIDOption.ViperKey, "test-client-id")
	profileViper.Set(profiles.PingOneWorkerClientSecretOption.ViperKey, "test-client-secret")
	profileViper.Set(profiles.PingOneWorkerEnvironmentIDOption.ViperKey, "test-environment-id")
	profileViper.Set(profiles.PingOneExportEnvironmentIDOption.ViperKey, "test-export-environment-id")
	profiles.SetProfileViperWithViper(profileViper, "testProfile")

	_ = PrintConfig()

	// Output:
	// pingctl:
	//     color: true
	//     outputformat: text
	// pingone:
	//     export:
	//         environmentid: test-export-environment-id
	//     region: test-region
	//     worker:
	//         clientid: test-client-id
	//         clientsecret: test-client-secret
	//         environmentid: test-environment-id
}

// Test printConfigFromKey() function
func Example_printConfigFromKey() {
	// set viper configuration key-value for testing
	profileViper := viper.New()
	profileViper.Set(profiles.PingOneRegionOption.ViperKey, "test-region")
	profileViper.Set(profiles.OutputOption.ViperKey, "text")
	profiles.SetProfileViperWithViper(profileViper, "testProfile")

	_ = printConfigFromKey(profiles.PingOneRegionOption.ViperKey)
	_ = printConfigFromKey(profiles.OutputOption.ViperKey)

	// Output:
	// test-region
	//
	// text
}
