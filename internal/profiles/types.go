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
		CobraParamName: "output",
		ViperKey:       "pingctl.output",
		EnvVar:         "PINGCTL_OUTPUT",
		Type:           ENUM_OUTPUT_FORMAT,
	}
	ColorOption = Option{
		CobraParamName: "color",
		ViperKey:       "pingctl.color",
		EnvVar:         "PINGCTL_COLOR",
		Type:           ENUM_BOOL,
	}
	ProfileOption = Option{
		CobraParamName: "profile",
		ViperKey:       "activeProfile",
		EnvVar:         "PINGCTL_ACTIVE_PROFILE",
		Type:           ENUM_STRING,
	}
	ExportEnvironmentIDOption = Option{
		CobraParamName: "pingone-export-environment-id",
		ViperKey:       "pingone.export.environmentID",
		EnvVar:         "PINGCTL_PINGONE_EXPORT_ENVIRONMENT_ID",
		Type:           ENUM_ID,
	}
	WorkerEnvironmentIDOption = Option{
		CobraParamName: "pingone-worker-environment-id",
		ViperKey:       "pingone.worker.environmentID",
		EnvVar:         "PINGCTL_PINGONE_WORKER_ENVIRONMENT_ID",
		Type:           ENUM_ID,
	}
	WorkerClientIDOption = Option{
		CobraParamName: "pingone-worker-client-id",
		ViperKey:       "pingone.worker.clientID",
		EnvVar:         "PINGCTL_PINGONE_WORKER_CLIENT_ID",
		Type:           ENUM_ID,
	}
	WorkerClientSecretOption = Option{
		CobraParamName: "pingone-worker-client-secret",
		ViperKey:       "pingone.worker.clientSecret",
		EnvVar:         "PINGCTL_PINGONE_WORKER_CLIENT_SECRET",
		Type:           ENUM_STRING,
	}
	RegionOption = Option{
		CobraParamName: "pingone-region",
		ViperKey:       "pingone.region",
		EnvVar:         "PINGCTL_PINGONE_REGION",
		Type:           ENUM_PINGONE_REGION,
	}
	ProfileDescriptionOption = Option{
		CobraParamName: "description",
		ViperKey:       "description",
		Type:           ENUM_STRING,
	}

	ConfigOptions = ConfigOpts{
		Options: []Option{
			OutputOption,
			ColorOption,
			ProfileOption,
			ExportEnvironmentIDOption,
			WorkerEnvironmentIDOption,
			WorkerClientIDOption,
			WorkerClientSecretOption,
			RegionOption,
			ProfileDescriptionOption,
		},
	}
)

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

func ExpandedProfileKeys() (keys []string) {
	leafKeys := ProfileKeys()
	for _, key := range leafKeys {
		keySplit := strings.Split(key, ".")
		for i := 0; i < len(keySplit); i++ {
			curKey := strings.Join(keySplit[:i+1], ".")
			if !slices.Contains(keys, curKey) {
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
