package profile_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Config Profile List Command Executes without issue
func TestConfigProfileListCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "list")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config Profile List Command fails when provided too many arguments
func TestConfigProfileListCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute '.*': command accepts \d arg\(s\), received 1$`
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "list", "extra-arg")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config Profile List Command executes successfully when provided the help flag
func TestConfigProfileListCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "profile", "list", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingctl(t, "config", "profile", "list", "-h")
	testutils.CheckExpectedError(t, err, nil)
}
