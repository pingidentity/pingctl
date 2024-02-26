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

type PingOneRegion string

// Verify that the custom type satisfies the pflag.Value interface
var (
	_ pflag.Value = (*MultiService)(nil)
	_ pflag.Value = (*ExportFormat)(nil)
	_ pflag.Value = (*PingOneRegion)(nil)
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
		return fmt.Errorf("unrecognized service %q. Must be one of: %q", service, serviceEnumPlatform)
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
		return fmt.Errorf("unrecognized export format %q. Must be one of: %q", format, connector.ENUMEXPORTFORMAT_HCL)
	}
	return nil
}

func (s *ExportFormat) Type() string {
	return "string"
}

func (s *ExportFormat) String() string {
	return string(*s)
}

// Implement pflag.Value interface for custom type in cobra pingone-region parameter

func (s *PingOneRegion) Set(region string) error {
	switch region {
	case connector.ENUMREGION_AP, connector.ENUMREGION_CA, connector.ENUMREGION_EU, connector.ENUMREGION_NA:
		*s = PingOneRegion(region)
	default:
		return fmt.Errorf("unrecognized PingOne Region: %q. Must be one of: %q, %q, %q, %q", region, connector.ENUMREGION_AP, connector.ENUMREGION_CA, connector.ENUMREGION_EU, connector.ENUMREGION_NA)
	}
	return nil
}

func (s *PingOneRegion) Type() string {
	return "string"
}

func (s *PingOneRegion) String() string {
	return string(*s)
}
