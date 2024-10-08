package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义接受数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为控制，则报错，是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" bind:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" bind:"required"`
}

func main() {
	// 1.创建路由
	// 默认使用了两个中间件Logger(), Recover()
	r := gin.Default()
	// JSON绑定
	r.POST("loginJSON", func(c *gin.Context) {
		// 声明接受的变量
		var json Login
		// 将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//判断用户名密码是否正确
		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8080")
	//  curl.exe http://127.0.0.1:8080/loginJSON -H "Content-Type: application/json" -d '{\"user\": \"root\", \"password\": \"admin\"}' -X POST
}
