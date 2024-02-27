package testutils

import (
	"os"
	"testing"
)

// Utility method to print log file if present.
func PrintLogs(t *testing.T) {
	t.Helper()

	logContent, err := os.ReadFile(os.Getenv("PINGCTL_LOG_PATH"))
	if err == nil {
		t.Logf("Captured Logs: %s", string(logContent[:]))
	}
}
