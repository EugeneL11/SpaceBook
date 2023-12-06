package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EugeneL11/SpaceBook/pkg"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

// DeleteComments deletes comments associated with a postID
func DeleteComments(postID gocql.UUID, cassandra *gocql.Session) bool {
	comments := []gocql.UUID{}
	if err := cassandra.Query("Select comments from post where postid = ?", postID).Iter().Scan(&comments); !err {

	}
	for i := range comments {
		if err2 := cassandra.Query("DELETE FROM COMMENT WHERE commentID = ?", comments[i]).Exec(); err2 != nil {
			fmt.Println("Error deleting comments:", err2)
			return false
		}
	}

	return true
}

// Deletes a Post given an ID
func DeletePost(postID gocql.UUID, cassandra *gocql.Session) bool {
	imagePaths := []string{}
	imagePath := ""
	if err := cassandra.Query("SELECT imagePaths FROM POST WHERE postID = ?", postID).Iter().Scan(&imagePaths); !err {
		fmt.Println("Error retrieving imagePaths:")
		return false
	}

	// Call DeleteImage for each imagePath
	for _, imagePath = range imagePaths {
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
		if err := cassandra.Query("DELETE FROM COMMENT WHERE commentID = ?", currComment).Exec(); err != nil {
			fmt.Println("Error deleting user comments:", err)
			return false
		}
	}
	if err := rows.Close(); err != nil {
		return false
	}
	for post, comments := range postToComment {
		var allcomments []gocql.UUID
		if err := cassandra.Query("SELECT comments FROM POST WHERE postID = ?", post).Iter().Scan(&allcomments); !err {
			fmt.Println("Error getting comment set:", err)
			return false
		}
		newcomments := pkg.RemoveSubset(allcomments, comments)
		if err2 := cassandra.Query("Update post set comments = ? where postID = ?", newcomments, post).Exec(); err2 != nil {
			fmt.Println("Error updating user comments:", err2)
			return false
		}
	}

	return true
}

func DeleteUserLikes(userID int, session *gocql.Session) bool {
	// Retrieve all post IDs where the likes set contains the specified userID
	var postIDs []gocql.UUID
	if iter := session.Query("SELECT postId FROM POST WHERE likes CONTAINS ? ALLOW FILTERING", userID).Iter(); iter != nil {
		for {
			var postID gocql.UUID
			if iter.Scan(&postID) {
				postIDs = append(postIDs, postID)
			} else {
				break
			}
		}
		if err := iter.Close(); err != nil {
			fmt.Println("Error retrieving user posts:", err)
			return false
		}
	}

	// Iterate over post IDs and update likes for each post
	for _, postID := range postIDs {
		// Retrieve the current set from Cassandra
		var currentLikes []int
		if err := session.Query("SELECT likes FROM POST WHERE postId = ?", postID).Scan(&currentLikes); err != nil {
			fmt.Println("Error retrieving current likes:", err)
			return false
		}

		// Modify the set (e.g., remove the specified userID)
		updatedLikes := pkg.RemoveFromSlice(currentLikes, userID)

		// Update the set in Cassandra
		if err := session.Query("UPDATE POST SET likes = ? WHERE postId = ?", updatedLikes, postID).Exec(); err != nil {
			fmt.Println("Error updating likes:", err)
			return false
		}
	}
	return true
}

func DeleteUserPosts(userID int, cassandra *gocql.Session) bool {
	// Retrieve postIDs made by the user
	var postIDs []gocql.UUID
	iter := cassandra.Query("SELECT postID FROM POST WHERE authorID = ? ALLOW FILTERING", userID).Iter()
	for {
		var postID gocql.UUID
		if iter.Scan(&postID) {
			postIDs = append(postIDs, postID)
		} else {
			break
		}
	}
	if err := iter.Close(); err != nil {
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

func DeleteUserDM(userID int, cassandra *gocql.Session) bool {
	user1Keys := []int{}
	user2Keys := []int{}
	chunkKeys := [][]gocql.UUID{}

	iter := cassandra.Query(`SELECT user1, user2, messageChunks FROM DMTABLE WHERE user1 = ? ALLOW FILTERING`, userID).Iter()
	var user1, user2 int
	chunkKey := []gocql.UUID{}
	for iter.Scan(&user1, &user2, &chunkKey) {
		user1Keys = append(user1Keys, user1)
		user2Keys = append(user2Keys, user2)
		chunkKeys = append(chunkKeys, chunkKey)

	}

	if err := iter.Close(); err != nil {
		fmt.Println("Error retrieving user dms:", err)
		return false
	}
	iter = cassandra.Query(`SELECT user1, user2, messagechunks FROM DMTABLE WHERE user2 = ? 
	ALLOW FILTERING`, userID).Iter()
	for iter.Scan(&user1, &user2, &chunkKey) {
		user1Keys = append(user1Keys, user1)
		user2Keys = append(user2Keys, user2)
		fmt.Println(user1, user2, chunkKey)
		chunkKeys = append(chunkKeys, chunkKey)
	}
	fmt.Println(chunkKeys)
	if err := iter.Close(); err != nil {
		fmt.Println("Error retrieving user posts:", err)
		return false
	}

	for i := 0; i < len(user1Keys); i++ {
		for x := 0; x < len(chunkKeys[i]); x++ {
			if err := cassandra.Query("Delete From DMsubset WHERE subsetID = ?", chunkKeys[i][x]).Exec(); err != nil {
				fmt.Println("Error deleting subsets:", err)
				return false
			}
			fmt.Println("hi")
		}
		if err := cassandra.Query("Delete From DMTABLE WHERE user1 = ? and user2 = ?", user1Keys[i], user2Keys[i]).Exec(); err != nil {
			fmt.Println("Error deleting user dms:", err)
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
	return true
}

// not tested
func DeleteUserFriends(userID int, postgres *sql.DB) bool {
	stmt, err := postgres.Prepare("DELETE FROM Orbit_buddies WHERE user1_id = $1 OR user2_id = $1")
	if err != nil {
		fmt.Println("Prep failed")
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID)
	if err != nil {
		fmt.Println("Query failed")
		return false
	}
	return true
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
		if path == "/images/utilities/pp.png" {

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

// Define some helper struct to make DeleteUserHandler more condensed
type CTXStatus struct {
	Status string `json:"status"`
}

// not tested
func DeleteUserHandler(ctx *gin.Context) {
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	postgres := ctx.MustGet("postgres").(*sql.DB)
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusOK, CTXStatus{Status: "error parsing input"})
		return
	}

	result := DeleteUserPosts(userID, cassandra)
	if !result {
		ctx.JSON(http.StatusOK, CTXStatus{Status: "failed to delete posts"})
		return
	}
	result = DeleteUserComments(userID, cassandra)
	if !result {
		ctx.JSON(http.StatusOK, CTXStatus{Status: "failed to delete comments"})
		return
	}
	result = DeleteUserLikes(userID, cassandra)
	if !result {
		ctx.JSON(http.StatusOK, CTXStatus{Status: "failed to delete likes"})
		return
	}
	result = DeleteUserDM(userID, cassandra)
	if !result {
		ctx.JSON(http.StatusOK, CTXStatus{Status: "failed to delete DMs"})
		return
	}
	result = DeleteUserFriends(userID, postgres)
	if !result {
		ctx.JSON(http.StatusOK, CTXStatus{Status: "failed to delete friends"})
		return
	}
	result = DeleteUserRequests(userID, postgres)
	if !result {
		ctx.JSON(http.StatusOK, CTXStatus{Status: "failed to delete friend requests"})
		return
	}
	result = DeleteUser(userID, postgres)
	if !result {
		ctx.JSON(http.StatusOK, CTXStatus{Status: "failed to delete user"})
		return
	}
	// send message
	ctx.JSON(http.StatusOK, CTXStatus{Status: "no error"})
}
