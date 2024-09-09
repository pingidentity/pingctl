package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

// Test MultiService NewMultiService function
func Test_MultiService_NewMultiService(t *testing.T) {
	multiService := customtypes.NewMultiService()

	if multiService == nil {
		t.Fatalf("NewMultiService returned nil")
	}

	if len(multiService.GetServices()) == 0 {
		t.Fatalf("NewMultiService returned empty Services")
	}
}

// Test MultiService Set function
func Test_MultiService_Set(t *testing.T) {
	multiService := customtypes.NewMultiService()

	service := customtypes.ENUM_SERVICE_PINGONE_MFA
	err := multiService.Set(service)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}

	services := multiService.GetServices()
	if len(services) != 1 {
		t.Errorf("GetServices returned: %v, expected: %v", services, service)
	}

	if services[0] != service {
		t.Errorf("GetServices returned: %v, expected: %v", services, service)
	}
}

// Test MultiService Set function with invalid value
func Test_MultiService_Set_InvalidValue(t *testing.T) {
	multiService := customtypes.NewMultiService()

	invalidValue := "invalid"
	expectedErrorPattern := `^unrecognized service '.*'. Must be one of: .*$`
	err := multiService.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test MultiService Set function with nil
func Test_MultiService_Set_Nil(t *testing.T) {
	var multiService *customtypes.MultiService

	service := customtypes.ENUM_SERVICE_PINGONE_MFA
	expectedErrorPattern := `^failed to set MultiService value: .* MultiService is nil$`
	err := multiService.Set(service)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test MultiService ContainsPingOneService function
func Test_MultiService_ContainsPingOneService(t *testing.T) {
	multiService := customtypes.NewMultiService()

	service := customtypes.ENUM_SERVICE_PINGONE_MFA
	err := multiService.Set(service)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}

	if !multiService.ContainsPingOneService() {
		t.Errorf("ContainsPingOneService returned false, expected true")
	}
}

// Test MultiService ContainsPingFederateService function
func Test_MultiService_ContainsPingFederateService(t *testing.T) {
	multiService := customtypes.NewMultiService()

	service := customtypes.ENUM_SERVICE_PINGFEDERATE
	err := multiService.Set(service)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}

	if !multiService.ContainsPingFederateService() {
		t.Errorf("ContainsPingFederateService returned false, expected true")
	}
}

// Test MultiService String function
func Test_MultiService_String(t *testing.T) {
	multiService := customtypes.NewMultiService()

	service := customtypes.ENUM_SERVICE_PINGONE_MFA
	err := multiService.Set(service)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}

	expected := service
	actual := multiService.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
