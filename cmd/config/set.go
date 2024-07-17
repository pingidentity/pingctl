package config

import (
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set pingctl configuration settings.",
		Long: `Set pingctl configuration settings.

Example command usage: 'pingctl config set pingctl.color=false'`,
		RunE: ConfigSetRunE,
	}

	return cmd
}
func ConfigSetRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Get Subcommand Called.")

	if err := config_internal.RunInternalConfigSet(args); err != nil {
		return err
	}

	return nil
}
