package handlers

import (
	"gordon-raptor/src/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cook(context *gin.Context) {
	var body dtos.CookDto
	if err := context.BindJSON(&body); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	context.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"body":    body,
	})
}
