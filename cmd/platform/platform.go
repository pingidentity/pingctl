package platform

import (
	"github.com/spf13/cobra"
)

func NewPlatformCommand() *cobra.Command {
	cmd := &cobra.Command{
		Long:  `Provides details and interactions with the connected Ping Platform.`,
		Short: "Provides details and interactions with the connected Ping Platform.",
		Use:   "platform",
	}

	cmd.AddCommand(NewExportCommand())

	return cmd
}
