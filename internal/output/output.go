package output

import (
	"encoding/json"

	"github.com/fatih/color"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cyan   = color.New(color.FgCyan).SprintfFunc()
	green  = color.New(color.FgGreen).SprintfFunc()
	red    = color.New(color.FgRed).SprintfFunc()
	white  = color.New(color.FgWhite).SprintfFunc()
	yellow = color.New(color.FgYellow).SprintfFunc()
)

type CommandOutputResult string

type CommandOutput struct {
	Command *cobra.Command
	Fields  map[string]interface{}
	Message string
	Result  CommandOutputResult
}

const (
	ENUMCOMMANDOUTPUTRESULT_NIL           CommandOutputResult = ""
	ENUMCOMMANDOUTPUTRESULT_SUCCESS       CommandOutputResult = "Success"
	ENUMCOMMANDOUTPUTRESULT_NOACTION_OK   CommandOutputResult = "No Action (OK)"
	ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN CommandOutputResult = "No Action (Warning)"
	ENUMCOMMANDOUTPUTRESULT_FAILURE       CommandOutputResult = "Failure"
)

func Format(cmdOut CommandOutput) {
	l := logger.Get()

	if cmdOut.Command == nil {
		l.Fatal().Msgf("Failed to output. Expected Command Field to be set.")
	}

	colorizeOutput := viper.GetBool("color")

	if !colorizeOutput {
		color.NoColor = true
	}

	outputFormat := viper.GetString("output")

	switch outputFormat {
	case "text":
		formatText(cmdOut)
	case "json":
		formatJson(cmdOut)
	default:
		l.Error().Msgf("Output format %q is not a recognized option. Defaulting to text output", outputFormat)
		formatText(cmdOut)
	}
}

func formatText(cmdOut CommandOutput) {
	switch cmdOut.Result {
	case ENUMCOMMANDOUTPUTRESULT_SUCCESS:
		cmdOut.Command.Println(green("%s - %s", cmdOut.Message, cmdOut.Result))
	case ENUMCOMMANDOUTPUTRESULT_NOACTION_OK:
		cmdOut.Command.Println(green("%s - %s", cmdOut.Message, cmdOut.Result))
	case ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN:
		cmdOut.Command.Println(yellow("%s - %s", cmdOut.Message, cmdOut.Result))
	case ENUMCOMMANDOUTPUTRESULT_FAILURE:
		cmdOut.Command.Println(red("%s - %s", cmdOut.Message, cmdOut.Result))
	case ENUMCOMMANDOUTPUTRESULT_NIL:
		cmdOut.Command.Println(white("%s", cmdOut.Message))
	default:
		cmdOut.Command.Println(white("%s", cmdOut.Message))
	}

	if cmdOut.Fields != nil {
		cmdOut.Command.Println(cyan("Additional Information:"))
		for k, v := range cmdOut.Fields {
			cmdOut.Command.Println(cyan("%s: %s", k, v))
		}
	}

}

func formatJson(cmdOut CommandOutput) {
	l := logger.Get()

	jsonOut, err := json.Marshal(cmdOut)

	if err != nil {
		l.Error().Err(err).Msgf("Failed to serialize output as JSON")
	}

	cmdOut.Command.Println(string(jsonOut))
}
