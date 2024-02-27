package cmd_test

import (
	"bytes"
	"os"
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
	executeErr := rootCmd.Execute()
	if executeErr != nil {
		logContent, err := os.ReadFile(os.Getenv("PINGCTL_LOG_PATH"))
		if err == nil {
			t.Logf("Captured Logs: %s", string(logContent[:]))
		}
		t.Fatal(executeErr)
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
	executeErr = rootCmd.Execute()
	if executeErr != nil {
		logContent, err := os.ReadFile(os.Getenv("PINGCTL_LOG_PATH"))
		if err == nil {
			t.Logf("Captured Logs: %s", string(logContent[:]))
		}
		t.Fatal(executeErr)
	}

	outputWithJSON := stdout.String()

	//expect both outputs to be the same
	if outputWithJSON != outputWithoutJSON {
		t.Errorf("Expected no change on output with json specified.\nOutput without JSON: %q\nOutput with JSON %q", outputWithoutJSON, outputWithJSON)
	}
}
