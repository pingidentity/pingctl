package testutils_terraform

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/common"
)

var (
	exportDir                   string
	terraformExecutableFilepath string
)

// Test --generate-config-out for a resource
func ValidateTerraformPlan(t *testing.T, resource connector.ExportableResource, ignoredErrors []string) {
	t.Helper()

	jsonOutputs := singleResourceTerraformPlanGenerateConfigOut(t, resource)

	for _, output := range jsonOutputs {
		if output["@level"] == "error" {
			// Ignore errors
			if ignoredErrors == nil || !slices.Contains(ignoredErrors, output["@message"].(string)) {
				t.Errorf("%v\n%v", output["@message"], output["diagnostic"])
			}
		}
	}
}

// Helper function to run terraform plan --generate-config-out on a single resource
func singleResourceTerraformPlanGenerateConfigOut(t *testing.T, resource connector.ExportableResource) (jsonOutput []map[string]interface{}) {
	t.Helper()

	dirEntries, err := os.ReadDir(exportDir)
	if err != nil {
		t.Fatalf("Failed to read directory entries: %v", err)
	}

	// Clear the export directory of all TF files not named main.tf
	re := regexp.MustCompile(`^.*\.tf$`)
	for _, de := range dirEntries {
		if de.Name() != "main.tf" && re.MatchString(de.Name()) {
			if err := os.RemoveAll(filepath.Join(exportDir, de.Name())); err != nil {
				t.Fatalf("Failed to remove directory entry: %v", err)
			}
		}
	}

	// Export the resource
	if err := common.WriteFiles([]connector.ExportableResource{resource}, connector.ENUMEXPORTFORMAT_HCL, exportDir, true); err != nil {
		t.Fatalf("Failed to export application resource: %v", err)
	}

	stdoutOutput := runTerraformPlanGenerateConfigOut(t, terraformExecutableFilepath, exportDir)

	stdoutLines := strings.Split(stdoutOutput, "\n")

	// Read through the lines, and output error types
	mappedLines := []map[string]interface{}{}
	for _, line := range stdoutLines {
		if line == "" {
			continue
		}

		var mapLine map[string]interface{}
		err := json.Unmarshal([]byte(line), &mapLine)
		if err != nil {
			t.Fatalf("Failed to unmarshal line: %v", err)
		}
		mappedLines = append(mappedLines, mapLine)
	}

	return mappedLines
}

// Helper function to run terraform plan --generate-config-out
func runTerraformPlanGenerateConfigOut(t *testing.T, terraformExecutableFilepath, exportDir string) string {
	// Create the os.exec Command
	terraformPlanCmd := exec.Command(terraformExecutableFilepath)
	// Add the arguments to the command
	terraformPlanCmd.Args = append(terraformPlanCmd.Args, "plan", "-generate-config-out=generated.tf", "-json")
	// Change directories for the command to the testing directory
	terraformPlanCmd.Dir = exportDir

	// Get stdout pipe
	stdout, err := terraformPlanCmd.StdoutPipe()
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Start the command
	if err := terraformPlanCmd.Start(); err != nil {
		t.Fatalf("Failed to start terraform plan command: %v", err)
	}

	// Read from stdout
	stdoutOutput, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatalf("Failed to read from stdout: %v", err)
	}

	// Wait for the command to finish
	if err := terraformPlanCmd.Wait(); err != nil {
		// If err is of type *exec.ExitError, ignore the error
		if _, ok := err.(*exec.ExitError); !ok {
			t.Fatalf("Failed to run terraform plan: %v", err)
		}
	}

	return string(stdoutOutput)
}

// Helper function to initialize the testing directory for terraform plan testing
func InitTerraform(t *testing.T) {
	t.Helper()

	// Create temporary directories for export files and terraform plan testing
	exportDir = t.TempDir()

	// Check if terraform is installed
	checkTerraformInstallPath(t)

	mainTFFileContents := fmt.Sprintf(`terraform {
	required_providers {
		pingone = {
		source = "pingidentity/pingone"
		version = "%s"
		}
	}
}
	
provider "pingone" {}
`, os.Getenv("PINGCTL_PINGONE_PROVIDER_VERSION"))

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

// Helper function to check the path of the terraform executable
func checkTerraformInstallPath(t *testing.T) {
	t.Helper()

	// Check if terraform is installed
	var err error
	terraformExecutableFilepath, err = exec.LookPath("terraform")
	if err != nil {
		t.Fatalf("Terraform is not installed: %v", err)
	}
}