package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

// not done, tested, or documented
func DeleteComments(postID gocql.UUID, cassandra *gocql.Session) bool {
	return true
}

// not done, not tested
func DeletePost(postID gocql.UUID, cassandra *gocql.Session) bool {
	// delete pictures
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

// not done
// not tested
func DeleteUserComments(userID int, cassandra *gocql.Session) bool {
	return true
}

// not done
// not tested
func DeleteUserLikes(userID int, cassandra *gocql.Session) bool {
	return true
}

// not done
// not tested
func DeleteUserPosts(userID int, cassandra *gocql.Session) bool {
	return true
}
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
	result := DeleteUserDM(userID, cassandra)
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
	result = DeleteUserPosts(userID, cassandra)
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
