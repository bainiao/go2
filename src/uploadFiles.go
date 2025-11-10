package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	dst := "/file/" + file.Filename
	c.SaveUploadedFile(file, dst)
	c.String(http.StatusOK, "%s uploaded success", file.Filename)
}

func UploadMultiFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["uploadp[]"]
	for _, file := range files {
		c.SaveUploadedFile(file, "/file/"+file.Filename)
	}
	c.String(http.StatusOK, "%d files uploaded success", len(files))
}
