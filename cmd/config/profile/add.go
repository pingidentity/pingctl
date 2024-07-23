package profile

import (
	"os"

	"github.com/pingidentity/pingctl/cmd/common"
	profile_internal "github.com/pingidentity/pingctl/internal/commands/config/profile"
	"github.com/pingidentity/pingctl/internal/logger"
	"github.com/pingidentity/pingctl/internal/profiles"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	profileName string
	description string
	setActive   bool

	setActiveFlag *pflag.Flag
)

func NewConfigProfileAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:                  common.ExactArgs(0),
		DisableFlagsInUseLine: true, // We write our own flags in @Use attribute
		Example: `pingctl config profile add
pingctl config profile add --name my-profile
pingctl config profile add --name my-profile --set-active
pingctl config profile add --name my-profile --description "My new profile"`,
		Long:  `Command to add a new configuration profile to pingctl.`,
		RunE:  ConfigProfileAddRunE,
		Short: "Command to add a new configuration profile to pingctl.",
		Use:   "add [flags]",
	}

	// Add flags that are not tracked in the viper configuration file
	cmd.Flags().StringVarP(&profileName, "name", "n", "", "Set the name of the new profile.")
	cmd.Flags().StringVarP(&description, profiles.ProfileDescriptionOption.CobraParamName, "d", "", "Set the description of the new profile.")
	cmd.Flags().BoolVarP(&setActive, "set-active", "s", false, "Set the new profile as the active profile for pingctl.")

	// create flag variable to determine if the boolean flag is default or changed value
	// If default, we will need to prompt the user to decide if they want to set the profile as active
	setActiveFlag = cmd.Flags().Lookup("set-active")
	return cmd
}

func ConfigProfileAddRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Profile Add Subcommand Called.")

	if err := profile_internal.RunInternalConfigProfileAdd(profileName, description, setActive, setActiveFlag.Changed, os.Stdin); err != nil {
		return err
	}

	return nil
}
