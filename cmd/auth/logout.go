package auth

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/noop"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/spf13/cobra"
)

func NewLogoutCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "logout",
		//TODO more fleshed-out descriptions
		Short: "Logout with Ping",
		Long:  "Logout with Ping",
		Run: func(cmd *cobra.Command, args []string) {
			// Just use the no-op connector for now by default
			authConnectors := []connector.Authenticatable{
				noop.Connector(),
			}
			err := authConnectors[0].Logout()
			if err != nil {
				output.Format(cmd, output.CommandOutput{
					Message: "Logout failed.",
					Fatal:   err,
					Result:  output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
				})
			}
		},
	}

	return cmd
}
