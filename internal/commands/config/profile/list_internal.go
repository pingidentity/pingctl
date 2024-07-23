package profile_internal

import (
	"fmt"

	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/profiles"
)

func RunInternalConfigProfileList() {
	activeProfileName := profiles.GetConfigActiveProfile()
	profileNames := profiles.ConfigProfileNames()
	listOutputString := "pingctl profiles:\n"

	for _, pName := range profileNames {
		if pName == activeProfileName {
			listOutputString += fmt.Sprintf("  * %s\n", pName)
		} else {
			listOutputString += fmt.Sprintf("    %s\n", pName)
		}
	}

	output.Print(output.Opts{
		Message: listOutputString,
		Result:  output.ENUM_RESULT_NIL,
	})
}
