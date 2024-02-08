package platform

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/noop"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewExportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "export",
		//TODO add command short and long description
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			l := logger.Get()
			var exportableConnector connector.Exportable
			// Just use the no-op connector for now by default
			exportableConnector = noop.Connector()
			//TODO selectable format and output location
			err := exportableConnector.Export(connector.ENUMEXPORTFORMAT_HCL, "/tmp")
			if err != nil {
				l.Fatal().Err(err).Msg("Export failed")
			}
		},
	}

	return cmd
}
