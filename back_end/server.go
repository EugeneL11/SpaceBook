package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	_ "github.com/lib/pq"
)

const PORT_NO = ":8080"

// Managing proxies: https://pkg.go.dev/github.com/gin-gonic/gin#section-readme
// TODO - Isolate queries to their own funcs with parameters
// Then call these queries in the route handler funcs to test functionality of queries
// Move route handlers and maybe queries to their own files (ex. routes.go, queries.go)
// TODO Get postgres variable into all handlers

// Connecting PostgreSQL
const (
	host     = "postgres" // service name in docker-compose.yml
	port     = 5432       // PostgreSQL default port
	user     = "postgres" // replace with your PostgreSQL username
	password = "postgres" // replace with your PostgreSQL password
	dbname   = "postgres" // replace with your PostgreSQL database name
)

const connStr = "user=" + user + " password=" + password + " dbname=" + dbname + " host=" + host + " sslmode=disable"

// Connecting CassandraDB (NoSQL)
const addr = "cassandra"

func main() {
	// Using Gin for the server, and settings for server:
	server := gin.Default()
	server.ForwardedByClientIP = true
	server.SetTrustedProxies([]string{"127.0.0.1", "client"}) // Add any other needed IPs
	setupRoutes(server)
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:8080", "*"}
	server.Use(cors.New(config))

	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// initialize(db)
	// defer db.Close()

	// Connecting to postgres:
	server.Use(func(ctx *gin.Context) {
		postgres, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		defer postgres.Close()
		ctx.Set("postgres", postgres)

		ctx.Next()
	})
	// Connecting to CassandraDB:
	server.Use(func(ctx *gin.Context) {
		cluster := gocql.NewCluster(addr)
		cluster.Keyspace = "sb_cassandra" // Name subject to change
		cluster.Consistency = gocql.Quorum

		session, err := cluster.CreateSession()
		if err != nil {
			panic(err)
		}
		defer session.Close()
		ctx.Set("cassandra", session)
		ctx.Next()
	})

	// One-time setup of DB:

	// Route handlers for API endpoints:
	server.Handler()
	// server.GET("/postgresTest", testPostgres)

	server.Run(PORT_NO)
}

// Expect handler to pass in Cassandra session using ctx.MustGet("cassandra").(*gocql.Session)
func testCassSelect(session *gocql.Session) {
	var val int
	if err := session.Query("SELECT * FROM test3 WHERE t = 1").Scan(&val); err != nil {
		panic(err)
	}
	fmt.Println(val)
}

// Expect handler to use ctx.MustGet("postgres").(*sql.DB) and pass in session
func testInsert(val string, postgres *sql.DB) {
	fmt.Println("Go!", postgres)
	stmt, err := postgres.Prepare("INSERT INTO Users (full_name,email,password) VALUES ($1,$2,22)")
	if err != nil {
		panic(err)
	}
	_, err2 := stmt.Exec(val, val)
	fmt.Println("Go2!")
	if err2 != nil {
		panic(err)
	}
}

func testInsertHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	val := ctx.Param("val")
	fmt.Println(val, "Go!")
	testInsert(val, postgres)

}

// func testSelect() {
// 	stmt, err := postgres.Prepare("SELECT * FROM test3 WHERE t = 1")
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
// 		// If multiple columns => rows.Scan(&val1, &val2, &val3)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("%d\n", val)
// 	}
// }

// func testUpdate() {
// 	stmt, err := postgres.Prepare("UPDATE test3 SET t = 2 WHERE t = 1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec()
// }

// func testDelete(e int) {
// 	stmt, err := postgres.Prepare("DELETE FROM test3 WHERE t = ")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stmt.Close()
// }

// func testConditionalSelect() {
// 	stmt, err := postgres.Prepare("SELECT * FROM test3 WHERE t < 5 AND t > 0")
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
// 	}
// }

// func testPostgres(ctx *gin.Context) {
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
