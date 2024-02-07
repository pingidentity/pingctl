package output

import (
	"encoding/json"
	"os"

	"github.com/fatih/color"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/spf13/viper"
)

var (
	green = color.New(color.FgGreen).PrintfFunc()
	white = color.New(color.FgWhite).PrintfFunc()
)

func Format(message string, fields map[string]interface{}) {
	l := logger.Get()

	colorizeOutput := viper.GetBool("color")

	if !colorizeOutput {
		color.NoColor = true
	}

	outputFormat := viper.GetString("output")

	switch outputFormat {
	case "text":
		formatText(message, fields)
	case "json":
		formatJson(message, fields)
	default:
		l.Error().Msgf("Output format %q is not a recognized option. Defaulting to text output", outputFormat)
		formatText(message, fields)
	}
}

func formatText(message string, fields map[string]interface{}) {
	white(message)

	for k, v := range fields {
		green("%s: %s\n", k, v)
	}
}

func formatJson(message string, fields map[string]interface{}) {
	l := logger.Get()

	if fields != nil {
		fields["message"] = message
	} else {
		fields = map[string]interface{}{
			"message": message,
		}
	}

	enc := json.NewEncoder(os.Stdout)

	if err := enc.Encode(fields); err != nil {
		l.Error().Err(err).Msgf("")
	}
}
