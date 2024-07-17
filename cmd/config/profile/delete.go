package profile

import (
	profile_internal "github.com/pingidentity/pingctl/internal/commands/config/profile"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigProfileDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Command to delete a configuration profile.",
		Long:  `Command to delete a configuration profile.`,
		RunE:  ConfigProfileDeleteRunE,
	}

	return cmd
}

func ConfigProfileDeleteRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Profile Delete Subcommand Called.")

	if err := profile_internal.RunInternalConfigProfileDelete(args); err != nil {
		return err
	}

	return nil
}
