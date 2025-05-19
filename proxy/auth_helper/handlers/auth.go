package handlers

import (
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/rs/zerolog/log"

	"github.com/thoughtgears/cloud-run-multi-container-nginx/proxy/auth_helper/services"
)

func AuthHandler(app *firebase.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		l := log.With().Str("handler", "AuthHandler").Logger()

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			l.Warn().Msg("Firebase Bearer token missing or malformed")
			http.Error(w, "Unauthorized: Authentication token missing or malformed", http.StatusUnauthorized)
			return
		}
		firebaseTokenString := authHeader[7:]

		verifiedToken, err := services.ValidateFirebaseToken(ctx, app, firebaseTokenString)
		if err != nil {
			tokenPrefix := firebaseTokenString
			if len(tokenPrefix) > 15 {
				tokenPrefix = tokenPrefix[:15] + "..."
			}
			l.Warn().Err(err).Str("token_prefix", tokenPrefix).Msg("Error validating Firebase ID token")
			http.Error(w, "Unauthorized: Invalid Firebase token", http.StatusUnauthorized) // Consider more specific errors based on err type
			return
		}

		l.Info().Str("uid", verifiedToken.UID).Msg("Firebase token validated successfully")
		w.Header().Set("X-Authenticated-User-Id", verifiedToken.UID)
		w.Header().Set("X-Auth-Method", "firebase")
		w.WriteHeader(http.StatusOK)
	})
}
