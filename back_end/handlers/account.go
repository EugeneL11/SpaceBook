package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/EugeneL11/SpaceBook/pkg"
	"github.com/gin-gonic/gin"
)

// Add new non-admin user to SQL db, returning error message if unsuccessful
func RegisterUser(fullName string, password string, email string, username string, postgres *sql.DB, user *User) string {
	stmt, err := postgres.Prepare("SELECT * FROM Users WHERE email = $1")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(email)
	if err2 != nil {
		return "unable to connect to db"
	}

	if rows.Next() {
		return "email taken"
	}
	stmt, err = postgres.Prepare("SELECT * FROM Users WHERE user_name = $1")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 = stmt.Query(username)
	if err2 != nil {
		return "unable to connect to db"
	}

	if rows.Next() {
		return "user name taken"
	}
	hashedPassword, err := pkg.HashPassword(password)
	if err != nil {
		// Unable to hash the password
		return "unable to hash password"
	}
	stmt, err = postgres.Prepare(`
    INSERT INTO Users (full_name, user_name, email, password, home_planet, profile_picture_path, isAdmin, bio)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)
	if err != nil {
		return "unable to connect to db"
	}
	_, err = stmt.Exec(fullName, username, email, hashedPassword, "Earth", "/images/utilities/pp.png", false, "test bio")
	if err != nil {
		fmt.Println("Could not execute insert into users")
		return "unable to connect to db"
	}
	stmt, err = postgres.Prepare("SELECT user_id FROM Users WHERE email = $1")
	if err != nil {
		return "unable to connect to db"
	}
	rows, err = stmt.Query(email)
	if err != nil {
		return "unable to connect to db"
	}
	if rows.Next() {
		err = rows.Scan(&user.User_id)
		if err != nil {
			return "unable to connect to db"
		}
	}
	user.Full_name = fullName
	user.User_name = username
	user.Home_planet = "Earth"
	user.Profile_picture_path = "Default"
	user.Email = email
	user.Admin = false

	return "no error"
}

// Route handler for registering a new user, returning error/user info in JSON
func RegisterHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	username := ctx.Param("username")
	password := ctx.Param("password")
	fullName := ctx.Param("fullname")
	email := ctx.Param("email")
	var user User
	err := RegisterUser(fullName, password, email, username, postgres, &user)
	if err == "unable to connect to db" || err == "unable to hash password" {
		ctx.JSON(http.StatusOK, ErrorUserResponse("unable to create account at this time"))
	} else if err == "user name taken" {
		ctx.JSON(http.StatusOK, ErrorUserResponse("user name not available"))
	} else if err == "email taken" {
		ctx.JSON(http.StatusOK, ErrorUserResponse("email already in use"))
	} else {
		ctx.JSON(http.StatusOK, GoodUserResponse(user))
	}
}

// Confirms whether provided username/password combo is consistent with SQL db
func LoginCorrect(username string, password string, postgres *sql.DB, user *User) bool {
	stmt, err := postgres.Prepare("SELECT password FROM users WHERE user_name = $1")
	if err != nil {
		return false
	}
	defer stmt.Close()

	// Query the database and get the hashed password stored for the user
	row := stmt.QueryRow(username)
	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		return false
	}

	// Compare the hashed password with the user provided password
	isCorrect := pkg.VerifyPassword(password, []byte(hashedPassword))

	if !isCorrect {
		// fmt.Println("Entered password", password, "incorrectly matches hashed password! :(")
		return false
	}

	// fmt.Println("Entered password", password, "correctly matches hashed password!")

	// Get user's information to give frontend
	stmt, err = postgres.Prepare("SELECT * FROM users WHERE user_name = $1 and password = $2")
	if err != nil {
		log.Panic(err)
		return false
	}
	defer stmt.Close()

	rows, err := stmt.Query(username, hashedPassword)
	if err != nil {
		log.Panic(err)
		return false
	}
	if rows.Next() {
		err := rows.Scan(&user.User_id, &user.Full_name, &user.User_name,
			&user.Email, nil, &user.Home_planet, &user.Profile_picture_path, &user.Admin, &user.Bio)
		fmt.Println(user.User_name)
		log.Println(err)
		return true

	} else {
		return false
	}
}

// Route handler for /login, returning a JSON with error/user info
func LoginHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	username := ctx.Param("username")
	password := ctx.Param("password")
	var user User
	correct := LoginCorrect(username, password, postgres, &user)
	if !correct {
		ctx.JSON(http.StatusOK, ErrorUserResponse("unable to find User"))
	} else {
		ctx.JSON(http.StatusOK, GoodUserResponse(user))
	}
}

// not tested
func UpdateUserProfile(
	user_id int, new_fullname string,
	new_home_planet string,
	bio string, postgres *sql.DB,
) string {
	stmt, err := postgres.Prepare(`
		UPDATE Users 
		SET full_name = $2,
		home_planet = $3,
		bio = $4
		WHERE user_id = $1
	`)
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id, new_fullname, new_home_planet, bio)
	if err != nil {
		return "unable to connect to db"
	}

	return "no error"
}

// not tested
func UpdateUserProfileHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	userID, err1 := strconv.Atoi(ctx.Param("userID"))
	newName := ctx.Param("newFullName")
	newPlanet := ctx.Param("newPlanet")
	newBio := ctx.Param("newBio")

	if err1 != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "unable to parse input",
		})
		return
	}
	status := UpdateUserProfile(userID, newName, newPlanet, newBio, postgres)
	ctx.JSON(http.StatusOK, gin.H{
		"status": status,
	})

}

func GetUserInfo(user_id int, postgres *sql.DB, userInfo *User) string {
	stmt, err := postgres.Prepare(`
		SELECT 
			user_id, full_name, user_name, 
			email, home_planet, 
			profile_picture_path, bio
		
		FROM Users 
		WHERE user_id = $1
	`)
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(user_id)
	if err2 != nil {
		return "unable to connect to db"
	}

	row := rows.Next()
	if row {
		err := rows.Scan(
			&userInfo.User_id, &userInfo.Full_name, &userInfo.User_name,
			&userInfo.Email, &userInfo.Home_planet, &userInfo.Profile_picture_path,
			&userInfo.Bio)
		if err != nil {
			return "unable to connect to db"
		}

		return "no error"
	} else {
		return "invalid user"
	}

	//return "no error"
}

func GetUserInfoHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	viewer, err1 := strconv.Atoi(ctx.Param("viewer"))
	viewed, err2 := strconv.Atoi(ctx.Param("viewed"))
	if err1 != nil || err2 != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":       "bad request",
			"user":         nil,
			"friendstatus": nil,
		})
		return
	}
	var user User
	result := GetUserInfo(viewed, postgres, &user)
	if result != "no error" {
		ctx.JSON(http.StatusOK, gin.H{
			"status":       "result",
			"user":         nil,
			"friendstatus": nil,
		})
		return
	}
	status := FriendStatus(viewer, viewed, postgres)
	if status == "unable to connect to db" {
		ctx.JSON(http.StatusOK, gin.H{
			"status":       status,
			"user":         nil,
			"friendstatus": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":       "no error",
		"user":         user,
		"friendstatus": status,
	})
}

// Used by GetUserInfoHandler to determine friend status
func FriendStatus(viewer int, viewed int, postgres *sql.DB) string {
	if viewer == viewed {
		return "own profile"
	}
	if viewer > viewed {
		viewer, viewed = viewed, viewer
	}
	stmt, err := postgres.Prepare("SELECT * FROM Orbit_buddies WHERE user1_id = $1 and user2_id = $2")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(viewer, viewed)
	if err2 != nil {
		return "unable to connect to db"
	}

	if rows.Next() {
		return "already friends"
	}
	stmt, err = postgres.Prepare("SELECT * FROM orbit_requests WHERE requested_buddy_id = $1 and requester_id = $2")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 = stmt.Query(viewed, viewer)
	if err2 != nil {
		return "unable to connect to db"
	}

	if rows.Next() {
		return "viewer sent request"
	}
	stmt, err = postgres.Prepare("SELECT * FROM orbit_requests WHERE requested_buddy_id = $1 and requester_id = $2")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 = stmt.Query(viewer, viewed)
	if err2 != nil {
		return "unable to connect to db"
	}

	if rows.Next() {
		return "viewed person sent request"
	}
	return "no requests"
}
