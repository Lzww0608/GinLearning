package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 1.异步
	r.GET("/long_async", func(c *gin.Context) {
		// 需要一个副本
		copyContext := c.Copy()
		// 异步副本
		go func() {
			time.Sleep(1 * time.Second)
			log.Println("异步执行: " + copyContext.Request.URL.Path)
		}()
	})
	// 2.同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(1 * time.Second)
		log.Println("同步执行: " + c.Request.URL.Path)
	})
	r.Run(":8080")
}
