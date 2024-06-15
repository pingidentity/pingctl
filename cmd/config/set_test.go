package config_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
	"github.com/pingidentity/pingctl/internal/testutils"
	"github.com/spf13/viper"
)

// Test Config Set Command Executes without issue
func TestConfigSetCmd_Execute(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "set", "pingctl.color=false"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		testutils.PrintLogs(t)
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}
}

// Test Config Set Command Fails when no arguments are provided
func TestConfigSetCmd_NoArgs(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "set"})

	// Execute the command
	err := rootCmd.Execute()
	if err == nil {
		testutils.PrintLogs(t)
		t.Fatalf("Expected error for providing no args, Captured StdOut: %q", stdout.String())
	}
}

// Test Config Set Command Fails when an invalid key is provided
func TestConfigSetCmd_InvalidKey(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "set", "pingctl.bogus=true"})

	// Execute the command
	err := rootCmd.Execute()
	if err == nil {
		testutils.PrintLogs(t)
		t.Fatalf("Expected error for providing invalid key, Captured StdOut: %q", stdout.String())
	}
}

// Test Config Set Command Fails when an invalid value type is provided
func TestConfigSetCmd_InvalidValueType(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "set", "pingctl.color=notboolean"})

	// Execute the command
	err := rootCmd.Execute()
	if err == nil {
		testutils.PrintLogs(t)
		t.Fatalf("Expected error for providing invalid value type, Captured StdOut: %q", stdout.String())
	}
}

// Test Config Set Command Fails when no value is provided
func TestConfigSetCmd_NoValueProvided(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "set", "pingone.export.environmentId="})

	// Execute the command
	err := rootCmd.Execute()
	if err == nil {
		testutils.PrintLogs(t)
		t.Fatalf("Expected error for providing no value, Captured StdOut: %q", stdout.String())
	}
}

// Test Config Set Command Executes normally when too many arguments are provided
func TestConfigSetCmd_TooManyArgs(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "set", "pingctl.color=false", "pingctl.bogus=true", "pingctl.bogus2=false"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		testutils.PrintLogs(t)
		t.Fatalf("Err: %s, Captured StdOut:\n%s", err.Error(), stdout.String())
	}
}

// Test Config Set Command for key 'pingone.worker.clientId' updates viper configuration
func TestConfigSetCmd_CheckViperConfig(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()
	viperKey := "pingone.worker.clientId"
	viperNewUUID := "12345678-1234-1234-1234-123456789012"
	viperOldValue := viper.GetString(viperKey)

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"config", "set", fmt.Sprintf("%s=%s", viperKey, viperNewUUID)})

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
	if viperNewValue != viperNewUUID {
		t.Fatalf("Expected viper configuration value to be %s, got %s", viperNewUUID, viperNewValue)
	}
}
