package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Using Gin for data validation simplifies the code and reduces the use of if-else statements
// Person Struct
type Person struct {
	// Cannot be empty and must be greater than 10
	Age      int    `form:"age" binding:"required,gt=10"`
	Name     string `form:"name" binding:"required"`
	Birthday string `form:"birthday" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/Lzww", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprint(err)) // Use 400 for Bad Request
			return
		}

		// Validate the birthday format separately
		_, err := time.Parse("2006-01-02", person.Birthday)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid birthday format, expected YYYY-MM-DD")
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("%#v", person))
	})
	r.Run(":8080")
}
