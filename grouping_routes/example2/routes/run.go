package routes

import "github.com/gin-gonic/gin"

var r = gin.Default()

func Run() {
	getRoutes()
	r.Run(":5000")
}

func getRoutes() {
	uni := r.Group("/university")
	addRoutes(uni)

	deg := r.Group("/degree")
	addDegreeRoutes(deg)
}
