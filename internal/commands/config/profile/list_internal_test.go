package profile_internal

import (
	"testing"

	"github.com/pingidentity/pingctl/internal/testing/testutils_viper"
)

// Test RunInternalConfigProfileList function
func Example_runInternalConfigProfileList() {
	testutils_viper.InitVipers(&testing.T{})

	RunInternalConfigProfileList()

	// Output:
	// pingctl profiles:
	//   * default
	//     production
}
