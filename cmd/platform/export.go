package platform

import (
	"github.com/pingidentity/pingctl/cmd/common"
	platform_internal "github.com/pingidentity/pingctl/internal/commands/platform"
	"github.com/pingidentity/pingctl/internal/configuration"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewExportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example: `pingctl platform export
pingctl platform export --output-directory dir --overwrite
pingctl platform export --export-format HCL
pingctl platform export --service pingone-platform --service pingone-sso
pingctl platform export --service pingone-platform --pingone-client-environment-id envID --pingone-worker-client-id clientID --pingone-worker-client-secret clientSecret --pingone-region region
pingctl platform export --service pingfederate --pingfederate-username user --pingfederate-password password
pingctl platform export --service pingfederate --pingfederate-client-id clientID --pingfederate-client-secret clientSecret --pingfederate-token-url tokenURL
pingctl platform export --service pingfederate --pingfederate-access-token accessToken
pingctl platform export --service pingfederate --x-bypass-external-validation=false --ca-certificate-pem-files "/path/to/cert.pem,/path/to/cert2.pem" --insecure-trust-all-tls=false`,
		Long:  `Export configuration-as-code packages for the Ping Platform.`,
		Short: "Export configuration-as-code packages for the Ping Platform.",
		RunE:  exportRunE,
		Use:   "export [flags]",
	}

	initGeneralExportFlags(cmd)
	initPingOneExportFlags(cmd)
	initPingFederateGeneralFlags(cmd)
	initPingFederateBasicAuthFlags(cmd)
	initPingFederateAccessTokenFlags(cmd)
	initPingFederateClientCredentialsFlags(cmd)
	markPingFederateFlagsExclusive(cmd)

	return cmd
}

func exportRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()

	l.Debug().Msgf("Platform Export Subcommand Called.")

	return platform_internal.RunInternalExport(cmd.Context(), cmd.Root().Version)
}

func initGeneralExportFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(configuration.PlatformExportExportFormatOption.Flag)
	cmd.Flags().AddFlag(configuration.PlatformExportServiceOption.Flag)
	cmd.Flags().AddFlag(configuration.PlatformExportOutputDirectoryOption.Flag)
	cmd.Flags().AddFlag(configuration.PlatformExportOverwriteOption.Flag)
}

func initPingOneExportFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(configuration.PlatformExportPingoneWorkerEnvironmentIDOption.Flag)
	cmd.Flags().AddFlag(configuration.PlatformExportPingoneExportEnvironmentIDOption.Flag)
	cmd.Flags().AddFlag(configuration.PlatformExportPingoneWorkerClientIDOption.Flag)
	cmd.Flags().AddFlag(configuration.PlatformExportPingoneWorkerClientSecretOption.Flag)
	cmd.Flags().AddFlag(configuration.PlatformExportPingoneRegionOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		configuration.PlatformExportPingoneWorkerEnvironmentIDOption.CobraParamName,
		configuration.PlatformExportPingoneWorkerClientIDOption.CobraParamName,
		configuration.PlatformExportPingoneWorkerClientSecretOption.CobraParamName,
		configuration.PlatformExportPingoneRegionOption.CobraParamName)
}

func initPingFederateGeneralFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(configuration.PlatformExportPingfederateHTTPSHostOption.Flag)
	cmd.Flags().AddFlag(configuration.PlatformExportPingfederateAdminAPIPathOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		configuration.PlatformExportPingfederateHTTPSHostOption.CobraParamName,
		configuration.PlatformExportPingfederateAdminAPIPathOption.CobraParamName)

	cmd.Flags().AddFlag(configuration.PlatformExportPingfederateXBypassExternalValidationHeaderOption.Flag)
	cmd.Flags().AddFlag(configuration.PlatformExportPingfederateCACertificatePemFilesOption.Flag)
	cmd.Flags().AddFlag(configuration.PlatformExportPingfederateInsecureTrustAllTLSOption.Flag)
}

func initPingFederateBasicAuthFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(configuration.PlatformExportPingfederateUsernameOption.Flag)
	cmd.Flags().AddFlag(configuration.PlatformExportPingfederatePasswordOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		configuration.PlatformExportPingfederateUsernameOption.CobraParamName,
		configuration.PlatformExportPingfederatePasswordOption.CobraParamName)
}

func initPingFederateAccessTokenFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(configuration.PlatformExportPingfederateAccessTokenOption.Flag)
}

func initPingFederateClientCredentialsFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(configuration.PlatformExportPingfederateClientIDOption.Flag)
	cmd.Flags().AddFlag(configuration.PlatformExportPingfederateClientSecretOption.Flag)
	cmd.Flags().AddFlag(configuration.PlatformExportPingfederateTokenURLOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		configuration.PlatformExportPingfederateClientIDOption.CobraParamName,
		configuration.PlatformExportPingfederateClientSecretOption.CobraParamName,
		configuration.PlatformExportPingfederateTokenURLOption.CobraParamName)

	cmd.Flags().AddFlag(configuration.PlatformExportPingfederateScopesOption.Flag)
}

func markPingFederateFlagsExclusive(cmd *cobra.Command) {
	// The username flag cannot be used with the access token or client credentials authentication methods
	cmd.MarkFlagsMutuallyExclusive(configuration.PlatformExportPingfederateUsernameOption.CobraParamName, configuration.PlatformExportPingfederateAccessTokenOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(configuration.PlatformExportPingfederateUsernameOption.CobraParamName, configuration.PlatformExportPingfederateClientIDOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(configuration.PlatformExportPingfederateUsernameOption.CobraParamName, configuration.PlatformExportPingfederateClientSecretOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(configuration.PlatformExportPingfederateUsernameOption.CobraParamName, configuration.PlatformExportPingfederateTokenURLOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(configuration.PlatformExportPingfederateUsernameOption.CobraParamName, configuration.PlatformExportPingfederateScopesOption.CobraParamName)

	// The password flag cannot be used with the access token or client credentials authentication methods
	cmd.MarkFlagsMutuallyExclusive(configuration.PlatformExportPingfederatePasswordOption.CobraParamName, configuration.PlatformExportPingfederateAccessTokenOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(configuration.PlatformExportPingfederatePasswordOption.CobraParamName, configuration.PlatformExportPingfederateClientIDOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(configuration.PlatformExportPingfederatePasswordOption.CobraParamName, configuration.PlatformExportPingfederateClientSecretOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(configuration.PlatformExportPingfederatePasswordOption.CobraParamName, configuration.PlatformExportPingfederateTokenURLOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(configuration.PlatformExportPingfederatePasswordOption.CobraParamName, configuration.PlatformExportPingfederateScopesOption.CobraParamName)

	// The access token flag cannot be used with the client credentials authentication method
	cmd.MarkFlagsMutuallyExclusive(configuration.PlatformExportPingfederateAccessTokenOption.CobraParamName, configuration.PlatformExportPingfederateClientIDOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(configuration.PlatformExportPingfederateAccessTokenOption.CobraParamName, configuration.PlatformExportPingfederateClientSecretOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(configuration.PlatformExportPingfederateAccessTokenOption.CobraParamName, configuration.PlatformExportPingfederateTokenURLOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(configuration.PlatformExportPingfederateAccessTokenOption.CobraParamName, configuration.PlatformExportPingfederateScopesOption.CobraParamName)

	// Client credential flag exclusivity is already defined above.
}
