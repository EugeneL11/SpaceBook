package handlers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

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
