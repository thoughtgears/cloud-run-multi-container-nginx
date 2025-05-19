// auth_helper/main.go
package main

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/thoughtgears/cloud-run-multi-container-nginx/proxy/auth_helper/config"
	"github.com/thoughtgears/cloud-run-multi-container-nginx/proxy/auth_helper/handlers"
)

var (
	cfg         config.Config
	firebaseApp *firebase.App
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.LevelFieldName = "severity"

	envconfig.MustProcess("", &cfg)

	var err error
	// Initialize Firebase Admin SDK.
	// This helper needs the roles/iam.serviceAccountTokenCreator and roles/firebaseauth.viewer
	// roles for the service account running this code.
	firebaseApp, err = firebase.NewApp(context.Background(), &firebase.Config{
		ProjectID: cfg.ProjectID,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Error initializing firebase app")
	}
}

func main() {
	port := "8081"

	http.HandleFunc("/get-token", handlers.GetTokenHandler)
	http.Handle("/auth", handlers.AuthHandler(firebaseApp))
	http.Handle("/auth/combined-flow", handlers.CombinedFlowHandler(firebaseApp))
	
	log.Info().Str("port", port).Msg("Listening on port")
	log.Fatal().Err(http.ListenAndServe(":"+port, nil)).Msg("Error starting server")
}
