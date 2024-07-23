package config

import (
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigUnsetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unset",
		Short: "Unset pingctl configuration settings.",
		Long: `Unset pingctl configuration settings.

Example command usage: 'pingctl config unset pingctl.color'`,
		RunE: ConfigUnsetRunE,
	}

	return cmd
}
func ConfigUnsetRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Get Subcommand Called.")

	if err := config_internal.RunInternalConfigUnset(args); err != nil {
		return err
	}

	return nil
}
