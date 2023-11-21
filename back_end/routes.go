package main

import (
	"github.com/EugeneL11/SpaceBook/handlers"
	"github.com/gin-gonic/gin"
)

func setupRoutes(server *gin.Engine) {
	setupTest(server)
	setupAccount(server)
	setupQueries(server)
}

func setupTest(server *gin.Engine) {
	server.GET("/ping", handlers.Pong)
	// /num/var1/var2
	server.GET("/num/:num1/:num2", handlers.Sum)
	server.GET("/testInsert/:val", handlers.TestInsertHandler)
	//server.POST("/user", handlers.Double)
	// server.GET("/postgresTest", TestPostgres)
}

func setupAccount(server *gin.Engine) {

	server.GET("/login/:username/:password", handlers.LoginHandler)
	server.GET("/register/:email/:password/:fullname/:username", handlers.RegisterHandler)

}

func setupQueries(server *gin.Engine) {
	server.GET("/getfriends/:user_id", handlers.GetFriendsHandler)
}
