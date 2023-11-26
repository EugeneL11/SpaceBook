package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	_ "github.com/lib/pq"
)

const PORT_NO = ":8080"

// Managing proxies: https://pkg.go.dev/github.com/gin-gonic/gin#section-readme

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
	// Using Gin for the server, and configure server settings:
	server := gin.Default()
	server.ForwardedByClientIP = true
	server.SetTrustedProxies([]string{"127.0.0.1"}) // Add any other needed IPs
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080", "http://client", "https://client", "https://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Connecting to postgres:
	server.Use(func(ctx *gin.Context) {
		postgres, err := sql.Open("postgres", connStr)
		if err != nil {
			fmt.Println("Failed to connect to postgres")
			panic(err)
		}
		defer postgres.Close()
		ctx.Set("postgres", postgres)
		ctx.Next()
	})

	// Connecting to CassandraDB:
	server.Use(func(ctx *gin.Context) {
		cluster := gocql.NewCluster(addr)
		cluster.Keyspace = "cassandra"
		cluster.Consistency = gocql.Quorum
		cluster.ProtoVersion = 4
		cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "cassandra"}

		session, err := cluster.CreateSession()
		if err != nil {
			panic(err)
		}
		defer session.Close()
		ctx.Set("cassandra", session)
		ctx.Next()
	})
	// certPath := "/etc/ssl/certs/localhost.crt"
	// keyPath := "/etc/ssl/private/localhost.key"
	setupRoutes(server)
	server.Static("/images", "./images")
	err := server.Run(PORT_NO)
	if err != nil {
		fmt.Println("Did not Go!")
		panic(err)
	}

}
