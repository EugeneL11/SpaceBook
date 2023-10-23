package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const PORT_NO = ":8080"

// Managing proxies: https://pkg.go.dev/github.com/gin-gonic/gin#section-readme
// CORS: https://github.com/gin-contrib/cors

func main() {
	// Using Gin for the server:
	server := gin.Default()
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{URLs Go here}
	server.GET("/ping", pong)
	server.PUT("/num")
	server.Use(cors.Default()) // This allows all origins
	// server.Use(cors.New(config))
	server.Run(PORT_NO)
}

func pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
