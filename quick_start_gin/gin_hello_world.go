package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Default port is :%s", port)
	}

	r := gin.Default()

	// 1.
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})
	// 2.
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "I am Taylor Swift")
	})
	// 3.
	name := "Sayali Pol"
	r.GET("/var", func(c *gin.Context) {
		c.String(200, name)
	})
	// 4.
	data := map[string]string{"ML": "Machine Learning",
		"AI":  "Artificial Intelligence",
		"EMB": "Embedded Systems"}

	r.GET("/course", func(c *gin.Context) {
		c.JSON(200, data)
	})

	r.Run()

}
