package Controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UploadMultipartFile(c *gin.Engine) {
	c.MaxMultipartMemory = 8 << 20 // 8MB 大小
	c.POST("/upload", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			context.String(http.StatusInternalServerError, "get err %s", err.Error())
			return
		}
		files := form.File["file"]
		for _, file := range files {
			log.Println(file.Filename)
			dst := "./" + file.Filename
			context.SaveUploadedFile(file, dst)
		}
		context.String(http.StatusOK, "upload file %s", len(files))
	})
	c.Run(":8080")
}
