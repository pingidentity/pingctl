package profiles_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
	"github.com/spf13/viper"
)

// Test SetProfileViperWithProfile function
func TestSetProfileViperWithProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Test SetProfileViperWithProfile function
	err := profiles.SetProfileViperWithProfile("default")
	if err != nil {
		t.Errorf("SetProfileViperWithProfile returned error: %v", err)
	}

	profileViper := profiles.GetProfileViper()
	if profileViper == nil {
		t.Errorf("GetProfileViper returned nil")
	}

	val := profileViper.Get(profiles.ProfileDescriptionOption.ViperKey)
	if val == nil {
		t.Errorf("Get returned nil")
	}
	if val != "default description" {
		t.Errorf("Get returned %s, expected 'default description'", val)
	}

	err = profiles.SetProfileViperWithProfile("production")
	if err != nil {
		t.Errorf("SetProfileViperWithProfile returned error: %v", err)
	}

	err = profiles.SetProfileViperWithProfile("test")
	if err == nil {
		t.Errorf("SetProfileViperWithProfile returned nil, expected error")
	}

	err = profiles.SetProfileViperWithProfile("invalid(*^&^%&%&^$)")
	if err == nil {
		t.Errorf("SetProfileViperWithProfile returned nil, expected error")
	}
}

// Test SetProfileViperWithViper function
func TestSetProfileViperWithViper(t *testing.T) {
	testutils_viper.InitVipers(t)

	testDesc := "test viper description"

	// Create new viper
	testViper := viper.New()
	testViper.Set(profiles.ProfileDescriptionOption.ViperKey, testDesc)

	// Test SetProfileViperWithViper function
	profiles.SetProfileViperWithViper(testViper, "test")
	profileViper := profiles.GetProfileViper()

	val := profileViper.Get(profiles.ProfileDescriptionOption.ViperKey)
	if val == nil {
		t.Errorf("Get returned nil")
	}
	if val != testDesc {
		t.Errorf("Get returned %s, expected %s", val, testDesc)
	}
}

// Test GetProfileViper function
func TestGetProfileViper(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Test GetProfileViper function
	profileViper := profiles.GetProfileViper()
	if profileViper == nil {
		t.Errorf("GetProfileViper returned nil")
	}

	val := profileViper.Get(profiles.ProfileDescriptionOption.ViperKey)
	if val == nil {
		t.Errorf("Get returned nil")
	}
	if val != "default description" {
		t.Errorf("Get returned %s, expected 'default description'", val)
	}
}

// Test CreateNewProfile function
func TestCreateNewProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Test CreateNewProfile function
	err := profiles.CreateNewProfile("test", "test description", false)
	if err != nil {
		t.Errorf("CreateNewProfile returned error: %v", err)
	}

	profileViper := profiles.GetProfileViper()
	if profileViper == nil {
		t.Errorf("GetProfileViper returned nil")
	}

	val := profileViper.Get(profiles.ProfileDescriptionOption.ViperKey)
	if val == nil {
		t.Errorf("Get returned nil")
	}
	if val != "default description" {
		t.Errorf("Get returned %s, expected 'default description'", val)
	}

	// Test CreateNewProfile function with setActive true
	err = profiles.CreateNewProfile("test2", "test description 2", true)
	if err != nil {
		t.Errorf("CreateNewProfile returned error: %v", err)
	}

	// Use the new profile
	err = profiles.SetProfileViperWithProfile("test2")
	if err != nil {
		t.Errorf("SetProfileViperWithProfile returned error: %v", err)
	}

	profileViper = profiles.GetProfileViper()
	if profileViper == nil {
		t.Errorf("GetProfileViper returned nil")
	}

	val = profileViper.Get(profiles.ProfileDescriptionOption.ViperKey)
	if val == nil {
		t.Errorf("Get returned nil")
	}
	if val != "test description 2" {
		t.Errorf("Get returned %s, expected 'test description 2'", val)
	}

	// CHeck profile names
	profileKeys := profiles.ConfigProfileNames()
	if len(profileKeys) != 4 {
		t.Errorf("ConfigProfileNames returned %d profiles, expected 4", len(profileKeys))
	}

	//Check active profile
	activeProfile := profiles.GetConfigActiveProfile()
	if activeProfile != "test2" {
		t.Errorf("GetConfigActiveProfile returned %s, expected 'test2'", activeProfile)
	}

}
