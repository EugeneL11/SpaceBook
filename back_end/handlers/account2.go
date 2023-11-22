package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

// not done
// not tested
func DeleteUserComments(userID int, cassandra gocql.Session) string {
	return "no error"
}

// not done
// not tested
func DeleteUserLikes(userID int, cassandra gocql.Session) string {
	return "no error"
}

// not done
// not tested
func DeleteUserPosts(userID int, cassandra gocql.Session) string {
	return "no error"
}

// not done
// not tested
func DeleteUserRequests(userID int, postgres *sql.DB) string {
	stmt, err := postgres.Prepare("DELETE FROM Orbit_requests WHERE requester_id = $1 OR requested_buddy_id = $1")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()
	_, err = stmt.Exec(user_id)
	if err != nil {
		return "unable to connect to db"
	}
	return "no error"
}

// not done
// not tested
func DeleteUserFriends(userID int, postgres *sql.DB) string {
	stmt, err := postgres.Prepare("DELETE FROM Orbit_buddies WHERE user_id1 = $1 OR user_id2 = $1")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()
	_, err = stmt.Exec(user_id)
	if err != nil {
		return "unable to connect to db"
	}
	return "no error"
}

// not done
// not tested
func DeleteUser(user_id int, postgres *sql.DB) string {
	// delete profile picture
	stmt, err := postgres.Prepare("DELETE FROM Users WHERE user_id = $1")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id)
	if err != nil {
		return "unable to connect to db"
	}

	return "no error"
}

// not done
// not tested
// not documented
func DeleteUserHandler(ctx *gin.Context) {

}
