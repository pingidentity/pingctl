package customtypes

import (
	"fmt"

	"github.com/spf13/pflag"
)

const (
	ENUM_PINGONE_REGION_AP string = "AsiaPacific"
	ENUM_PINGONE_REGION_CA string = "Canada"
	ENUM_PINGONE_REGION_EU string = "Europe"
	ENUM_PINGONE_REGION_NA string = "NorthAmerica"
)

type PingOneRegion string

// Verify that the custom type satisfies the pflag.Value interface
var _ pflag.Value = (*PingOneRegion)(nil)

// Implement pflag.Value interface for custom type in cobra pingone-region parameter

func (s *PingOneRegion) Set(region string) error {
	switch region {
	case ENUM_PINGONE_REGION_AP, ENUM_PINGONE_REGION_CA, ENUM_PINGONE_REGION_EU, ENUM_PINGONE_REGION_NA:
		*s = PingOneRegion(region)
	default:
		return fmt.Errorf("unrecognized PingOne Region: '%s'. Must be one of: %s", region, PingOneRegionValidValues())
	}
	return nil
}

func (s *PingOneRegion) Type() string {
	return "string"
}

func (s *PingOneRegion) String() string {
	return string(*s)
}

func PingOneRegionValidValues() string {
	return fmt.Sprintf("'%s', '%s', '%s', '%s'", ENUM_PINGONE_REGION_AP, ENUM_PINGONE_REGION_CA, ENUM_PINGONE_REGION_EU, ENUM_PINGONE_REGION_NA)
}
