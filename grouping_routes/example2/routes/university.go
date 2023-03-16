package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addRoutes(rg *gin.RouterGroup) {
	uniName := rg.Group("/un")			//url: http://localhost:5000/university/un/

	uniName.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Students !!!")
	})

	uniName.GET("/BU", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"University": "Boston University",
		})
	})

	uniName.GET("/HU", func(c *gin.Context) {
		c.JSON(200, "Harvard University")
	})
}
