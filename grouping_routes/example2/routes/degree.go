package routes

import "github.com/gin-gonic/gin"

func addDegreeRoutes(rg *gin.RouterGroup) {
	degree := rg.Group("/deg")

	degree.GET("/", func(c *gin.Context) {
		c.String(200, "Degree Programs !!")
	})

	degree.GET("/CS", func(c *gin.Context) {
		c.JSON(200, "Computer Science")
	})

	degree.GET("/DSA", func(c *gin.Context) { //http://localhost:5000/degree/deg/DSA
		c.String(200, "Data Structures and Algorithm")
	})
}
