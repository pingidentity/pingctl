package customtypes_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

// Test custom type MultiService NewMultiService() method
func TestMultiService_NewMultiService(t *testing.T) {
	multiService := customtypes.NewMultiService()
	if multiService == nil {
		t.Error("Expected non-nil MultiService object")
	}
}

// Test custom type MultiService GetServices() method
func TestMultiService_GetServices(t *testing.T) {
	multiService := customtypes.NewMultiService()
	expectedService := customtypes.ENUM_SERVICE_PINGONE_PLATFORM
	err := multiService.Set(expectedService)
	testutils.CheckExpectedError(t, err, nil)

	services := multiService.GetServices()
	if services == nil || *services == nil {
		t.Fatal("Expected non-nil services slice")
	}

	if len(*services) != 1 {
		t.Errorf("Expected 1 service but got %d", len(*services))
	}

	if (*services)[0] != expectedService {
		t.Errorf("Expected service '%s' but got '%s'", expectedService, (*services)[0])
	}
}

// Test custom type MultiService Set() method with a valid value
func TestMultiService_SetValid(t *testing.T) {
	multiService := customtypes.NewMultiService()
	err := multiService.Set(customtypes.ENUM_SERVICE_PINGONE_PLATFORM)
	testutils.CheckExpectedError(t, err, nil)
}

// Test custom type MultiService Set() method with an invalid value
func TestMultiService_SetInvalid(t *testing.T) {
	expectedErrorPattern := `unrecognized service 'INVALID'. Must be one of: [a-z\-]+`
	multiService := customtypes.NewMultiService()
	err := multiService.Set("INVALID")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test custom type MultiService Type() method
func TestMultiService_Type(t *testing.T) {
	multiService := customtypes.NewMultiService()
	typeValue := multiService.Type()
	if typeValue != "string" {
		t.Errorf("Expected 'string' but got '%s'", typeValue)
	}
}

// Test custom type MultiService String() method
func TestMultiService_String(t *testing.T) {
	expectedServicesStr := strings.Join(*customtypes.NewMultiService().GetServices(), ", ")
	multiService := customtypes.NewMultiService()
	stringValue := multiService.String()
	if stringValue != expectedServicesStr {
		t.Errorf("Expected '%s' but got '%s'", expectedServicesStr, stringValue)
	}
}

// Test custom type MultiService MultiServiceValidValues() method
func TestMultiService_MultiServiceValidValues(t *testing.T) {
	expectedValues := *customtypes.NewMultiService().GetServices()
	validValues := customtypes.MultiServiceValidValues()

	if len(validValues) != len(expectedValues) {
		t.Errorf("Expected %d valid values but got %d", len(expectedValues), len(validValues))
	}

	for _, expectedValue := range expectedValues {
		if !slices.Contains(validValues, expectedValue) {
			t.Errorf("Expected value '%s' is not in %v", expectedValue, validValues)
		}
	}
}
