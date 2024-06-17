package config

import (
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Config Subcommand...")
}

func NewConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Command to get, set, and unset pingctl configuration settings.",
		Long:  `Command to get, set, and unset pingctl configuration settings.`,
	}

	cmd.AddCommand(NewConfigGetCommand())
	cmd.AddCommand(NewConfigSetCommand())
	cmd.AddCommand(NewConfigUnsetCommand())

	return cmd
}
