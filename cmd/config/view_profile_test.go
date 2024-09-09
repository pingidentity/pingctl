package config_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Config Set Command Executes without issue
func TestConfigViewProfileCmd_Execute(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "config", "view-profile")
	testutils.CheckExpectedError(t, err, nil)
}
