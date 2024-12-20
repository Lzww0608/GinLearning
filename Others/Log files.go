package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	gin.DisableConsoleColor()

	// Logging to a file
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// gin.,DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run()
}
