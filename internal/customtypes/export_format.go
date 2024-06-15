package customtypes

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/spf13/pflag"
)

type ExportFormat string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*ExportFormat)(nil)

// Implement pflag.Value interface for custom type in cobra export-format parameter

func (s *ExportFormat) Set(format string) error {
	switch format {
	case connector.ENUMEXPORTFORMAT_HCL:
		*s = ExportFormat(format)
	default:
		return fmt.Errorf("unrecognized export format %q. Must be one of: %q", format, ExportFormatValidValues())
	}
	return nil
}

func (s *ExportFormat) Type() string {
	return "string"
}

func (s *ExportFormat) String() string {
	return string(*s)
}

func ExportFormatValidValues() string {
	return fmt.Sprintf("'%s'", connector.ENUMEXPORTFORMAT_HCL)
}
