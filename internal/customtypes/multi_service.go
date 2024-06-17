package customtypes

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

const (
	ENUM_SERVICE_PLATFORM string = "pingone-platform"
	ENUM_SERVICE_SSO      string = "pingone-sso"
	ENUM_SERVICE_MFA      string = "pingone-mfa"
	ENUM_SERVICE_PROTECT  string = "pingone-protect"
)

type MultiService struct {
	services          *map[string]bool
	isDefaultServices bool
}

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*MultiService)(nil)

// Implement pflag.Value interface for custom type in cobra MultiService parameter

func NewMultiService() *MultiService {
	return &MultiService{
		services: &map[string]bool{
			ENUM_SERVICE_MFA:      true,
			ENUM_SERVICE_PLATFORM: true,
			ENUM_SERVICE_PROTECT:  true,
			ENUM_SERVICE_SSO:      true,
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
	case ENUM_SERVICE_PLATFORM:
		(*s.services)[ENUM_SERVICE_PLATFORM] = true
	case ENUM_SERVICE_SSO:
		(*s.services)[ENUM_SERVICE_SSO] = true
	case ENUM_SERVICE_MFA:
		(*s.services)[ENUM_SERVICE_MFA] = true
	case ENUM_SERVICE_PROTECT:
		(*s.services)[ENUM_SERVICE_PROTECT] = true
	default:
		return fmt.Errorf("unrecognized service '%s'. Must be one of: %s", service, strings.Join(MultiServiceValidValues(), ", "))
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

func MultiServiceValidValues() []string {
	return []string{ENUM_SERVICE_PLATFORM, ENUM_SERVICE_SSO, ENUM_SERVICE_MFA, ENUM_SERVICE_PROTECT}
}
