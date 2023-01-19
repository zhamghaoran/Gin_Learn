package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorld(c gin.RouterGroup) {
	c.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "hello,world"})
	})
}
