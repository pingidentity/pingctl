package platform_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils/testutils_command"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test Platform Command Executes without issue
func TestPlatformCmd_Execute(t *testing.T) {
	err := testutils_command.ExecutePingctl("platform")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Platform Command fails when provided invalid flag
func TestPlatformCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_command.ExecutePingctl("platform", "--invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Command --help, -h flag
func TestPlatformCmd_HelpFlag(t *testing.T) {
	err := testutils_command.ExecutePingctl("platform", "--help")
	testutils_helpers.CheckExpectedError(t, err, nil)

	err = testutils_command.ExecutePingctl("platform", "-h")
	testutils_helpers.CheckExpectedError(t, err, nil)
}
