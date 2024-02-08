package auth_test

import (
	"bytes"
	"testing"

	"github.com/pingidentity/pingctl/cmd/auth"
)

// Test Auth Login Command Executes without issue
func TestAuthLoginCmd_Execute(t *testing.T) {
	// Create the command
	loginCmd := auth.NewLoginCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	loginCmd.SetOut(&stdout)
	loginCmd.SetErr(&stdout)

	// Execute the command
	err := loginCmd.Execute()
	if err != nil {
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}
}

// Test Auth Logout Command Executes without issue
func TestAuthLogoutCmd_Execute(t *testing.T) {
	// Create the command
	logoutCmd := auth.NewLogoutCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	logoutCmd.SetOut(&stdout)
	logoutCmd.SetErr(&stdout)

	// Execute the command
	err := logoutCmd.Execute()
	if err != nil {
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}
}
