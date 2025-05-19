package services

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

const metadataURLBase = "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/identity"

// FetchGCPIdentityToken retrieves a GCP identity token for the specified audience from the GCP metadata server.
// This is used to authenticate backend requests for our Cloud Run Services.
// The audience parameter must not be empty.
// If the audience is empty, an error is returned.
//
// Parameters:
// - ctx: The context for the operation.
// - audience: The audience for which the token is requested.
// - httpClient: An optional HTTP client to use for the request. If nil, a default client is created.
//
// Returns:
// - string: The GCP identity token if successful.
// - error: An error if the request fails or if the audience is empty.
func FetchGCPIdentityToken(ctx context.Context, audience string, httpClient *http.Client) (string, error) {
	if audience == "" {
		return "", fmt.Errorf("audience cannot be empty")
	}

	if httpClient == nil { // Default client if none provided
		httpClient = &http.Client{Timeout: 5 * time.Second}
	}

	fullMetadataURL := fmt.Sprintf("%s?audience=%s", metadataURLBase, audience)
	req, err := http.NewRequestWithContext(ctx, "GET", fullMetadataURL, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request for GCP metadata server: %w", err)
	}
	req.Header.Add("Metadata-Flavor", "Google")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error requesting GCP token from metadata server (audience %s): %w", audience, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("error from GCP metadata server (audience: %s): status %d, body: %s", audience, resp.StatusCode, string(bodyBytes))
	}

	tokenBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading GCP token response body (audience: %s): %w", audience, err)
	}
	return string(tokenBytes), nil
}
