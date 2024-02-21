package platform

import (
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewPlatformCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "platform",
		//TODO add command short and long description
		Short: "",
		Long:  ``,
	}

	cmd.AddCommand(NewExportCommand())

	return cmd
}

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Platform Subcommand...")
}
