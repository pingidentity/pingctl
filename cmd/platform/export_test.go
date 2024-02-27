package platform_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
)

// Test Platform Export Command Executes without issue
func TestPlatformExportCmd_Execute(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"platform", "export", "--output-directory", os.Getenv("TMPDIR"), "--overwrite"})

	// Execute the command
	executeErr := rootCmd.Execute()
	if executeErr != nil {
		logContent, err := os.ReadFile(os.Getenv("PINGCTL_LOG_PATH"))
		if err == nil {
			t.Logf("Captured Logs: %q", string(logContent[:]))
		}
		t.Fatalf("Export Command failed. Make sure to have PingOne env variables set if test is failing.\nErr: %q, Captured StdOut: %q", executeErr, stdout.String())
	}
}
