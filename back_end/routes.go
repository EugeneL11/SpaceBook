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
	server.GET("/getuserinfo/:viewer/:viewed", handlers.GetUserInfoHandler)
	server.POST("/uploadprofileimage/:userID", handlers.ProfilePicHandler)
	server.DELETE("/deleteuser/:userID", handlers.DeleteUserHandler)
}

func setupFriends(server *gin.Engine) {
	server.GET("/friends/:user_id", handlers.GetFriendsHandler)
	server.GET("/friendrequests/:user_id", handlers.GetFriendRequestsHandler)
	server.GET("/search/:user_id/:searchTerm", handlers.SearchPeopleHandler)
	server.POST("/sendfriendreq/:sender_user_id/:receiver_user_id", handlers.SendFriendRequestHandler)
	server.DELETE("/rejectfriendreq/:rejecter_id/:rejectee_id", handlers.RejectFriendRequestHandler)
	server.DELETE("/removefriend/:id1/:id2", handlers.RemoveFriendHandler)
}

func setupPosts(server *gin.Engine) {
	server.GET("/homepageposts/:user_id", handlers.HomepageHandler)
	server.GET("/postdetails/:postID/:userID", handlers.PostDetailsHandler)
	server.POST("/makepost/:user_id/:caption", handlers.MakePostHandler)
	server.POST("/uploadpostimage/:postID", handlers.UploadImagePost)
	server.GET("/gethomepageposts/:user_id", handlers.HomepageHandler)
	server.POST("/makecomment/:postID/:userID/:comment", handlers.CommentHandler)
	server.PUT("/likepost/:postID/:userID", handlers.LikePostHandler)
	server.DELETE("deletepost/:postID", handlers.DeletePostHandler)
}

func setupDMs(server *gin.Engine) {
	// only returns most recent 20-40 messages at size 1, returns 40-60 messages at size 2 etc
	// will return whether or all the messages have been returned
	server.GET("/getmessages/:user1/:user2/subset_size", handlers.GetMessagesHandler)
	server.PUT("/sendmessage/:sender/:reciever/:message", handlers.SendDMHandler)
	server.GET("/getallchats/:userID", handlers.GetDMHandler)
	server.POST("/createdm/:user1/:user2", handlers.CreateNewDMHandler)

	// gets all friends who you do not have a dm with
	// idk wtf to name this
	server.GET("/getallnewDM/:userID", handlers.NewDMListHandler)
}
