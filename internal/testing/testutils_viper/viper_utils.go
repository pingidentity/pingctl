package testutils_viper

import (
	"fmt"
	"os"
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
)

var (
	configFileContents        string
	defaultConfigFileContents string = fmt.Sprintf(`activeProfile: default
default:
    description: "default description"
    pingctl:
        color: true
        outputFormat: text
    pingone:
        export:
            environmentid: ""
        region: %s
        worker:
            clientid: %s
            clientsecret: %s
            environmentid: %s
    pingfederate:
        accesstokenauth:
            accesstoken: ""
        adminapipath: ""
        basicauth:
            password: "%s"
            username: "%s"
        caCertificatePemFiles: "%s"
        clientcredentialsauth:
            clientid: "%s"
            clientsecret: "%s"
            scopes: "%s"
            tokenurl: "%s"
        httpshost: "%s"
        insecureTrustAllTLS: true
        xBypassExternalValidationHeader: true
production:
    description: "test profile description"
    pingctl:
        color: true
        outputFormat: text
    pingone:
        export:
            environmentid: ""
        region: ""
        worker:
            clientid: ""
            clientsecret: ""
            environmentid: ""
    pingfederate:
        accesstokenauth:
            accesstoken: ""
        adminapipath: ""
        basicauth:
            password: ""
            username: ""
        caCertificatePemFiles: []
        clientcredentialsauth:
            clientid: ""
            clientsecret: ""
            scopes: []
            tokenurl: ""
        httpshost: ""
        insecureTrustAllTLS: false
        xBypassExternalValidationHeader: false`,
		os.Getenv(profiles.PingOneRegionOption.EnvVar),
		os.Getenv(profiles.PingOneWorkerClientIDOption.EnvVar),
		os.Getenv(profiles.PingOneWorkerClientSecretOption.EnvVar),
		os.Getenv(profiles.PingOneWorkerEnvironmentIDOption.EnvVar),
		os.Getenv(profiles.PingFederatePasswordOption.EnvVar),
		os.Getenv(profiles.PingFederateUsernameOption.EnvVar),
		os.Getenv(profiles.PingFederateCACertificatePemFilesOption.EnvVar),
		os.Getenv(profiles.PingFederateClientIDOption.EnvVar),
		os.Getenv(profiles.PingFederateClientSecretOption.EnvVar),
		os.Getenv(profiles.PingFederateScopesOption.EnvVar),
		os.Getenv(profiles.PingFederateTokenURLOption.EnvVar),
		os.Getenv(profiles.PingFederateHttpsHostOption.EnvVar))
)

func CreateConfigFile(t *testing.T) string {
	t.Helper()

	if configFileContents == "" {
		configFileContents = defaultConfigFileContents
	}

	configFilepath := t.TempDir() + "/config.yaml"
	if err := os.WriteFile(configFilepath, []byte(configFileContents), 0600); err != nil {
		t.Fatalf("Failed to create config file: %s", err)
	}

	return configFilepath
}

func configureMainViper(t *testing.T) {
	t.Helper()

	// Create and write to a temporary config file
	configFilepath := CreateConfigFile(t)
	// Give main viper instance a file location to write to
	mainViper := profiles.GetMainViper()
	mainViper.SetConfigFile(configFilepath)
	if err := mainViper.ReadInConfig(); err != nil {
		t.Fatal(err)
	}

	if err := profiles.SetProfileViperWithProfile("default"); err != nil {
		t.Fatal(err)
	}
}

func InitVipers(t *testing.T) {
	t.Helper()

	configFileContents = defaultConfigFileContents

	configureMainViper(t)
}

func InitVipersCustomFile(t *testing.T, fileContents string) {
	t.Helper()

	configFileContents = fileContents
	configureMainViper(t)
}
