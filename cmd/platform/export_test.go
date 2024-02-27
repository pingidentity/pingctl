package platform_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
	"github.com/pingidentity/pingctl/internal/testutils"
)

// Test Platform Export Command Executes without issue
func TestPlatformExportCmd_Execute(t *testing.T) {
	// Create the command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)

	rootCmd.SetArgs([]string{"platform", "export", "--output-directory", os.TempDir(), "--overwrite"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		testutils.PrintLogs(t)
		t.Fatalf("Export Command failed. Make sure to have PingOne env variables set if test is failing.\nErr: %q, Captured StdOut: %s", err, stdout.String())
	}
}
