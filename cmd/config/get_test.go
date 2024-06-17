package config_test

import (
	"bytes"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
	"github.com/pingidentity/pingctl/internal/testutils"
)

// Test Config Get Command Executes without issue
func TestConfigGetCmd_Execute(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "get"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		testutils.PrintLogs(t)
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}
}

// Test Config Get Command Executes when provided a full key
func TestConfigGetCmd_FullKey(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "get", "pingone.worker.clientId"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		testutils.PrintLogs(t)
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}
}

// Test Config Get Command Executes when provided a partial key
func TestConfigGetCmd_PartialKey(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "get", "pingone"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		testutils.PrintLogs(t)
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}
}

// Test Config Get Command fails when provided an invalid key
func TestConfigGetCmd_InvalidKey(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "get", "pingctl.bogus"})

	// Execute the command
	err := rootCmd.Execute()
	if err == nil {
		testutils.PrintLogs(t)
		t.Fatalf("Expected error for providing invalid key, Captured StdOut: %q", stdout.String())
	}
}

// Test Config Get Command Executes normally when too many arguments are provided
func TestConfigGetCmd_TooManyArgs(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "get", "pingctl.color", "pingctl.bogus", "pingctl.bogus2"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		testutils.PrintLogs(t)
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}
}
