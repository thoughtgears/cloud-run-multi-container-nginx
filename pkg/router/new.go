package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/thoughtgears/cloud-run-multi-container-nginx/pkg/router/middleware"
)

type Router struct {
	Engine *gin.Engine
	port   string
}

func NewRouter(port string) *Router {
	var newRouter Router

	if port == "" {
		newRouter.port = "8080"
	} else {
		newRouter.port = port
	}

	newRouter.Engine = gin.New()
	newRouter.Engine.Use(middleware.Logger())
	newRouter.Engine.Use(gin.Recovery())
	_ = newRouter.Engine.SetTrustedProxies(nil)

	newRouter.Engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Service is running",
		})
	})

	return &newRouter
}
