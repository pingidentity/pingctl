package connector

const (
	ENUMEXPORTFORMAT_HCL = "HCL"
)

// A connector that allows exporting configuration
type Exportable interface {
	Export(format, outputDir string) error
}
