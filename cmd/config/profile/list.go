package profile

import (
	"github.com/pingidentity/pingctl/cmd/common"
	profile_internal "github.com/pingidentity/pingctl/internal/commands/config/profile"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigProfileListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Long:                  `Command to list all configuration profiles.`,
		RunE:                  ConfigProfileListRunE,
		Short:                 "Command to list all configuration profiles.",
		Use:                   "list [flags]",
	}

	return cmd
}

func ConfigProfileListRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Profile List Subcommand Called.")

	profile_internal.RunInternalConfigProfileList()

	return nil
}
