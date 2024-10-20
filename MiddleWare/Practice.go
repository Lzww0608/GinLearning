package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 定义程序计时中间件，然后定义2个路由，执行函数后应该打印统计的执行时间

func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	// 统计时间
	since := time.Since(start)
	fmt.Println("程序用时", since)
}

func main() {
	r := gin.Default()
	r.Use(myTime)

	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	r.Run(":8080")
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(1 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(2 * time.Second)
}
