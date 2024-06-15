package config_test

import (
	"bytes"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
	"github.com/pingidentity/pingctl/internal/testutils"
	"github.com/spf13/viper"
)

// Test Config Unset Command Executes without issue
func TestConfigUnsetCmd_Execute(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "unset", "pingctl.color"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		testutils.PrintLogs(t)
		t.Fatalf("Err: %q, Captured StdOut:\n%s", err, stdout.String())
	}
}

// Test Config Unset Command Fails when no arguments are provided
func TestConfigUnsetCmd_NoArgs(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "unset"})

	// Execute the command
	err := rootCmd.Execute()
	if err == nil {
		testutils.PrintLogs(t)
		t.Fatalf("Expected error for providing no args, Captured StdOut: %q", stdout.String())
	}
}

// Test Config Unset Command Fails when an invalid key is provided
func TestConfigUnsetCmd_InvalidKey(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "unset", "pingctl.bogus"})

	// Execute the command
	err := rootCmd.Execute()
	if err == nil {
		testutils.PrintLogs(t)
		t.Fatalf("Expected error for providing invalid key, Captured StdOut: %q", stdout.String())
	}
}

// Test Config Unset Command Executes normally when too many arguments are provided
func TestConfigUnsetCmd_TooManyArgs(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "unset", "pingctl.color", "pingctl.bogus", "pingctl.bogus2"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		testutils.PrintLogs(t)
		t.Fatalf("Err: %s, Captured StdOut:\n%s", err.Error(), stdout.String())
	}
}

// Test Config Unset Command for key 'pingone.worker.clientId' updates viper configuration
func TestConfigUnsetCmd_CheckViperConfig(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()
	viperKey := "pingone.worker.clientId"
	viperOldValue := viper.GetString(viperKey)

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "unset", viperKey})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		testutils.PrintLogs(t)
		t.Fatalf("Err: %s, Captured StdOut:\n%s", err.Error(), stdout.String())
	}

	// Check the viper configuration
	viperNewValue := viper.GetString(viperKey)
	if viperNewValue == viperOldValue {
		t.Fatalf("Expected viper configuration to be changed for key %s. Old: %s, New: %s", viperKey, viperOldValue, viperNewValue)
	}
	if viper.GetString(viperKey) != "" {
		t.Fatalf("Expected viper configuration to be empty string for key %s, got %s", viperKey, viperNewValue)
	}
}
