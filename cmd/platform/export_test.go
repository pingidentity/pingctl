package platform_test

import (
	"os"
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils/testutils_command"
	"github.com/pingidentity/pingctl/internal/testutils/testutils_helpers"
)

// Test Platform Export Command Executes without issue
func TestPlatformExportCmd_Execute(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_command.ExecutePingctl(t, "platform", "export", "--output-directory", outputDir, "--overwrite")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command fails when provided invalid flag
func TestPlatformExportCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_command.ExecutePingctl(t, "platform", "export", "--invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --help, -h flag
func TestPlatformExportCmd_HelpFlag(t *testing.T) {
	err := testutils_command.ExecutePingctl(t, "platform", "export", "--help")
	testutils_helpers.CheckExpectedError(t, err, nil)

	err = testutils_command.ExecutePingctl(t, "platform", "export", "-h")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --service flag
func TestPlatformExportCmd_ServiceFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_command.ExecutePingctl(t, "platform", "export", "--output-directory", outputDir, "--overwrite", "--service", "pingone-protect")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --service flag with invalid service
func TestPlatformExportCmd_ServiceFlagInvalidService(t *testing.T) {
	expectedErrorPattern := `^invalid argument "invalid" for "--service" flag: unrecognized service 'invalid'\. Must be one of: [a-z-\s,]+$`
	err := testutils_command.ExecutePingctl(t, "platform", "export", "--service", "invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --export-format flag
func TestPlatformExportCmd_ExportFormatFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_command.ExecutePingctl(t, "platform", "export", "--output-directory", outputDir, "--export-format", "HCL", "--overwrite", "--service", "pingone-protect")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --export-format flag with invalid format
func TestPlatformExportCmd_ExportFormatFlagInvalidFormat(t *testing.T) {
	expectedErrorPattern := `^invalid argument "invalid" for "--export-format" flag: unrecognized export format 'invalid'\. Must be one of: [A-Z]+$`
	err := testutils_command.ExecutePingctl(t, "platform", "export", "--export-format", "invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --output-directory flag
func TestPlatformExportCmd_OutputDirectoryFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_command.ExecutePingctl(t, "platform", "export", "--output-directory", outputDir, "--overwrite", "--service", "pingone-protect")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --output-directory flag with invalid directory
func TestPlatformExportCmd_OutputDirectoryFlagInvalidDirectory(t *testing.T) {
	expectedErrorPattern := `^failed to create 'platform export' output directory '\/invalid': mkdir \/invalid: .+$`
	err := testutils_command.ExecutePingctl(t, "platform", "export", "--output-directory", "/invalid")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --overwrite flag
func TestPlatformExportCmd_OverwriteFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_command.ExecutePingctl(t, "platform", "export", "--output-directory", outputDir, "--overwrite", "--service", "pingone-protect")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --overwrite flag false with existing directory
// where the directory already contains a file
func TestPlatformExportCmd_OverwriteFlagFalseWithExistingDirectory(t *testing.T) {
	outputDir := t.TempDir()

	_, err := os.Create(outputDir + "/file")
	if err != nil {
		t.Errorf("Error creating file in output directory: %v", err)
	}

	expectedErrorPattern := `^'platform export' output directory '[A-Za-z0-9_\-\/]+' is not empty\. Use --overwrite to overwrite existing export data$`
	err = testutils_command.ExecutePingctl(t, "platform", "export", "--output-directory", outputDir, "--service", "pingone-protect", "--overwrite=false")
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --overwrite flag true with existing directory
// where the directory already contains a file
func TestPlatformExportCmd_OverwriteFlagTrueWithExistingDirectory(t *testing.T) {
	outputDir := t.TempDir()

	_, err := os.Create(outputDir + "/file")
	if err != nil {
		t.Errorf("Error creating file in output directory: %v", err)
	}

	err = testutils_command.ExecutePingctl(t, "platform", "export", "--output-directory", outputDir, "--service", "pingone-protect", "--overwrite")
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command with
// --pingone-worker-environment-id flag
// --pingone-worker-client-id flag
// --pingone-worker-client-secret flag
// --pingone-region flag
func TestPlatformExportCmd_PingOneWorkerEnvironmentIdFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_command.ExecutePingctl(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--service", "pingone-protect",
		"--pingone-worker-environment-id", os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID"),
		"--pingone-worker-client-id", os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_ID"),
		"--pingone-worker-client-secret", os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_SECRET"),
		"--pingone-region", os.Getenv("PINGCTL_PINGONE_REGION"))
	testutils_helpers.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command fails when not provided required pingone flags together
func TestPlatformExportCmd_PingOneWorkerEnvironmentIdFlagRequiredTogether(t *testing.T) {
	expectedErrorPattern := `^if any flags in the group \[pingone-worker-environment-id pingone-worker-client-id pingone-worker-client-secret pingone-region] are set they must all be set; missing \[pingone-region pingone-worker-client-id pingone-worker-client-secret]$`
	err := testutils_command.ExecutePingctl(t, "platform", "export",
		"--pingone-worker-environment-id", os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID"))
	testutils_helpers.CheckExpectedError(t, err, &expectedErrorPattern)
}
