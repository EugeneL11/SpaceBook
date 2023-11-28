package handlers

import (
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

func DeleteImage(filepath string) error {
	err := os.Remove(filepath)
	return err
}

// straight from the big gpt
func generateUniqueFilename() string {
	// Generate a unique identifier (UUID)
	uniqueID := uuid.New()

	// Get the current timestamp
	timestamp := time.Now().Unix()

	// Combine the unique identifier and timestamp to create a unique filename
	uniqueFilename := fmt.Sprintf("%s_%d.txt", uniqueID, timestamp)

	return uniqueFilename
}

func UploadPic(file multipart.File, dir string) (bool, string) {
	// make random somehow
	filename := filepath.Join("images", dir, generateUniqueFilename())

	// Create the file on the server
	out, err := os.Create(filename)
	if err != nil {
		return false, ""
	}
	defer out.Close()

	// Copy the file data to the server file
	_, err = io.Copy(out, file)
	if err != nil {
		return false, ""
	}
	return true, filename
}

// not done
// not tested
func UpdateProfilePath(userID int, newPath string, postgres *sql.DB) bool {

	return true
}

// not done
// not tested
// not documented
// TODO: Rename files to match user ID
func ProfilePicHandler(ctx *gin.Context) {
	// Parse the form data, limit to 10 MB
	userID, err := strconv.Atoi(ctx.Param("userID"))
	postgres := ctx.MustGet("postgres").(*sql.DB)
	if err != nil {
		ctx.String(400, "Bad Request")
		return
	}
	err = ctx.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		ctx.String(400, "Bad Request")
		return
	}

	// Get the file from the form data
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.String(400, "Bad Request")
		return
	}
	defer file.Close()

	// Create a unique filename for the uploaded file
	success, file_name := UploadPic(file, "users")
	if !success {
		ctx.String(500, "Internal Server Error")
		return
	}
	if !UpdateProfilePath(userID, file_name, postgres) {
		ctx.String(500, "Internal Server Error")
	}

	ctx.String(200, fmt.Sprintf("File %s uploaded successfully!", header.Filename))
}

func UpdatePostPath(postID gocql.UUID, path string, cassandra *gocql.Session) bool {
	pathSlice := []string{path}

	updateStmt := cassandra.Query(`
		UPDATE post
		SET imagePaths = imagePaths + ?
		WHERE postID = ?`, pathSlice, postID)

	if err := updateStmt.Exec(); err != nil {
		return false
	}

	return true
}

// not documented
func UploadImagePost(ctx *gin.Context) {
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	err := ctx.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		ctx.String(400, "Bad Request")
		return
	}

	// Get the file from the form data
	file, header, err := ctx.Request.FormFile("image")
	postID := ctx.Param("postID")
	if err != nil {
		ctx.String(400, "Bad Request")
		return
	}
	defer file.Close()
	success, filename := UploadPic(file, "posts")
	if !success {
		ctx.String(400, "Bad Request")
		return
	}
	uuid, err := gocql.ParseUUID(postID)
	if err != nil {
		ctx.String(400, "Bad Request")
		return
	}
	success = UpdatePostPath(uuid, filename, cassandra)
	if !success {
		ctx.String(400, "Bad Rdsdfsdfdsequest")
		return
	}

	ctx.String(200, fmt.Sprintf("File %s uploaded successfully!", header.Filename))
}
