package feedback_internal_test

import (
	feedback_internal "github.com/pingidentity/pingctl/internal/commands/feedback"
)

func ExamplePrintFeedbackMessage() {
	feedback_internal.PrintFeedbackMessage()

	//Output:
	//Thank you for participating in early adoption of the refreshed Ping Identity universal CLI!
	//
	//We appreciate your feedback and suggestions for improvement regarding your experiences with the CLI.
	//
	//Please visit the following URL in your browser to fill out a short, anonymous survey that will help guide our development efforts and improve the CLI for all users:
	//
	//	https://forms.gle/xLz6ao4Ts86Zn2yt9
	//
	//If you encounter any bugs while using the tool, please open an issue on the project's GitHub repository's issue tracker:
	//
	//	https://github.com/pingidentity/pingctl/issues/new
	//
}
