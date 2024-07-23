package platform_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Platform Command Executes without issue
func TestPlatformCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "platform")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Command fails when provided invalid flag
func TestPlatformCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Command --help, -h flag
func TestPlatformCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "platform", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingctl(t, "platform", "-h")
	testutils.CheckExpectedError(t, err, nil)
}
