package testutils_helpers

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/sso"
)

// Helper function to run terraform plan --generate-config-out on a single resource
func TestSingleResourceTerraformPlanGenerateConfigOut(t *testing.T, resource connector.ExportableResource, sdkClientInfo *connector.SDKClientInfo) {
	t.Helper()

	// Get an instance of the PingOne SSO Connector
	ssoConnector := sso.SSOConnector(sdkClientInfo.Context, sdkClientInfo.ApiClient, sdkClientInfo.ApiClientId, sdkClientInfo.ExportEnvironmentID)

	// Create temporary directories for export files and terraform plan testing
	exportDir := createTempExportDir(t, resource.ResourceType())

	// Check if terraform is installed
	terraformExecutableFilepath := checkTerraformInstallPath(t)

	// Terraform Initialize the testing directory for terraform plan testing
	initTerraformInDir(t, exportDir, terraformExecutableFilepath)

	// Export the resource
	if err := ssoConnector.ExportSingle(connector.ENUMEXPORTFORMAT_HCL, exportDir, true, resource); err != nil {
		t.Fatalf("Failed to export application resource: %v", err)
	}

	stderrOutput := runTerraformPlanGenerateConfigOut(t, terraformExecutableFilepath, exportDir)

	// if stderrOutput is not empty, then fail the test
	if stderrOutput != "" {
		t.Errorf("Failed to run terraform plan --generate-config-out: %v", stderrOutput)
	}

	// Cleanup temporary directories
	cleanupTempExportDir(t, exportDir)
}

// Helper function to run terraform plan --generate-config-out
func runTerraformPlanGenerateConfigOut(t *testing.T, terraformExecutableFilepath, exportDir string) string {
	// Create the os.exec Command
	terraformPlanCmd := exec.Command(terraformExecutableFilepath)
	// Add the arguments to the command
	terraformPlanCmd.Args = append(terraformPlanCmd.Args, "plan", "-generate-config-out=generated.tf")
	// Change directories for the command to the testing directory
	terraformPlanCmd.Dir = exportDir

	// Get stderr pipe
	stderr, err := terraformPlanCmd.StderrPipe()
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Start the command
	if err := terraformPlanCmd.Start(); err != nil {
		t.Fatalf("Failed to start terraform plan command: %v", err)
	}

	// Read from stderr
	stderrOutput, _ := io.ReadAll(stderr)

	// Wait for the command to finish
	if err := terraformPlanCmd.Wait(); err != nil {
		// If err is of type *exec.ExitError, ignore the error
		if _, ok := err.(*exec.ExitError); !ok {
			t.Fatalf("Failed to run terraform plan: %v", err)
		}
	}

	return string(stderrOutput)
}

// Helper function to check the path of the terraform executable
func checkTerraformInstallPath(t *testing.T) string {
	t.Helper()

	// Check if terraform is installed
	terraformExecutableFilepath, err := exec.LookPath("terraform")
	if err != nil {
		t.Fatalf("Terraform is not installed: %v", err)
	}

	return terraformExecutableFilepath
}

// Helper function to initialize the testing directory for terraform plan testing
func initTerraformInDir(t *testing.T, exportDir string, terraformExecutableFilepath string) {
	t.Helper()

	const mainTFFileContents = `terraform {
		required_providers {
		  pingone = {
			source = "pingidentity/pingone"
			version = "0.29.1"
		  }
		}
	  }
	  
	  provider "pingone" {
		# Configuration options
	  }`

	// Write main.tf to testing directory
	mainTFFilepath := filepath.Join(exportDir, "main.tf")
	if err := os.WriteFile(mainTFFilepath, []byte(mainTFFileContents), 0600); err != nil {
		t.Fatalf("Failed to write main.tf to testing directory: %v", err)
	}

	// Run terraform init in testing directory
	initCmd := exec.Command(terraformExecutableFilepath)
	initCmd.Args = append(initCmd.Args, "init")
	initCmd.Dir = exportDir

	// Run the command
	combinedOutput, err := initCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run terraform init: %v\n%s", err, combinedOutput)
	}
}

// Helper function to create temporary directories for export files and terraform plan testing
func createTempExportDir(t *testing.T, resourceName string) string {
	t.Helper()

	exportDir := os.TempDir() + "/pingctlTestConnectorExport" + resourceName

	// Clean up the directories if they already exists
	cleanupTempExportDir(t, exportDir)

	if err := os.MkdirAll(exportDir, os.ModePerm); err != nil {
		t.Fatalf("Failed to create temporary export directory: %v", err)
	}

	return exportDir
}

// Helper function to clean up temporary directories
func cleanupTempExportDir(t *testing.T, exportDir string) {
	t.Helper()

	if err := os.RemoveAll(exportDir); err != nil {
		t.Fatalf("Failed to remove temporary export directory: %v", err)
	}
}
