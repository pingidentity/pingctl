package platform_test

import (
	"os"
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
	"github.com/pingidentity/pingctl/internal/testing/testutils_cobra"
)

// Test Platform Export Command Executes without issue
func TestPlatformExportCmd_Execute(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingctl(t, "platform", "export", "--output-directory", outputDir, "--overwrite")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command fails when provided too many arguments
func TestPlatformExportCmd_TooManyArgs(t *testing.T) {
	expectedErrorPattern := `^failed to execute 'pingctl platform export': command accepts 0 arg\(s\), received 1$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "export", "extra-arg")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command fails when provided invalid flag
func TestPlatformExportCmd_InvalidFlag(t *testing.T) {
	expectedErrorPattern := `^unknown flag: --invalid$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "export", "--invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --help, -h flag
func TestPlatformExportCmd_HelpFlag(t *testing.T) {
	err := testutils_cobra.ExecutePingctl(t, "platform", "export", "--help")
	testutils.CheckExpectedError(t, err, nil)

	err = testutils_cobra.ExecutePingctl(t, "platform", "export", "-h")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --service flag
func TestPlatformExportCmd_ServiceFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingctl(t, "platform", "export", "--output-directory", outputDir, "--overwrite", "--service", "pingone-protect")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --service flag with invalid service
func TestPlatformExportCmd_ServiceFlagInvalidService(t *testing.T) {
	expectedErrorPattern := `^invalid argument "invalid" for "-s, --service" flag: unrecognized service 'invalid'\. Must be one of: [a-z-\s,]+$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "export", "--service", "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --export-format flag
func TestPlatformExportCmd_ExportFormatFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingctl(t, "platform", "export", "--output-directory", outputDir, "--export-format", "HCL", "--overwrite", "--service", "pingone-protect")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --export-format flag with invalid format
func TestPlatformExportCmd_ExportFormatFlagInvalidFormat(t *testing.T) {
	expectedErrorPattern := `^invalid argument "invalid" for "-e, --export-format" flag: unrecognized export format 'invalid'\. Must be one of: [A-Z]+$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "export", "--export-format", "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --output-directory flag
func TestPlatformExportCmd_OutputDirectoryFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingctl(t, "platform", "export", "--output-directory", outputDir, "--overwrite", "--service", "pingone-protect")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command --output-directory flag with invalid directory
func TestPlatformExportCmd_OutputDirectoryFlagInvalidDirectory(t *testing.T) {
	expectedErrorPattern := `^failed to create 'platform export' output directory '\/invalid': mkdir \/invalid: .+$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "export", "--output-directory", "/invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --overwrite flag
func TestPlatformExportCmd_OverwriteFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingctl(t, "platform", "export", "--output-directory", outputDir, "--overwrite", "--service", "pingone-protect")
	testutils.CheckExpectedError(t, err, nil)
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
	err = testutils_cobra.ExecutePingctl(t, "platform", "export", "--output-directory", outputDir, "--service", "pingone-protect", "--overwrite=false")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command --overwrite flag true with existing directory
// where the directory already contains a file
func TestPlatformExportCmd_OverwriteFlagTrueWithExistingDirectory(t *testing.T) {
	outputDir := t.TempDir()

	_, err := os.Create(outputDir + "/file")
	if err != nil {
		t.Errorf("Error creating file in output directory: %v", err)
	}

	err = testutils_cobra.ExecutePingctl(t, "platform", "export", "--output-directory", outputDir, "--service", "pingone-protect", "--overwrite")
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command with
// --pingone-worker-environment-id flag
// --pingone-worker-client-id flag
// --pingone-worker-client-secret flag
// --pingone-region flag
func TestPlatformExportCmd_PingOneWorkerEnvironmentIdFlag(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingctl(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--service", "pingone-protect",
		"--pingone-worker-environment-id", os.Getenv(profiles.PingOneWorkerEnvironmentIDOption.EnvVar),
		"--pingone-worker-client-id", os.Getenv(profiles.PingOneWorkerClientIDOption.EnvVar),
		"--pingone-worker-client-secret", os.Getenv(profiles.PingOneWorkerClientSecretOption.EnvVar),
		"--pingone-region", os.Getenv(profiles.PingOneRegionOption.EnvVar))
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command fails when not provided required pingone flags together
func TestPlatformExportCmd_PingOneWorkerEnvironmentIdFlagRequiredTogether(t *testing.T) {
	expectedErrorPattern := `^if any flags in the group \[pingone-worker-environment-id pingone-worker-client-id pingone-worker-client-secret pingone-region] are set they must all be set; missing \[pingone-region pingone-worker-client-id pingone-worker-client-secret]$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "export",
		"--pingone-worker-environment-id", os.Getenv(profiles.PingOneWorkerEnvironmentIDOption.EnvVar))
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export command with PingFederate Basic Auth flags
func TestPlatformExportCmd_PingFederateBasicAuthFlags(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingctl(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--service", "pingfederate",
		"--pingfederate-username", os.Getenv(profiles.PingFederateUsernameOption.EnvVar),
		"--pingfederate-password", os.Getenv(profiles.PingFederatePasswordOption.EnvVar))
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command fails when not provided required PingFederate Basic Auth flags together
func TestPlatformExportCmd_PingFederateBasicAuthFlagsRequiredTogether(t *testing.T) {
	expectedErrorPattern := `^if any flags in the group \[pingfederate-username pingfederate-password] are set they must all be set; missing \[pingfederate-password]$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "export",
		"--pingfederate-username", "Administrator")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command fails when provided invalid PingFederate Basic Auth flags
func TestPlatformExportCmd_PingFederateBasicAuthFlagsInvalid(t *testing.T) {
	outputDir := t.TempDir()

	expectedErrorPattern := `^failed to export 'pingfederate' service: failed to export resource .*\. err: .* Request for resource '.*' was not successful\.\s+Response Code: 401 Unauthorized\s+Response Body: {{"resultId":"invalid_credentials","message":"The credentials you provided were not recognized\."}}\s+Error: 401 Unauthorized$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--service", "pingfederate",
		"--pingfederate-username", "Administrator",
		"--pingfederate-password", "invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export command with PingFederate Client Credentials Auth flags
func TestPlatformExportCmd_PingFederateClientCredentialsAuthFlags(t *testing.T) {
	outputDir := t.TempDir()

	err := testutils_cobra.ExecutePingctl(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--service", "pingfederate",
		"--pingfederate-client-id", os.Getenv(profiles.PingFederateClientIDOption.EnvVar),
		"--pingfederate-client-secret", os.Getenv(profiles.PingFederateClientSecretOption.EnvVar),
		"--pingfederate-scopes", os.Getenv(profiles.PingFederateScopesOption.EnvVar),
		"--pingfederate-token-url", os.Getenv(profiles.PingFederateTokenURLOption.EnvVar))
	testutils.CheckExpectedError(t, err, nil)
}

// Test Platform Export Command fails when not provided required PingFederate Client Credentials Auth flags together
func TestPlatformExportCmd_PingFederateClientCredentialsAuthFlagsRequiredTogether(t *testing.T) {
	expectedErrorPattern := `^if any flags in the group \[pingfederate-client-id pingfederate-client-secret pingfederate-token-url] are set they must all be set; missing \[pingfederate-client-secret pingfederate-token-url]$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "export",
		"--pingfederate-client-id", "test")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command fails when provided invalid PingFederate Client Credentials Auth flags
func TestPlatformExportCmd_PingFederateClientCredentialsAuthFlagsInvalid(t *testing.T) {
	outputDir := t.TempDir()

	expectedErrorPattern := `^failed to export 'pingfederate' service: failed to export resource .*\. err: .* Request for resource '.*' was not successful\. Response is nil\. Error: oauth2: "invalid_client" "Invalid client or client credentials\."$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--service", "pingfederate",
		"--pingfederate-client-id", "test",
		"--pingfederate-client-secret", "invalid",
		"--pingfederate-token-url", "https://localhost:9031/as/token.oauth2")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export Command fails when provided invalid PingFederate OAuth2 Token URL
func TestPlatformExportCmd_PingFederateClientCredentialsAuthFlagsInvalidTokenURL(t *testing.T) {
	outputDir := t.TempDir()

	expectedErrorPattern := `(?s)^failed to export 'pingfederate' service: failed to export resource.*\. err:.*Request for resource '.*' was not successful\. Response is nil\. Error: oauth2: cannot fetch token: 404 Not Found\s+Response: \<!DOCTYPE html\>\s+.*$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "export",
		"--output-directory", outputDir,
		"--overwrite",
		"--service", "pingfederate",
		"--pingfederate-client-id", os.Getenv(profiles.PingFederateClientIDOption.EnvVar),
		"--pingfederate-client-secret", os.Getenv(profiles.PingFederateClientSecretOption.EnvVar),
		"--pingfederate-token-url", "https://localhost:9031/as/invalid")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export command fails when basic auth flags are provided with client credentials auth flags
func TestPlatformExportCmd_PingFederateClientCredentialsAuthFlagsWithUsername(t *testing.T) {
	expectedErrorPattern := `^if any flags in the group \[.*\] are set none of the others can be; \[.*\] were all set$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "export",
		"--pingfederate-client-id", os.Getenv(profiles.PingFederateClientIDOption.EnvVar),
		"--pingfederate-client-secret", os.Getenv(profiles.PingFederateClientSecretOption.EnvVar),
		"--pingfederate-token-url", os.Getenv(profiles.PingFederateTokenURLOption.EnvVar),
		"--pingfederate-username", os.Getenv(profiles.PingFederateUsernameOption.EnvVar),
		"--pingfederate-password", os.Getenv(profiles.PingFederatePasswordOption.EnvVar))

	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export command fails when access token flags are provided with client credentials auth flags
func TestPlatformExportCmd_PingFederateClientCredentialsAuthFlagsWithAccessToken(t *testing.T) {
	expectedErrorPattern := `^if any flags in the group \[.*\] are set none of the others can be; \[.*\] were all set$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "export",
		"--pingfederate-client-id", os.Getenv(profiles.PingFederateClientIDOption.EnvVar),
		"--pingfederate-client-secret", os.Getenv(profiles.PingFederateClientSecretOption.EnvVar),
		"--pingfederate-token-url", os.Getenv(profiles.PingFederateTokenURLOption.EnvVar),
		"--pingfederate-access-token", "token")

	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Platform Export command fails with invalid basic auth flags while there is valid client credentials in config.
// This is because cobra/viper model prioritizes flags over config values.
func TestPlatformExportCmd_PingFederateClientCredentialsAuthFlagsWithInvalidBasicAuth(t *testing.T) {
	expectedErrorPattern := `^failed to export 'pingfederate' service: failed to export resource .*\. err: .* Request for resource '.*' was not successful\.\s+Response Code: 401 Unauthorized\s+Response Body: {{"resultId":"invalid_credentials","message":"The credentials you provided were not recognized\."}}\s+Error: 401 Unauthorized$`
	err := testutils_cobra.ExecutePingctl(t, "platform", "export",
		"--pingfederate-username", os.Getenv(profiles.PingFederateUsernameOption.EnvVar),
		"--pingfederate-password", "invalid",
		"--service", "pingfederate",
		"--overwrite")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}
