package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// returns all friends (including the info needed to display them) given a user id
func GetFriends(user_id int, postgres *sql.DB) ([]UserPreview, string) {
	stmt, err := postgres.Prepare(`
		SELECT 
			u.user_id, u.full_name, u.user_name, 
			u.profile_picture_path
		FROM Users as u
		WHERE EXISTS (
			SELECT * 
			FROM Orbit_Buddies 
			WHERE user1_id = $1 and user2_id= u.user_id
		) or exists (
			SELECT * 
			FROM Orbit_Buddies 
			WHERE user2_id = $1 and user1_id = u.user_id
		)
	`)
	if err != nil {
		return nil, "unable to connect to db 1"
	}
	defer stmt.Close()

	rows, err := stmt.Query(user_id)
	if err != nil {
		fmt.Println(err)
		return nil, "unable to connect to db 2"
	}
	defer rows.Close()

	var mySlice []UserPreview
	for rows.Next() {
		var newUser UserPreview
		err := rows.Scan(
			&newUser.UserID, &newUser.Full_name, &newUser.User_name,
			&newUser.Profile_picture_path,
		)
		if err != nil {
			fmt.Println(err)
			return nil, "unable to connect to db 3"
		}
		mySlice = append(mySlice, newUser)
	}

	return mySlice, "no error"
}

// processes get all friends relations given a user id
func GetFriendsHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	user_id, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		return
	}

	users, err2 := GetFriends(user_id, postgres)
	if err2 != "no error" {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err2,
			"users": nil,
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
			"users": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "no error",
			"users": users,
		})
	}
}

// removes a friend relation given the two users
func RemoveFriend(user1_id int, user2_id int, postgres *sql.DB) string {
	if user1_id > user2_id {
		temp := user1_id
		user1_id = user2_id
		user2_id = temp
	}

	stmt, err := postgres.Prepare(`
		DELETE FROM Orbit_Buddies
		WHERE user1_id = $1 AND user2_id = $2
	`)
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	_, err = stmt.Exec(user1_id, user2_id)
	if err != nil {
		return "unable to connect to db"
	}

	return "no error"
}

// processes requets to remove friends
func RemoveFriendHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	user1_id, err1 := strconv.Atoi(ctx.Param("id1"))
	user2_id, err2 := strconv.Atoi(ctx.Param("id2"))
	if err1 != nil {
		log.Panic(err1)
	} else if err2 != nil {
		log.Panic(err2)
	}

	result := RemoveFriend(user1_id, user2_id, postgres)
	ctx.JSON(http.StatusOK, gin.H{
		"status": result,
	})
}

// either creates a request from one user to another or accepts a request
func SendFriendRequest(sender_id int, receiver_id int, postgres *sql.DB) string {
	stmt, err := postgres.Prepare("Select * from Orbit_Requests WHERE requester_id = $1 AND requested_buddy_id = $2")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(receiver_id, sender_id)
	if err2 != nil {
		return "unable to connect to db"
	}
	if rows.Next() {
		stmt, err = postgres.Prepare("DELETE FROM Orbit_Requests WHERE requester_id = $1 AND requested_buddy_id = $2")
		if err != nil {
			return "unable to connect to db"
		}
		defer stmt.Close()

		_, err2 = stmt.Exec(receiver_id, sender_id)
		if err2 != nil {
			return "unable to connect to db"
		}
		stmt, err = postgres.Prepare("Insert Into Orbit_Buddies (user1_id,user2_id) Values ($1,$2)")
		if err != nil {
			return "unable to connect to db"
		}
		defer stmt.Close()

		if sender_id > receiver_id {
			temp := sender_id
			sender_id = receiver_id
			receiver_id = temp
		}

		_, err2 = stmt.Exec(sender_id, receiver_id)
		if err2 != nil {
			return "unable to connect to db"
		}

		return "no error"

	} else {
		stmt, err = postgres.Prepare("Insert Into Orbit_Requests (requester_id,requested_buddy_id) Values ($1,$2)")
		if err != nil {
			return "unable to connect to db"
		}
		defer stmt.Close()

		_, err2 = stmt.Exec(sender_id, receiver_id)
		if err2 != nil {
			return "unable to connect to db"
		}

	}

	return "no error"
}

// processes the request
func SendFriendRequestHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	sender, err1 := strconv.Atoi(ctx.Param("sender_user_id"))
	reciever, err2 := strconv.Atoi(ctx.Param("receiver_user_id"))
	if err1 != nil {
		log.Panic(err1)
	} else if err2 != nil {
		log.Panic(err2)
	}

	result := SendFriendRequest(sender, reciever, postgres)
	ctx.JSON(http.StatusOK, gin.H{
		"status": result,
	})
}

func RejectFriendRequest(rejecter_id int, rejectee_id int, postgres *sql.DB) string {
	stmt, err := postgres.Prepare(`
		DELETE FROM Orbit_Requests 
		WHERE requester_id = $1 AND requested_buddy_id = $2 
	`)
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(rejecter_id, rejectee_id)
	if err2 != nil {
		return "unable to connect to db"
	}

	return "no error"
}

func RejectFriendRequestHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	rejecter, err1 := strconv.Atoi(ctx.Param("rejecter_id"))
	rejectee, err2 := strconv.Atoi(ctx.Param("rejectee_id"))
	if err1 != nil {
		log.Panic(err1)
	} else if err2 != nil {
		log.Panic(err2)
	}

	result := RejectFriendRequest(rejecter, rejectee, postgres)
	ctx.JSON(http.StatusOK, gin.H{
		"status": result,
	})

}

// returns all requests to a friend given a user id
func GetFriendRequests(user_id int, postgres *sql.DB) ([]UserPreview, string) {
	stmt, err := postgres.Prepare(`
	SELECT u.user_id, u.full_name, u.user_name, u.profile_picture_path
	FROM users u
	WHERE EXISTS (
		SELECT 1
		FROM orbit_requests
		WHERE requested_buddy_id = $1
		AND requester_id = u.user_id
	)`)
	if err != nil {
		return nil, "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(user_id)
	if err2 != nil {
		return nil, "unable to connect to db"
	}

	var mySlice []UserPreview
	for rows.Next() {
		var newUser UserPreview
		err := rows.Scan(
			&newUser.UserID, &newUser.Full_name, &newUser.User_name,
			&newUser.Profile_picture_path,
		)
		if err != nil {
			return nil, "unable to connect to db"
		}
		mySlice = append(mySlice, newUser)
	}

	return mySlice, "no error"

}

// processes request of getting all friend requests
func GetFriendRequestsHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	user, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":   "bad request",
			"requests": nil,
		})
	}

	requests, result := GetFriendRequests(user, postgres)
	if result == "unable to connect to db" {
		ctx.JSON(http.StatusOK, gin.H{
			"status":   result,
			"requests": nil,
		})
	} else {
		var status string
		if len(requests) == 0 {
			status = "no requests"
		} else {
			status = "pending request"
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status":   status,
			"requests": requests,
		})

	}
}

// gets a list of people based on a prefix, cannot return own account
func SearchPeople(userID int, searchTerm string, postgres *sql.DB) (string, []UserPreview) {
	searchTerm = strings.ToLower(searchTerm)
	stmt, err := postgres.Prepare(`SELECT user_id, full_name, user_name, profile_picture_path FROM USERS
	WHERE (user_name LIKE $1 OR user_name = $2) and not user_id = $3
	LIMIT 20`)
	if err != nil {
		return "unable to connect to db", nil
	}
	SQLsearchTerm := searchTerm + "%"
	var users []UserPreview
	row, err := stmt.Query(SQLsearchTerm, searchTerm, userID)
	if err != nil {
		return "unable to connect to db", nil
	}
	for row.Next() {
		var newUser UserPreview
		err := row.Scan(
			&newUser.UserID, &newUser.Full_name, &newUser.User_name,
			&newUser.Profile_picture_path,
		)
		if err != nil {
			return "unable to connect to db", nil
		}
		users = append(users, newUser)
	}
	return "no error", users
}

// processes request to see all users with a certain prefix
func SearchPeopleHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error":        "bad request",
			"userPreviews": nil,
		})
	}
	searchTerm := ctx.Param("searchTerm")
	errMsg, users := SearchPeople(userID, searchTerm, postgres)
	if len(users) == 0 {
		errMsg = "no users found"
	}
	// Return JSON with "error" and "userPreviews", which is a nested JSON
	if errMsg == "no error" {

		ctx.JSON(http.StatusOK, gin.H{
			"error":        "no error",
			"userPreviews": users,
		})

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"error":        errMsg,
			"userPreviews": nil,
		})
	}
}
