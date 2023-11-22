package handlers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
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
// not doucumeted
func PostHandler(ctx *gin.Context) {

}
