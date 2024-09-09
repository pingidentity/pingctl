package testutils_cobra

import (
	"bytes"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
	"github.com/pingidentity/pingctl/internal/configuration"
	testutils_viper "github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// ExecutePingctl executes the pingctl command with the provided arguments
// and returns the error if any
func ExecutePingctl(t *testing.T, args ...string) (err error) {
	t.Helper()

	// Reset options for testing individual executions of pingctl
	configuration.InitAllOptions()

	root := cmd.NewRootCommand()

	// Add config location to the root command
	configFilepath := testutils_viper.CreateConfigFile(t)
	args = append([]string{"--config", configFilepath}, args...)
	root.SetArgs(args)

	return root.Execute()
}

// ExecutePingctlCaptureCobraOutput executes the pingctl command with
// the provided arguments and returns the output and error if any
// Note: The returned output will only contain cobra module specific output
// such as usage, help, and cobra errors
// It will NOT include internal/output/output.go output
// nor with it contain zerolog logs
func ExecutePingctlCaptureCobraOutput(t *testing.T, args ...string) (output string, err error) {
	t.Helper()

	root := cmd.NewRootCommand()

	// Add config location to the root command
	configFilepath := testutils_viper.CreateConfigFile(t)
	args = append([]string{"--config", configFilepath}, args...)
	root.SetArgs(args)

	// Create byte buffer to capture output
	var stdout bytes.Buffer
	root.SetOut(&stdout)
	root.SetErr(&stdout)

	return stdout.String(), root.Execute()
}
