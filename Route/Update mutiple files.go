package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建路由
	//默认使用了2个中间件Logger(), Recover()
	r := gin.Default()
	//限制表单上传大小，8MB，默认为32MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		}
		//获取所有图片
		files := form.File["files"]
		//遍历所有图片
		for _, file := range files {
			//逐个存储
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}
		c.String(200, fmt.Sprintf("%d files uploaded", len(files)))
	})

	r.Run(":8080")
}
