package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

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
	// hashedPassword, err := pkg.GeneratePasswordHash(password)
	hashedPassword := 22
	// fmt.Println(hashedPassword)
	// if err != nil {
	// 	// Unable to hash the password
	// 	return "unable to hash password"
	// }
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

func DeleteUser(user_id int, postgres *sql.DB) string {

	stmt, err := postgres.Prepare("DELETE FROM Users WHERE user_id = $1")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id)
	if err != nil {
		return "unable to connect to db"
	}
	// delete their posts, likes, and comments!!

	return "no error"
}

func LoginCorrect(email string, password string, postgres *sql.DB, user *User) string {
	// hashedPassword, err := pkg.GeneratePasswordHash(password)
	// if err != nil {
	// 	// Could not hash password
	// 	return false
	// }
	hashedPassword := 22
	stmt, err := postgres.Prepare("Select * from Users WHERE user_name = $1 AND password = $2")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(email, hashedPassword)
	if err2 != nil {
		return "unable to connect to db"
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.User_id, &user.Full_name, &user.User_name,
			&user.Email, &user.Password, &user.Home_planet, &user.Profile_picture_path, &user.Admin)

		if err != nil {
			return "unable to connect to db"
		}

		return "no errer"
	} else {
		return "invalid login credentials"
	}
}
func UpdateUserProfile(
	user_id int, new_username string, new_fullname string,
	new_email string, new_home_planet string,
	new_profile_pic_path string, bio string, postgres *sql.DB,
) string {
	stmt, err := postgres.Prepare(`
		UPDATE Users 
		SET user_name = $2, full_name = $3,
		email = $4, planet = $5,
		profile_picture_path = $6, bio = $7
		WHERE user_id = $1
	`)
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id, new_username, new_fullname, new_email, new_home_planet, new_profile_pic_path, bio)
	if err != nil {
		return "unable to connect to db"
	}

	return "no error"
}

func GetUserInfo(user_id int, postgres *sql.DB, userInfo *API_UserInfo) (API_UserInfo, string) {
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
		return nil, "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(user_id)
	if err2 != nil {
		return nil, "unable to connect to db"
	}

	row := rows.next()
	if row {
		err := rows.Scan(
			&userInfo.User_id, &userInfo.Full_name, &userInfo.User_name,
			&userInfo.Email, &userInfo.Home_planet, &userInfo.Profile_picture_path,
			&userInfo.Admin, &userInfo.Bio)
		if err != nil {
			return nil, "unable to connect to db"
		}

		return userInfo, "no error"
	} else {
		return nil, "invalid user"
	}

	//return "no error"
}

func LoginHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	username := ctx.Param("username")
	password := ctx.Param("password")
	var user User
	err := LoginCorrect(username, password, postgres, &user)
	if err != "no error" {
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

func DeleteUserHandler(ctx *gin.Context) {

}

func UpdateUserProfileHandler(ctx *gin.Context) {

}

func GetUserInfoHandler(ctx *gin.Context) {

}
