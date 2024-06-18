package testutils_command

import (
	"bytes"

	"github.com/pingidentity/pingctl/cmd"
)

// ExecutePingctl executes the pingctl command with the provided arguments
// and returns the error if any
func ExecutePingctl(args ...string) (err error) {
	root := cmd.NewRootCommand()
	root.SetArgs(args)

	return root.Execute()
}

// ExecutePingctlCaptureCobraOutput executes the pingctl command with
// the provided arguments and returns the output and error if any
// Note: The returned output will only contain cobra module specific output
// such as usage, help, and cobra errors
// It will NOT include internal/output/output.go output
// nor with it contain zerolog logs
func ExecutePingctlCaptureCobraOutput(args ...string) (output string, err error) {
	root := cmd.NewRootCommand()
	root.SetArgs(args)

	// Create byte buffer to capture output
	var stdout bytes.Buffer
	root.SetOut(&stdout)
	root.SetErr(&stdout)

	return stdout.String(), root.Execute()
}
