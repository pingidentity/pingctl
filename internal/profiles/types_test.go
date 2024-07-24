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
	if !slices.Contains(keys, profiles.WorkerClientIDOption.ViperKey) {
		t.Errorf("Expected key %s in the list", profiles.WorkerClientIDOption.ViperKey)
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
	if !slices.Contains(keys, profiles.WorkerClientIDOption.ViperKey) {
		t.Errorf("Expected key %s in the list", profiles.WorkerClientIDOption.ViperKey)
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
	optType, ok := profiles.OptionTypeFromViperKey(profiles.WorkerClientIDOption.ViperKey)
	if !ok {
		t.Errorf("Expected key %s to be found", profiles.WorkerClientIDOption.ViperKey)
	}

	// Check the type of the option
	if optType != profiles.WorkerClientIDOption.Type {
		t.Errorf("Expected type %s, but got %s", profiles.WorkerClientIDOption.Type, optType)
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
	val := profiles.GetDefaultValue(profiles.ENUM_BOOL)
	b, ok := val.(bool)
	if !ok || b != false {
		t.Errorf("Expected value %v, but got %v", false, val)
	}

	// Test GetDefaultValue function with string type
	val = profiles.GetDefaultValue(profiles.ENUM_STRING)
	s, ok := val.(string)
	if !ok || s != "" {
		t.Errorf("Expected value %v, but got %v", "", val)
	}

	// Test GetDefaultValue function with UUID type
	val = profiles.GetDefaultValue(profiles.ENUM_ID)
	s, ok = val.(string)
	if !ok || s != "" {
		t.Errorf("Expected value %v, but got %v", "", val)
	}

	// Test GetDefaultValue function with output format type
	val = profiles.GetDefaultValue(profiles.ENUM_OUTPUT_FORMAT)
	o, ok := val.(customtypes.OutputFormat)
	if !ok || o != customtypes.OutputFormat("text") {
		t.Errorf("Expected value %v, but got %v", "text", val)
	}

	// Test GetDefaultValue function with pingone region type
	val = profiles.GetDefaultValue(profiles.ENUM_PINGONE_REGION)
	p, ok := val.(customtypes.PingOneRegion)
	if !ok || p != customtypes.PingOneRegion("") {
		t.Errorf("Expected value %v, but got %v", "", val)
	}

	// Test GetDefaultValue function with random type
	val = profiles.GetDefaultValue("random")
	if val != nil {
		t.Errorf("Expected value %v, but got %v", nil, val)
	}
}
