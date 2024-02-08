package cmd_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/pingidentity/pingctl/cmd"
)

// Test Feedback Command Executes without issue
func TestFeedbackCmd_Execute(t *testing.T) {
	// Create the feedback command
	feedbackCmd := cmd.NewFeedbackCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	feedbackCmd.SetOut(&stdout)
	feedbackCmd.SetErr(&stdout)

	// Execute the feedback command
	err := feedbackCmd.Execute()
	if err != nil {
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}
}

func TestFeedbackCmd_Message(t *testing.T) {
	// Create the feedback command
	feedbackCmd := cmd.NewFeedbackCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	feedbackCmd.SetOut(&stdout)
	feedbackCmd.SetErr(&stdout)

	// Execute the feedback command
	err := feedbackCmd.Execute()
	if err != nil {
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}

	// Make sure output matches expected message
	if stdout.String() != (cmd.FeedbackMessage + "\n") {
		t.Errorf("Expected Feedback message output to equal %q\n Actual Output: %q", cmd.FeedbackMessage, stdout.String())
	}
}

func TestFeedbackCmd_ValidJSON(t *testing.T) {
	// Create the root command
	rootCmd := cmd.NewRootCommand()

	// Redirect stdout to a buffer to capture the output
	var stdout bytes.Buffer
	rootCmd.SetOut(&stdout)
	rootCmd.SetErr(&stdout)
	rootCmd.SetArgs([]string{"--output=json", "feedback"})

	// Execute the root command
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Err: %q, Captured StdOut: %q", err, stdout.String())
	}

	// Make sure output json is valid
	if !json.Valid(stdout.Bytes()) {
		t.Errorf("The output JSON %q is not valid json", stdout.String())
	}
}
