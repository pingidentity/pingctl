package testutils_viper

import (
	"fmt"
	"os"
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
)

var (
	// use spaces for tabs in the config file contents below,
	// as tabs are not supported in YAML
	configFileContents = fmt.Sprintf(`activeProfile: default
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
            environmentid: ""`,
		os.Getenv(profiles.RegionOption.EnvVar),
		os.Getenv(profiles.WorkerClientIDOption.EnvVar),
		os.Getenv(profiles.WorkerClientSecretOption.EnvVar),
		os.Getenv(profiles.WorkerEnvironmentIDOption.EnvVar))
)

func CreateConfigFile(t *testing.T) string {
	t.Helper()

	configFilepath := t.TempDir() + "/config.yaml"
	if err := os.WriteFile(configFilepath, []byte(configFileContents), 0600); err != nil {
		t.Fatalf("Failed to create config file: %s", err)
	}

	return configFilepath
}

func InitVipers(t *testing.T) {
	t.Helper()

	// // Create and write to a temporary config file
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

func InitVipersCustomFile(t *testing.T, fileContents string) {
	t.Helper()

	configFileContents = fileContents
	InitVipers(t)
}
