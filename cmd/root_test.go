package cmd_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/testutils"
)

// Test Root Command Executes without issue
func TestRootCmd_Execute(t *testing.T) {
	err := testutils.ExecutePingctl()
	if err != nil {
		t.Errorf("Error executing root command: %s", err.Error())
	}
}

// Test Root Command Executes fails when provided additional arguments
func TestRootCmd_TooManyArgs(t *testing.T) {
	err := testutils.ExecutePingctl("arg1", "arg2", "arg3")
	if err == nil {
		t.Errorf("Expected error executing root command")
	}
}

// Test Root Command --help, -h flag
func TestRootCmd_HelpFlag(t *testing.T) {
	err := testutils.ExecutePingctl("--help")
	if err != nil {
		t.Errorf("Error executing root command: %s", err.Error())
	}

	err = testutils.ExecutePingctl("-h")
	if err != nil {
		t.Errorf("Error executing root command: %s", err.Error())
	}
}

// Test Root Command fails with invalid flag
func TestRootCmd_InvalidFlag(t *testing.T) {
	err := testutils.ExecutePingctl("--invalid")
	if err == nil {
		t.Errorf("Expected error executing root command")
	}
}

// Test Root Command Executes when provided the --version, -v flag
func TestRootCmd_VersionFlag(t *testing.T) {
	err := testutils.ExecutePingctl("--version")
	if err != nil {
		t.Errorf("Error executing root command: %s", err.Error())
	}

	err = testutils.ExecutePingctl("-v")
	if err != nil {
		t.Errorf("Error executing root command: %s", err.Error())
	}
}

// Test Root Command Executes when provided the --output flag
func TestRootCmd_OutputFlag(t *testing.T) {
	for _, outputFormat := range customtypes.OutputFormatValidValues() {
		err := testutils.ExecutePingctl("--output", outputFormat)
		if err != nil {
			t.Errorf("Error executing root command: %s", err.Error())
		}
	}
}

// Test Root Command fails when provided an invalid value for the --output flag
func TestRootCmd_InvalidOutputFlag(t *testing.T) {
	err := testutils.ExecutePingctl("--output", "invalid")
	if err == nil {
		t.Errorf("Expected error executing root command")
	}
}

// Test Root Command fails when provided no value for the --output flag
func TestRootCmd_NoValueOutputFlag(t *testing.T) {
	err := testutils.ExecutePingctl("--output")
	if err == nil {
		t.Errorf("Expected error executing root command")
	}
}

// Test Root Command Executes output does not change with output=text vs output=json
func TestRootCmd_OutputFlagTextVsJSON(t *testing.T) {
	textOutput, err := testutils.ExecutePingctlCaptureCobraOutput("--output=text")
	if err != nil {
		t.Errorf("Error executing root command: %s", err.Error())
	}

	jsonOutput, err := testutils.ExecutePingctlCaptureCobraOutput("--output=json")
	if err != nil {
		t.Errorf("Error executing root command: %s", err.Error())
	}

	if textOutput != jsonOutput {
		t.Errorf("Expected text and json output to be the same")
	}
}

// Test Root Command Executes when provided the --color flag
func TestRootCmd_ColorFlag(t *testing.T) {
	err := testutils.ExecutePingctl("--color=true")
	if err != nil {
		t.Errorf("Error executing root command: %s", err.Error())
	}

	err = testutils.ExecutePingctl("--color=false")
	if err != nil {
		t.Errorf("Error executing root command: %s", err.Error())
	}
}

// Test Root Command fails when provided an invalid value for the --color flag
func TestRootCmd_InvalidColorFlag(t *testing.T) {
	err := testutils.ExecutePingctl("--color=invalid")
	if err == nil {
		t.Errorf("Expected error executing root command")
	}
}

// Test Root Command Executes when provided the --config flag
func TestRootCmd_ConfigFlag(t *testing.T) {
	err := testutils.ExecutePingctl("--config", "config.yaml")
	if err != nil {
		t.Errorf("Error executing root command: %s", err.Error())
	}
}

// Test Root Command fails when provided no value for the --config flag
func TestRootCmd_NoValueConfigFlag(t *testing.T) {
	err := testutils.ExecutePingctl("--config")
	if err == nil {
		t.Errorf("Expected error executing root command")
	}
}
