package handlers

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

// not done
// not tested
func GetHomePagePost(userID int, date time.Time, postgres *sql.DB, cassandra *gocql.Session) string {

	return "no error"
}

// not done
// not tested
// not documented
func HomepageHandler(ctx *gin.Context) {

}

// not done
// not tested
func GetPostDetails(postID int, viewingUser int) string {
	return "no error"
}

// not done
// not tested
// not documented
func PostDetailsHandler(ctx *gin.Context) {

}

// not done
// not tested
func LikePost(postID int, userID int) string {
	return "no error"
}

// not done
// not tested
// not documented
func LikePostHandler(ctx *gin.Context) {

}

// not done
// not tested
func UnlikePost(userID int, postID int) string {
	return "no error"
}

// not done
// not tested
// not documented
func UnlikePostHandler(ctx *gin.Context) {

}
func CommentPost(comment string, userID int, postID int) string {
	return "no error"
}

// not done
// not tested
// not documented
func CommentHandler(ctx *gin.Context) {

}

// not done
// not tested
func DeletePost(postID int, cassandra *gocql.Session) string {
	//cassandra.Query()
	return "no error"
}

// not done
// not tested
// not documented
func DeletePostHandler(ctx *gin.Context) {

}
