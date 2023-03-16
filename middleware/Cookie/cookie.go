package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CookieTool() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("person"); err == nil {
			if cookie == "Taylor" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "No cokie found"})
	}
}

func main() {
	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		// syntax : c.SetCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool)
		c.SetCookie("person", "Taylor", 30, "/", "localhost", false, true)
		c.String(http.StatusOK, "Hello Taylor Swift")
	})

	r.GET("/profile", CookieTool(), func(c *gin.Context) {
		c.JSON(200, gin.H{"Details": "Hey, Taylor Swift"})
	})

	r.Run(":5000")
}
