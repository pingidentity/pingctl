package common

import (
	"fmt"
	"net/http"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/pingctl/internal/logger"
)

func HandleClientResponse(response *http.Response, err error, apiFunctionName string, resourceType string) error {
	l := logger.Get()
	defer response.Body.Close()

	if err != nil {
		l.Error().Err(err).Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return err
	}

	if response.StatusCode == 404 {
		l.Error().Msgf("%s Request was not successful. Resources %s not found", apiFunctionName, resourceType)
		l.Error().Err(err).Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return fmt.Errorf("failed to fetch %s resources via %s()", resourceType, apiFunctionName)
	}

	if response.StatusCode >= 300 {
		l.Error().Msgf("%s Request was not successful", apiFunctionName)
		l.Error().Err(err).Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return fmt.Errorf("failed to fetch %s resources via %s()", resourceType, apiFunctionName)
	}
	return nil
}

// Executes the function apiExecuteFunc for the ManagementAPIClient
// Handles err and response if Client call failed
// Returns embedded data if not nil
// Treats nil embedded data as an error
func GetManagementEmbedded(apiExecuteFunc func() (*management.EntityArray, *http.Response, error), apiFunctionName string, resourceType string) (*management.EntityArrayEmbedded, error) {
	l := logger.Get()

	entityArray, response, err := apiExecuteFunc()

	err = HandleClientResponse(response, err, apiFunctionName, resourceType)
	if err != nil {
		return nil, err
	}

	if entityArray == nil {
		l.Error().Msgf("Returned %s() entityArray is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", resourceType, apiFunctionName)
	}

	embedded, embeddedOk := entityArray.GetEmbeddedOk()
	if !embeddedOk {
		l.Error().Msgf("Returned %s() embedded data is nil.", apiFunctionName)
		l.Error().Msgf("%s Response Code: %s\nResponse Body: %s", apiFunctionName, response.Status, response.Body)
		return nil, fmt.Errorf("failed to fetch %s resources via %s()", resourceType, apiFunctionName)
	}

	return embedded, nil
}
