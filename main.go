package main

import (
	"Gin_Learn/Controller"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	c := ginInit()
	group := c.RouterGroup
	Controller.HelloWorld(group)
	Controller.ShouldBind(group)
	Controller.UploadFile(group)
	err := c.Run(":8080")
	if err != nil {
		return
	}
}
