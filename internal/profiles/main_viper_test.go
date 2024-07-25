package profiles_test

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test GetMainViper function
func TestGetMainViper(t *testing.T) {
	testutils_viper.InitVipers(t)

	v := profiles.GetMainViper()
	if v == nil {
		t.Errorf("GetMainViper returned nil")
	}
}

// Test GetConfigActiveProfile function
func TestGetConfigActiveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	profile := profiles.GetConfigActiveProfile()
	if profile == "" {
		t.Errorf("GetConfigActiveProfile returned empty string")
	}

	if profile != "default" {
		t.Errorf("GetConfigActiveProfile returned %s, expected 'default'", profile)
	}
}

// Test SetConfigActiveProfile function
func TestSetConfigActiveProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	pName := "test"

	err := profiles.SetConfigActiveProfile(pName)
	if err != nil {
		t.Errorf("SetConfigActiveProfile returned error: %v", err)
	}

	profile := profiles.GetConfigActiveProfile()
	if profile != pName {
		t.Errorf("GetConfigActiveProfile returned '%s', expected '%s'", profile, pName)
	}
}

// Test ConfigProfileNames function
func TestConfigProfileNames(t *testing.T) {
	testutils_viper.InitVipers(t)

	profileKeys := profiles.ConfigProfileNames()
	if len(profileKeys) == 0 {
		t.Errorf("ConfigProfileNames returned empty slice")
	}

	if len(profileKeys) != 2 {
		t.Errorf("ConfigProfileNames returned %d profiles, expected 2", len(profileKeys))
	}

	if profileKeys[0] != "default" {
		t.Errorf("ConfigProfileNames returned %s, expected 'default'", profileKeys[0])
	}

	if profileKeys[1] != "production" {
		t.Errorf("ConfigProfileNames returned %s, expected 'production'", profileKeys[1])
	}
}

// Test ValidateNewProfileName function
func TestValidateNewProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := profiles.ValidateNewProfileName("")
	if err == nil {
		t.Errorf("ValidateNewProfileName returned nil, expected error")
	}

	err = profiles.ValidateNewProfileName("default")
	if err == nil {
		t.Errorf("ValidateNewProfileName returned nil, expected error")
	}

	err = profiles.ValidateNewProfileName("production")
	if err == nil {
		t.Errorf("ValidateNewProfileName returned nil, expected error")
	}

	err = profiles.ValidateNewProfileName("test")
	if err != nil {
		t.Errorf("ValidateNewProfileName returned error: %v", err)
	}

	err = profiles.ValidateNewProfileName("invalid(*^&^%&%&^$)")
	if err == nil {
		t.Errorf("ValidateNewProfileName returned nil, expected error")
	}
}

// Test ValidateExistingProfileName function
func TestValidateExistingProfileName(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := profiles.ValidateExistingProfileName("")
	if err == nil {
		t.Errorf("ValidateExistingProfileName returned nil, expected error")
	}

	err = profiles.ValidateExistingProfileName("default")
	if err != nil {
		t.Errorf("ValidateExistingProfileName returned error: %v", err)
	}

	err = profiles.ValidateExistingProfileName("production")
	if err != nil {
		t.Errorf("ValidateExistingProfileName returned error: %v", err)
	}

	err = profiles.ValidateExistingProfileName("test")
	if err == nil {
		t.Errorf("ValidateExistingProfileName returned nil, expected error")
	}

	err = profiles.ValidateExistingProfileName("invalid(*^&^%&%&^$)")
	if err == nil {
		t.Errorf("ValidateExistingProfileName returned nil, expected error")
	}
}

// Test ValidateProfileNameFormat function
func TestValidateProfileNameFormat(t *testing.T) {
	err := profiles.ValidateProfileNameFormat("")
	if err == nil {
		t.Errorf("ValidateProfileNameFormat returned nil, expected error")
	}

	err = profiles.ValidateProfileNameFormat("default")
	if err != nil {
		t.Errorf("ValidateProfileNameFormat returned error: %v", err)
	}

	err = profiles.ValidateProfileNameFormat("production")
	if err != nil {
		t.Errorf("ValidateProfileNameFormat returned error: %v", err)
	}

	err = profiles.ValidateProfileNameFormat("test")
	if err != nil {
		t.Errorf("ValidateProfileNameFormat returned error: %v", err)
	}

	err = profiles.ValidateProfileNameFormat("invalid(*^&^%&%&^$)")
	if err == nil {
		t.Errorf("ValidateProfileNameFormat returned nil, expected error")
	}
}

// Test DeleteConfigProfile function
func TestDeleteConfigProfile(t *testing.T) {
	testutils_viper.InitVipers(t)

	err := profiles.DeleteConfigProfile("")
	if err == nil {
		t.Errorf("DeleteConfigProfile returned nil, expected error")
	}

	err = profiles.DeleteConfigProfile("default")
	if err == nil {
		t.Errorf("DeleteConfigProfile returned nil, expected error")
	}

	err = profiles.DeleteConfigProfile("production")
	if err != nil {
		t.Errorf("DeleteConfigProfile returned error: %v", err)
	}

	err = profiles.DeleteConfigProfile("test")
	if err == nil {
		t.Errorf("DeleteConfigProfile returned nil, expected error")
	}

	err = profiles.DeleteConfigProfile("invalid(*^&^%&%&^$)")
	if err == nil {
		t.Errorf("DeleteConfigProfile returned nil, expected error")
	}

	profileKeys := profiles.ConfigProfileNames()
	if len(profileKeys) != 1 {
		t.Errorf("ConfigProfileNames returned %d profiles, expected 1", len(profileKeys))
	}

	if profileKeys[0] != "default" {
		t.Errorf("ConfigProfileNames returned %s, expected 'default'", profileKeys[0])
	}
}

// Test SaveProfileViperToFile function
func TestSaveProfileViperToFile(t *testing.T) {
	testutils_viper.InitVipers(t)

	// Create a new profile
	err := profiles.CreateNewProfile("test", "test", true)
	if err != nil {
		t.Errorf("CreateNewProfile returned error: %v", err)
	}

	// Use the new profile
	err = profiles.SetConfigActiveProfile("test")
	if err != nil {
		t.Errorf("SetConfigActiveProfile returned error: %v", err)
	}

	err = profiles.SetProfileViperWithProfile("test")
	if err != nil {
		t.Errorf("SetProfileViperWithProfile returned error: %v", err)
	}

	// Save the new profile to file
	err = profiles.SaveProfileViperToFile()
	if err != nil {
		t.Errorf("SaveProfileViperToFile returned error: %v", err)
	}

	// Check if the new profile was saved to file
	profileKeys := profiles.ConfigProfileNames()
	if len(profileKeys) != 3 {
		t.Errorf("ConfigProfileNames returned %d profiles, expected 3", len(profileKeys))
	}

	if profileKeys[0] != "default" {
		t.Errorf("ConfigProfileNames returned %s, expected 'default'", profileKeys[0])
	}

	if profileKeys[1] != "production" {
		t.Errorf("ConfigProfileNames returned %s, expected 'production'", profileKeys[1])
	}

	if profileKeys[2] != "test" {
		t.Errorf("ConfigProfileNames returned %s, expected 'test'", profileKeys[2])
	}
}
