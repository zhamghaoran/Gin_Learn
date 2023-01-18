package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloWorld() {
	c := gin.Default()
	c.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "hello,world"})
	})
	c.Run(":8080")
}
