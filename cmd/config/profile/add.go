package profile

import (
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
		Use:   "add",
		Short: "Command to add a new configuration profile to pingctl.",
		Long:  `Command to add a new configuration profile to pingctl.`,
		RunE:  ConfigProfileAddRunE,
	}

	// Add flags that are not tracked in the viper configuration file
	cmd.Flags().StringVar(&profileName, "name", "", "(Optional) Set the name of the new profile.")
	cmd.Flags().StringVar(&description, profiles.ProfileDescriptionOption.CobraParamName, "", "(Optional) Set the description of the new profile.")
	cmd.Flags().BoolVar(&setActive, "set-active", false, "(Optional) Set the new profile as the active profile for pingctl.")

	// create flag variable to determine if the boolean flag is default or changed value
	// If default, we will need to prompt the user to decide if they want to set the profile as active
	setActiveFlag = cmd.Flags().Lookup("set-active")
	return cmd
}

func ConfigProfileAddRunE(cmd *cobra.Command, args []string) error {
	l := logger.Get()
	l.Debug().Msgf("Config Profile Add Subcommand Called.")

	if err := profile_internal.RunInternalConfigProfileAdd(profileName, description, setActive, setActiveFlag.Changed); err != nil {
		return err
	}

	return nil
}
