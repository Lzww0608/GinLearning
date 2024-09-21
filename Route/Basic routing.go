package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// POST 请求的处理函数
func getting(c *gin.Context) {
	// 从请求中获取表单数据或JSON数据
	name := c.DefaultPostForm("name", "Guest")
	c.JSON(http.StatusOK, gin.H{
		"message": "POST received",
		"name":    name,
	})
}

// PUT 请求的处理函数
func putting(c *gin.Context) {
	// 假设处理的是某个更新操作
	id := c.DefaultQuery("id", "0")
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT received",
		"id":      id,
	})
}

func main() {
	r := gin.Default()

	// GET 请求
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})

	// POST 请求
	r.POST("/xxxpost", getting)

	// PUT 请求
	r.PUT("/xxput", putting)

	// 监听端口为8080
	r.Run(":8080")
}
