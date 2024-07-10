package testutils_command

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
	"github.com/pingidentity/pingctl/internal/profiles"
)

var (
	configFileContents = fmt.Sprintf(`activeProfile: default
default:
    pingctl:
        color: true
        output: text
    pingone:
        export:
            environmentid: ""
        region: %s
        worker:
            clientid: %s
            clientsecret: %s
            environmentid: %s`,
		os.Getenv(profiles.RegionOption.EnvVar),
		os.Getenv(profiles.WorkerClientIDOption.EnvVar),
		os.Getenv(profiles.WorkerClientSecretOption.EnvVar),
		os.Getenv(profiles.WorkerEnvironmentIDOption.EnvVar))
)

func createConfigFile(t *testing.T) string {
	t.Helper()

	configFilepath := t.TempDir() + "/config.yaml"
	if err := os.WriteFile(configFilepath, []byte(configFileContents), 0600); err != nil {
		t.Fatalf("Failed to create config file: %s", err)
	}

	return configFilepath
}

// ExecutePingctl executes the pingctl command with the provided arguments
// and returns the error if any
func ExecutePingctl(t *testing.T, args ...string) (err error) {
	t.Helper()

	root := cmd.NewRootCommand()

	// Add config location to the root command
	configFilepath := createConfigFile(t)
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
	configFilepath := createConfigFile(t)
	args = append([]string{"--config", configFilepath}, args...)
	root.SetArgs(args)

	// Create byte buffer to capture output
	var stdout bytes.Buffer
	root.SetOut(&stdout)
	root.SetErr(&stdout)

	return stdout.String(), root.Execute()
}
