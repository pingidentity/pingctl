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

	err := RunInternalConfigGet([]string{})
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigGet function with args that are set
func Test_RunInternalConfigGet_WithArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigGet([]string{profiles.ColorOption.ViperKey})
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigGet function with args that are not set
func Test_RunInternalConfigGet_WithArgs_NotSet(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigGet([]string{profiles.WorkerClientIDOption.ViperKey})
	testutils.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigGet function with invalid key
func Test_RunInternalConfigGet_InvalidKey(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedErrorPattern := `^unable to get configuration: value 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := RunInternalConfigGet([]string{"pingctl.invalid"})
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigGet function with too many args
func Test_RunInternalConfigGet_TooManyArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := RunInternalConfigGet([]string{profiles.WorkerClientIDOption.ViperKey, profiles.WorkerClientSecretOption.ViperKey})
	testutils.CheckExpectedError(t, err, nil)
}

// Test parseGetArgs function
func Test_parseGetArgs(t *testing.T) {
	_, err := parseGetArgs([]string{profiles.WorkerClientIDOption.ViperKey})
	testutils.CheckExpectedError(t, err, nil)
}

// Test parseGetArgs function with no args
func Test_parseGetArgs_NoArgs(t *testing.T) {
	testutils_viper.InitVipers(t)

	_, err := parseGetArgs([]string{})
	testutils.CheckExpectedError(t, err, nil)
}

// Test parseGetArgs function with too many args
func Test_parseGetArgs_TooManyArgs(t *testing.T) {
	_, err := parseGetArgs([]string{profiles.WorkerClientIDOption.ViperKey, profiles.WorkerClientSecretOption.ViperKey})
	testutils.CheckExpectedError(t, err, nil)
}

// Test PrintConfig() function
func Example_printConfig() {
	// set viper configuration key-value for testing
	profileViper := viper.New()
	profileViper.Set(profiles.ColorOption.ViperKey, true)
	profileViper.Set(profiles.OutputOption.ViperKey, "text")
	profileViper.Set(profiles.RegionOption.ViperKey, "test-region")
	profileViper.Set(profiles.WorkerClientIDOption.ViperKey, "test-client-id")
	profileViper.Set(profiles.WorkerClientSecretOption.ViperKey, "test-client-secret")
	profileViper.Set(profiles.WorkerEnvironmentIDOption.ViperKey, "test-environment-id")
	profileViper.Set(profiles.ExportEnvironmentIDOption.ViperKey, "test-export-environment-id")
	profiles.SetProfileViperWithViper(profileViper, "testProfile")

	_ = PrintConfig()

	// Output:
	// pingctl:
	//     color: true
	//     output: text
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
	profileViper.Set(profiles.RegionOption.ViperKey, "test-region")
	profileViper.Set(profiles.OutputOption.ViperKey, "text")
	profiles.SetProfileViperWithViper(profileViper, "testProfile")

	_ = printConfigFromKey(profiles.RegionOption.ViperKey)
	_ = printConfigFromKey(profiles.OutputOption.ViperKey)

	// Output:
	// test-region
	//
	// text
}
