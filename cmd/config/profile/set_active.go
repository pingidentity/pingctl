package profile

import (
	profile_internal "github.com/pingidentity/pingctl/internal/commands/config/profile"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigProfileSetActiveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-active",
		Short: "Command to set the active configuration profile.",
		Long:  `Command to set the active configuration profile.`,
		RunE:  ConfigProfileSetActiveRunE,
	}

	return cmd
}

func ConfigProfileSetActiveRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Profile set-active Subcommand Called.")

	if err := profile_internal.RunInternalConfigProfileSetActive(args); err != nil {
		return err
	}

	return nil
}
