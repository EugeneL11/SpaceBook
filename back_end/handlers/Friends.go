package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// not tested
func GetFriends(user_id int, postgres *sql.DB) ([]API_UserInfo, string) {
	stmt, err := postgres.Prepare(`
		SELECT (
			user_id, full_name, user_name,
			email, home_planet, 
			profile_picture_path, isAdmin, bio
		)
		FROM Users
		WHERE user_id IN (
			SELECT user2_id 
			FROM Orbit_Buddies 
			WHERE user1_id = $1
			UNION 
			SELECT user1_id 
			FROM Orbit_Buddies 
			WHERE user2_id = $1
		)
	`)
	if err != nil {
		return nil, "unable to connect to db"
	}
	defer stmt.Close()

	rows, err := stmt.Query(user_id)
	if err != nil {
		return nil, "unable to connect to db"
	}
	defer rows.Close()

	var mySlice []API_UserInfo
	for rows.Next() {
		var newUser API_UserInfo
		err := rows.Scan(
			&newUser.User_id, &newUser.Full_name, &newUser.User_name,
			&newUser.Email, &newUser.Home_planet, &newUser.Profile_picture_path,
			&newUser.Admin, &newUser.Bio,
		)
		if err != nil {
			return nil, "unable to connect to db"
		}
		mySlice = append(mySlice, newUser)
	}

	return mySlice, "no error"
}

// not tested
// not documented
func GetFriendsHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	user_id, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		return
	}

	var users []API_UserInfo

	users, err2 := GetFriends(user_id, postgres)
	ctx.JSON(http.StatusOK, gin.H{
		"error": err2,
	})

	usersJson, err := json.Marshal(users)
	log.Println(string(usersJson))

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
			"users": usersJson,
		})
	}
}

// not tested
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

// not done
// not tested
// not documented
func RemoveFriendHandler(ctx *gin.Context) {

}

// not tested
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

// not done
// not tested
// not documented
func SendFriendRequestHandler(ctx *gin.Context) {

}

// not tested
func RejectFriendRequest(rejecter_id int, rejectee_id int, postgres *sql.DB) string {
	stmt, err := postgres.Prepare(`
		DELETE FROM Orbit_Requests 
		WHERE requested_buddy_id = $1 AND requester_id = $2 
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

// not done
// not tested
// not documented
func RejectFriendRequestHandler(ctx *gin.Context) {

}

// not tested
func GetFriendRequests(user_id int, postgres *sql.DB) ([]API_UserInfo, string) {
	stmt, err := postgres.Prepare(`
		SELECT requester_id 
		FROM Orbit_Requests 
		WHERE requested_buddy_id = $1
	`)
	if err != nil {
		return nil, "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(user_id)
	if err2 != nil {
		return nil, "unable to connect to db"
	}

	var mySlice []API_UserInfo
	for rows.Next() {
		var newUser API_UserInfo
		err := rows.Scan(
			&newUser.User_id, &newUser.Full_name, &newUser.User_name,
			&newUser.Email, &newUser.Home_planet, &newUser.Profile_picture_path,
			&newUser.Admin, &newUser.Bio,
		)
		if err != nil {
			return nil, "unable to connect to db"
		}
		mySlice = append(mySlice, newUser)
	}

	return mySlice, "no error"

}

// not done
// not tested
// not documented
func GetFriendRequestsHandler(ctx *gin.Context) {

}

func SearchPeople(userID int, postgres *sql.DB) {

}

func SearchPeopleHandler(ctx *gin.Context) {

}
