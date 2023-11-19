package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

// TODO gotta separate JSON afterwards

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

func RemoveFriendHandler(ctx *gin.Context) {

}

func SendFriendRequestHandler(ctx *gin.Context) {
	
}

func RejectFriendRequestHandler(ctx *gin.Context) {
	
}

func GetFriendRequestsHandler(ctx *gin.Context) {
	
}

func UpdateUserProfileHandler(ctx *gin.Context) {
	
}

func GetUserInfoHandler(ctx *gin.Context) {

}


/*
	func Double(ctx *gin.Context) {
		var inputUser User
		if err := ctx.ShouldBindJSON(&inputUser); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newID := inputUser.ID + inputUser.ID
		newName := inputUser.Name + inputUser.Name

		responseUser := User{
			ID:   newID,
			Name: newName,
		}

		ctx.JSON(http.StatusOK, responseUser)
	}
*/
func Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Sum(ctx *gin.Context) {
	num1, err1 := strconv.Atoi(ctx.Param("num1"))
	num2, err2 := strconv.Atoi(ctx.Param("num2"))
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
	} else if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
	}
	sum := num1 + num2
	ctx.String(http.StatusOK, strconv.Itoa(sum))
}

// Expect handler to pass in Cassandra session using ctx.MustGet("cassandra").(*gocql.Session)
func TestCassSelect(session *gocql.Session) {
	var val int
	if err := session.Query("SELECT * FROM test3 WHERE t = 1").Scan(&val); err != nil {
		panic(err)
	}
	fmt.Println(val)
}

func TestInsertHandler(ctx *gin.Context) {
	fmt.Println("ssadasd")
	postgres := ctx.MustGet("postgres").(*sql.DB)
	fmt.Println("ssadasd2")
	fmt.Println(postgres)
	ctx.String(http.StatusOK, strconv.Itoa(1))
	val := ctx.Param("val")
	fmt.Println(val, "Go!")
	TestInsert(val, postgres)

}

// func TestPostgres(ctx *gin.Context) {
// 	stmt, err := postgres.Prepare("SELECT * FROM test3")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stmt.Close()

// 	rows, err := stmt.Query()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var val int
// 		err := rows.Scan(&val)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("%d\n", val)
// 		ctx.JSON(http.StatusOK, gin.H{
// 			"value": val,
// 		})
// 	}
// }
