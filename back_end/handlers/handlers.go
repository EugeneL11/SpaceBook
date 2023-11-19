package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

// TODO gotta separate JSON afterwards

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
