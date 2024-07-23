package platform

import (
	"github.com/spf13/cobra"
)

func NewPlatformCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "platform",
		Short: "Provides details and interactions with the connected Ping Platform.",
		Long:  `Provides details and interactions with the connected Ping Platform.`,
	}

	cmd.AddCommand(NewExportCommand())

	return cmd
}
