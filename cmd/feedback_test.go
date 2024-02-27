package cmd_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
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
		logContent, err := os.ReadFile(os.Getenv("PINGCTL_LOG_PATH"))
		if err == nil {
			t.Logf("Captured Logs: %s", string(logContent[:]))
		}
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}
}
