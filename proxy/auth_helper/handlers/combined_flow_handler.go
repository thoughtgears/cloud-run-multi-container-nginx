package handlers

import (
	"net/http"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/rs/zerolog/log"

	"github.com/thoughtgears/cloud-run-multi-container-nginx/proxy/auth_helper/services"
)

var combinedFlowHttpClient = &http.Client{Timeout: 5 * time.Second}

func CombinedFlowHandler(app *firebase.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		l := log.With().Str("handler", "CombinedFlowHandler").Logger()

		// --- Stage 1: Validate Client Credentials (e.g., Firebase) ---
		authHeader := r.Header.Get("Authorization") // Passed by Nginx
		if authHeader == "" || !strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			l.Warn().Msg("Firebase Bearer token missing or malformed")
			http.Error(w, "Unauthorized: Authentication token missing or malformed", http.StatusUnauthorized)
			return
		}
		firebaseTokenString := authHeader[7:]

		verifiedFirebaseToken, err := services.ValidateFirebaseToken(ctx, app, firebaseTokenString)
		if err != nil {
			tokenPrefix := firebaseTokenString
			if len(tokenPrefix) > 15 {
				tokenPrefix = tokenPrefix[:15] + "..."
			}
			l.Warn().Err(err).Str("token_prefix", tokenPrefix).Msg("Error validating Firebase ID token")
			http.Error(w, "Unauthorized: Invalid Firebase token", http.StatusUnauthorized)
			return
		}
		l.Info().Str("uid", verifiedFirebaseToken.UID).Msg("Firebase token validated successfully")

		gcpAudience := r.URL.Query().Get("gcp_audience")

		gcpToken, err := services.FetchGCPIdentityToken(ctx, gcpAudience, combinedFlowHttpClient)
		if err != nil {
			l.Warn().Err(err).Str("gcp_audience", gcpAudience).Msg("Error fetching GCP token")
			if strings.Contains(err.Error(), "audience cannot be empty") {
				http.Error(w, "Bad Request: Missing gcp_audience", http.StatusBadRequest)
			} else if strings.Contains(err.Error(), "GCP metadata server error") {
				http.Error(w, "Bad Gateway: Error from GCP metadata server", http.StatusBadGateway)
			} else {
				http.Error(w, "Internal Server Error: Could not fetch GCP token", http.StatusInternalServerError)
			}
			return
		}
		l.Info().Str("gcp_audience", gcpAudience).Msg("GCP token fetched successfully")

		w.Header().Set("X-Gcp-Token", gcpToken)
		w.Header().Set("X-Firebase-User-Id", verifiedFirebaseToken.UID)
		w.Header().Set("X-Auth-Method", "firebase-auth")
		w.WriteHeader(http.StatusOK)
	})
}
