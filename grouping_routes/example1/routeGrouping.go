package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	v1 := r.Group("/v1")
	v2 := r.Group("/v2")

	v1.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Route v1")
	})

	v2.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Route v2")
	})

	r.Run()
}
