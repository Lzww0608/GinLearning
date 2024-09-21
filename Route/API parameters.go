package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()

	// 定义带有路径参数的路由
	r.GET("/:name/*action", func(c *gin.Context) {
		name := c.Param("name")     // 获取路径中的 name 参数
		action := c.Param("action") // 获取路径中的 action 参数
		// 截取字符串，去除前后的 "/"
		action = strings.Trim(action, "/")
		// 返回结果
		c.String(http.StatusOK, name+" is "+action)
	})

	// 监听端口为8080
	r.Run(":8080")
}
