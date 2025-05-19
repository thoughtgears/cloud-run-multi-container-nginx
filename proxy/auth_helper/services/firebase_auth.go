package services

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

// ValidateFirebaseToken validates a Firebase ID token and returns the decoded token if valid.
// It uses the Firebase Admin SDK to verify the token and extract user information.
//
// Parameters:
// - ctx: The context for the operation.
// - app: The Firebase app instance.
// - idTokenString: The Firebase ID token string to validate.
//
// Returns:
// - *auth.Token: The decoded token if valid.
// - error: An error if the token is invalid or if there was an issue verifying it.
func ValidateFirebaseToken(ctx context.Context, app *firebase.App, idTokenString string) (*auth.Token, error) {
	if app == nil {
		return nil, fmt.Errorf("firebase app not initialized")
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Firebase auth client: %w", err)
	}

	verifiedToken, err := client.VerifyIDToken(ctx, idTokenString)
	if err != nil {
		return nil, fmt.Errorf("error verifying Firebase ID token: %w", err)
	}
	return verifiedToken, nil
}
