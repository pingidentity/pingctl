package auth

import (
	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/spf13/cobra"
)

func NewLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "login",
		//TODO more fleshed-out descriptions
		Short: "Login with Ping",
		Long:  "Login with Ping",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Just use the no-op connector for now by default
			authConnectors := []connector.Authenticatable{}
			err := authConnectors[0].Login()
			if err != nil {
				output.Format(cmd, output.CommandOutput{
					Message: "Login failed.",
					Result:  output.ENUMCOMMANDOUTPUTRESULT_FAILURE,
				})
				return err
			}
			return nil
		},
	}

	return cmd
}
