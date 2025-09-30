package main

import (
	"fmt"
	"gordon-raptor/src/internal/app"
	"gordon-raptor/src/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	server, err := app.NewApp(cfg)
	if err != nil {
		fmt.Println("Error creating app:", err)
		return
	}

	port := fmt.Sprintf(":%v", cfg.Port)
	server.Run(port)
}
