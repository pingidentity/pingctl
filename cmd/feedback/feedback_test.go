package feedback_test

import (
	"regexp"
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils"
)

// Test Feedback Command Executes without issue
func TestFeedbackCmd_Execute(t *testing.T) {
	err := testutils.ExecutePingctl("feedback")
	if err != nil {
		t.Errorf("Error executing feedback command: %s", err.Error())
	}
}

// Test Feedback Command Executes without issue when provided additional arguments
func TestFeedbackCmd_TooManyArgs(t *testing.T) {
	err := testutils.ExecutePingctl("feedback", "arg1", "arg2", "arg3")
	if err != nil {
		t.Errorf("Error executing feedback command: %s", err.Error())
	}
}

// Test Feedback Command help flag
func TestFeedbackCmd_HelpFlag(t *testing.T) {
	err := testutils.ExecutePingctl("feedback", "--help")
	if err != nil {
		t.Errorf("Error executing feedback command: %s", err.Error())
	}
}

// Test Feedback Command fails with invalid flag
func TestFeedbackCmd_InvalidFlag(t *testing.T) {
	regex := regexp.MustCompile(`^unknown flag: --invalid$`)
	err := testutils.ExecutePingctl("feedback", "--invalid")

	if !regex.MatchString(err.Error()) {
		t.Errorf("Platform command error message did not match expected regex\n\nerror message: '%v'\n\nregex pattern %s", err, regex.String())
	}
}
