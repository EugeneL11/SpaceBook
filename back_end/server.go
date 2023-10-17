package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const PORT_NO = ":8080"

func main() {
	// Using Gin for the server:
	server := gin.Default()
	server.GET("/ping", pong)
	server.PUT("/num")
	server.Run(PORT_NO)
}

func pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
