package feedback

import (
	feedback_internal "github.com/pingidentity/pingctl/internal/commands/feedback"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
)

func NewFeedbackCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "feedback",
		Short: "Information on tool feedback",
		Long: `A command to provide the user information
on how to give feedback or get help with the tool
through the use of the GitHub repository's issue tracker.`,
		RunE: feedbackRunE,
	}

	return cmd
}

func feedbackRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Running Feedback Subcommand with args %s", args)

	feedback_internal.PrintFeedbackMessage()

	return nil
}
