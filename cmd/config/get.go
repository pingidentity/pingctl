package config

import (
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Config Get Subcommand...")
}

func NewConfigGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get pingctl configuration settings.",
		Long: `Get pingctl configuration settings.

Example command usage: 'pingctl config get pingctl.color'`,
		RunE: ConfigGetRunE,
	}

	return cmd
}

func ConfigGetRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Get Subcommand Called.")

	if err := config_internal.RunInternalConfigGet(args); err != nil {
		return err
	}

	return nil
}
