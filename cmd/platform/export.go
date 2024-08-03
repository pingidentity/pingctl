package platform

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingctl/cmd/common"
	platform_internal "github.com/pingidentity/pingctl/internal/commands/platform"
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/spf13/cobra"
)

var (
	multiService customtypes.MultiService = *customtypes.NewMultiService()

	exportFormat    customtypes.ExportFormat = connector.ENUMEXPORTFORMAT_HCL
	pingoneRegion   customtypes.PingOneRegion
	outputDir       string
	overwriteExport bool
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
pingctl platform export --service pingfederate --pingfederate-access-token accessToken`,
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

	basicAuthFlagsUsed := false
	accessTokenAuthFlagsUsed := false

	//Check if basic auth flags are used
	if cmd.Flags().Lookup(profiles.PingFederateUsernameOption.CobraParamName).Changed {
		basicAuthFlagsUsed = true
	}

	//Check if access token auth flags are used
	if cmd.Flags().Lookup(profiles.PingFederateAccessTokenOption.CobraParamName).Changed {
		accessTokenAuthFlagsUsed = true
	}

	return platform_internal.RunInternalExport(cmd.Context(), cmd.Root().Version, outputDir, string(exportFormat), overwriteExport, &multiService, basicAuthFlagsUsed, accessTokenAuthFlagsUsed)
}

func initGeneralExportFlags(cmd *cobra.Command) {
	// Add flags that are not tracked in the viper configuration file
	cmd.Flags().VarP(&exportFormat, "export-format", "e", fmt.Sprintf("Specifies export format\nAllowed: %q", connector.ENUMEXPORTFORMAT_HCL))
	cmd.Flags().VarP(&multiService, "service", "s", fmt.Sprintf("Specifies service(s) to export. Allowed services: %s", multiService.String()))
	cmd.Flags().StringVarP(&outputDir, "output-directory", "d", "", "Specifies output directory for export (Default: Present working directory)")
	cmd.Flags().BoolVarP(&overwriteExport, "overwrite", "o", false, "Overwrite existing generated exports if set.")
}

func initPingOneExportFlags(cmd *cobra.Command) {
	cmd.Flags().String(profiles.PingOneWorkerEnvironmentIDOption.CobraParamName, "", fmt.Sprintf("The ID of the PingOne environment that contains the worker client used to authenticate.  Also configurable via environment variable %s", profiles.PingOneWorkerEnvironmentIDOption.EnvVar))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.PingOneWorkerEnvironmentIDOption,
		Flag:   cmd.Flags().Lookup(profiles.PingOneWorkerEnvironmentIDOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.PingOneWorkerEnvironmentIDOption)

	cmd.Flags().String(profiles.PingOneExportEnvironmentIDOption.CobraParamName, "", fmt.Sprintf("The ID of the PingOne environment to export. Also configurable via environment variable %s (Default: The PingOne worker environment ID)", profiles.PingOneExportEnvironmentIDOption.EnvVar))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.PingOneExportEnvironmentIDOption,
		Flag:   cmd.Flags().Lookup(profiles.PingOneExportEnvironmentIDOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.PingOneExportEnvironmentIDOption)

	cmd.Flags().String(profiles.PingOneWorkerClientIDOption.CobraParamName, "", fmt.Sprintf("The ID of the PingOne worker client used to authenticate.  Also configurable via environment variable %s", profiles.PingOneWorkerClientIDOption.EnvVar))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.PingOneWorkerClientIDOption,
		Flag:   cmd.Flags().Lookup(profiles.PingOneWorkerClientIDOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.PingOneWorkerClientIDOption)

	cmd.Flags().String(profiles.PingOneWorkerClientSecretOption.CobraParamName, "", fmt.Sprintf("The PingOne worker client secret used to authenticate.  Also configurable via environment variable %s", profiles.PingOneWorkerClientSecretOption.EnvVar))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.PingOneWorkerClientSecretOption,
		Flag:   cmd.Flags().Lookup(profiles.PingOneWorkerClientSecretOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.PingOneWorkerClientSecretOption)

	cmd.Flags().Var(&pingoneRegion, profiles.PingOneRegionOption.CobraParamName, fmt.Sprintf("The region of the PingOne service(s). Allowed: %s.  Also configurable via environment variable %s", strings.Join(customtypes.PingOneRegionValidValues(), ", "), profiles.PingOneRegionOption.EnvVar))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.PingOneRegionOption,
		Flag:   cmd.Flags().Lookup(profiles.PingOneRegionOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.PingOneRegionOption)

	cmd.MarkFlagsRequiredTogether(profiles.PingOneWorkerEnvironmentIDOption.CobraParamName, profiles.PingOneWorkerClientIDOption.CobraParamName, profiles.PingOneWorkerClientSecretOption.CobraParamName, profiles.PingOneRegionOption.CobraParamName)
}

func initPingFederateGeneralFlags(cmd *cobra.Command) {
	cmd.Flags().String(profiles.PingFederateHttpsHostOption.CobraParamName, "", fmt.Sprintf("The PingFederate HTTPS host used to communicate with PingFederate's API.  Also configurable via environment variable %s", profiles.PingFederateHttpsHostOption.EnvVar))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.PingFederateHttpsHostOption,
		Flag:   cmd.Flags().Lookup(profiles.PingFederateHttpsHostOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.PingFederateHttpsHostOption)

	cmd.Flags().String(profiles.PingFederateAdminApiPathOption.CobraParamName, "/pf-admin-api/v1", fmt.Sprintf("The PingFederate API URL path used to communicate with PingFederate's API.  Also configurable via environment variable %s", profiles.PingFederateAdminApiPathOption.EnvVar))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.PingFederateAdminApiPathOption,
		Flag:   cmd.Flags().Lookup(profiles.PingFederateAdminApiPathOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.PingFederateAdminApiPathOption)

	cmd.MarkFlagsRequiredTogether(profiles.PingFederateHttpsHostOption.CobraParamName, profiles.PingFederateAdminApiPathOption.CobraParamName)
}

func initPingFederateBasicAuthFlags(cmd *cobra.Command) {
	cmd.Flags().String(profiles.PingFederateUsernameOption.CobraParamName, "", fmt.Sprintf("The PingFederate username used to authenticate.  Also configurable via environment variable %s", profiles.PingFederateUsernameOption.EnvVar))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.PingFederateUsernameOption,
		Flag:   cmd.Flags().Lookup(profiles.PingFederateUsernameOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.PingFederateUsernameOption)

	cmd.Flags().String(profiles.PingFederatePasswordOption.CobraParamName, "", fmt.Sprintf("The PingFederate password used to authenticate.  Also configurable via environment variable %s", profiles.PingFederatePasswordOption.EnvVar))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.PingFederatePasswordOption,
		Flag:   cmd.Flags().Lookup(profiles.PingFederatePasswordOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.PingFederatePasswordOption)

	// When either the username or password flag is used, both must be used
	cmd.MarkFlagsRequiredTogether(profiles.PingFederateUsernameOption.CobraParamName, profiles.PingFederatePasswordOption.CobraParamName)
}

func initPingFederateAccessTokenFlags(cmd *cobra.Command) {
	cmd.Flags().String(profiles.PingFederateAccessTokenOption.CobraParamName, "", fmt.Sprintf("The PingFederate access token used to authenticate.  Also configurable via environment variable %s", profiles.PingFederateAccessTokenOption.EnvVar))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.PingFederateAccessTokenOption,
		Flag:   cmd.Flags().Lookup(profiles.PingFederateAccessTokenOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.PingFederateAccessTokenOption)
}

func initPingFederateClientCredentialsFlags(cmd *cobra.Command) {
	cmd.Flags().String(profiles.PingFederateClientIDOption.CobraParamName, "", fmt.Sprintf("The PingFederate OAuth client ID used to authenticate.  Also configurable via environment variable %s", profiles.PingFederateClientIDOption.EnvVar))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.PingFederateClientIDOption,
		Flag:   cmd.Flags().Lookup(profiles.PingFederateClientIDOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.PingFederateClientIDOption)

	cmd.Flags().String(profiles.PingFederateClientSecretOption.CobraParamName, "", fmt.Sprintf("The PingFederate OAuth client secret used to authenticate.  Also configurable via environment variable %s", profiles.PingFederateClientSecretOption.EnvVar))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.PingFederateClientSecretOption,
		Flag:   cmd.Flags().Lookup(profiles.PingFederateClientSecretOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.PingFederateClientSecretOption)

	cmd.Flags().String(profiles.PingFederateTokenURLOption.CobraParamName, "", fmt.Sprintf("The PingFederate OAuth token URL used to authenticate.  Also configurable via environment variable %s", profiles.PingFederateTokenURLOption.EnvVar))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.PingFederateTokenURLOption,
		Flag:   cmd.Flags().Lookup(profiles.PingFederateTokenURLOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.PingFederateTokenURLOption)

	// When any of the above flags are used, all must be used
	cmd.MarkFlagsRequiredTogether(profiles.PingFederateClientIDOption.CobraParamName, profiles.PingFederateClientSecretOption.CobraParamName, profiles.PingFederateTokenURLOption.CobraParamName)

	cmd.Flags().String(profiles.PingFederateScopesOption.CobraParamName, "", fmt.Sprintf("The PingFederate OAuth scopes used to authenticate. Multiple scopes can be defined as a comma-separated string.  Also configurable via environment variable %s", profiles.PingFederateScopesOption.EnvVar))
	profiles.AddFlagBinding(profiles.Binding{
		Option: profiles.PingFederateScopesOption,
		Flag:   cmd.Flags().Lookup(profiles.PingFederateScopesOption.CobraParamName),
	})
	profiles.AddEnvVarBinding(profiles.PingFederateScopesOption)
}

func markPingFederateFlagsExclusive(cmd *cobra.Command) {
	// The username flag cannot be used with the access token or client credentials authentication methods
	cmd.MarkFlagsMutuallyExclusive(profiles.PingFederateUsernameOption.CobraParamName, profiles.PingFederateAccessTokenOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(profiles.PingFederateUsernameOption.CobraParamName, profiles.PingFederateClientIDOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(profiles.PingFederateUsernameOption.CobraParamName, profiles.PingFederateClientSecretOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(profiles.PingFederateUsernameOption.CobraParamName, profiles.PingFederateTokenURLOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(profiles.PingFederateUsernameOption.CobraParamName, profiles.PingFederateScopesOption.CobraParamName)

	// The password flag cannot be used with the access token or client credentials authentication methods
	cmd.MarkFlagsMutuallyExclusive(profiles.PingFederatePasswordOption.CobraParamName, profiles.PingFederateAccessTokenOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(profiles.PingFederatePasswordOption.CobraParamName, profiles.PingFederateClientIDOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(profiles.PingFederatePasswordOption.CobraParamName, profiles.PingFederateClientSecretOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(profiles.PingFederatePasswordOption.CobraParamName, profiles.PingFederateTokenURLOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(profiles.PingFederatePasswordOption.CobraParamName, profiles.PingFederateScopesOption.CobraParamName)

	// The access token flag cannot be used with the client credentials authentication method
	cmd.MarkFlagsMutuallyExclusive(profiles.PingFederateAccessTokenOption.CobraParamName, profiles.PingFederateClientIDOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(profiles.PingFederateAccessTokenOption.CobraParamName, profiles.PingFederateClientSecretOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(profiles.PingFederateAccessTokenOption.CobraParamName, profiles.PingFederateTokenURLOption.CobraParamName)
	cmd.MarkFlagsMutuallyExclusive(profiles.PingFederateAccessTokenOption.CobraParamName, profiles.PingFederateScopesOption.CobraParamName)

	// Client credential flag exclusivity is already defined above.
}
