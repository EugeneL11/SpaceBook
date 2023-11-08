package main

import "github.com/gin-gonic/gin"
import "github.com/EugeneL11/SpaceBook/handlers"

func setupRoutes(server *gin.Engine) {
	server.GET("/ping", handlers.Pong)
	server.GET("/num/:num1/:num2", handlers.Sum)
	server.GET("/testInsert/:val", handlers.TestInsertHandler)
	server.POST("/user", handlers.Double)
	// server.GET("/postgresTest", TestPostgres)
}
