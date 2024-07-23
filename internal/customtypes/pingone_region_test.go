package customtypes_test

import (
	"slices"
	"testing"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

// Test the custom type PingOneRegion Set method with a valid value
func TestPingOneRegion_SetValid(t *testing.T) {
	pingOneRegion := customtypes.PingOneRegion("AsiaPacific")
	err := pingOneRegion.Set("Europe")
	testutils.CheckExpectedError(t, err, nil)
}

// Test the custom type PingOneRegion Set method with an invalid value
func TestPingOneRegion_SetInvalid(t *testing.T) {
	expectedErrorPattern := `^unrecognized PingOne Region: 'INVALID'. Must be one of: [A-Za-z\s,]+$`
	pingOneRegion := customtypes.PingOneRegion("AsiaPacific")
	err := pingOneRegion.Set("INVALID")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test the custom type PingOneRegion Type method
func TestPingOneRegion_Type(t *testing.T) {
	pingOneRegion := customtypes.PingOneRegion("AsiaPacific")
	typeValue := pingOneRegion.Type()
	if typeValue != "string" {
		t.Errorf("Expected 'string' but got '%s'", typeValue)
	}
}

// Test the custom type PingOneRegion String method
func TestPingOneRegion_String(t *testing.T) {
	pingOneRegion := customtypes.PingOneRegion("AsiaPacific")
	stringValue := pingOneRegion.String()
	if stringValue != "AsiaPacific" {
		t.Errorf("Expected 'AsiaPacific' but got '%s'", stringValue)
	}
}

// Test the custom type PingOneRegion PingOneRegionValidValues method
func TestPingOneRegion_PingOneRegionValidValues(t *testing.T) {
	expectedValues := []string{
		customtypes.ENUM_PINGONE_REGION_AP,
		customtypes.ENUM_PINGONE_REGION_CA,
		customtypes.ENUM_PINGONE_REGION_EU,
		customtypes.ENUM_PINGONE_REGION_NA}
	validValues := customtypes.PingOneRegionValidValues()

	if len(validValues) != len(expectedValues) {
		t.Errorf("Expected %d valid values but got %d", len(expectedValues), len(validValues))
	}

	for _, expectedValue := range expectedValues {
		if !slices.Contains(validValues, expectedValue) {
			t.Errorf("Expected value '%s' is not in %v", expectedValue, validValues)
		}
	}
}
