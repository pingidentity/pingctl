package connector

import (
	"context"

	sdk "github.com/patrickcping/pingone-go-sdk-v2/pingone"
)

type ImportBlock struct {
	ResourceType string
	ResourceName string
	ResourceID   string
}

type SDKClientInfo struct {
	Context             context.Context
	ApiClient           *sdk.Client
	ExportEnvironmentID string
}

// A connector that allows exporting configuration
type ExportableResource interface {
	ExportAll() (*[]ImportBlock, error)
	ResourceType() string
}
