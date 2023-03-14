package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func main() {
	university := r.Group("/university")

	university.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello students")
	})

	university.GET("/BU", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Boston University, Massachusetts")
	})

	university.GET("/BU/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, " Hello candidate "+id)
	})

	r.Run()
}
