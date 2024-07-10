package config_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils/testutils_command"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test Config Command Executes without issue
func TestConfigCmd_Execute(t *testing.T) {
	err := testutils_command.ExecutePingctl(t, "config")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Config Command fails when provided invalid flag
func TestConfigCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_command.ExecutePingctl(t, "config", "--invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Command --help, -h flag
func TestConfigCmd_HelpFlag(t *testing.T) {
	err := testutils_command.ExecutePingctl(t, "config", "--help")
	testutils_helpers.CheckExpectedError(t, err, nil)

	err = testutils_command.ExecutePingctl(t, "config", "-h")
	testutils_helpers.CheckExpectedError(t, err, nil)
}
