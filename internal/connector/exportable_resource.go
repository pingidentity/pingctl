package connector

import (
	"context"
	"regexp"
	"strings"

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

func (b *ImportBlock) Sanitize() {

	// Replace spaces with underscores
	b.ResourceName = strings.ReplaceAll(b.ResourceName, " ", "_")
	b.ResourceType = strings.ReplaceAll(b.ResourceType, " ", "_")
	// Remove all non-Alphanumeric characters/non-underscores
	re := regexp.MustCompile(`[^a-zA-Z0-9_]+`)
	b.ResourceName = re.ReplaceAllString(b.ResourceName, "")
	b.ResourceType = re.ReplaceAllString(b.ResourceType, "")
	// Make everything lowercase
	b.ResourceName = strings.ToLower(b.ResourceName)
	b.ResourceType = strings.ToLower(b.ResourceType)
}
