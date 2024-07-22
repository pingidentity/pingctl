package config_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Config Get Command Executes without issue
func TestConfigGetCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "get")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Get Command Executes when provided a full key
func TestConfigGetCmd_FullKey(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "get", profiles.WorkerClientIDOption.ViperKey)
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Get Command Executes when provided a partial key
func TestConfigGetCmd_PartialKey(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "get", "pingone")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Get Command fails when provided an invalid key
func TestConfigGetCmd_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^unable to get configuration: value 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := testutils_cobra.ExecutePingctl(t, "config", "get", "pingctl.invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Get Command Executes normally when too many arguments are provided
func TestConfigGetCmd_TooManyArgs(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "get", profiles.ColorOption.ViperKey, profiles.OutputOption.ViperKey)
	testutils.CheckExpectedError(t, err, nil)
}
