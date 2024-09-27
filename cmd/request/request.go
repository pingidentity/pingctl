package request

import (
	"fmt"

	"github.com/pingidentity/pingctl/cmd/common"
	request_internal "github.com/pingidentity/pingctl/internal/commands/request"
	"github.com/pingidentity/pingctl/internal/configuration/options"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

const (
	commandExamples = `  pingctl request --service pingone environments
  pingctl request --service pingone --http-method GET environments/{{environmentID}}
  pingctl request --service pingone --http-method POST --data {{raw-data}} environments
  pingctl request --service pingone --http-method POST --data @{{filepath}} environments
  pingctl request --service pingone --http-method DELETE environments/{{environmentID}}`

	profileConfigurationFormat = `Profile Configuration Format:
request:
    data: @<Filepath> OR <RawData>
    http-method: <Method>
    service: <Service>
service:
    pingone:
        regionCode: <Code>
        authentication:
            type: <Type>
            worker:
                clientID: <ID>
                clientSecret: <Secret>
                environmentID: <ID>`
)

func NewRequestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(1),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example:               fmt.Sprintf("%s\n\n%s", commandExamples, profileConfigurationFormat),
		Long:                  `Send a custom request to a Ping Service.`,
		RunE:                  requestRunE,
		Short:                 "Send a custom request to a Ping Service.",
		Use:                   "request [flags] API_URI",
	}

	cmd.Flags().AddFlag(options.RequestHTTPMethodOption.Flag)
	cmd.Flags().AddFlag(options.RequestServiceOption.Flag)
	cmd.Flags().AddFlag(options.RequestDataOption.Flag)

	return cmd
}

func requestRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Request Subcommand Called.")

	if err := request_internal.RunInternalRequest(args[0]); err != nil {
		return err
	}

	return nil
}
