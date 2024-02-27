package auth_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
)

// Test Auth Login Command Executes without issue
func TestAuthLoginCmd_Execute(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"auth", "login"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		logContent, err := os.ReadFile(os.Getenv("PINGCTL_LOG_PATH"))
		if err == nil {
			t.Logf("Captured Logs: %s", string(logContent[:]))
		}
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}
}

// Test Auth Logout Command Executes without issue
func TestAuthLogoutCmd_Execute(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"auth", "logout"})

	// Execute the command
	executeErr := rootCmd.Execute()
	if executeErr != nil {
		logContent, err := os.ReadFile(os.Getenv("PINGCTL_LOG_PATH"))
		if err == nil {
			t.Logf("Captured Logs: %s", string(logContent[:]))
		}
		t.Fatalf("Err: %q, Captured StdOut: %q", executeErr, stdout.String())
	}
}
