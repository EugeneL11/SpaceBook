package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/EugeneL11/SpaceBook/pkg"
	"github.com/gin-gonic/gin"
)

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
	_, err = stmt.Exec(fullName, username, email, hashedPassword, "Earth", "default", false, "test bio")
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

func RegisterHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	username := ctx.Param("username")
	password := ctx.Param("password")
	fullName := ctx.Param("fullname")
	email := ctx.Param("email")
	var user User
	err := RegisterUser(fullName, password, email, username, postgres, &user)
	if err == "unable to connect to db" || err == "unable to hash password" {
		ctx.JSON(http.StatusOK, gin.H{
			"error":                "unable to create account at this time",
			"id":                   "null",
			"username":             "null",
			"admin":                "null",
			"full_name":            "null",
			"Email":                "null",
			"Home_planet":          "null",
			"Profile_picture_path": "null",
		})
	} else if err == "user name taken" {
		ctx.JSON(http.StatusOK, gin.H{
			"error":                "user name not availible",
			"id":                   "null",
			"username":             "null",
			"admin":                "null",
			"full_name":            "null",
			"Email":                "null",
			"Home_planet":          "null",
			"Profile_picture_path": "null",
		})
	} else if err == "email taken" {
		ctx.JSON(http.StatusOK, gin.H{
			"error":                "email already in use",
			"id":                   "null",
			"username":             "null",
			"admin":                "null",
			"full_name":            "null",
			"Email":                "null",
			"Home_planet":          "null",
			"Profile_picture_path": "null",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"error":                "no error!",
			"id":                   user.User_id,
			"username":             user.User_name,
			"admin":                user.Admin,
			"full_name":            user.Full_name,
			"Email":                user.Email,
			"Home_planet":          user.Home_planet,
			"Profile_picture_path": user.Profile_picture_path,
		})
	}
}

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
			&user.Email, &user.Password, &user.Home_planet, &user.Profile_picture_path, &user.Admin, &user.Bio)
		fmt.Println(user.User_name)
		log.Println(err)
		return true

	} else {
		return false
	}
}

func LoginHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	username := ctx.Param("username")
	password := ctx.Param("password")
	var user User
	correct := LoginCorrect(username, password, postgres, &user)
	// Incorrect Login
	if !correct {
		ctx.JSON(http.StatusOK, gin.H{
			"error":                "unable to find User",
			"id":                   "null",
			"username":             "null",
			"admin":                "false",
			"full_name":            "null",
			"Email":                `null"`,
			"Home_planet":          `null`,
			"Profile_picture_path": "null",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"error":                "no error!",
			"id":                   user.User_id,
			"username":             user.User_name,
			"admin":                user.Admin,
			"full_name":            user.Full_name,
			"Email":                user.Email,
			"Home_planet":          user.Home_planet,
			"Profile_picture_path": user.Profile_picture_path,
		})
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

// not done
// not tested
// not documented
func UpdateUserProfileHandler(ctx *gin.Context) {

}

// not tested
func GetUserInfo(user_id int, postgres *sql.DB, userInfo *API_UserInfo) string {
	stmt, err := postgres.Prepare(`
		SELECT (
			user_id, full_name, user_name, 
			email, home_planet, 
			profile_picture_path, isAdmin, bio
		)
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
			&userInfo.Admin, &userInfo.Bio)
		if err != nil {
			return "unable to connect to db"
		}

		return "no error"
	} else {
		return "invalid user"
	}

	//return "no error"
}

// not done
// not tested
// not documented
func GetUserInfoHandler(ctx *gin.Context) {

}
