package config_test

import (
	"regexp"
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
	regex := regexp.MustCompile(`^unknown flag: --invalid$`)
	err := testutils.ExecutePingctl("config", "--invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Config command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
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
