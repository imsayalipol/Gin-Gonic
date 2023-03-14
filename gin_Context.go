package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// create a new Gin engine
	r := gin.Default()

	// define a handler function
	handler := func(c *gin.Context) {
		// print some information about the request context
		fmt.Printf("Method: %s\n", c.Request.Method)
		fmt.Printf("URL: %s\n", c.Request.URL)
		fmt.Printf("User-Agent: %s\n", c.Request.UserAgent())
		fmt.Printf("Content-Type: %s\n", c.Request.Header.Get("Content-Type"))

		// send a response
		c.String(200, "Hello, world!")
	}

	// add the handler to the engine
	r.GET("/", handler)

	// start the server
	r.Run()
}
