package cmd_test

import (
	"regexp"
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
	regex := regexp.MustCompile(`^unknown command "arg1" for "pingctl"$`)
	err := testutils.ExecutePingctl("arg1")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Root command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
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
	regex := regexp.MustCompile(`^unknown flag: --invalid$`)
	err := testutils.ExecutePingctl("--invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Root command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
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
	regex := regexp.MustCompile(`^invalid argument "invalid" for "--output" flag: unrecognized Output Format: 'invalid'\. Must be one of: [a-z\s,]+$`)
	err := testutils.ExecutePingctl("--output", "invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Root command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}

// Test Root Command fails when provided no value for the --output flag
func TestRootCmd_NoValueOutputFlag(t *testing.T) {
	regex := regexp.MustCompile(`^flag needs an argument: --output$`)
	err := testutils.ExecutePingctl("--output")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Root command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
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
	regex := regexp.MustCompile(`^invalid argument "invalid" for "--color" flag: strconv\.ParseBool: parsing "invalid": invalid syntax$`)
	err := testutils.ExecutePingctl("--color=invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Root command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
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
	regex := regexp.MustCompile(`^flag needs an argument: --config$`)
	err := testutils.ExecutePingctl("--config")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Root command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}
