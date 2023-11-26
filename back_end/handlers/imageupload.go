package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

// not tested
func DeleteImage(filepath string) error {
	err := os.Remove(filepath)
	return err
}

func UploadPic(file multipart.File, header *multipart.FileHeader, dir string) (bool, string) {
	filename := filepath.Join("images", dir, header.Filename)

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
func UpdateProfilePath() string {
	return "no error"
}

// not done
// not tested
// not documented
func ProfilePicHandler(ctx *gin.Context) {
	// Parse the form data, limit to 10 MB
	err := ctx.Request.ParseMultipartForm(10 << 20)
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
	filename := filepath.Join("images", header.Filename)

	// Create the file on the server
	out, err := os.Create(filename)
	if err != nil {
		ctx.String(500, "Internal Server Error")
		return
	}
	defer out.Close()

	// Copy the file data to the server file
	_, err = io.Copy(out, file)
	if err != nil {
		ctx.String(500, "Internal Server Error")
		return
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
		ctx.String(200, "Bad Request")
		return
	}
	defer file.Close()
	success, filename := UploadPic(file, header, "posts")
	if !success {
		ctx.String(200, "Bad Request")
		return
	}
	uuid, err := gocql.ParseUUID(postID)
	if err != nil {
		ctx.String(200, "Bad Request")
		return
	}
	success = UpdatePostPath(uuid, filename, cassandra)
	if !success {
		ctx.String(200, "Bad Rdsdfsdfdsequest")
		return
	}

	// Create a unique filename for the uploaded file

	ctx.String(200, fmt.Sprintf("File %s uploaded successfully!", header.Filename))
}
