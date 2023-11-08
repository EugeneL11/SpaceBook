package main

import "github.com/gin-gonic/gin"
import "github.com/EugeneL11/SpaceBook/back_end/handlers"

func setupRoutes(server *gin.Engine) {
	server.GET("/ping", handlers.pong)
	server.PUT("/num")
	server.GET("/num/:num1/:num2", handlers.sum)
	server.GET("/testInsert/:val", testInsertHandler)
	server.POST("/user", handlers.double)
}
