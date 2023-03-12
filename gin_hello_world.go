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

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "I am Taylor Swift")
	})

	name := "Sayali Pol"
	r.GET("/var", func(c *gin.Context) {
		c.String(200, name)
	})

	

	r.Run()

}
