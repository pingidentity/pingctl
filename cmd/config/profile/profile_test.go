package profile_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils/testutils_command"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test Config Command Executes without issue
func TestConfigProfileCmd_Execute(t *testing.T) {
	err := testutils_command.ExecutePingctl(t, "config", "profile")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Config Command fails when provided invalid flag
func TestConfigProfileCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_command.ExecutePingctl(t, "config", "profile", "--invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Command --help, -h flag
func TestConfigProfileCmd_HelpFlag(t *testing.T) {
	err := testutils_command.ExecutePingctl(t, "config", "profile", "--help")
	testutils_helpers.CheckExpectedError(t, err, nil)

	err = testutils_command.ExecutePingctl(t, "config", "profile", "-h")
	testutils_helpers.CheckExpectedError(t, err, nil)
}
