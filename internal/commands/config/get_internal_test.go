package config_internal

import (
	"slices"
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
	"github.com/spf13/viper"
)

// Test RunInternalConfigGet function
func Test_RunInternalConfigGet_NoArgs(t *testing.T) {
	err := RunInternalConfigGet([]string{})
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigGet function with args that are set
func Test_RunInternalConfigGet_WithArgs(t *testing.T) {
	viper.Set("pingone.worker.clientId", "test-client-id")

	err := RunInternalConfigGet([]string{"pingone.worker.clientId"})
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigGet function with args that are not set
func Test_RunInternalConfigGet_WithArgs_NotSet(t *testing.T) {
	err := RunInternalConfigGet([]string{"pingone.worker.clientId"})
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test RunInternalConfigGet function with invalid key
func Test_RunInternalConfigGet_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^unable to get configuration: value 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := RunInternalConfigGet([]string{"pingctl.invalid"})
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test RunInternalConfigGet function with too many args
func Test_RunInternalConfigGet_TooManyArgs(t *testing.T) {
	err := RunInternalConfigGet([]string{"pingone.worker.clientId", "pingone.worker.clientSecret"})
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test parseGetArgs function
func Test_parseGetArgs(t *testing.T) {
	_, err := parseGetArgs([]string{"pingone.worker.clientId"})
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test parseGetArgs function with no args
func Test_parseGetArgs_NoArgs(t *testing.T) {
	_, err := parseGetArgs([]string{})
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test parseGetArgs function with too many args
func Test_parseGetArgs_TooManyArgs(t *testing.T) {
	_, err := parseGetArgs([]string{"pingone.worker.clientId", "pingone.worker.clientSecret"})
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test PrintConfig() function
func ExamplePrintConfig() {
	// set viper configuration key-value for testing
	viper.Set("pingctl.color", true)
	viper.Set("pingctl.output", "text")
	viper.Set("pingone.region", "test-region")
	viper.Set("pingone.worker.clientId", "test-client-id")
	viper.Set("pingone.worker.clientSecret", "test-client-secret")
	viper.Set("pingone.worker.environmentId", "test-environment-id")
	viper.Set("pingone.export.environmentId", "test-export-environment-id")

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
	viper.Set("pingone.region", "test-region")
	viper.Set("pingctl.output", "text")

	_ = printConfigFromKey("pingone.region")
	_ = printConfigFromKey("pingctl.output")

	// Output:
	// test-region
	//
	// text
}

// Test getValidGetKeys() function
func Test_getValidGetKeys(t *testing.T) {
	expectedNumKeys := 11

	allKeys := getValidGetKeys()

	if allKeys == nil {
		t.Errorf("Error getting valid keys, expected non-nil value")
	}

	if len(allKeys) == 0 {
		t.Errorf("Error getting valid keys, expected non-empty value")
	}

	if len(allKeys) != expectedNumKeys {
		t.Errorf("Error getting valid keys, expected %d keys, got %d", expectedNumKeys, len(allKeys))
	}

	if !slices.Contains(allKeys, "pingone.worker") {
		t.Errorf("Error getting valid keys, expected 'pingone.worker' key")
	}
}
