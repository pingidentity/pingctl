package platform

import (
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Platform Subcommand...")
}

func NewPlatformCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "platform",
		Short: "Provides details and interactions with the connected Ping Platform.",
		Long:  `Provides details and interactions with the connected Ping Platform.`,
	}

	cmd.AddCommand(NewExportCommand())

	return cmd
}
