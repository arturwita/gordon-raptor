package routes

import (
	"gordon-raptor/src/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/ping", handlers.Ping)
		api.POST("/cook", handlers.Cook)
	}
}
