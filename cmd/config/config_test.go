package config_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils/testutils_command"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test Config Command Executes without issue
func TestConfigCmd_Execute(t *testing.T) {
	expectedErrorPattern := "" //No error expected
	err := testutils_command.ExecutePingctl("config")
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test Config Command fails when provided invalid flag
func TestConfigCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_command.ExecutePingctl("config", "--invalid")
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}

// Test Config Command --help, -h flag
func TestConfigCmd_HelpFlag(t *testing.T) {
	expectedErrorPattern := "" //No error expected
	err := testutils_command.ExecutePingctl("config", "--help")
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)

	err = testutils_command.ExecutePingctl("config", "-h")
	testutils_helpers.CheckExpectedError(t, err, expectedErrorPattern)
}
