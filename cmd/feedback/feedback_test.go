package feedback_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils/testutils_command"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test Feedback Command Executes without issue
func TestFeedbackCmd_Execute(t *testing.T) {
	err := testutils_command.ExecutePingctl(t, "feedback")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Feedback Command Executes without issue when provided additional arguments
func TestFeedbackCmd_TooManyArgs(t *testing.T) {
	err := testutils_command.ExecutePingctl(t, "feedback", "arg1", "arg2", "arg3")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Feedback Command help flag
func TestFeedbackCmd_HelpFlag(t *testing.T) {
	err := testutils_command.ExecutePingctl(t, "feedback", "--help")
	testutils_helpers.CheckExpectedError(t, err, nil)

	err = testutils_command.ExecutePingctl(t, "feedback", "-h")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Feedback Command fails with invalid flag
func TestFeedbackCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_command.ExecutePingctl(t, "feedback", "--invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}
