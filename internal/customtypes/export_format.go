package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/spf13/pflag"
)

type ExportFormat string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*ExportFormat)(nil)

// Implement pflag.Value interface for custom type in cobra export-format parameter

func (ef *ExportFormat) Set(format string) error {
	if ef == nil {
		return fmt.Errorf("failed to set Export Format value: %s. Export Format is nil", format)
	}

	switch format {
	case connector.ENUMEXPORTFORMAT_HCL:
		*ef = ExportFormat(format)
	default:
		return fmt.Errorf("unrecognized export format '%s'. Must be one of: %s", format, strings.Join(ExportFormatValidValues(), ", "))
	}
	return nil
}

func (ef ExportFormat) Type() string {
	return "string"
}

func (ef ExportFormat) String() string {
	return string(ef)
}

func ExportFormatValidValues() []string {
	exportFormats := []string{
		connector.ENUMEXPORTFORMAT_HCL,
	}

	slices.Sort(exportFormats)

	return exportFormats
}
