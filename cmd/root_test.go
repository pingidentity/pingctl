package cmd_test

import (
	"bytes"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
)

// Test Root Command Executes without issue
func TestRootCmd_Execute(t *testing.T) {
	// Create the root command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	// Execute the root command
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}
}

// Test Root Command Executes output does not change with output=json
func TestRootCmd_JSONOutput(t *testing.T) {
	// Create the root command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	// Execute the root command
	err := rootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	outputWithoutJSON := stdout.String()

	// Create the root command
	rootCmd = cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	stdout = bytes.Buffer{}
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)
	rootCmd.SetArgs([]string{"--output=json"})

	// Execute the root command
	err = rootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}

	outputWithJSON := stdout.String()

	//expect both outputs to be the same
	if outputWithJSON != outputWithoutJSON {
		t.Errorf("Expected no change on output with json specified.\nOutput without JSON: %q\nOutput with JSON %q", outputWithoutJSON, outputWithJSON)
	}
}
