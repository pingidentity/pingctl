package config

import (
	"github.com/pingidentity/pingctl/cmd/common"
	config_internal "github.com/pingidentity/pingctl/internal/commands/config"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

var ()

func NewConfigGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.RangeArgs(0, 1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example: `pingctl config get
pingctl config get pingone
pingctl config get pingctl.color
pingctl config get pingone.export.environmentID`,
		Use:   "get [flags] [key]",
		Short: "Get pingctl configuration settings.",
		Long:  `Get pingctl configuration settings.`,
		RunE:  ConfigGetRunE,
	}

	return cmd
}

func ConfigGetRunE(cmd *cobra.Command, args []string) error {
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
