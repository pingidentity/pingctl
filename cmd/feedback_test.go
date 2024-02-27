package cmd_test

import (
	"bytes"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
	"github.com/pingidentity/pingctl/internal/testutils"
)

// Test Feedback Command Executes without issue
func TestFeedbackCmd_Execute(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"feedback"})

	// Execute the root command
	err := rootCmd.Execute()
	if err != nil {
		testutils.PrintLogs(t)
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}
}
