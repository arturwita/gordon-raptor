package handlers

import (
	"fmt"
	"gordon-raptor/src/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cook(context *gin.Context) {
	var body dtos.CookDto
	_ := context.BindJSON(&body)

	context.JSON(http.StatusCreated, gin.H{
		"message": "pong",
		"body":    body,
	})
}
