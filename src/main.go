package main

import (
	"fmt"
	cfg "gordon-raptor/src/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := cfg.LoadConfig()
	if err != nil {
		return
	}

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port := fmt.Sprintf(":%v", config.Port)

	router.Run(port)
}
