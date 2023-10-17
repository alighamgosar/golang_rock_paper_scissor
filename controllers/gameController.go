package controllers

import (
	"github.com/alighamgosar/rock_paper_scissors_golang/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddTempGame(c *gin.Context) {

	Random_ID := utils.GenerateAndCheckUniqueRandID(&gorm.DB{})
	// _ = Random_ID

	// var tempGame struct {
	// 	ID          string
	// 	Player_1_ID uint
	// 	Player_2_ID uint
	// }

	// c.Bind(&tempGame)
	c.JSON(200, gin.H{
		"random_id": Random_ID,
	})
	// c.Status(200)

}

func AddTempGame1(c *gin.Context) {

	c.Status(200)

}
