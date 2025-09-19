package main

import (
	"fmt"
	"gordon-raptor/src/config"
	"gordon-raptor/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}

	router := gin.Default()
	router.SetTrustedProxies([]string{cfg.TrustedProxy})

	routes.RegisterRoutes(router)

	port := fmt.Sprintf(":%v", cfg.Port)
	router.Run(port)
}
