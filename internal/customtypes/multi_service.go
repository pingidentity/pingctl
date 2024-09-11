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

type MultiService map[string]bool

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*MultiService)(nil)

// Implement pflag.Value interface for custom type in cobra MultiService parameter

func NewMultiService() *MultiService {
	ms := map[string]bool{
		ENUM_SERVICE_PINGFEDERATE:     true,
		ENUM_SERVICE_PINGONE_PLATFORM: true,
		ENUM_SERVICE_PINGONE_SSO:      true,
		ENUM_SERVICE_PINGONE_MFA:      true,
		ENUM_SERVICE_PINGONE_PROTECT:  true,
	}

	return (*MultiService)(&ms)
}

func (ms MultiService) GetServices() []string {
	enabledExportServices := []string{}

	if ms == nil {
		return enabledExportServices
	}

	for k, v := range ms {
		if v {
			enabledExportServices = append(enabledExportServices, k)
		}
	}

	slices.Sort(enabledExportServices)

	return enabledExportServices
}

func (ms *MultiService) Set(services string) error {
	if ms == nil {
		return fmt.Errorf("failed to set MultiService value: %s. MultiService is nil", services)
	}

	*ms = map[string]bool{}

	serviceList := strings.Split(services, ",")

	for _, service := range serviceList {
		switch service {
		case ENUM_SERVICE_PINGFEDERATE:
			(*ms)[ENUM_SERVICE_PINGFEDERATE] = true
		case ENUM_SERVICE_PINGONE_PLATFORM:
			(*ms)[ENUM_SERVICE_PINGONE_PLATFORM] = true
		case ENUM_SERVICE_PINGONE_SSO:
			(*ms)[ENUM_SERVICE_PINGONE_SSO] = true
		case ENUM_SERVICE_PINGONE_MFA:
			(*ms)[ENUM_SERVICE_PINGONE_MFA] = true
		case ENUM_SERVICE_PINGONE_PROTECT:
			(*ms)[ENUM_SERVICE_PINGONE_PROTECT] = true
		default:
			return fmt.Errorf("unrecognized service '%s'. Must be one of: %s", service, strings.Join(MultiServiceValidValues(), ", "))
		}
	}
	return nil
}

func (ms MultiService) ContainsPingOneService() bool {
	if ms == nil {
		return false
	}

	return ms[ENUM_SERVICE_PINGONE_PLATFORM] ||
		ms[ENUM_SERVICE_PINGONE_SSO] ||
		ms[ENUM_SERVICE_PINGONE_MFA] ||
		ms[ENUM_SERVICE_PINGONE_PROTECT]
}

func (ms MultiService) ContainsPingFederateService() bool {
	if ms == nil {
		return false
	}

	return ms[ENUM_SERVICE_PINGFEDERATE]
}

func (ms MultiService) Type() string {
	return "string"
}

func (ms MultiService) String() string {
	if ms == nil {
		return ""
	}

	enabledExportServices := ms.GetServices()

	if len(enabledExportServices) == 0 {
		return ""
	}

	return strings.Join(enabledExportServices, ",")
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
