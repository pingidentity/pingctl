package connector

import (
	"context"
	"fmt"
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
	// Remove all non-Alphanumeric characters/non-underscores
	b.ResourceName = regexp.MustCompile(`[^a-zA-Z0-9_]+`).ReplaceAllString(b.ResourceName, "")
	// Make everything lowercase
	b.ResourceName = strings.ToLower(b.ResourceName)
}

func (b *ImportBlock) Equals(a ImportBlock) bool {
	if a.ResourceType != b.ResourceType {
		return false
	}

	if a.ResourceName != b.ResourceName {
		return false
	}

	if a.ResourceID != b.ResourceID {
		return false
	}

	return true
}

func (b *ImportBlock) String() string {
	pattern := `import {
		to = %s.%s
		id = "%s"
	}`
	return fmt.Sprintf(pattern, b.ResourceType, b.ResourceName, b.ResourceID)
}
