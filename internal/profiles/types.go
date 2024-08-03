package profiles

import (
	"slices"
	"strings"

	"github.com/pingidentity/pingctl/internal/customtypes"
)

type ConfigOpts struct {
	Options []Option
}

type Option struct {
	CobraParamName string
	ViperKey       string
	EnvVar         string
	Type           OptionType
}

type OptionType string

// Variable type enums
const (
	ENUM_BOOL           OptionType = "ENUM_BOOL"
	ENUM_ID             OptionType = "ENUM_ID"
	ENUM_OUTPUT_FORMAT  OptionType = "ENUM_OUTPUT_FORMAT"
	ENUM_PINGONE_REGION OptionType = "ENUM_PINGONE_REGION"
	ENUM_STRING         OptionType = "ENUM_STRING"
)

var (
	OutputOption = Option{
		CobraParamName: "output-format",
		ViperKey:       "pingctl.outputFormat",
		EnvVar:         "PINGCTL_OUTPUT_FORMAT",
		Type:           ENUM_OUTPUT_FORMAT,
	}
	ColorOption = Option{
		CobraParamName: "color",
		ViperKey:       "pingctl.color",
		EnvVar:         "PINGCTL_COLOR",
		Type:           ENUM_BOOL,
	}
	ProfileOption = Option{
		CobraParamName: "active-profile",
		ViperKey:       "activeProfile",
		EnvVar:         "PINGCTL_ACTIVE_PROFILE",
		Type:           ENUM_STRING,
	}
	ProfileDescriptionOption = Option{
		CobraParamName: "description",
		ViperKey:       "description",
		Type:           ENUM_STRING,
	}
	PingOneExportEnvironmentIDOption = Option{
		CobraParamName: "pingone-export-environment-id",
		ViperKey:       "pingone.export.environmentID",
		EnvVar:         "PINGCTL_PINGONE_EXPORT_ENVIRONMENT_ID",
		Type:           ENUM_ID,
	}
	PingOneWorkerEnvironmentIDOption = Option{
		CobraParamName: "pingone-worker-environment-id",
		ViperKey:       "pingone.worker.environmentID",
		EnvVar:         "PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID",
		Type:           ENUM_ID,
	}
	PingOneWorkerClientIDOption = Option{
		CobraParamName: "pingone-worker-client-id",
		ViperKey:       "pingone.worker.clientID",
		EnvVar:         "PINGCTL_PINGONE_WORKER_CLIENT_ID",
		Type:           ENUM_ID,
	}
	PingOneWorkerClientSecretOption = Option{
		CobraParamName: "pingone-worker-client-secret",
		ViperKey:       "pingone.worker.clientSecret",
		EnvVar:         "PINGCTL_PINGONE_WORKER_CLIENT_SECRET",
		Type:           ENUM_STRING,
	}
	PingOneRegionOption = Option{
		CobraParamName: "pingone-region",
		ViperKey:       "pingone.region",
		EnvVar:         "PINGCTL_PINGONE_REGION",
		Type:           ENUM_PINGONE_REGION,
	}
	PingFederateUsernameOption = Option{
		CobraParamName: "pingfederate-username",
		ViperKey:       "pingfederate.basicAuth.username",
		EnvVar:         "PINGCTL_PINGFEDERATE_USERNAME",
		Type:           ENUM_STRING,
	}
	PingFederatePasswordOption = Option{
		CobraParamName: "pingfederate-password",
		ViperKey:       "pingfederate.basicAuth.password",
		EnvVar:         "PINGCTL_PINGFEDERATE_PASSWORD",
		Type:           ENUM_STRING,
	}
	PingFederateHttpsHostOption = Option{
		CobraParamName: "pingfederate-https-host",
		ViperKey:       "pingfederate.httpsHost",
		EnvVar:         "PINGCTL_PINGFEDERATE_HTTPS_HOST",
		Type:           ENUM_STRING,
	}
	PingFederateAdminApiPathOption = Option{
		CobraParamName: "pingfederate-admin-api-path",
		ViperKey:       "pingfederate.adminApiPath",
		EnvVar:         "PINGCTL_PINGFEDERATE_ADMIN_API_PATH",
		Type:           ENUM_STRING,
	}
	PingFederateClientIDOption = Option{
		CobraParamName: "pingfederate-client-id",
		ViperKey:       "pingfederate.clientCredentialsAuth.clientID",
		EnvVar:         "PINGCTL_PINGFEDERATE_CLIENT_ID",
		Type:           ENUM_STRING,
	}
	PingFederateClientSecretOption = Option{
		CobraParamName: "pingfederate-client-secret",
		ViperKey:       "pingfederate.clientCredentialsAuth.clientSecret",
		EnvVar:         "PINGCTL_PINGFEDERATE_CLIENT_SECRET",
		Type:           ENUM_STRING,
	}
	PingFederateTokenURLOption = Option{
		CobraParamName: "pingfederate-token-url",
		ViperKey:       "pingfederate.clientCredentialsAuth.tokenURL",
		EnvVar:         "PINGCTL_PINGFEDERATE_TOKEN_URL",
		Type:           ENUM_STRING,
	}
	PingFederateScopesOption = Option{
		CobraParamName: "pingfederate-scopes",
		ViperKey:       "pingfederate.clientCredentialsAuth.scopes",
		EnvVar:         "PINGCTL_PINGFEDERATE_SCOPES",
		Type:           ENUM_STRING,
	}
	PingFederateAccessTokenOption = Option{
		CobraParamName: "pingfederate-access-token",
		ViperKey:       "pingfederate.accessTokenAuth.accessToken",
		EnvVar:         "PINGCTL_PINGFEDERATE_ACCESS_TOKEN",
		Type:           ENUM_STRING,
	}

	ConfigOptions = ConfigOpts{
		Options: []Option{
			OutputOption,
			ColorOption,
			ProfileOption,
			PingOneExportEnvironmentIDOption,
			PingOneWorkerEnvironmentIDOption,
			PingOneWorkerClientIDOption,
			PingOneWorkerClientSecretOption,
			PingOneRegionOption,
			ProfileDescriptionOption,
			PingFederateUsernameOption,
			PingFederatePasswordOption,
			PingFederateHttpsHostOption,
			PingFederateAdminApiPathOption,
			PingFederateClientIDOption,
			PingFederateClientSecretOption,
			PingFederateTokenURLOption,
			PingFederateScopesOption,
			PingFederateAccessTokenOption,
		},
	}
)

// Return a list of all viper keys from Options defined in @ConfigOptions
func ProfileKeys() (keys []string) {
	for _, option := range ConfigOptions.Options {
		if option.ViperKey == ProfileOption.ViperKey {
			continue
		}

		keys = append(keys, option.ViperKey)
	}

	slices.Sort(keys)
	return keys
}

// Return a list of all viper keys from Options defined in @ConfigOptions
// Including all substrings of parent keys.
// For example, the option key export.environmentID adds the keys
// 'export' and 'export.environmentID' to the list.
func ExpandedProfileKeys() (keys []string) {
	leafKeys := ProfileKeys()
	for _, key := range leafKeys {
		keySplit := strings.Split(key, ".")
		for i := 0; i < len(keySplit); i++ {
			curKey := strings.Join(keySplit[:i+1], ".")
			if !slices.ContainsFunc(keys, func(v string) bool {
				return strings.EqualFold(v, curKey)
			}) {
				keys = append(keys, curKey)
			}
		}
	}

	slices.Sort(keys)
	return keys
}

func OptionTypeFromViperKey(key string) (optType OptionType, ok bool) {
	for _, opt := range ConfigOptions.Options {
		if strings.EqualFold(opt.ViperKey, key) {
			return opt.Type, true
		}
	}
	return "", false
}

func GetDefaultValue(optType OptionType) (val any) {
	switch optType {
	case ENUM_BOOL:
		return false
	case ENUM_ID:
		return ""
	case ENUM_OUTPUT_FORMAT:
		return customtypes.OutputFormat("text")
	case ENUM_PINGONE_REGION:
		return customtypes.PingOneRegion("")
	case ENUM_STRING:
		return ""
	}
	return nil
}
