package config_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils"
)

// Test Config Get Command Executes without issue
func TestConfigGetCmd_Execute(t *testing.T) {
	err := testutils.ExecutePingctl("config", "get")
	if err != nil {
		t.Errorf("Error executing config get command: %s", err.Error())
	}
}

// Test Config Get Command Executes when provided a full key
func TestConfigGetCmd_FullKey(t *testing.T) {
	err := testutils.ExecutePingctl("config", "get", "pingone.worker.clientId")
	if err != nil {
		t.Errorf("Error executing config get command: %s", err.Error())
	}
}

// Test Config Get Command Executes when provided a partial key
func TestConfigGetCmd_PartialKey(t *testing.T) {
	err := testutils.ExecutePingctl("config", "get", "pingone")
	if err != nil {
		t.Errorf("Error executing config get command: %s", err.Error())
	}
}

// Test Config Get Command fails when provided an invalid key
func TestConfigGetCmd_InvalidKey(t *testing.T) {
	err := testutils.ExecutePingctl("config", "get", "pingctl.invalid")
	if err == nil {
		t.Errorf("Expected error executing config get command")
	}
}

// Test Config Get Command Executes normally when too many arguments are provided
func TestConfigGetCmd_TooManyArgs(t *testing.T) {
	err := testutils.ExecutePingctl("config", "get", "pingctl.color", "pingctl.arg1", "pingctl.arg2")
	if err != nil {
		t.Errorf("Error executing config get command: %s", err.Error())
	}
}
