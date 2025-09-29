package main

import (
	"fmt"
	"gordon-raptor/src/internal/config"
	"gordon-raptor/src/internal/di"
	"gordon-raptor/src/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	deps, err := di.NewDIContainer(cfg)
	if err != nil {
		fmt.Println("Error creating DI container:", err)
		return
	}

	server := gin.Default()
	server.SetTrustedProxies([]string{cfg.TrustedProxy})

	router.RegisterRoutes(deps)(server)

	port := fmt.Sprintf(":%v", cfg.Port)
	server.Run(port)
}
