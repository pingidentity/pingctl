package customtypes

import (
	"fmt"

	"github.com/spf13/pflag"
)

const (
	ENUM_OUTPUT_FORMAT_TEXT string = "text"
	ENUM_OUTPUT_FORMAT_JSON string = "json"
)

type OutputFormat string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*OutputFormat)(nil)

// Implement pflag.Value interface for custom type in cobra pingctl-output parameter

func (s *OutputFormat) Set(outputFormat string) error {
	switch outputFormat {

	case ENUM_OUTPUT_FORMAT_TEXT, ENUM_OUTPUT_FORMAT_JSON:
		*s = OutputFormat(outputFormat)
	default:
		return fmt.Errorf("unrecognized Output Format: '%s'. Must be one of: %s", outputFormat, OutputFormatValidValues())
	}
	return nil
}

func (s *OutputFormat) Type() string {
	return "string"
}

func (s *OutputFormat) String() string {
	return string(*s)
}

func OutputFormatValidValues() string {
	return fmt.Sprintf("'%s', '%s'", ENUM_OUTPUT_FORMAT_TEXT, ENUM_OUTPUT_FORMAT_JSON)
}
