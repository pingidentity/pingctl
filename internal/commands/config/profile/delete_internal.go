package profile_internal

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/profiles"
)

func RunInternalConfigProfileDelete(pName string) (err error) {
	err = profiles.DeleteConfigProfile(pName)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %v", err)
	}

	output.Print(output.Opts{
		Message: fmt.Sprintf("Profile '%s' deleted successfully", pName),
		Result:  output.ENUM_RESULT_SUCCESS,
	})

	return nil
}
