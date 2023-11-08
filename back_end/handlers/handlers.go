package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func double(ctx *gin.Context) {
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

func pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func sum(ctx *gin.Context) {
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
