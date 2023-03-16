package routes

import "github.com/gin-gonic/gin"

func addDegreeRoutes(rg *gin.RouterGroup) {
	degree := rg.Group("/degree")

	degree.GET("/CS", func(c *gin.Context) {
		c.JSON(200, "Computer Science")
	})
}
