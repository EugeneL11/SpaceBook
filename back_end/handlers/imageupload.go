package handlers

import (
	"database/sql"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

// removes an image from our file system
func DeleteImage(filepath string) error {

	err := os.Remove(filepath[1:])
	return err
}

// taken from chatGPT
// extracts the file extension
func getFileExtension(file multipart.File) string {
	fileHeader := make([]byte, 512) // Read the first 512 bytes to detect the file type
	_, err := file.Read(fileHeader)
	if err != nil {
		fmt.Println("Error reading file header:", err)
		return ""
	}

	fileType := http.DetectContentType(fileHeader)
	switch fileType {
	case "image/jpeg":
		return "jpg"
	case "image/png":
		return "png"
	// Add more cases for other file types if needed
	default:
		// If the file type is not recognized, you can use the file name to extract the extension
		_, fileHeaderParams, _ := mime.ParseMediaType(fileType)
		return path.Ext(fileHeaderParams["name"])
	}
}

// taken from chatGPT
// ensures a unique name for each file
func generateUniqueFilename(ext string) string {
	// Generate a unique identifier (UUID)
	uniqueID := uuid.New()

	// Get the current timestamp
	timestamp := time.Now().Unix()

	// Combine the unique identifier and timestamp to create a unique filename
	uniqueFilename := fmt.Sprintf("%s_%d.%s", uniqueID, timestamp, ext)

	return uniqueFilename
}

// uploads an image to our file
func UploadPic(file multipart.File, header *multipart.FileHeader, dir string) (bool, string) {
	// make random somehow
	fileExt := filepath.Ext(header.Filename)

	filename := filepath.Join("images", dir, generateUniqueFilename(fileExt))
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

// Updates the the users profile picture to point to a correct path in file system
func UpdateProfilePath(userID int, newPath string, postgres *sql.DB) bool {
	stmt, err := postgres.Prepare("Select Profile_picture_path from Users WHERE user_id = $1")
	if err != nil {
		fmt.Println("1")
		return false
	}
	defer stmt.Close()

	row, err := stmt.Query(userID)
	if err != nil {
		fmt.Println("2")
		return false
	}

	if row.Next() {
		var path string
		row.Scan(&path)
		if path == "/images/utilities/pp.png" {

		} else if DeleteImage(path) != nil {
			fmt.Println("3")
			return false
		}
	} else {
		fmt.Println("4")
		return false
	}
	stmt, err = postgres.Prepare("Update Users set Profile_picture_path = $1 WHERE user_id = $2")
	if err != nil {
		fmt.Println("5")
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(newPath, userID)
	if err != nil {
		fmt.Println("6")
		return false
	}
	return true
}

// Handles API call for changing a user's pfp
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
	success, file_name := UploadPic(file, header, "users")
	if !success {
		ctx.String(500, "Internal Server Error")
		return
	}
	if !UpdateProfilePath(userID, "/"+file_name, postgres) {
		ctx.String(500, "Internal Server Error Whasupp")
	}

	ctx.String(200, fmt.Sprintf("File %s uploaded successfully!", header.Filename))
}

// Adds a path to a post object
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

// Uploads an image with an associated post
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
	success, filename := UploadPic(file, header, "posts")
	if !success {
		ctx.String(400, "Bad Request")
		return
	}
	uuid, err := gocql.ParseUUID(postID)
	if err != nil {
		ctx.String(400, "Bad Request")
		return
	}
	success = UpdatePostPath(uuid, "/"+filename, cassandra)
	if !success {
		ctx.String(400, "Bad Request")
		return
	}

	ctx.String(200, fmt.Sprintf("File %s uploaded successfully!", header.Filename))
}
