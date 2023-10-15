package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Nickname string `gorm:"size:15"`
}
