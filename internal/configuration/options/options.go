package options

import "github.com/spf13/pflag"

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

// 'pingctl config' command options
var (
	ConfigProfileOption     Option
	ConfigNameOption        Option
	ConfigDescriptionOption Option

	ConfigAddProfileDescriptionOption Option
	ConfigAddProfileNameOption        Option
	ConfigAddProfileSetActiveOption   Option

	ConfigDeleteProfileOption Option

	ConfigViewProfileOption Option

	ConfigSetActiveProfileOption Option

	ConfigGetProfileOption Option

	ConfigSetProfileOption Option

	ConfigUnsetProfileOption Option
)

// 'pingctl platform export' command options
var (
	PlatformExportExportFormatOption    Option
	PlatformExportServiceOption         Option
	PlatformExportOutputDirectoryOption Option
	PlatformExportOverwriteOption       Option

	PlatformExportPingoneWorkerEnvironmentIDOption Option
	PlatformExportPingoneExportEnvironmentIDOption Option
	PlatformExportPingoneWorkerClientIDOption      Option
	PlatformExportPingoneWorkerClientSecretOption  Option
	PlatformExportPingoneRegionOption              Option

	PlatformExportPingfederateHTTPSHostOption                       Option
	PlatformExportPingfederateAdminAPIPathOption                    Option
	PlatformExportPingfederateXBypassExternalValidationHeaderOption Option
	PlatformExportPingfederateCACertificatePemFilesOption           Option
	PlatformExportPingfederateInsecureTrustAllTLSOption             Option
	PlatformExportPingfederateUsernameOption                        Option
	PlatformExportPingfederatePasswordOption                        Option
	PlatformExportPingfederateAccessTokenOption                     Option
	PlatformExportPingfederateClientIDOption                        Option
	PlatformExportPingfederateClientSecretOption                    Option
	PlatformExportPingfederateTokenURLOption                        Option
	PlatformExportPingfederateScopesOption                          Option
)

// Generic viper profile options
var (
	ProfileDescriptionOption Option
)

// Options
var (
	RootActiveProfileOption Option
	RootColorOption         Option
	RootConfigOption        Option
	RootOutputFormatOption  Option
)
