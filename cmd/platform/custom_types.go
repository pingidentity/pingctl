package platform

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/spf13/pflag"
)

const (
	serviceEnumPlatform string = "pingone-platform"
	serviceEnumSSO      string = "pingone-sso"
	serviceEnumMFA      string = "pingone-mfa"
)

type MultiService struct {
	services          *map[string]bool
	isDefaultServices bool
}

type ExportFormat string

type PingOneRegion string

// Verify that the custom type satisfies the pflag.Value interface
var (
	_ pflag.Value = (*MultiService)(nil)
	_ pflag.Value = (*ExportFormat)(nil)
	_ pflag.Value = (*PingOneRegion)(nil)
)

// Implement pflag.Value interface for custom type in cobra MultiService parameter

func NewMultiService() *MultiService {
	return &MultiService{
		services: &map[string]bool{
			serviceEnumPlatform: true,
			serviceEnumSSO:      true,
			serviceEnumMFA:      true,
		},
		isDefaultServices: true,
	}
}

func (s *MultiService) GetServices() *[]string {
	enabledExportServices := []string{}

	for k, v := range *s.services {
		if v {
			enabledExportServices = append(enabledExportServices, k)
		}
	}

	return &enabledExportServices
}

func (s *MultiService) Set(service string) error {
	// If the user is defining services to export, remove default services from map
	if s.isDefaultServices {
		s.services = &map[string]bool{}
		s.isDefaultServices = false
	}

	switch service {
	case serviceEnumPlatform:
		(*s.services)[serviceEnumPlatform] = true
	case serviceEnumSSO:
		(*s.services)[serviceEnumSSO] = true
	case serviceEnumMFA:
		(*s.services)[serviceEnumMFA] = true
	default:
		return fmt.Errorf("unrecognized service %q. Must be one of: %q, %q, %q", service, serviceEnumPlatform, serviceEnumSSO, serviceEnumMFA)
	}
	return nil
}

func (s *MultiService) Type() string {
	return "string"
}

func (s *MultiService) String() string {
	enabledExportServices := *s.GetServices()

	if len(enabledExportServices) == 0 {
		return "[]"
	}

	return fmt.Sprintf("[ %s ]", strings.Join(enabledExportServices, ", "))
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
