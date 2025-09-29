package main

import (
	"fmt"
	"gordon-raptor/src/pkg/config"
	"gordon-raptor/src/pkg/di"
	"gordon-raptor/src/routes"

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

	router := gin.Default()
	router.SetTrustedProxies([]string{cfg.TrustedProxy})

	routes.RegisterRoutes(deps)(router)

	port := fmt.Sprintf(":%v", cfg.Port)
	router.Run(port)
}
