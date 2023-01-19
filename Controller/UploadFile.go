package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(c gin.RouterGroup) {
	c.POST("/upload", func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			context.String(http.StatusInternalServerError, "get err %s", err.Error())
		}
		dst := "./" + file.Filename
		context.SaveUploadedFile(file, dst)
		context.JSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "upload Success",
		})
	})
}
