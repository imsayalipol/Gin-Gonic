package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Add a middleware function that prints a message before and after calling the next middleware
	r.Use(func(c *gin.Context) {
		// Do some processing before calling the next middleware
		fmt.Println("Before calling the next middleware...")

		// Call the next middleware
		c.Next()

		// Do some processing after the next middleware has completed
		fmt.Println("After calling the next middleware...")
	})

	// Add a handler function that returns a simple message
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Start the server on port 8080
	r.Run(":8080")
}
