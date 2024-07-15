package cmd_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_command"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test Root Command Executes without issue
func TestRootCmd_Execute(t *testing.T) {
	err := testutils_command.ExecutePingctl(t)
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Root Command Executes fails when provided additional arguments
func TestRootCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^unknown command "arg1" for "pingctl"$`
	err := testutils_command.ExecutePingctl(t, "arg1")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command --help, -h flag
func TestRootCmd_HelpFlag(t *testing.T) {
	err := testutils_command.ExecutePingctl(t, "--help")
	testutils_helpers.CheckExpectedError(t, err, nil)

	err = testutils_command.ExecutePingctl(t, "-h")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Root Command fails with invalid flag
func TestRootCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_command.ExecutePingctl(t, "--invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command Executes when provided the --version, -v flag
func TestRootCmd_VersionFlag(t *testing.T) {
	err := testutils_command.ExecutePingctl(t, "--version")
	testutils_helpers.CheckExpectedError(t, err, nil)

	err = testutils_command.ExecutePingctl(t, "-v")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Root Command Executes when provided the --output flag
func TestRootCmd_OutputFlag(t *testing.T) {
	for _, outputFormat := range customtypes.OutputFormatValidValues() {
		err := testutils_command.ExecutePingctl(t, "--output", outputFormat)
		testutils_helpers.CheckExpectedError(t, err, nil)
	}
}

// Test Root Command fails when provided an invalid value for the --output flag
func TestRootCmd_InvalidOutputFlag(t *testing.T) {
	expectedErrorPattern := `^invalid argument "invalid" for "--output" flag: unrecognized Output Format: 'invalid'\. Must be one of: [a-z\s,]+$`
	err := testutils_command.ExecutePingctl(t, "--output", "invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command fails when provided no value for the --output flag
func TestRootCmd_NoValueOutputFlag(t *testing.T) {
	expectedErrorPattern := `^flag needs an argument: --output$`
	err := testutils_command.ExecutePingctl(t, "--output")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command Executes output does not change with output=text vs output=json
func TestRootCmd_OutputFlagTextVsJSON(t *testing.T) {
	textOutput, err := testutils_command.ExecutePingctlCaptureCobraOutput(t, "--output=text")
	testutils_helpers.CheckExpectedError(t, err, nil)

	jsonOutput, err := testutils_command.ExecutePingctlCaptureCobraOutput(t, "--output=json")
	testutils_helpers.CheckExpectedError(t, err, nil)

	if textOutput != jsonOutput {
		t.Errorf("Expected text and json output to be the same")
	}
}

// Test Root Command Executes when provided the --color flag
func TestRootCmd_ColorFlag(t *testing.T) {
	err := testutils_command.ExecutePingctl(t, "--color=true")
	testutils_helpers.CheckExpectedError(t, err, nil)

	err = testutils_command.ExecutePingctl(t, "--color=false")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Root Command fails when provided an invalid value for the --color flag
func TestRootCmd_InvalidColorFlag(t *testing.T) {
	expectedErrorPattern := `^invalid argument "invalid" for "--color" flag: strconv\.ParseBool: parsing "invalid": invalid syntax$`
	err := testutils_command.ExecutePingctl(t, "--color=invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Root Command Executes when provided the --config flag
func TestRootCmd_ConfigFlag(t *testing.T) {
	err := testutils_command.ExecutePingctl(t, "--config", "config.yaml")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Root Command fails when provided no value for the --config flag
func TestRootCmd_NoValueConfigFlag(t *testing.T) {
	expectedErrorPattern := `^flag needs an argument: --config$`
	err := testutils_command.ExecutePingctl(t, "--config")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}
