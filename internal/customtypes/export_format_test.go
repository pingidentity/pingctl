package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

// Test the custom type ExportFormat Set method with a valid value
func TestExportFormat_SetValid(t *testing.T) {
	exportFormat := customtypes.ExportFormat("HCL")
	err := exportFormat.Set("HCL")
	testutils.CheckExpectedError(t, err, nil)
}

// Test the custom type ExportFormat Set method with an invalid value
func TestExportFormat_SetInvalid(t *testing.T) {
	expectedErrorPattern := `unrecognized export format 'INVALID'. Must be one of: [A-Z]+`
	exportFormat := customtypes.ExportFormat("HCL")
	err := exportFormat.Set("INVALID")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test the custom type ExportFormat Type method
func TestExportFormat_Type(t *testing.T) {
	exportFormat := customtypes.ExportFormat("HCL")
	typeValue := exportFormat.Type()
	if typeValue != "string" {
		t.Errorf("Expected 'string' but got '%s'", typeValue)
	}
}

// Test the custom type ExportFormat String method
func TestExportFormat_String(t *testing.T) {
	exportFormat := customtypes.ExportFormat("HCL")
	stringValue := exportFormat.String()
	if stringValue != "HCL" {
		t.Errorf("Expected 'HCL' but got '%s'", stringValue)
	}
}

// Test the custom type ExportFormat ExportFormatValidValues method
func TestExportFormat_ExportFormatValidValues(t *testing.T) {
	expectedValues := []string{"HCL"}
	validValues := customtypes.ExportFormatValidValues()

	if len(validValues) != len(expectedValues) {
		t.Errorf("Expected %d valid values but got %d", len(expectedValues), len(validValues))
	}

	for i, expectedValue := range expectedValues {
		if validValues[i] != expectedValue {
			t.Errorf("Expected '%s' but got '%s'", expectedValue, validValues[i])
		}
	}
}
