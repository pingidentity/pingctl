package profiles_test

import (
	"slices"
	"testing"

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
