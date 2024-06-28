package viperconfig

import (
	"slices"
	"strings"
)

type ConfigCobraParam string

const (
	RootOutputParamName ConfigCobraParam = "output"
	RootColorParamName  ConfigCobraParam = "color"

	ExportPingoneExportEnvironmentIdParamName ConfigCobraParam = "pingone-export-environment-id"
	ExportPingoneWorkerEnvironmentIdParamName ConfigCobraParam = "pingone-worker-environment-id"
	ExportPingoneWorkerClientIdParamName      ConfigCobraParam = "pingone-worker-client-id"
	ExportPingoneWorkerClientSecretParamName  ConfigCobraParam = "pingone-worker-client-secret"
	ExportPingoneRegionParamName              ConfigCobraParam = "pingone-region"
)

type ConfigType string

// Variable type enums
const (
	ENUM_BOOL           ConfigType = "ENUM_BOOL"
	ENUM_ID             ConfigType = "ENUM_ID"
	ENUM_OUTPUT_FORMAT  ConfigType = "ENUM_OUTPUT_FORMAT"
	ENUM_PINGONE_REGION ConfigType = "ENUM_PINGONE_REGION"
	ENUM_STRING         ConfigType = "ENUM_STRING"
)

// Struct to hold cobra param name, viper config key, env var, and variable type
// for each pingctl configuration option
type ConfigOption struct {
	ViperConfigKey string
	EnvVar         string
	VariableType   ConfigType
}

// Define map from cobraParamName to ConfigOption struct for pingctl cobra commands
var ConfigOptions = map[ConfigCobraParam]ConfigOption{
	RootOutputParamName: {
		ViperConfigKey: "pingctl.output",
		EnvVar:         "PINGCTL_OUTPUT",
		VariableType:   ENUM_OUTPUT_FORMAT,
	},
	RootColorParamName: {
		ViperConfigKey: "pingctl.color",
		EnvVar:         "PINGCTL_COLOR",
		VariableType:   ENUM_BOOL,
	},
	ExportPingoneExportEnvironmentIdParamName: {
		ViperConfigKey: "pingone.export.environmentId",
		EnvVar:         "PINGCTL_PINGONE_EXPORT_ENVIRONMENT_ID",
		VariableType:   ENUM_ID,
	},
	ExportPingoneWorkerEnvironmentIdParamName: {
		ViperConfigKey: "pingone.worker.environmentId",
		EnvVar:         "PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID",
		VariableType:   ENUM_ID,
	},
	ExportPingoneWorkerClientIdParamName: {
		ViperConfigKey: "pingone.worker.clientId",
		EnvVar:         "PINGCTL_PINGONE_WORKER_CLIENT_ID",
		VariableType:   ENUM_ID,
	},
	ExportPingoneWorkerClientSecretParamName: {
		ViperConfigKey: "pingone.worker.clientSecret",
		EnvVar:         "PINGCTL_PINGONE_WORKER_CLIENT_SECRET",
		VariableType:   ENUM_STRING,
	},
	ExportPingoneRegionParamName: {
		ViperConfigKey: "pingone.region",
		EnvVar:         "PINGCTL_PINGONE_REGION",
		VariableType:   ENUM_PINGONE_REGION,
	},
}

func GetCobraParamNames() []string {
	cobraParamNames := []string{}
	for cobraParamName := range ConfigOptions {
		cobraParamNames = append(cobraParamNames, string(cobraParamName))
	}

	return cobraParamNames
}

func GetViperConfigKeys() []string {
	viperConfigKeys := []string{}
	for _, configOption := range ConfigOptions {
		viperConfigKeys = append(viperConfigKeys, configOption.ViperConfigKey)
	}

	return viperConfigKeys
}

func GetEnvVars() []string {
	envVars := []string{}
	for _, configOption := range ConfigOptions {
		envVars = append(envVars, configOption.EnvVar)
	}

	return envVars
}

func IsValidViperKey(viperKey string) bool {
	// The only valid configuration keys are those returned by GetViperConfigKeys()
	validViperKeys := GetViperConfigKeys()
	return slices.ContainsFunc(validViperKeys, func(v string) bool {
		return strings.EqualFold(v, viperKey)
	})
}

func GetValueTypeFromViperKey(viperKey string) (ConfigType, bool) {
	for _, configOption := range ConfigOptions {
		if strings.EqualFold(viperKey, configOption.ViperConfigKey) {
			return configOption.VariableType, true
		}
	}

	return "", false
}
