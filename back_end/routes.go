package main

import (
	"github.com/EugeneL11/SpaceBook/handlers"
	"github.com/gin-gonic/gin"
)

func setupRoutes(server *gin.Engine) {

	setupAccount(server)
	setupFriends(server)
	setupPosts(server)
	setupDMs(server)

}

func setupAccount(server *gin.Engine) {

	server.GET("/login/:username/:password", handlers.LoginHandler)
	server.POST("/register/:email/:password/:fullname/:username", handlers.RegisterHandler)
	server.POST("/upload", handlers.ImageHandler)
	server.GET("/getuserinfo/:viewer/:viewed", handlers.GetUserInfoHandler)
	// server.POST("/uploadprofileimg", handlers.ProfileImageHandler)
}

func setupFriends(server *gin.Engine) {
	server.GET("/getfriends/:user_id", handlers.GetFriendsHandler)
	server.GET("/getfriendreqs/:user_id", handlers.GetFriendRequestsHandler)
	server.GET("/search/:user_id/:searchTerm", handlers.SearchPeopleHandler)
	server.POST("/sendfriendreq/:sender_user_id/:receiver_user_id", handlers.SendFriendRequestHandler)
	server.DELETE("/rejectfriendreq/:rejecter_id/:rejectee_id", handlers.RejectFriendRequestHandler)
	server.DELETE("/removefriend/:id1/:id2", handlers.RemoveFriendHandler)
}

func setupPosts(server *gin.Engine) {
	server.POST("/makepost/:user_id/:caption", handlers.MakePostHandler)
	server.POST("uploadpostimage/:postID")
	server.GET("/gethomepageposts/:user_id", handlers.HomepageHandler)
	server.POST("/makecomment/:postID/:userID/:comment", handlers.CommentHandler)
}

func setupDMs(server *gin.Engine) {

}
