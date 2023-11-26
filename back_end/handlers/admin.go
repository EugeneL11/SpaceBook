package handlers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

// DeleteComments deletes comments associated with a postID
func DeleteComments(postID gocql.UUID, cassandra *gocql.Session) bool {
	if err := cassandra.Query("DELETE FROM COMMENT WHERE postID = ? ALLOW FILTERING", postID).Exec(); err != nil {
		fmt.Println("Error deleting comments:", err)
		return false
	}
	return true
}

// DeletePost deletes a post and associated data
func DeletePost(postID gocql.UUID, cassandra *gocql.Session) bool {
	var imagePaths []string
	if err := cassandra.Query("SELECT imagePaths FROM POST WHERE postID = ?", postID).Iter().Scan(&imagePaths); !err {
		fmt.Println("Error retrieving imagePaths:", err)
		return false
	}

	// Call DeleteImage for each imagePath
	for _, imagePath := range imagePaths {
		if DeleteImage(imagePath) != nil {
			fmt.Println("Couldn't delete image", imagePath)
			return false
		}
	}
	// Delete the post
	if err2 := cassandra.Query("DELETE FROM POST WHERE postID = ?", postID).Exec(); err2 != nil {
		fmt.Println("Error deleting post:", err2)
		return false
	}

	return true
}

// not done
// not tested
// not documented
func DeletePostHandler(ctx *gin.Context) {
	postID, err := gocql.ParseUUID(ctx.Param("postID"))
	if err != nil {
		// send messgage
		return
	}
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	result := DeleteComments(postID, cassandra)
	if !result {
		// send message
		return
	}
	result = DeletePost(postID, cassandra)
	if !result {
		// send message
		return
	}

}

// DeleteUserComments deletes comments made by a user
func DeleteUserComments(userID int, cassandra *gocql.Session) bool {
	postToComment := make(map[gocql.UUID][]gocql.UUID)
	rows := cassandra.Query("SELECT commentID, postID FROM COMMENT WHERE commenter = ? ALLOW FILTERING", userID).Iter()
	for {
		var currPost, currComment gocql.UUID
		if !rows.Scan(&currComment, &currPost) {
			break
		}
		postToComment[currPost] = append(postToComment[currPost], currComment)
	}
	for key, value := range postToComment {
		if err := cassandra.Query("Update post set comments = comments - ? where postID = ?", value, key).Exec(); err != nil {
			fmt.Println("Error deleting user comments:", err)
			return false
		}
	}
	if err := rows.Close(); err != nil {
		return false
	}
	if err := cassandra.Query("DELETE FROM COMMENT WHERE commenter = ?", userID).Exec(); err != nil {
		fmt.Println("Error deleting user comments:", err)
		return false
	}
	return true
}

// not done not tested
func DeleteUserLikes(userID int, cassandra *gocql.Session) bool {
	if err := cassandra.Query("UPDATE POST SET likes = likes - ? WHERE authorID = ? ALLOW FILTERING", userID, userID).Exec(); err != nil {
		fmt.Println("Error deleting user likes:", err)
		return false
	}
	return true
}

// not tested
func DeleteUserPosts(userID int, cassandra *gocql.Session) bool {
	// Retrieve postIDs made by the user
	var postIDs []gocql.UUID
	if err := cassandra.Query("SELECT postID FROM POST WHERE authorID = ? ALLOW FILTERING", userID).Iter().Scan(&postIDs); !err {
		fmt.Println("Error retrieving user posts:", err)
		return false
	}

	// Delete associated comments, images, and posts
	for _, postID := range postIDs {
		if !DeletePost(postID, cassandra) {
			return false
		}
	}

	return true
}

// not tested
func DeleteUserDM(userID int, cassandra *gocql.Session) bool {

	return true
}

// not tested
func DeleteUserRequests(userID int, postgres *sql.DB) bool {
	stmt, err := postgres.Prepare("DELETE FROM Orbit_requests WHERE requester_id = $1 OR requested_buddy_id = $1")
	if err != nil {
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID)
	if err != nil {
		return false
	}
	return false
}

// not tested
func DeleteUserFriends(userID int, postgres *sql.DB) bool {
	stmt, err := postgres.Prepare("DELETE FROM Orbit_buddies WHERE user_id1 = $1 OR user_id2 = $1")
	if err != nil {
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID)
	if err != nil {
		return false
	}
	return false
}

// not done not tested
func DeleteUser(user_id int, postgres *sql.DB) bool {
	// delete profile picture
	stmt, err := postgres.Prepare("DELETE FROM Users WHERE user_id = $1")
	if err != nil {
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id)
	if err != nil {
		return false
	}

	return true
}

// not done
// not tested
// not documented
func DeleteUserHandler(ctx *gin.Context) {
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	postrges := ctx.MustGet("postrges").(*sql.DB)
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		// send message
		return
	}
	result := DeleteUserPosts(userID, cassandra)
	if !result {
		// send message
		return
	}
	result = DeleteUserComments(userID, cassandra)
	if !result {
		// send message
		return
	}
	result = DeleteUserLikes(userID, cassandra)
	if !result {
		// send message
		return
	}
	result = DeleteUserDM(userID, cassandra)
	if !result {
		// send message
		return
	}
	result = DeleteUserFriends(userID, postrges)
	if !result {
		// send message
		return
	}
	result = DeleteUserRequests(userID, postrges)
	if !result {
		// send message
		return
	}
	result = DeleteUser(userID, postrges)
	if !result {
		// send message
		return
	}
	// send messgae
}
