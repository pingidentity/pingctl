package config

import (
	"github.com/pingidentity/pingctl/cmd/common"
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewConfigGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.RangeArgs(0, 1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example: `pingctl config get
pingctl config get pingone
pingctl config get pingctl.color
pingctl config get pingone.export.environmentID`,
		Long:  `Get pingctl configuration settings.`,
		RunE:  configGetRunE,
		Short: "Get pingctl configuration settings.",
		Use:   "get [flags] [key]",
	}

	return cmd
}

func configGetRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Get Subcommand Called.")

	key := ""
	if len(args) > 0 {
		key = args[0]
	}

	if err := config_internal.RunInternalConfigGet(key); err != nil {
		return err
	}

	return nil
}
