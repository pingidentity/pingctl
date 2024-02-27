package platform_test

import (
	"bytes"
	"os"
	"path/filepath"
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
	err := os.Setenv("PINGCTL_LOG_LEVEL", "DEBUG")
	if err != nil {
		t.Fatalf(err.Error())
	}

	logFilepath := filepath.Join(os.Getenv("TMPDIR"), filepath.Base("export_test.log"))
	err = os.Setenv("PINGCTL_LOG_PATH", logFilepath)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// Execute the command
	err = rootCmd.Execute()
	if err != nil {
		logContents, err := os.ReadFile(logFilepath)
		if err == nil {
			t.Logf("Captured Logs: %q", logContents)
		}
		t.Fatalf("Export Command failed. Make sure to have PingOne env variables set if test is failing.\nErr: %q, Captured StdOut: %q", err, stdout.String())
	}
}
