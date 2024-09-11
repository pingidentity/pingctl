package platform

import (
	"github.com/pingidentity/pingctl/cmd/common"
	platform_internal "github.com/pingidentity/pingctl/internal/commands/platform"
	"github.com/pingidentity/pingctl/internal/configuration/options"
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
	cmd.Flags().AddFlag(options.PlatformExportExportFormatOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportServiceOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportOutputDirectoryOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportOverwriteOption.Flag)
}

func initPingOneExportFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PlatformExportPingoneWorkerEnvironmentIDOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportPingoneExportEnvironmentIDOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportPingoneWorkerClientIDOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportPingoneWorkerClientSecretOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportPingoneRegionOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		options.PlatformExportPingoneWorkerEnvironmentIDOption.CobraParamName,
		options.PlatformExportPingoneWorkerClientIDOption.CobraParamName,
		options.PlatformExportPingoneWorkerClientSecretOption.CobraParamName,
		options.PlatformExportPingoneRegionOption.CobraParamName)
}

func initPingFederateGeneralFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PlatformExportPingfederateHTTPSHostOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportPingfederateAdminAPIPathOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		options.PlatformExportPingfederateHTTPSHostOption.CobraParamName,
		options.PlatformExportPingfederateAdminAPIPathOption.CobraParamName)

	cmd.Flags().AddFlag(options.PlatformExportPingfederateXBypassExternalValidationHeaderOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportPingfederateCACertificatePemFilesOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportPingfederateInsecureTrustAllTLSOption.Flag)
}

func initPingFederateBasicAuthFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PlatformExportPingfederateUsernameOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportPingfederatePasswordOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		options.PlatformExportPingfederateUsernameOption.CobraParamName,
		options.PlatformExportPingfederatePasswordOption.CobraParamName)
}

func initPingFederateAccessTokenFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PlatformExportPingfederateAccessTokenOption.Flag)
}

func initPingFederateClientCredentialsFlags(cmd *cobra.Command) {
	cmd.Flags().AddFlag(options.PlatformExportPingfederateClientIDOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportPingfederateClientSecretOption.Flag)
	cmd.Flags().AddFlag(options.PlatformExportPingfederateTokenURLOption.Flag)

	cmd.MarkFlagsRequiredTogether(
		options.PlatformExportPingfederateClientIDOption.CobraParamName,
		options.PlatformExportPingfederateClientSecretOption.CobraParamName,
		options.PlatformExportPingfederateTokenURLOption.CobraParamName)

	cmd.Flags().AddFlag(options.PlatformExportPingfederateScopesOption.Flag)
}

func markPingFederateFlagsExclusive(cmd *cobra.Command) {
	// The username flag cannot be used with the access token or client credentials authentication methods
	cmd.MarkFlagsMutuallyExclusive(options.PlatformExportPingfederateUsernameOption.CobraParamName, options.PlatformExportPingfederateAccessTokenOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(options.PlatformExportPingfederateUsernameOption.CobraParamName, options.PlatformExportPingfederateClientIDOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(options.PlatformExportPingfederateUsernameOption.CobraParamName, options.PlatformExportPingfederateClientSecretOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(options.PlatformExportPingfederateUsernameOption.CobraParamName, options.PlatformExportPingfederateTokenURLOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(options.PlatformExportPingfederateUsernameOption.CobraParamName, options.PlatformExportPingfederateScopesOption.CobraParamName)

	// The password flag cannot be used with the access token or client credentials authentication methods
	cmd.MarkFlagsMutuallyExclusive(options.PlatformExportPingfederatePasswordOption.CobraParamName, options.PlatformExportPingfederateAccessTokenOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(options.PlatformExportPingfederatePasswordOption.CobraParamName, options.PlatformExportPingfederateClientIDOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(options.PlatformExportPingfederatePasswordOption.CobraParamName, options.PlatformExportPingfederateClientSecretOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(options.PlatformExportPingfederatePasswordOption.CobraParamName, options.PlatformExportPingfederateTokenURLOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(options.PlatformExportPingfederatePasswordOption.CobraParamName, options.PlatformExportPingfederateScopesOption.CobraParamName)

	// The access token flag cannot be used with the client credentials authentication method
	cmd.MarkFlagsMutuallyExclusive(options.PlatformExportPingfederateAccessTokenOption.CobraParamName, options.PlatformExportPingfederateClientIDOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(options.PlatformExportPingfederateAccessTokenOption.CobraParamName, options.PlatformExportPingfederateClientSecretOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(options.PlatformExportPingfederateAccessTokenOption.CobraParamName, options.PlatformExportPingfederateTokenURLOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(options.PlatformExportPingfederateAccessTokenOption.CobraParamName, options.PlatformExportPingfederateScopesOption.CobraParamName)

	// Client credential flag exclusivity is already defined above.
}
