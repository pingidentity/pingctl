package profile

import (
	profile_internal "github.com/pingidentity/pingctl/internal/commands/config/profile"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigProfileDescribeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Command to describe a configuration profiles.",
		Long:  `Command to describe a configuration profiles.`,
		RunE:  ConfigProfileDescribeRunE,
	}

	return cmd
}

func ConfigProfileDescribeRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Profile Describe Subcommand Called.")

	if err := profile_internal.RunInternalConfigProfileDescribe(args); err != nil {
		return err
	}

	return nil
}
