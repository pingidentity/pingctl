package config_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Config list-profiles Command Executes without issue
func TestConfigListProfilesCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "list-profiles")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Config list-profiles Command fails when provided too many arguments
func TestConfigListProfilesCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingctl config list-profiles': command accepts 0 arg\(s\), received 1$`
	err := testutils_cobra.ExecutePingctl(t, "config", "list-profiles", "extra-arg")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config list-profiles Command fails when provided an invalid flag
func TestConfigListProfilesCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingctl(t, "config", "list-profiles", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Config list-profiles Command --help, -h flag
func TestConfigListProfilesCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "list-profiles", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingctl(t, "config", "list-profiles", "-h")
	testutils.CheckExpectedError(t, err, nil)
}
