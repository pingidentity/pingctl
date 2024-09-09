package configuration

import (
	"github.com/pingidentity/pingctl/internal/customtypes"
)

// Options
var (
	ProfileDescriptionOption Option
)

func initProfilesOptions() {
	initDescriptionOption()
}

func initDescriptionOption() {
	ProfileDescriptionOption = Option{
		CobraParamName:  "",  // No cobra param name
		CobraParamValue: nil, // No cobra param value
		DefaultValue:    new(customtypes.String),
		EnvVar:          "",  // No environment variable
		Flag:            nil, // No flag
		Type:            ENUM_STRING,
		ViperKey:        "description",
	}
}
