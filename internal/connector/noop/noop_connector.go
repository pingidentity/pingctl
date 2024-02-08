package noop

import "github.com/pingidentity/pingctl/internal/connector"

// Verify that the connector satisfies the expected interfaces
var (
	_ connector.Exportable      = &NoopConnector{}
	_ connector.Authenticatable = &NoopConnector{}
)

// A sample connector implementation that doesn't do anything
type NoopConnector struct{}

// Utility method for creating a NoopConnector
func Connector() *NoopConnector {
	return &NoopConnector{}
}

func (c *NoopConnector) Export(format, outputDir string) error {
	//no-op
	println("No op export")
	return nil
}

func (c *NoopConnector) Login() error {
	//no-op
	println("No op login")
	return nil
}

func (c *NoopConnector) Logout() error {
	//no-op
	println("No op logout")
	return nil
}
