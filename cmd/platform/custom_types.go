package platform

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/spf13/pflag"
)

const (
	serviceEnumPlatform = "pingone-platform"
)

type MultiService struct {
	services *[]string
}

type ExportFormat string

// Verify that the custom type satisfies the pflag.Value interface
var (
	_ pflag.Value = (*MultiService)(nil)
	_ pflag.Value = (*ExportFormat)(nil)
)

// Implement pflag.Value interface for custom type in cobra service parameter

func (s *MultiService) Set(service string) error {
	switch service {
	case serviceEnumPlatform:
		if *s.services == nil {
			s.services = &[]string{}
		}
		*s.services = append(*s.services, service)
	default:
		return fmt.Errorf("unrecognized service %q", service)
	}
	return nil
}

func (s *MultiService) Type() string {
	return "string"
}

func (s *MultiService) String() string {
	return fmt.Sprintf("[ %s ]", strings.Join(*s.services, ", "))
}

// Implement pflag.Value interface for custom type in cobra export-format parameter

func (s *ExportFormat) Set(format string) error {
	switch format {
	case connector.ENUMEXPORTFORMAT_HCL:
		*s = ExportFormat(format)
	default:
		return fmt.Errorf("unrecognized export format %q", format)
	}
	return nil
}

func (s *ExportFormat) Type() string {
	return "string"
}

func (s *ExportFormat) String() string {
	return string(*s)
}
