package configuration

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

type OptionType string

// OptionType enums
const (
	ENUM_BOOL           OptionType = "ENUM_BOOL"
	ENUM_EXPORT_FORMAT  OptionType = "ENUM_EXPORT_FORMAT"
	ENUM_UUID           OptionType = "ENUM_UUID"
	ENUM_MULTI_SERVICE  OptionType = "ENUM_MULTI_SERVICE"
	ENUM_OUTPUT_FORMAT  OptionType = "ENUM_OUTPUT_FORMAT"
	ENUM_PINGONE_REGION OptionType = "ENUM_PINGONE_REGION"
	ENUM_STRING         OptionType = "ENUM_STRING"
	ENUM_STRING_SLICE   OptionType = "ENUM_STRING_SLICE"
)

type Option struct {
	CobraParamName  string
	CobraParamValue pflag.Value
	DefaultValue    pflag.Value
	EnvVar          string
	Flag            *pflag.Flag
	Type            OptionType
	ViperKey        string
}

func Options() []Option {
	return []Option{
		PlatformExportExportFormatOption,
		PlatformExportServiceOption,
		PlatformExportOutputDirectoryOption,
		PlatformExportOverwriteOption,
		PlatformExportPingoneWorkerEnvironmentIDOption,
		PlatformExportPingoneExportEnvironmentIDOption,
		PlatformExportPingoneWorkerClientIDOption,
		PlatformExportPingoneWorkerClientSecretOption,
		PlatformExportPingoneRegionOption,
		PlatformExportPingfederateHTTPSHostOption,
		PlatformExportPingfederateAdminAPIPathOption,
		PlatformExportPingfederateXBypassExternalValidationHeaderOption,
		PlatformExportPingfederateCACertificatePemFilesOption,
		PlatformExportPingfederateInsecureTrustAllTLSOption,
		PlatformExportPingfederateUsernameOption,
		PlatformExportPingfederatePasswordOption,
		PlatformExportPingfederateAccessTokenOption,
		PlatformExportPingfederateClientIDOption,
		PlatformExportPingfederateClientSecretOption,
		PlatformExportPingfederateTokenURLOption,
		PlatformExportPingfederateScopesOption,

		RootActiveProfileOption,
		RootColorOption,
		RootConfigOption,
		RootOutputFormatOption,

		ProfileDescriptionOption,

		ConfigProfileOption,
		ConfigNameOption,
		ConfigDescriptionOption,
		ConfigAddProfileDescriptionOption,
		ConfigAddProfileNameOption,
		ConfigAddProfileSetActiveOption,
		ConfigDeleteProfileOption,
		ConfigViewProfileOption,
		ConfigSetActiveProfileOption,
		ConfigGetProfileOption,
		ConfigSetProfileOption,
		ConfigUnsetProfileOption,
	}
}

func init() {
	InitAllOptions()
}

func ViperKeys() (keys []string) {
	for _, opt := range Options() {
		if opt.ViperKey != "" {
			keys = append(keys, opt.ViperKey)
		}
	}

	slices.Sort(keys)
	return keys
}

func ValidateViperKey(viperKey string) error {
	validKeys := ViperKeys()
	for _, vKey := range validKeys {
		if vKey == viperKey {
			return nil
		}
	}

	validKeysStr := strings.Join(validKeys, ", ")
	return fmt.Errorf("key '%s' is not recognized as a valid configuration key. Valid keys: %s", viperKey, validKeysStr)
}

// Return a list of all viper keys from Options
// Including all substrings of parent keys.
// For example, the option key export.environmentID adds the keys
// 'export' and 'export.environmentID' to the list.
func ExpandedViperKeys() (keys []string) {
	leafKeys := ViperKeys()
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

func ValidateParentViperKey(viperKey string) error {
	validKeys := ExpandedViperKeys()
	for _, vKey := range validKeys {
		if vKey == viperKey {
			return nil
		}
	}

	validKeysStr := strings.Join(validKeys, ", ")
	return fmt.Errorf("key '%s' is not recognized as a valid configuration key. Valid keys: %s", viperKey, validKeysStr)
}

func OptionFromViperKey(viperKey string) (opt Option, err error) {
	for _, opt := range Options() {
		if strings.EqualFold(opt.ViperKey, viperKey) {
			return opt, nil
		}
	}
	return opt, fmt.Errorf("failed to get option: no option found for viper key: %s", viperKey)
}

func InitAllOptions() {
	initConfigOptions()
	initPlatformExportOptions()
	initProfilesOptions()
	initRootOptions()
}
