package main

import (
	"github.com/EugeneL11/SpaceBook/handlers"
	"github.com/gin-gonic/gin"
)

func setupRoutes(server *gin.Engine) {

	setupAccount(server)
	setupFriends(server)
}

func setupAccount(server *gin.Engine) {

	server.GET("/login/:username/:password", handlers.LoginHandler)
	server.POST("/register/:email/:password/:fullname/:username", handlers.RegisterHandler)
	server.POST("/upload", handlers.ImageHandler)
	// server.POST("/uploadprofileimg", handlers.ProfileImageHandler)
}
func setupFriends(server *gin.Engine) {
	server.GET("/getfriends/:user_id", handlers.GetFriendsHandler)
}
