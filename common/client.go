package common

import (
	"net/http"
	"net/url"

	"github.com/alwitt/httpmq-go/api"
)

// RequestIDHeader is expected HTTP header field carrying the httpmq request ID
const RequestIDHeader = "Httpmq-Request-ID"

/*
DefineAPIClient create an API client object targeting a specific httpmq server URL

 @param serverUrl string - httpmq server URL
 @param transport *http.Client - HTTP client used by the API
 @param userAgent *string - user agent to use if not nil
 @param debug bool - whether to enable API debug logging
 @return client object
*/
func DefineAPIClient(
	serverURL string, transport *http.Client, userAgent *string, debug bool,
) (*api.APIClient, error) {
	parsed, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	baseConfig := api.NewConfiguration()
	baseConfig.Host = parsed.Host
	baseConfig.Scheme = parsed.Scheme
	baseConfig.Debug = debug
	if userAgent != nil {
		baseConfig.UserAgent = *userAgent
	}
	baseConfig.Servers[0].URL = serverURL
	baseConfig.HTTPClient = transport
	// Create the client object
	client := api.NewAPIClient(baseConfig)
	return client, nil
}
