package profile_internal

import (
	"fmt"
	"strings"

	"github.com/pingidentity/pingctl/internal/output"
	"github.com/pingidentity/pingctl/internal/profiles"
)

func RunInternalConfigProfileList(args []string) {
	if len(args) > 0 {
		output.Print(output.Opts{
			Message: fmt.Sprintf("'pingctl config profile list' does not take additional arguments. Ignoring extra arguments: %s", strings.Join(args, " ")),
			Result:  output.ENUM_RESULT_NOACTION_WARN,
		})
	}

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
