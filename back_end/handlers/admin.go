package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
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

// not tested
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

// not tested
func DeletePostHandler(ctx *gin.Context) {
	postID, err := gocql.ParseUUID(ctx.Param("postID"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "error parsing input",
		})
		return
	}
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	result := DeleteComments(postID, cassandra)
	var status string
	if !result {
		status = "error deleting comments"
	}
	result = DeletePost(postID, cassandra)
	if !result {
		status = "error deleting post"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}

// DeleteUserComments deletes all comments made by a user
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

// not tested
func DeleteUserLikes(userID int, cassandra *gocql.Session) bool {
	userIDSlice := []int{userID}
	if err := cassandra.Query("UPDATE POST SET likes = likes - ? WHERE authorID = ? ALLOW FILTERING", userIDSlice, userID).Exec(); err != nil {
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
	var user1Keys []int
	var user2Keys []int
	var chunkKeys [][]gocql.UUID

	iter := cassandra.Query(`SELECT user1, user2, messageChunks FROM DMTABLE WHERE user2 = ? or user1 = ? 
	ALLOW FILTERING`, userID, userID).Iter()
	var user1, user2 int
	var chunkKey []gocql.UUID
	for iter.Scan(&user1, &user2, &chunkKey) {
		user1Keys = append(user1Keys, user1)
		user2Keys = append(user2Keys, user2)
		chunkKeys = append(chunkKeys, chunkKey)
	}

	if err := iter.Close(); err != nil {
		fmt.Println("Error retrieving user posts:", err)
		return false
	}
	for i := 0; i < len(user1Keys); i++ {
		for x := 0; x < len(chunkKeys[i]); x++ {
			if err := cassandra.Query("Delete DMsubset WHERE subsetID = ?", chunkKey[i][x]).Exec(); err != nil {
				fmt.Println("Error deleting user likes:", err)
				return false
			}
		}
		if err := cassandra.Query("Delete DMTABLE WHERE user1 = ? and user2 = ?", user1Keys[i], user2Keys[i]).Exec(); err != nil {
			fmt.Println("Error deleting user likes:", err)
			return false
		}
	}
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

// not tested
func DeleteUser(user_id int, postgres *sql.DB) bool {
	stmt, err := postgres.Prepare("Select Profile_picture_path from Users WHERE user_id = $1")
	if err != nil {
		return false
	}
	defer stmt.Close()

	row, err := stmt.Query(user_id)
	if err != nil {
		return false
	}

	if row.Next() {
		var path string
		row.Scan(&path)
		if path == "/images/utilties/pp.png" {

		} else if DeleteImage(path) != nil {
			return false
		}
	} else {
		return false
	}
	stmt, err = postgres.Prepare("DELETE FROM Users WHERE user_id = $1")
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
func DeleteUserHandler(ctx *gin.Context) {
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	postgres := ctx.MustGet("postgres").(*sql.DB)
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
	result = DeleteUserFriends(userID, postgres)
	if !result {
		// send message
		return
	}
	result = DeleteUserRequests(userID, postgres)
	if !result {
		// send message
		return
	}
	result = DeleteUser(userID, postgres)
	if !result {
		// send message
		return
	}
	// send messgae
}
