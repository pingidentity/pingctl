package feedback_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Feedback Command Executes without issue
func TestFeedbackCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "feedback")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Feedback Command Executes without issue when provided additional arguments
func TestFeedbackCmd_TooManyArgs(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "feedback", "arg1", "arg2", "arg3")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Feedback Command help flag
func TestFeedbackCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "feedback", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingctl(t, "feedback", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Feedback Command fails with invalid flag
func TestFeedbackCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingctl(t, "feedback", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
