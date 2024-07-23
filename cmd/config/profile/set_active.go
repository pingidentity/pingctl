package profile

import (
	"github.com/pingidentity/pingctl/cmd/common"
	profile_internal "github.com/pingidentity/pingctl/internal/commands/config/profile"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigProfileSetActiveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:    common.ExactArgs(1),
		Example: `pingctl config profile set-active my-profile`,
		Long:    `Command to set the active configuration profile.`,
		RunE:    ConfigProfileSetActiveRunE,
		Short:   "Command to set the active configuration profile.",
		Use:     "set-active [flags] profile",
	}

	return cmd
}

func ConfigProfileSetActiveRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Profile set-active Subcommand Called.")

	if err := profile_internal.RunInternalConfigProfileSetActive(args[0]); err != nil {
		return err
	}

	return nil
}
