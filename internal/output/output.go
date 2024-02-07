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

func Format(cmd *cobra.Command, output CommandOutput) {
	l := logger.Get()

	if cmd == nil {
		l.Fatal().Msgf("Failed to output. Expected cmd to be set.")
	}

	colorizeOutput := viper.GetBool("color")

	if !colorizeOutput {
		color.NoColor = true
	}

	outputFormat := viper.GetString("output")

	switch outputFormat {
	case "text":
		formatText(cmd, output)
	case "json":
		formatJson(cmd, output)
	default:
		l.Error().Msgf("Output format %q is not a recognized option. Defaulting to text output", outputFormat)
		formatText(cmd, output)
	}
}

func formatText(cmd *cobra.Command, output CommandOutput) {
	switch output.Result {
	case ENUMCOMMANDOUTPUTRESULT_SUCCESS:
		cmd.Println(green("%s - %s", output.Message, output.Result))
	case ENUMCOMMANDOUTPUTRESULT_NOACTION_OK:
		cmd.Println(green("%s - %s", output.Message, output.Result))
	case ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN:
		cmd.Println(yellow("%s - %s", output.Message, output.Result))
	case ENUMCOMMANDOUTPUTRESULT_FAILURE:
		cmd.Println(red("%s - %s", output.Message, output.Result))
	case ENUMCOMMANDOUTPUTRESULT_NIL:
		cmd.Println(white("%s", output.Message))
	default:
		cmd.Println(white("%s", output.Message))
	}

	if output.Fields != nil {
		cmd.Println(cyan("Additional Information:"))
		for k, v := range output.Fields {
			cmd.Println(cyan("%s: %s", k, v))
		}
	}

}

func formatJson(cmd *cobra.Command, output CommandOutput) {
	l := logger.Get()

	jsonOut, err := json.Marshal(output)

	if err != nil {
		l.Error().Err(err).Msgf("Failed to serialize output as JSON")
	}

	cmd.Println(string(jsonOut))
}
