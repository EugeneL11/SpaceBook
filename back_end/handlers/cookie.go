package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

// Returns a user based on the cookie
func GetCookie(ID gocql.UUID, cassandra *gocql.Session, postgres *sql.DB, user *User) bool {
	var userID int
	if err := cassandra.Query("Select userID from Cookie where machineID = ?", ID).Scan(&userID); err != nil {
		return false
	}
	if GetUserInfo(userID, postgres, user) != "no error" {
		return false
	}
	return true
}

// Processes request to get a user given a cookie
func GetCookieHandler(ctx *gin.Context) {
	cookieID := ctx.Param("CookieID")

	ID, err := gocql.ParseUUID(cookieID)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "no user",
			"user":   nil,
		})
		fmt.Print("Invalid cookie value. Setting to default.")
		return
	}

	fmt.Print("Cookie exists. ID:", ID.String())

	// Continue with the rest of your logic (e.g., querying the database)
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	postgres := ctx.MustGet("postgres").(*sql.DB)
	var user User
	exist := GetCookie(ID, cassandra, postgres, &user)
	if !exist {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "no user",
			"user":   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "user found",
		"user":   user,
	})
}

// Associates a cookie ID with a user
func SetupCookie(ID gocql.UUID, userID int, cassandra *gocql.Session) bool {
	if err := cassandra.Query("Insert Into Cookie (machineID, userID) Values (? , ?)", ID, userID).Exec(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Processes request to associate cookie with a user
func SetCookieHandler(ctx *gin.Context) {
	cookieID := ctx.Param("CookieID")
	ID, err := gocql.ParseUUID(cookieID)
	if err != nil {
		ctx.SetCookie("CookieID", "123", 3600, "/", "localhost", false, false)
		ctx.JSON(http.StatusOK, gin.H{
			"status": "no user",
			"user":   nil,
		})
		return
	}

	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	user, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.String(200, "bad request")
		return
	}
	res := SetupCookie(ID, user, cassandra)
	if !res {
		ctx.String(200, "failed")
		return
	}
	ctx.String(200, "success")

}

// Removes saved information given a cookie
func RemoveCookie(ID gocql.UUID, cassandra *gocql.Session) bool {
	if err := cassandra.Query("Delete From Cookie Where MachineID = ?", ID).Exec(); err != nil {
		return false
	}
	return true
}

// Processes request to remove a cookie
func RemoveCookieHandler(ctx *gin.Context) {
	cookieID := ctx.Param("CookieID")
	ID, err := gocql.ParseUUID(cookieID)
	if err != nil {
		ctx.SetCookie("CookieID", "123", 3600, "/", "localhost", false, false)
		ctx.JSON(http.StatusOK, gin.H{
			"status": "no user",
			"user":   nil,
		})
		return
	}

	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	res := RemoveCookie(ID, cassandra)
	if !res {
		ctx.String(200, "failed")
		return
	}
	ctx.String(200, "success")

}

// Gives the server a cookie to store
func CreateCookieHandler(ctx *gin.Context) {
	// Generate a unique ID for the cookie
	cookieID := gocql.TimeUUID()

	ctx.String(200, cookieID.String())
}
