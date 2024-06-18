package platform_test

import (
	"regexp"
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils"
)

// Test Platform Command Executes without issue
func TestPlatformCmd_Execute(t *testing.T) {
	err := testutils.ExecutePingctl("platform")
	if err != nil {
		t.Errorf("Error executing platform command: %s", err.Error())
	}
}

// Test Platform Command fails when provided invalid flag
func TestPlatformCmd_InvalidFlag(t *testing.T) {
	regex := regexp.MustCompile(`^unknown flag: --invalid$`)
	err := testutils.ExecutePingctl("platform", "--invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Platform command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test Platform Command --help, -h flag
func TestPlatformCmd_HelpFlag(t *testing.T) {
	err := testutils.ExecutePingctl("platform", "--help")
	if err != nil {
		t.Errorf("Error executing platform command: %s", err.Error())
	}

	err = testutils.ExecutePingctl("platform", "-h")
	if err != nil {
		t.Errorf("Error executing platform command: %s", err.Error())
	}
}
