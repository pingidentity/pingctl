package customtypes

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/pflag"
)

const (
	ENUM_SERVICE_PINGONE_PLATFORM string = "pingone-platform"
	ENUM_SERVICE_PINGONE_SSO      string = "pingone-sso"
	ENUM_SERVICE_PINGONE_MFA      string = "pingone-mfa"
	ENUM_SERVICE_PINGONE_PROTECT  string = "pingone-protect"
	ENUM_SERVICE_PINGFEDERATE     string = "pingfederate"
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
			ENUM_SERVICE_PINGFEDERATE:     true,
			ENUM_SERVICE_PINGONE_PLATFORM: true,
			ENUM_SERVICE_PINGONE_SSO:      true,
			ENUM_SERVICE_PINGONE_MFA:      true,
			ENUM_SERVICE_PINGONE_PROTECT:  true,
		},
		isDefaultServices: true,
	}
}

func (s *MultiService) GetServices() *[]string {
	enabledExportServices := []string{}

	if s == nil {
		return &enabledExportServices
	}

	for k, v := range *s.services {
		if v {
			enabledExportServices = append(enabledExportServices, k)
		}
	}

	slices.Sort(enabledExportServices)

	return &enabledExportServices
}

func (s *MultiService) Set(service string) error {
	if s == nil {
		return fmt.Errorf("MultiService is nil")
	}

	// If the user is defining services to export, remove default services from map
	if s.isDefaultServices {
		s.services = &map[string]bool{}
		s.isDefaultServices = false
	}

	switch service {
	case ENUM_SERVICE_PINGFEDERATE:
		(*s.services)[ENUM_SERVICE_PINGFEDERATE] = true
	case ENUM_SERVICE_PINGONE_PLATFORM:
		(*s.services)[ENUM_SERVICE_PINGONE_PLATFORM] = true
	case ENUM_SERVICE_PINGONE_SSO:
		(*s.services)[ENUM_SERVICE_PINGONE_SSO] = true
	case ENUM_SERVICE_PINGONE_MFA:
		(*s.services)[ENUM_SERVICE_PINGONE_MFA] = true
	case ENUM_SERVICE_PINGONE_PROTECT:
		(*s.services)[ENUM_SERVICE_PINGONE_PROTECT] = true
	default:
		return fmt.Errorf("unrecognized service '%s'. Must be one of: %s", service, strings.Join(MultiServiceValidValues(), ", "))
	}
	return nil
}

func (s *MultiService) ContainsPingOneService() bool {
	if s == nil {
		return false
	}

	return (*s.services)[ENUM_SERVICE_PINGONE_PLATFORM] ||
		(*s.services)[ENUM_SERVICE_PINGONE_SSO] ||
		(*s.services)[ENUM_SERVICE_PINGONE_MFA] ||
		(*s.services)[ENUM_SERVICE_PINGONE_PROTECT]
}

func (s *MultiService) ContainsPingFederateService() bool {
	if s == nil {
		return false
	}

	return (*s.services)[ENUM_SERVICE_PINGFEDERATE]
}

func (s *MultiService) Type() string {
	return "string"
}

func (s *MultiService) String() string {
	if s == nil {
		return "[]"
	}

	enabledExportServices := *s.GetServices()

	if len(enabledExportServices) == 0 {
		return "[]"
	}

	slices.Sort(enabledExportServices)

	return strings.Join(enabledExportServices, ", ")
}

func MultiServiceValidValues() []string {
	allServices := []string{
		ENUM_SERVICE_PINGFEDERATE,
		ENUM_SERVICE_PINGONE_PLATFORM,
		ENUM_SERVICE_PINGONE_SSO,
		ENUM_SERVICE_PINGONE_MFA,
		ENUM_SERVICE_PINGONE_PROTECT,
	}

	slices.Sort(allServices)

	return allServices
}
