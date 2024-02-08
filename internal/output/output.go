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
	l := logger.Get()

	var resultFormat string
	var resultColor func(format string, a ...interface{}) string
	switch output.Result {
	case ENUMCOMMANDOUTPUTRESULT_SUCCESS:
		resultFormat = "%s - %s"
		resultColor = green
	case ENUMCOMMANDOUTPUTRESULT_NOACTION_OK:
		resultFormat = "%s - %s"
		resultColor = green
	case ENUMCOMMANDOUTPUTRESULT_NOACTION_WARN:
		resultFormat = "%s - %s"
		resultColor = yellow
	case ENUMCOMMANDOUTPUTRESULT_FAILURE:
		resultFormat = "%s - %s"
		resultColor = red
	case ENUMCOMMANDOUTPUTRESULT_NIL:
		resultFormat = "%s%s"
		resultColor = white
	default:
		resultFormat = "%s%s"
		resultColor = white
	}

	cmd.Println(resultColor(resultFormat, output.Message, output.Result))
	l.Info().Msgf("%s", resultColor(resultFormat, output.Message, output.Result))

	if output.Fields != nil {
		cmd.Println(cyan("Additional Information:"))
		for k, v := range output.Fields {
			cmd.Println(cyan("%s: %s", k, v))
			l.Info().Msgf("%s", cyan("%s: %s", k, v))
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
	l.Info().Msgf("%s", string(jsonOut))
}
