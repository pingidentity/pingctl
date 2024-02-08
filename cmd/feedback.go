package cmd

import (
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/spf13/cobra"
)

var FeedbackMessage = `
Thank you for participating in early adoption of the refreshed Ping Identity universal CLI.

We appreciate your feedback and information regarding your expirience with the CLI.

Please visit the following URL in your browser and let us know of any feedback or issues related
to the tool:

	https://github.com/pingidentity/pingctl/issues/new
`

func NewFeedbackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "feedback",
		Short: "Information on tool feedback",
		Long: `A command to provide the user information
			on how to give feedback or get help with the tool
			through the use of the GitHub repository's issue tracker.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			l := logger.Get()

			l.Debug().Msgf("Feedback command called.")

			output.Format(cmd, output.CommandOutput{
				Message: FeedbackMessage,
				Result:  output.ENUMCOMMANDOUTPUTRESULT_NIL,
			})

			return nil
		},
	}

	return cmd
}
