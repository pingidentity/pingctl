package connector

type ImportBlock struct {
	ResourceType string
	ResourceName string
	ResourceID   string
}

// A connector that allows exporting configuration
type ExportableResource interface {
	ExportAll() (*[]ImportBlock, error)
	ResourceType() string
}
