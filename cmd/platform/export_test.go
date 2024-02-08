package platform_test

import (
	"bytes"
	"testing"

	"github.com/pingidentity/pingctl/cmd/platform"
)

// Test Platform Export Command Executes without issue
func TestPlatformExportCmd_Execute(t *testing.T) {
	// Create the command
	exportCmd := platform.NewExportCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	exportCmd.SetOut(&stdout)
	exportCmd.SetErr(&stdout)

	// Execute the command
	err := exportCmd.Execute()
	if err != nil {
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}
}
