package main

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {

	var login Login

	r := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		err := c.ShouldBindJSON(&login)
		if err == nil {
			c.JSON(200, gin.H{
				"User":     login.User,
				"Password": login.Password,
			})

		} else {
			c.JSON(400, gin.H{
				"message": "Invalid Input",
				"err":     err,
			})
		}

	})

	r.POST("/upload", func(c *gin.Context) {

		// single file
		file, _ := c.FormFile("file")

		os.MkdirAll("uploads", os.ModePerm)

		// Upload the file to specific dst.
		dst := filepath.Join("uploads", file.Filename)

		err := c.SaveUploadedFile(file, dst)

		if err == nil {
			c.JSON(200, gin.H{
				"message":  "Uploaded Successfully",
				"filename": file.Filename,
			})
		} else {
			c.JSON(400, gin.H{
				"message": "Failed",
				"err":     err,
			})
		}

	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
