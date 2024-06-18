package config_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils/testutils_command"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test Config Get Command Executes without issue
func TestConfigGetCmd_Execute(t *testing.T) {
	err := testutils_command.ExecutePingctl("config", "get")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Config Get Command Executes when provided a full key
func TestConfigGetCmd_FullKey(t *testing.T) {
	err := testutils_command.ExecutePingctl("config", "get", "pingone.worker.clientId")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Config Get Command Executes when provided a partial key
func TestConfigGetCmd_PartialKey(t *testing.T) {
	err := testutils_command.ExecutePingctl("config", "get", "pingone")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Config Get Command fails when provided an invalid key
func TestConfigGetCmd_InvalidKey(t *testing.T) {
	expectedErrorPattern := `^unable to get configuration: value 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`
	err := testutils_command.ExecutePingctl("config", "get", "pingctl.invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Get Command Executes normally when too many arguments are provided
func TestConfigGetCmd_TooManyArgs(t *testing.T) {
	err := testutils_command.ExecutePingctl("config", "get", "pingctl.color", "pingctl.output")
	testutils_helpers.CheckExpectedError(t, err, nil)
}
