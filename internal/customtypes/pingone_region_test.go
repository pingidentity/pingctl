package customtypes_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/testing/testutils"
)

// Test PingOneRegion Set function
func Test_PingOneRegion_Set(t *testing.T) {
	pingoneRegion := new(customtypes.PingOneRegion)

	err := pingoneRegion.Set(customtypes.ENUM_PINGONE_REGION_CA)
	if err != nil {
		t.Errorf("Set returned error: %v", err)
	}
}

// Test Set function fails with invalid value
func Test_PingOneRegion_Set_InvalidValue(t *testing.T) {
	pingoneRegion := new(customtypes.PingOneRegion)

	invalidValue := "invalid"

	expectedErrorPattern := `^unrecognized PingOne Region: '.*'\. Must be one of: .*$`
	err := pingoneRegion.Set(invalidValue)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test Set function fails with nil
func Test_PingOneRegion_Set_Nil(t *testing.T) {
	var pingoneRegion *customtypes.PingOneRegion

	val := customtypes.ENUM_PINGONE_REGION_CA

	expectedErrorPattern := `^failed to set PingOne Region value: .* PingOne Region is nil$`
	err := pingoneRegion.Set(val)
	testutils.CheckExpectedError(t, err, &expectedErrorPattern)
}

// Test String function
func Test_PingOneRegion_String(t *testing.T) {
	pingoneRegion := customtypes.PingOneRegion(customtypes.ENUM_PINGONE_REGION_CA)

	expected := customtypes.ENUM_PINGONE_REGION_CA
	actual := pingoneRegion.String()
	if actual != expected {
		t.Errorf("String returned: %s, expected: %s", actual, expected)
	}
}
