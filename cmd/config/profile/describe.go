package profile

import (
	"github.com/pingidentity/pingctl/cmd/common"
	profile_internal "github.com/pingidentity/pingctl/internal/commands/config/profile"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigProfileDescribeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:    common.ExactArgs(1),
		Example: `pingctl config profile describe my-profile`,
		Long:    `Command to describe a configuration profile.`,
		RunE:    ConfigProfileDescribeRunE,
		Short:   "Command to describe a configuration profile.",
		Use:     "describe [flags] profile",
	}

	return cmd
}

func ConfigProfileDescribeRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Profile Describe Subcommand Called.")

	if err := profile_internal.RunInternalConfigProfileDescribe(args[0]); err != nil {
		return err
	}

	return nil
}
