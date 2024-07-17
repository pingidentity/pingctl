package profile

import (
	profile_internal "github.com/pingidentity/pingctl/internal/commands/config/profile"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigProfileListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Command to list all configuration profiles.",
		Long:  `Command to list all configuration profiles.`,
		RunE:  ConfigProfileListRunE,
	}

	return cmd
}

func ConfigProfileListRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Profile List Subcommand Called.")

	profile_internal.RunInternalConfigProfileList(args)

	return nil
}
