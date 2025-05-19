package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/thoughtgears/cloud-run-multi-container-nginx/proxy/auth_helper/services"
)

var defaultHttpClient = &http.Client{Timeout: time.Second * 5} // Example

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := log.With().Str("handler", "GetTokenHandler").Logger()

	audience := r.URL.Query().Get("audience")
	// The service function will validate if audience is empty

	gcpToken, err := services.FetchGCPIdentityToken(ctx, audience, defaultHttpClient) // Pass httpClient
	if err != nil {
		l.Warn().Err(err).Str("audience", audience).Msg("Error fetching GCP token")
		// Determine status code based on error type if possible
		// For now, assuming 500 for simplicity, but service could return a typed error
		if strings.Contains(err.Error(), "audience cannot be empty") {
			http.Error(w, "Bad Request: Missing audience", http.StatusBadRequest)
		} else if strings.Contains(err.Error(), "GCP metadata server error") {
			http.Error(w, "Bad Gateway: Error from GCP metadata server", http.StatusBadGateway)
		} else {
			http.Error(w, "Internal Server Error: Could not fetch GCP token", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("X-Auth-Token", gcpToken)
	w.Header().Set("Content-Type", "application/json")
	l.Info().Str("audience", audience).Msg("GCP Token captured and set in X-Auth-Token")
}
