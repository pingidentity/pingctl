package platform_test

import (
	"os"
	"testing"

	"github.com/pingidentity/pingctl/internal/testutils"
)

// Test Platform Export Command Executes without issue
func TestPlatformExportCmd_Execute(t *testing.T) {
	outputDir := os.TempDir() + "/pingctlTestPlatformExportExecute"
	err := testutils.ExecutePingctl("platform", "export", "--output-directory", outputDir, "--overwrite")
	if err != nil {
		t.Errorf("Error executing platform export command: %v", err)
	}

	// Empty output directory
	err = os.RemoveAll(outputDir)
	if err != nil {
		t.Errorf("Error removing output directory: %v", err)
	}
}

// Test Platform Export Command fails when provided invalid flag
func TestPlatformExportCmd_InvalidFlag(t *testing.T) {
	err := testutils.ExecutePingctl("platform", "export", "--invalid")
	if err == nil {
		t.Errorf("Expected error executing platform export command")
	}
}

// Test Platform Export Command --help, -h flag
func TestPlatformExportCmd_HelpFlag(t *testing.T) {
	err := testutils.ExecutePingctl("platform", "export", "--help")
	if err != nil {
		t.Errorf("Error executing platform export command: %v", err)
	}

	err = testutils.ExecutePingctl("platform", "export", "-h")
	if err != nil {
		t.Errorf("Error executing platform export command: %v", err)
	}
}

// Test Platform Export Command --service flag
func TestPlatformExportCmd_ServiceFlag(t *testing.T) {
	outputDir := os.TempDir() + "/pingctlTestPlatformExportServiceFlag"
	err := testutils.ExecutePingctl("platform", "export", "--output-directory", outputDir, "--service", "pingone-protect", "--overwrite")
	if err != nil {
		t.Errorf("Error executing platform export command: %v", err)
	}

	// Empty output directory
	err = os.RemoveAll(outputDir)
	if err != nil {
		t.Errorf("Error removing output directory: %v", err)
	}
}

// Test Platform Export Command --service flag with invalid service
func TestPlatformExportCmd_ServiceFlagInvalidService(t *testing.T) {
	err := testutils.ExecutePingctl("platform", "export", "--service", "invalid")
	if err == nil {
		t.Errorf("Expected error executing platform export command")
	}
}

// Test Platform Export Command --export-format flag
func TestPlatformExportCmd_ExportFormatFlag(t *testing.T) {
	outputDir := os.TempDir() + "/pingctlTestPlatformExportExportFormatFlag"
	err := testutils.ExecutePingctl("platform", "export", "--output-directory", outputDir, "--export-format", "HCL", "--overwrite", "--service", "pingone-protect")
	if err != nil {
		t.Errorf("Error executing platform export command: %v", err)
	}

	// Empty output directory
	err = os.RemoveAll(outputDir)
	if err != nil {
		t.Errorf("Error removing output directory: %v", err)
	}
}

// Test Platform Export Command --export-format flag with invalid format
func TestPlatformExportCmd_ExportFormatFlagInvalidFormat(t *testing.T) {
	err := testutils.ExecutePingctl("platform", "export", "--export-format", "invalid")
	if err == nil {
		t.Errorf("Expected error executing platform export command")
	}
}

// Test Platform Export Command --output-directory flag
func TestPlatformExportCmd_OutputDirectoryFlag(t *testing.T) {
	outputDir := os.TempDir() + "/pingctlTestPlatformExportOutputDirectoryFlag"
	err := testutils.ExecutePingctl("platform", "export", "--output-directory", outputDir, "--overwrite", "--service", "pingone-protect")
	if err != nil {
		t.Errorf("Error executing platform export command: %v", err)
	}

	// Empty output directory
	err = os.RemoveAll(outputDir)
	if err != nil {
		t.Errorf("Error removing output directory: %v", err)
	}
}

// Test Platform Export Command --output-directory flag with invalid directory
func TestPlatformExportCmd_OutputDirectoryFlagInvalidDirectory(t *testing.T) {
	err := testutils.ExecutePingctl("platform", "export", "--output-directory", "/invalid")
	if err == nil {
		t.Errorf("Expected error executing platform export command")
	}
}

// Test Platform Export Command --overwrite flag
func TestPlatformExportCmd_OverwriteFlag(t *testing.T) {
	outputDir := os.TempDir() + "/pingctlTestPlatformExportOverwriteFlag"
	err := testutils.ExecutePingctl("platform", "export", "--output-directory", outputDir, "--overwrite", "--service", "pingone-protect")
	if err != nil {
		t.Errorf("Error executing platform export command: %v", err)
	}

	// Empty output directory
	err = os.RemoveAll(outputDir)
	if err != nil {
		t.Errorf("Error removing output directory: %v", err)
	}
}

// Test Platform Export Command --overwrite flag false with existing directory
// where the directory already contains a file
func TestPlatformExportCmd_OverwriteFlagFalseWithExistingDirectory(t *testing.T) {
	outputDir := os.TempDir() + "/pingctlTestPlatformExportOverwriteFlagFalseWithExistingDirectory"
	err := os.Mkdir(outputDir, 0755)
	if err != nil {
		t.Errorf("Error creating output directory: %v", err)
	}
	_, err = os.Create(outputDir + "/file")
	if err != nil {
		t.Errorf("Error creating file in output directory: %v", err)
	}

	err = testutils.ExecutePingctl("platform", "export", "--output-directory", outputDir, "--service", "pingone-protect", "--overwrite=false")
	if err == nil {
		t.Errorf("Expected error executing platform export command")
	}

	// Empty output directory
	err = os.RemoveAll(outputDir)
	if err != nil {
		t.Errorf("Error removing output directory: %v", err)
	}
}

// Test Platform Export Command --overwrite flag true with existing directory
// where the directory already contains a file
func TestPlatformExportCmd_OverwriteFlagTrueWithExistingDirectory(t *testing.T) {
	outputDir := os.TempDir() + "/pingctlTestPlatformExportOverwriteFlagTrueWithExistingDirectory"
	err := os.Mkdir(outputDir, 0755)
	if err != nil {
		t.Errorf("Error creating output directory: %v", err)
	}
	_, err = os.Create(outputDir + "/file")
	if err != nil {
		t.Errorf("Error creating file in output directory: %v", err)
	}

	err = testutils.ExecutePingctl("platform", "export", "--output-directory", outputDir, "--service", "pingone-protect", "--overwrite")
	if err != nil {
		t.Errorf("Error executing platform export command: %v", err)
	}

	// Empty output directory
	err = os.RemoveAll(outputDir)
	if err != nil {
		t.Errorf("Error removing output directory: %v", err)
	}
}

// Test Platform Export Command with
// --pingone-worker-environment-id flag
// --pingone-worker-client-id flag
// --pingone-worker-client-secret flag
// --pingone-region flag
func TestPlatformExportCmd_PingOneWorkerEnvironmentIdFlag(t *testing.T) {
	outputDir := os.TempDir() + "/pingctlTestPlatformExportPingOneWorkerEnvironmentIdFlag"
	err := testutils.ExecutePingctl("platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--service", "pingone-protect",
		"--pingone-worker-environment-id", os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID"),
		"--pingone-worker-client-id", os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_ID"),
		"--pingone-worker-client-secret", os.Getenv("PINGCTL_PINGONE_WORKER_CLIENT_SECRET"),
		"--pingone-region", os.Getenv("PINGCTL_PINGONE_REGION"))
	if err != nil {
		t.Errorf("Error executing platform export command: %v", err)
	}

	// Empty output directory
	err = os.RemoveAll(outputDir)
	if err != nil {
		t.Errorf("Error removing output directory: %v", err)
	}
}

// Test Platform Export Command fails when not provided required pingone flags together
func TestPlatformExportCmd_PingOneWorkerEnvironmentIdFlagRequiredTogether(t *testing.T) {
	outputDir := os.TempDir() + "/pingctlTestPlatformExportPingOneWorkerEnvironmentIdFlagRequiredTogether"
	err := testutils.ExecutePingctl("platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--service", "pingone-protect",
		"--pingone-worker-environment-id", os.Getenv("PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID"))
	if err == nil {
		t.Errorf("Expected error executing platform export command")
	}

	// Empty output directory
	err = os.RemoveAll(outputDir)
	if err != nil {
		t.Errorf("Error removing output directory: %v", err)
	}
}
