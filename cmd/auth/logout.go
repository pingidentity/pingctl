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
		RunE: func(cmd *cobra.Command, args []string) error {
			// Just use the no-op connector for now by default
			authConnectors := []connector.Authenticatable{
				noop.Connector(),
			}
			err := authConnectors[0].Logout()
			if err != nil {
				output.Format(cmd, output.CommandOutput{
					Message: "Logout failed.",
					Result:  output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
				})
				return err
			}
			return nil
		},
	}

	return cmd
}
