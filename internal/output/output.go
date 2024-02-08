package output

import (
	"encoding/json"

	"github.com/fatih/color"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	boldRed = color.New(color.FgRed).Add(color.Bold).SprintfFunc()
	cyan    = color.New(color.FgCyan).SprintfFunc()
	green   = color.New(color.FgGreen).SprintfFunc()
	red     = color.New(color.FgRed).SprintfFunc()
	white   = color.New(color.FgWhite).SprintfFunc()
	yellow  = color.New(color.FgYellow).SprintfFunc()
)

type CommandOutputResult string

type CommandOutput struct {
	Fields  map[string]interface{}
	Message string
	Warn    string
	Error   error
	Fatal   error
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
		formatText(cmd, CommandOutput{
			Message: "",
			Warn:    "Output format is not recognized. Defaulting to \"text\" output",
			Result:  ENUMCOMMANDOUTPUTRESULT_NIL,
		})
		formatText(cmd, output)
	}
}

func formatText(cmd *cobra.Command, output CommandOutput) {
	l := logger.Get()

	var resultFormat string
	var resultColor func(format string, a ...interface{}) string

	// Determine message color and format based on status
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

	// Supply the user a formatted message and a result status if any.
	cmd.Println(resultColor(resultFormat, output.Message, output.Result))
	l.Info().Msgf(resultColor(resultFormat, output.Message, output.Result))

	// Output and log any additional key/value pairs supplied to the user.
	if output.Fields != nil {
		cmd.Println(cyan("Additional Information:"))
		for k, v := range output.Fields {
			cmd.Println(cyan("%s: %s", k, v))
			l.Info().Msgf("%s: %s", k, v)
		}
	}

	// Inform the user of a warning and log the warning
	if output.Warn != "" {
		cmd.Println(yellow("Warn: %s", output.Warn))
		l.Warn().Msgf(output.Warn)
	}

	// Inform the user of an error and log the error
	if output.Error != nil {
		cmd.Println(red("Error: %s", output.Error.Error()))
		l.Error().Msgf(output.Error.Error())
	}

	// Inform the user of a fatal error and log the fatal error. This exits the program.
	if output.Fatal != nil {
		cmd.Println(boldRed("Fatal: %s", output.Fatal.Error()))
		l.Fatal().Msgf(output.Fatal.Error())
	}

}

func formatJson(cmd *cobra.Command, output CommandOutput) {
	l := logger.Get()

	// Convert the CommandOutput struct to JSON
	jsonOut, err := json.Marshal(output)
	if err != nil {
		l.Error().Err(err).Msgf("Failed to serialize output as JSON")
	}

	// Output the JSON as uncolored string
	cmd.Println(string(jsonOut))

	// Log the serialized JSON as info.
	l.Info().Msgf(string(jsonOut))

	// Log the warning if exists
	if output.Warn != "" {
		l.Warn().Msgf(output.Warn)
	}

	// Log the error if exists
	if output.Error != nil {
		l.Error().Msgf(output.Error.Error())
	}

	// Log the fatal error if exists. This exits the program.
	if output.Fatal != nil {
		l.Fatal().Msgf(output.Fatal.Error())
	}
}
