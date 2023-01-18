package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func ShouldBind() {
	c := gin.Default()
	c.POST("/login", func(context *gin.Context) {
		var user User
		err := context.ShouldBind(&user)
		if err != nil {
			context.JSON(http.StatusInternalServerError, err)
		}
		if user.Password == "123456" && user.Username == "root" {
			context.JSON(http.StatusOK, gin.H{"status": "login successfully"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "unauthorized"})
	})
	c.Run(":8080")
}
