package customtypes_test

import (
	"slices"
	"testing"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

// Test the custom type OutputFormat Set method with a valid value
func TestOutputFormat_SetValid(t *testing.T) {
	outputFormat := customtypes.OutputFormat("text")
	err := outputFormat.Set("json")
	testutils.CheckExpectedError(t, err, nil)
}

// Test the custom type OutputFormat Set method with an invalid value
func TestOutputFormat_SetInvalid(t *testing.T) {
	expectedErrorPattern := `unrecognized Output Format: 'INVALID'. Must be one of: [a-z\s,]+`
	outputFormat := customtypes.OutputFormat("text")
	err := outputFormat.Set("INVALID")
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test the custom type OutputFormat Type method
func TestOutputFormat_Type(t *testing.T) {
	outputFormat := customtypes.OutputFormat("text")
	typeValue := outputFormat.Type()
	if typeValue != "string" {
		t.Errorf("Expected 'string' but got '%s'", typeValue)
	}
}

// Test the custom type OutputFormat String method
func TestOutputFormat_String(t *testing.T) {
	outputFormat := customtypes.OutputFormat("text")
	stringValue := outputFormat.String()
	if stringValue != "text" {
		t.Errorf("Expected 'text' but got '%s'", stringValue)
	}
}

// Test the custom type OutputFormat OutputFormatValidValues method
func TestOutputFormat_OutputFormatValidValues(t *testing.T) {
	expectedValues := []string{customtypes.ENUM_OUTPUT_FORMAT_TEXT, customtypes.ENUM_OUTPUT_FORMAT_JSON}
	validValues := customtypes.OutputFormatValidValues()

	if len(validValues) != len(expectedValues) {
		t.Errorf("Expected %d valid values but got %d", len(expectedValues), len(validValues))
	}

	for _, expectedValue := range expectedValues {
		if !slices.Contains(validValues, expectedValue) {
			t.Errorf("Expected value '%s' is not in %v", expectedValue, validValues)
		}
	}
}
