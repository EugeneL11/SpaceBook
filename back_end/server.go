package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const PORT_NO = ":8080"

// Managing proxies: https://pkg.go.dev/github.com/gin-gonic/gin#section-readme
// CORS: https://github.com/gin-contrib/cors
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Using Gin for the server:
	server := gin.Default()
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{URLs Go here}
	server.GET("/ping", pong)
	server.PUT("/num")
	server.Use(cors.Default()) // This allows all origins
	// server.Use(cors.New(config))
	server.GET("/num/:num1/:num2", sum)
	server.POST("/user", double)
	server.Run(PORT_NO)
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
