package configuration

import (
	"github.com/pingidentity/pingctl/internal/customtypes"
	"github.com/spf13/pflag"
)

// Options
var (
	ConfigProfileOption     Option
	ConfigNameOption        Option
	ConfigDescriptionOption Option

	ConfigAddProfileDescriptionOption Option
	ConfigAddProfileNameOption        Option
	ConfigAddProfileSetActiveOption   Option

	ConfigDeleteProfileOption Option

	ConfigViewProfileOption Option

	ConfigSetActiveProfileOption Option

	ConfigGetProfileOption Option

	ConfigSetProfileOption Option

	ConfigUnsetProfileOption Option
)

func initConfigOptions() {
	initConfigProfileOption()
	initConfigNameOption()
	initConfigDescriptionOption()

	initAddProfileDescriptionOption()
	initAddProfileNameOption()
	initAddProfileSetActiveOption()

	initDeleteProfileOption()

	initViewProfileOption()

	initSetActiveProfileOption()

	initGetProfileOption()

	initSetProfileOption()

	initUnsetProfileOption()
}

func initConfigProfileOption() {
	cobraParamName := "profile"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	ConfigProfileOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "p",
			Usage:     "The name of the profile to update.",
			Value:     cobraValue,
			DefValue:  "The active profile",
		},
		Type:     ENUM_STRING,
		ViperKey: "", // No viper key
	}
}

func initConfigNameOption() {
	cobraParamName := "name"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	ConfigNameOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "n",
			Usage:     "The new name for the profile.",
			Value:     cobraValue,
			DefValue:  "",
		},
		Type:     ENUM_STRING,
		ViperKey: "", // No viper key
	}
}

func initConfigDescriptionOption() {
	cobraParamName := "description"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	ConfigDescriptionOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "d",
			Usage:     "The new description for the profile.",
			Value:     cobraValue,
			DefValue:  "",
		},
		Type:     ENUM_STRING,
		ViperKey: "", // No viper key
	}
}

func initAddProfileDescriptionOption() {
	cobraParamName := "description"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	ConfigAddProfileDescriptionOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "d",
			Usage:     "The description of the new configuration profile.",
			Value:     cobraValue,
			DefValue:  "",
		},
		Type:     ENUM_STRING,
		ViperKey: "", // No viper key
	}
}

func initAddProfileNameOption() {
	cobraParamName := "name"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	ConfigAddProfileNameOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "n",
			Usage:     "The name of the new configuration profile.",
			Value:     cobraValue,
			DefValue:  "",
		},
		Type:     ENUM_STRING,
		ViperKey: "", // No viper key
	}
}

func initAddProfileSetActiveOption() {
	cobraParamName := "set-active"
	cobraValue := new(customtypes.Bool)
	defaultValue := customtypes.Bool(false)

	ConfigAddProfileSetActiveOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "s",
			Usage:     "Set the new configuration profile as the active profile for pingctl.",
			Value:     cobraValue,
			DefValue:  "false",
		},
		Type:     ENUM_BOOL,
		ViperKey: "", // No viper key
	}
}

func initDeleteProfileOption() {
	cobraParamName := "profile"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	ConfigDeleteProfileOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "p",
			Usage:     "The configuration profile to delete.",
			Value:     cobraValue,
			DefValue:  "",
		},
		Type:     ENUM_STRING,
		ViperKey: "", // No viper key
	}
}

func initViewProfileOption() {
	cobraParamName := "profile"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	ConfigViewProfileOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "p",
			Usage:     "The configuration profile name to view.",
			Value:     cobraValue,
			DefValue:  "The active profile",
		},
		Type:     ENUM_STRING,
		ViperKey: "", // No viper key
	}
}

func initSetActiveProfileOption() {
	cobraParamName := "profile"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	ConfigSetActiveProfileOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "p",
			Usage:     "The configuration profile to set as the active profile.",
			Value:     cobraValue,
			DefValue:  "",
		},
		Type:     ENUM_STRING,
		ViperKey: "", // No viper key
	}
}

func initGetProfileOption() {
	cobraParamName := "profile"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	ConfigGetProfileOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "p",
			Usage:     "The configuration profile used to get the configuration value from.",
			Value:     cobraValue,
			DefValue:  "The active profile",
		},
		Type:     ENUM_STRING,
		ViperKey: "", // No viper key
	}
}

func initSetProfileOption() {
	cobraParamName := "profile"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	ConfigSetProfileOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "p",
			Usage:     "The configuration profile used to set the configuration value.",
			Value:     cobraValue,
			DefValue:  "The active profile",
		},
		Type:     ENUM_STRING,
		ViperKey: "", // No viper key
	}
}

func initUnsetProfileOption() {
	cobraParamName := "profile"
	cobraValue := new(customtypes.String)
	defaultValue := customtypes.String("")

	ConfigUnsetProfileOption = Option{
		CobraParamName:  cobraParamName,
		CobraParamValue: cobraValue,
		DefaultValue:    &defaultValue,
		EnvVar:          "", // No environment variable
		Flag: &pflag.Flag{
			Name:      cobraParamName,
			Shorthand: "p",
			Usage:     "The configuration profile used to unset the configuration value.",
			Value:     cobraValue,
			DefValue:  "The active profile",
		},
		Type:     ENUM_STRING,
		ViperKey: "", // No viper key
	}
}
