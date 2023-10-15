package models

import "gorm.io/gorm"

type Set struct {
	gorm.Model
	Set_Number         int    `gorm:"type:int;check:Set_Number BETWEEN 1 AND 3"`
	Game_ID            uint   `gorm:"not null"`
	Player_1_Selection string `gorm:"not null"`
	Player_2_Selection string `gorm:"not null"`
	WinnerID           uint   `gorm:"not null"`
}
