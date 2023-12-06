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
	setupCookies(server)
}

func setupAccount(server *gin.Engine) {

	server.POST("/login/:username/:password", handlers.LoginHandler)
	server.POST("/register/:email/:password/:fullname/:username", handlers.RegisterHandler)
	server.POST("/getuserinfo/:viewer/:viewed", handlers.GetUserInfoHandler)
	server.POST("/uploadprofileimage/:userID", handlers.ProfilePicHandler)
	server.DELETE("/deleteuser/:userID", handlers.DeleteUserHandler)
	server.PUT("/updateuserprofile/:userID/:newFullName/:newPlanet/:newBio", handlers.UpdateUserProfileHandler)
}

func setupFriends(server *gin.Engine) {
	server.POST("/friends/:user_id", handlers.GetFriendsHandler)
	server.POST("/friendrequests/:user_id", handlers.GetFriendRequestsHandler)
	server.POST("/search/:user_id/:searchTerm", handlers.SearchPeopleHandler)
	server.POST("/sendfriendreq/:sender_user_id/:receiver_user_id", handlers.SendFriendRequestHandler)
	server.DELETE("/rejectfriendreq/:rejecter_id/:rejectee_id", handlers.RejectFriendRequestHandler)
	server.DELETE("/removefriend/:id1/:id2", handlers.RemoveFriendHandler)
}

func setupPosts(server *gin.Engine) {
	server.POST("/homepageposts/:user_id", handlers.HomepageHandler)
	server.POST("/postdetails/:postID/:userID", handlers.PostDetailsHandler)
	server.POST("/makepost/:user_id/:caption", handlers.MakePostHandler)
	server.POST("/uploadpostimage/:postID", handlers.UploadImagePost)
	server.POST("/gethomepageposts/:user_id", handlers.HomepageHandler)
	server.POST("/makecomment/:postID/:userID/:comment", handlers.CommentHandler)
	server.PUT("/likepost/:postID/:userID", handlers.LikePostHandler)
	server.DELETE("deletepost/:postID", handlers.DeletePostHandler)
}

func setupDMs(server *gin.Engine) {
	// only returns most recent 20-40 messages at size 1, returns 40-60 messages at size 2 etc
	// will return whether or all the messages have been returned
	server.POST("/getmessages/:user1/:user2/:subset_size", handlers.GetMessagesHandler)
	server.POST("/newdm/:user1/:user2", handlers.CreateNewDMHandler)
	server.POST("/senddm/:senderID/:receiverID/:message", handlers.SendDMHandler)
	server.POST("/userdms/:userID", handlers.GetDMHandler)
	// gets all friends who you do not have a dm with
	server.POST("/getallnewdm/:userID", handlers.NewDMListHandler)
}
func setupCookies(server *gin.Engine) {
	server.POST("/createcookie", handlers.CreateCookieHandler)
	server.POST("/setcookie/:userID", handlers.SetCookieHandler)
	server.POST("/getcookie", handlers.GetCookieHandler)
	server.DELETE("/removecookie", handlers.RemoveCookieHandler)
}
