//go:build tools
// +build tools

package tools

// Manage tool dependencies via go.mod.
//
//nolint:all
import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/pavius/impi"
)
