package main

import (
	"github.com/EugeneL11/SpaceBook/handlers"
	"github.com/gin-gonic/gin"
)

func setupRoutes(server *gin.Engine) {
	server.GET("/ping", handlers.Pong)
	// /num/var1/var2
	server.GET("/num/:num1/:num2", handlers.Sum)
	server.GET("/testInsert/:val", handlers.TestInsertHandler)
	//server.POST("/user", handlers.Double)
	// server.GET("/postgresTest", TestPostgres)
}
