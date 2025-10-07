package app

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gordon-raptor/src/internal/config"
	"gordon-raptor/src/internal/di"
	"gordon-raptor/src/internal/middlewares"
	"gordon-raptor/src/internal/routes"
)

func NewApp(cfg *config.Config) (*gin.Engine, error) {
	deps, err := di.NewDIContainer(cfg)
	if err != nil {
		fmt.Println("Error creating DI container:", err)
		return nil, err
	}

	server := gin.Default()
	server.Use(middlewares.CORSMiddleware(cfg))
	server.SetTrustedProxies([]string{cfg.TrustedProxy})

	routes.RegisterRoutes(server, deps)

	return server, nil
}
