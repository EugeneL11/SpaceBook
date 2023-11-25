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

// testing- will delete later
func ImageHandler(c *gin.Context) {
	// Parse the form data, limit to 10 MB
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.String(400, "Bad Request")
		return
	}

	// Get the file from the form data
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.String(400, "Bad Request")
		return
	}
	defer file.Close()

	// Create a unique filename for the uploaded file
	filename := filepath.Join("images", header.Filename)

	// Create the file on the server
	out, err := os.Create(filename)
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
	defer out.Close()

	// Copy the file data to the server file
	_, err = io.Copy(out, file)
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}

	c.String(200, fmt.Sprintf("File %s uploaded successfully!", header.Filename))
}

// not tested
func DeleteImage(filepath string) error {
	err := os.Remove(filepath)
	return err
}

// not done
// not tested
func UploadProfilePic() {

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

// not done
// not tested
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
func UpdatePostPath(postID gocql.UUID, path string, cassandra *gocql.Session) bool {

	updateStmt := cassandra.Query(`UPDATE post
			SET imagePaths += ?
			WHERE postID = ?`)

	if err := updateStmt.Bind(path, postID).Exec(); err != nil {
		return false
	}

	return true

}

// not done
// not tested
// not doucumeted
func PostHandler(ctx *gin.Context) {
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
	success = UpdatePostPath(uuid, filename, cassandra)
	if !success {
		return
	}

	// Create a unique filename for the uploaded file

	ctx.String(200, fmt.Sprintf("File %s uploaded successfully!", header.Filename))
}
