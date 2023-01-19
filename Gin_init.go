package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func ginInit() *gin.Engine {
	r := gin.Default()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) // 同时将日志输出到文件和控制台
	return r
}
