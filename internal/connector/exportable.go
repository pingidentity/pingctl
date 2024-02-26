package connector

import (
	_ "embed"
)

const (
	ENUMEXPORTFORMAT_HCL = "HCL"

	ENUMREGION_AP = "AsiaPacific"
	ENUMREGION_CA = "Canada"
	ENUMREGION_EU = "Europe"
	ENUMREGION_NA = "NorthAmerica"
)

// Embed import block template needed for export generation
//
//go:embed templates/hcl_import_block.template
var HCLImportBlockTemplate string

// A connector that allows exporting configuration
type Exportable interface {
	Export(format, outputDir string, overwriteExport bool) error
	ConnectorServiceName() string
}
