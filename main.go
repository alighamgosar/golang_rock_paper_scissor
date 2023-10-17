package main

import (
	"github.com/alighamgosar/rock_paper_scissors_golang/controllers"
	"github.com/alighamgosar/rock_paper_scissors_golang/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConntectToDB()
}
func main() {
	r := gin.Default()

	r.PUT("/addGame", controllers.AddTempGame)
	r.GET("/addGame", controllers.AddTempGame1)

	r.Run()
}
