package platform_test

import (
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
	err := testutils.ExecutePingctl("platform", "--invalid")
	if err == nil {
		t.Errorf("Expected error executing platform command")
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
