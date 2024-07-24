package profile

import (
	"github.com/pingidentity/pingctl/cmd/common"
	profile_internal "github.com/pingidentity/pingctl/internal/commands/config/profile"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigProfileDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               `pingctl config profile delete my-profile`,
		Long:                  `Command to delete a configuration profile.`,
		RunE:                  configProfileDeleteRunE,
		Short:                 "Command to delete a configuration profile.",
		Use:                   "delete [flags] profile",
	}

	return cmd
}

func configProfileDeleteRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Profile Delete Subcommand Called.")

	if err := profile_internal.RunInternalConfigProfileDelete(args[0]); err != nil {
		return err
	}

	return nil
}
