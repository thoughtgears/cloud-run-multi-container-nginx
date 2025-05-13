package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"

	"github.com/thoughtgears/cloud-run-multi-container-nginx/pkg/router"
)

var Config struct {
	Port string `envconfig:"PORT" default:"8080"`
}

func init() {
	envconfig.MustProcess("", &Config)
}

func main() {
	r := router.NewRouter(Config.Port)

	// Register ping route
	r.Engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	log.Fatal().Err(r.Run()).Msg("Failed to run server")
}
