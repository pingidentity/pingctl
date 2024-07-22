package profile_internal

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test RunInternalConfigProfileList function with no args
func Example_runInternalConfigProfileList_noArgs() {
	testutils_viper.InitVipers(&testing.T{})

	RunInternalConfigProfileList([]string{})

	// Output:
	// pingctl profiles:
	//   * default
	//     production
}

// Test RunInternalConfigProfileList function with args
func Example_runInternalConfigProfileList_withArgs() {
	testutils_viper.InitVipers(&testing.T{})

	RunInternalConfigProfileList([]string{"arg1", "arg2"})

	// Output:
	// 'pingctl config profile list' does not take additional arguments. Ignoring extra arguments: arg1 arg2 - No Action (Warning)
	// pingctl profiles:
	//   * default
	//     production
}
