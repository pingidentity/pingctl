package customtypes

import (
	"fmt"
	"slices"
	"strings"

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

func (p *PingOneRegion) Set(region string) error {
	if p == nil {
		return fmt.Errorf("failed to set PingOne Region value: %s. PingOne Region is nil", region)
	}
	switch region {
	case ENUM_PINGONE_REGION_AP, ENUM_PINGONE_REGION_CA, ENUM_PINGONE_REGION_EU, ENUM_PINGONE_REGION_NA:
		*p = PingOneRegion(region)
	default:
		return fmt.Errorf("unrecognized PingOne Region: '%s'. Must be one of: %s", region, strings.Join(PingOneRegionValidValues(), ", "))
	}
	return nil
}

func (p PingOneRegion) Type() string {
	return "string"
}

func (p PingOneRegion) String() string {
	return string(p)
}

func PingOneRegionValidValues() []string {
	pingoneRegions := []string{
		ENUM_PINGONE_REGION_AP,
		ENUM_PINGONE_REGION_CA,
		ENUM_PINGONE_REGION_EU,
		ENUM_PINGONE_REGION_NA}

	slices.Sort(pingoneRegions)

	return pingoneRegions
}
