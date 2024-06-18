package config_test

import (
	"regexp"
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
	regex := regexp.MustCompile(`^unable to get configuration: value 'pingctl\.invalid' is not recognized as a valid configuration key\. Valid keys: [A-Za-z\.\s,]+$`)
	err := testutils.ExecutePingctl("config", "get", "pingctl.invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Config Get command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test Config Get Command Executes normally when too many arguments are provided
func TestConfigGetCmd_TooManyArgs(t *testing.T) {
	err := testutils.ExecutePingctl("config", "get", "pingctl.color", "pingctl.arg1", "pingctl.arg2")
	if err != nil {
		t.Errorf("Error executing config get command: %s", err.Error())
	}
}
