package main

import (
	"fmt"
	"gordon-raptor/src/config"
	"gordon-raptor/src/pkg/di"
	"gordon-raptor/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: Error loading .env file:", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}

	deps, err := di.DIContainerFactory(cfg)
	if err != nil {
		return
	}

	router := gin.Default()
	router.SetTrustedProxies([]string{cfg.TrustedProxy})

	routes.RegisterRoutesFactory(deps.RecipeRepository)(router)

	port := fmt.Sprintf(":%v", cfg.Port)
	router.Run(port)
}
