package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func GetCookie(IP string, cassandra *gocql.Session, postgres *sql.DB, user *User) bool {
	var userID int
	if err := cassandra.Query("Select userID from Cookie where machineID = ?", IP).Scan(&userID); err != nil {
		return false
	}
	if GetUserInfo(userID, postgres, user) != "no error" {
		return false
	}
	return true
}
func GetCookieHandler(ctx *gin.Context) {
	IP := ctx.ClientIP()
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	postgres := ctx.MustGet("postgres").(*sql.DB)
	var user User
	exist := GetCookie(IP, cassandra, postgres, &user)
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
	return
}
func SetupCookie(IP string, userID int, cassandra *gocql.Session) bool {
	if err := cassandra.Query("Insert Into Cookie (machineID, userID) Values (? , ?)", IP, userID).Exec(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func SetCookieHandler(ctx *gin.Context) {
	IP := ctx.ClientIP()

	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	user, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.String(200, "bad request")
		return
	}
	res := SetupCookie(IP, user, cassandra)
	if !res {
		ctx.String(200, "failed")
		return
	}
	ctx.String(200, "success")

}
func RemoveCookie(IP string, cassandra *gocql.Session) bool {
	if err := cassandra.Query("Delete From Cookie Where MachineID = ?", IP).Exec(); err != nil {
		return false
	}
	return true
}
func RemoveCookieHandler(ctx *gin.Context) {
	IP := ctx.ClientIP()

	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	res := RemoveCookie(IP, cassandra)
	if !res {
		ctx.String(200, "failed")
		return
	}
	ctx.String(200, "success")

}
