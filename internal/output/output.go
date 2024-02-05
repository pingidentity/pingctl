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
)

func Format(fields map[string]interface{}) {
	l := logger.Get()

	colorizeOutput := viper.GetBool("color")

	if !colorizeOutput {
		color.NoColor = true
	}

	outputFormat := viper.GetString("output")

	switch outputFormat {
	case "text":
		formatText(fields)
	case "json":
		formatJson(fields)
	default:
		l.Error().Msgf("Output format %q is not a recognized option. Defaulting to text output", outputFormat)
		formatText(fields)
	}
}

func formatText(fields map[string]interface{}) {
	for k, v := range fields {
		green("%s: %s\n", k, v)
	}
}

func formatJson(fields map[string]interface{}) {
	l := logger.Get()

	l.Info().Msg("Yep")

	enc := json.NewEncoder(os.Stdout)

	if err := enc.Encode(fields); err != nil {
		l.Error().Err(err).Msgf("")
	}
}
