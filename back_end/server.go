package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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

const (
	host     = "localhost" // replace with your PostgreSQL host
	port     = 5432        // PostgreSQL default port
	user     = "postgres"  // replace with your PostgreSQL username
	password = ""          // replace with your PostgreSQL password
	dbname   = "test"      // replace with your PostgreSQL database name
)

const connStr = "user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

var postgres *sql.DB

func main() {
	// Using Gin for the server:
	server := gin.Default()
	server.ForwardedByClientIP = true
	server.SetTrustedProxies([]string{"127.0.0.1"}) // Add any other needed IPs
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	postgres, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer postgres.Close()

	err = postgres.Ping()
	if err != nil {
		panic(err)
	}

	// Use a prepared statement to prevent SQL injection
	// Execute the prepared statement with user input
	server.GET("/postgresTest", testPostgres)
	server.GET("/ping", pong)
	server.PUT("/num")
	// server.Use(cors.Default()) // This allows all origins
	server.Use(cors.New(config))
	server.GET("/num/:num1/:num2", sum)
	server.POST("/user", double)
	server.Run(PORT_NO)
}

func testInsert() {
	stmt, err := postgres.Prepare("INSERT INTO test3 (t) VALUES (4)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
}

func testSelect() {
	stmt, err := postgres.Prepare("SELECT * FROM test3 WHERE t = 1")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var val int
		err := rows.Scan(&val)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d\n", val)
	}
}

func testUpdate() {
	stmt, err := postgres.Prepare("UPDATE test3 SET t = 2 WHERE t = 1")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
}

func testDelete() {
	stmt, err := postgres.Prepare("DELETE FROM test3 WHERE t = 1")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
}

func testConditionalSelect() {
	stmt, err := postgres.Prepare("SELECT * FROM test3 WHERE t < 5 AND t > 0")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var val int
		err := rows.Scan(&val)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d\n", val)
	}
}

func testPostgres(ctx *gin.Context) {
	stmt, err := postgres.Prepare("SELECT * FROM test3")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var val int
		err := rows.Scan(&val)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d\n", val)
		ctx.JSON(http.StatusOK, gin.H{
			"value": val,
		})
	}
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
