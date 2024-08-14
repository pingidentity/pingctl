package profiles_test

import (
	"slices"
	"testing"

	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test ProfileKeys function
func TestProfileKeys(t *testing.T) {
	testutils_viper.InitVipers(t)

	expectedNumKeys := len(profiles.ConfigOptions.Options) - 1

	// Test ProfileKeys function
	keys := profiles.ProfileKeys()
	if len(keys) != expectedNumKeys {
		t.Errorf("Expected %d keys, but got %d", expectedNumKeys, len(keys))
	}

	// Make sure profile option is not in the list
	for _, key := range keys {
		if key == profiles.ProfileOption.ViperKey {
			t.Errorf("Profile option key should not be in the list")
		}
	}

	// Check random profile key is in the list
	if !slices.Contains(keys, profiles.PingOneWorkerClientIDOption.ViperKey) {
		t.Errorf("Expected key %s in the list", profiles.PingOneWorkerClientIDOption.ViperKey)
	}

	// Make sure the list is sorted
	if !slices.IsSorted(keys) {
		t.Errorf("Keys are not sorted")
	}
}

// Test ExpandedProfileKeys function
func TestExpandedProfileKeys(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Test ExpandedProfileKeys function
	keys := profiles.ExpandedProfileKeys()

	// Make sure keys is not empty
	if len(keys) == 0 {
		t.Errorf("Keys should not be empty")
	}

	// Make sure profile option is not in the list
	for _, key := range keys {
		if key == profiles.ProfileOption.ViperKey {
			t.Errorf("Profile option key should not be in the list")
		}
	}

	// Check random profile key is in the list
	if !slices.Contains(keys, profiles.PingOneWorkerClientIDOption.ViperKey) {
		t.Errorf("Expected key %s in the list", profiles.PingOneWorkerClientIDOption.ViperKey)
	}

	// Check random parent profile key is in the list
	if !slices.Contains(keys, "pingctl") {
		t.Errorf("Expected key %s in the list", "pingctl")
	}

	// Make sure the list is sorted
	if !slices.IsSorted(keys) {
		t.Errorf("Keys are not sorted")
	}
}

// Test OptionTypeFromViperKey function
func TestOptionTypeFromViperKey(t *testing.T) {
	// Test OptionTypeFromViperKey function
	optType, ok := profiles.OptionTypeFromViperKey(profiles.PingOneWorkerClientIDOption.ViperKey)
	if !ok {
		t.Errorf("Expected key %s to be found", profiles.PingOneWorkerClientIDOption.ViperKey)
	}

	// Check the type of the option
	if optType != profiles.PingOneWorkerClientIDOption.Type {
		t.Errorf("Expected type %s, but got %s", profiles.PingOneWorkerClientIDOption.Type, optType)
	}

	// Check random key is not found
	_, ok = profiles.OptionTypeFromViperKey("random")
	if ok {
		t.Errorf("Expected key %s to not be found", "random")
	}
}

// Test GetDefaultValue function
func TestGetDefaultValue(t *testing.T) {
	// Test GetDefaultValue function with bool type
	val, err := profiles.GetDefaultValue(profiles.ENUM_BOOL)
	if err != nil {
		t.Errorf("Expected error %v, but got %v", nil, err)
	}
	b, ok := val.(bool)
	if !ok || b != false {
		t.Errorf("Expected value %v, but got %v", false, val)
	}

	// Test GetDefaultValue function with string type
	val, err = profiles.GetDefaultValue(profiles.ENUM_STRING)
	if err != nil {
		t.Errorf("Expected error %v, but got %v", nil, err)
	}
	s, ok := val.(string)
	if !ok || s != "" {
		t.Errorf("Expected value %v, but got %v", "", val)
	}

	// Test GetDefaultValue function with string slice type
	val, err = profiles.GetDefaultValue(profiles.ENUM_STRING_SLICE)
	if err != nil {
		t.Errorf("Expected error %v, but got %v", nil, err)
	}
	ss, ok := val.([]string)
	if !ok || ss != nil {
		t.Errorf("Expected value %v, but got %v", "", val)
	}

	// Test GetDefaultValue function with UUID type
	val, err = profiles.GetDefaultValue(profiles.ENUM_ID)
	if err != nil {
		t.Errorf("Expected error %v, but got %v", nil, err)
	}
	s, ok = val.(string)
	if !ok || s != "" {
		t.Errorf("Expected value %v, but got %v", "", val)
	}

	// Test GetDefaultValue function with output format type
	val, err = profiles.GetDefaultValue(profiles.ENUM_OUTPUT_FORMAT)
	if err != nil {
		t.Errorf("Expected error %v, but got %v", nil, err)
	}
	o, ok := val.(customtypes.OutputFormat)
	if !ok || o != customtypes.OutputFormat("text") {
		t.Errorf("Expected value %v, but got %v", "text", val)
	}

	// Test GetDefaultValue function with pingone region type
	val, err = profiles.GetDefaultValue(profiles.ENUM_PINGONE_REGION)
	if err != nil {
		t.Errorf("Expected error %v, but got %v", nil, err)
	}
	p, ok := val.(customtypes.PingOneRegion)
	if !ok || p != customtypes.PingOneRegion("") {
		t.Errorf("Expected value %v, but got %v", "", val)
	}

	// Test GetDefaultValue function with random type
	_, err = profiles.GetDefaultValue("random")
	if err == nil {
		t.Errorf("Expected error, but got %v", nil)
	}
}
