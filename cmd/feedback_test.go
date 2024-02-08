package cmd_test

import (
	"bytes"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
)

// Test Root Command Executes without issue
func TestFeedbackCmd_Execute(t *testing.T) {
	// Create the root command
	feedbackCmd := cmd.NewFeedbackCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	feedbackCmd.SetOut(&stdout)
	feedbackCmd.SetErr(&stdout)

	// Execute the root command
	err := feedbackCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
}
