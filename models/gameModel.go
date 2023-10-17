package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Random_ID      string
	Player_1_ID    uint  `gorm:"not null"`
	Player_2_ID    uint  `gorm:"not null"`
	Player_1_Score uint  `gorm:"not null"`
	Player_2_Score uint  `gorm:"not null"`
	Winner_ID      uint  `gorm:"not null"`
	Set            []Set `gorm:"foreignKey:id"`
}
