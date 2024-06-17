package config_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils"
)

// Test Config Command Executes without issue
func TestConfigCmd_Execute(t *testing.T) {
	err := testutils.ExecutePingctl("config")
	if err != nil {
		t.Errorf("Error executing config command: %s", err.Error())
	}
}

// Test Config Command fails when provided invalid flag
func TestConfigCmd_InvalidFlag(t *testing.T) {
	err := testutils.ExecutePingctl("config", "--invalid")
	if err == nil {
		t.Errorf("Expected error executing config command")
	}
}

// Test Config Command --help, -h flag
func TestConfigCmd_HelpFlag(t *testing.T) {
	err := testutils.ExecutePingctl("config", "--help")
	if err != nil {
		t.Errorf("Error executing config command: %s", err.Error())
	}

	err = testutils.ExecutePingctl("config", "-h")
	if err != nil {
		t.Errorf("Error executing config command: %s", err.Error())
	}
}
