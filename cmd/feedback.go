package cmd

import (
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/output"
	"github.com/spf13/cobra"
)

func init() {
	l := logger.Get()

	l.Debug().Msgf("Initializing Feedback Subcommand...")
}

func NewFeedbackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "feedback",
		Short: "Information on tool feedback",
		Long: `A command to provide the user information
			on how to give feedback or get help with the tool
			through the use of the GitHub repository's issue tracker.`,
		RunE: FeedbackRunE,
	}

	return cmd
}

func FeedbackRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()

	l.Debug().Msgf("Feedback command called.")

	feedbackMessage := `Thank you for participating in early adoption of the refreshed Ping Identity universal CLI!

We appreciate your feedback and suggestions for improvement regarding your experiences with the CLI.

Please visit the following URL in your browser to fill out a short, anonymous survey that will help guide our development efforts and improve the CLI for all users:

	https://forms.gle/xLz6ao4Ts86Zn2yt9

If you encounter any bugs while using the tool, please open an issue on the project's GitHub repository's issue tracker:

	https://github.com/pingidentity/pingctl/issues/new

`
	output.Format(cmd, output.CommandOutput{
		Message: feedbackMessage,
		Result:  output.ENUMCOMMANDOUTPUTRESULT_NIL,
	})

	return nil
}
