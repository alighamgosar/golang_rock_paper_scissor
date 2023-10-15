package main

import (
	"github.com/alighamgosar/rock_paper_scissors_golang/initializers"
	"github.com/alighamgosar/rock_paper_scissors_golang/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConntectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Tournament{})
	initializers.DB.AutoMigrate(&models.Game{})
	initializers.DB.AutoMigrate(&models.Set{})
	initializers.DB.AutoMigrate(&models.Player{})

}
